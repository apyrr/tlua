package lsp

import (
	"context"

	"github.com/apyrr/tlua/internal/collections"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/ls/lsconv"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/project"
)

// workspaceDiagnosticsCancelledError is the answer to a `workspace/diagnostic`
// pull that was cancelled partway through.
//
// It must be reported as ServerCancelled carrying DiagnosticServerCancellationData
// rather than as a plain RequestCancelled: vscode-languageclient shuts its
// workspace pull loop down after five consecutive error responses, and the only
// errors exempted from that count are the ones carrying this data (see
// pullWorkspace in the client's diagnostic.js). Without the data, cancelling a
// few pulls in a row -- which is exactly what typing does -- would silently
// stop workspace diagnostics for the rest of the session.
type workspaceDiagnosticsCancelledError struct{}

func (workspaceDiagnosticsCancelledError) Error() string {
	return "workspace diagnostics request cancelled"
}

func (workspaceDiagnosticsCancelledError) Unwrap() error {
	return lsproto.ErrorCodeServerCancelled
}

func (workspaceDiagnosticsCancelledError) ResponseErrorData() any {
	return lsproto.DiagnosticServerCancellationData{RetriggerRequest: true}
}

var _ responseErrorData = workspaceDiagnosticsCancelledError{}

// handleWorkspaceDiagnostic answers a `workspace/diagnostic` pull: the client
// asks for the diagnostics of every file in the workspace, handing back the
// result IDs it already holds, and gets a full report for the files whose
// diagnostics it does not have and an `unchanged` marker for the rest.
//
// All of the state that makes the common case cheap lives in
// project.WorkspaceDiagCache; this function only decides what to send.
//
// TODO: workspace-wide config discovery. Only projects the ProjectCollection
// already knows about are pulled. A tluaconfig project none of whose files has
// ever been opened is not discovered -- WithSnapshotLoadingProjectTree loads
// ancestor trees through DidRequestProjectTrees, which iterates the configured
// projects that exist (internal/project/projectcollectionbuilder.go) rather
// than searching the workspace for config files. Invisible in single-project
// workspaces (and at parity with tsserver's geterrForProject, which is also
// loaded-projects-only); matters for multi-config monorepos. A fix would
// enumerate tluaconfig.json files under the workspace folders and load them
// through the existing configured-project machinery on the background queue --
// this handler and the cache need no changes, they iterate whatever exists.
func (s *Server) handleWorkspaceDiagnostic(ctx context.Context, params *lsproto.WorkspaceDiagnosticParams) (lsproto.WorkspaceDiagnosticResponse, error) {
	s.workspacePullMu.Lock()
	defer s.workspacePullMu.Unlock()
	// A pull the client abandoned while it waited behind the previous one has
	// nothing left to do; answer before pruning, reconciling or walking the
	// project tree on its behalf.
	if ctx.Err() != nil {
		return nil, workspaceDiagnosticsCancelledError{}
	}

	// The request ID stays on the context: files are checked one after another,
	// so keeping the checkers affine to this request is what makes the loop
	// reuse a single checker instead of one per file.
	ctx = core.WithCheckerLifetime(ctx, core.CheckerLifetimeDiagnostics)

	var items []lsproto.WorkspaceFullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport
	var cancelled bool
	// A project whose program is not built in this snapshot contributes no
	// reports and, crucially, no `seen` entries: running the stale-URI cleanup
	// then would clear every diagnostic of that project for one pull and
	// resend them all on the next -- flicker for nothing. Staleness for one
	// cycle is the better trade; the next pull cleans up properly.
	var unreportedProject bool

	s.session.WithSnapshotLoadingProjectTree(ctx, nil, func(snapshot *project.Snapshot) {
		cache := s.session.WorkspaceDiagnosticsCache()
		cache.Prune(snapshot)
		// The cached items are already converted to LSP severities, so a change
		// to the preference that decides them invalidates every entry.
		cache.ReconcilePreferences(snapshot.UserPreferences().ReportStyleChecksAsWarnings.IsTrue())

		previous := make(map[lsproto.DocumentUri]string, len(params.PreviousResultIds))
		for _, previousResult := range params.PreviousResultIds {
			previous[previousResult.Uri] = previousResult.Value
		}
		// A file can belong to more than one project; the first project that
		// reports it wins. Projects() is sorted, so which one that is stays
		// stable across pulls.
		seen := collections.Set[lsproto.DocumentUri]{}

		// The version a report is tagged with is only meaningful for documents
		// the editor owns; for everything else the content on disk is the
		// truth and the protocol wants null.
		versionOf := func(fileName string) lsproto.IntegerOrNull {
			if handle := snapshot.GetFile(fileName); handle != nil && handle.IsOverlay() {
				return lsproto.IntegerOrNull{Integer: new(handle.Version())}
			}
			return lsproto.IntegerOrNull{}
		}

		// With validation off there is nothing to report, but the stale-URI
		// cleanup below still runs: every result the client holds is cleared.
		validationOff := snapshot.UserPreferences().EnableValidation.IsFalse()
		if !validationOff {
		projects:
			for _, proj := range snapshot.ProjectCollection.Projects() {
				program := proj.GetProgram()
				if program == nil {
					unreportedProject = true
					continue
				}
				cache.Reconcile(proj)
				languageService := project.NewLanguageServiceForProject(proj, snapshot)
				if languageService == nil {
					unreportedProject = true
					continue
				}
				projectKey := proj.Id()
				for _, file := range program.GetSourceFiles() {
					if !project.IsWorkspaceDiagnosticFile(program, file) {
						continue
					}
					if ctx.Err() != nil {
						cancelled = true
						break projects
					}
					uri := lsconv.FileNameToDocumentURI(file.FileName())
					if !seen.AddIfAbsent(uri) {
						continue
					}
					version := versionOf(file.FileName())
					previousID, clientHasResult := previous[uri]

					status, cached := cache.Status(projectKey, file.Path())
					if cached && status.Valid {
						if clientHasResult && previousID == status.ResultID {
							items = append(items, unchangedWorkspaceReport(status.ResultID, uri, version))
							continue
						}
						if len(status.Items) == 0 && !status.Reported && !clientHasResult {
							// A clean file the client has never been told
							// about: sending an empty report would only make
							// it start tracking one.
							continue
						}
						items = append(items, fullWorkspaceReport(status.ResultID, status.Items, uri, version))
						cache.MarkReported(projectKey, file.Path(), status.ResultID)
						continue
					}

					diagnostics := languageService.ProvideFileDiagnostics(ctx, file, false /*includeSuggestions*/)
					if ctx.Err() != nil {
						// The diagnostics of a cancelled check are not
						// trustworthy; drop them rather than caching them.
						cancelled = true
						break projects
					}

					var resultID string
					if cached {
						// The generation and hash the diagnostics were computed
						// against: if the program moved underneath us, the
						// store is refused and this report goes out uncached.
						storedID, changed, stored := cache.StoreComputed(
							projectKey,
							file.Path(),
							status.Generation,
							status.FileHash,
							diagnostics,
						)
						if stored {
							resultID = storedID
							if !changed && clientHasResult && previousID == storedID {
								items = append(items, unchangedWorkspaceReport(storedID, uri, version))
								continue
							}
						}
					}
					if len(diagnostics) == 0 && !status.Reported && !clientHasResult {
						continue
					}
					items = append(items, fullWorkspaceReport(resultID, diagnostics, uri, version))
					if resultID != "" {
						cache.MarkReported(projectKey, file.Path(), resultID)
					}
				}
			}
		}

		if cancelled {
			return
		}

		// Files the client still holds a result for but that no project reports
		// on any more (deleted, excluded from the program, or in a project that
		// went away): an empty full report is what clears them. Skipped while a
		// project could not report (program still building): its files are
		// absent from `seen` for that reason, not because they are gone, and
		// clearing them would flicker the Problems panel. The validation-off
		// wipe is intentional and still runs.
		if unreportedProject && !validationOff {
			return
		}
		for _, previousResult := range params.PreviousResultIds {
			if seen.Has(previousResult.Uri) {
				continue
			}
			items = append(items, fullWorkspaceReport(
				"",  /*resultID*/
				nil, /*diagnostics*/
				previousResult.Uri,
				lsproto.IntegerOrNull{},
			))
			// The cache is deliberately left alone: Reconcile owns entry
			// lifecycle, and a file that left the program was already dropped
			// by its file-set rebuild. Dropping entries here would be worse
			// than redundant — on the validation-off path the program has not
			// changed, so Reconcile's generation fast path would never
			// recreate them, and every later pull would recompute those files
			// and send full reports with no result ID, forever. A retired ID
			// can never be answered `unchanged` anyway: the clearing report
			// above carries no result ID, so the client stops sending one.
		}
	})

	if cancelled {
		return nil, workspaceDiagnosticsCancelledError{}
	}
	return &lsproto.WorkspaceDiagnosticReport{Items: items}, nil
}

// fullWorkspaceReport builds a full report for one file. An empty resultID
// means the report is not cached: the client simply has no previous ID to send
// back for this file next time.
func fullWorkspaceReport(
	resultID string,
	diagnostics []*lsproto.Diagnostic,
	uri lsproto.DocumentUri,
	version lsproto.IntegerOrNull,
) lsproto.WorkspaceFullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport {
	if diagnostics == nil {
		diagnostics = []*lsproto.Diagnostic{}
	}
	report := &lsproto.WorkspaceFullDocumentDiagnosticReport{
		Items:   diagnostics,
		Uri:     uri,
		Version: version,
	}
	if resultID != "" {
		report.ResultId = &resultID
	}
	return lsproto.WorkspaceFullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport{
		FullDocumentDiagnosticReport: report,
	}
}

func unchangedWorkspaceReport(
	resultID string,
	uri lsproto.DocumentUri,
	version lsproto.IntegerOrNull,
) lsproto.WorkspaceFullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport {
	return lsproto.WorkspaceFullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport{
		UnchangedDocumentDiagnosticReport: &lsproto.WorkspaceUnchangedDocumentDiagnosticReport{
			ResultId: resultID,
			Uri:      uri,
			Version:  version,
		},
	}
}
