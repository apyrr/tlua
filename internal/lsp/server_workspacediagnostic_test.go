package lsp_test

import (
	"context"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/apyrr/tlua/internal/bundled"
	"github.com/apyrr/tlua/internal/json"
	"github.com/apyrr/tlua/internal/ls/lsconv"
	"github.com/apyrr/tlua/internal/ls/lsutil"
	"github.com/apyrr/tlua/internal/lsp"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil/lsptestutil"
	"github.com/apyrr/tlua/internal/vfs/iovfs"
	"github.com/apyrr/tlua/internal/vfs/vfstest"
	"gotest.tools/v3/assert"
)

// The workspace pull is defined over a whole project, so these tests share one
// three-file layout. Every file is a value-only module (no bare top-level
// `type`/`interface`, which tlua would hoist into the global scope and which
// would therefore invalidate the whole project on every edit), and `a` requires
// `b` so that editing `b` exercises reverse-dependency invalidation.
const (
	workspaceDiagRoot = "/root"

	// `require` takes a Lua module name resolved from the project root; a
	// path-shaped specifier such as "./b" is an error in tlua.
	workspaceDiagFileA = "local b = require(\"b\")\nlocal total = b.value\nreturn { total = total }\n"
	workspaceDiagFileB = "local x: number = \"oops\"\nreturn { value = 1 }\n"
	workspaceDiagFileC = "local y = 2\nreturn { y = y }\n"

	// The replacement content for `b` in the edit test: still exactly one error,
	// but a different one, and the module's shape is unchanged so `a` stays clean.
	workspaceDiagFileBEdited = "local x: string = 42\nreturn { value = 1 }\n"
)

func workspaceDiagFiles() map[string]string {
	return map[string]string{
		"/root/tluaconfig.json": `{ "compilerOptions": { "noEmit": true } }`,
		"/root/a.tlua":          workspaceDiagFileA,
		"/root/b.tlua":          workspaceDiagFileB,
		"/root/c.tlua":          workspaceDiagFileC,
	}
}

// initWorkspaceDiagClient brings up a real LSP session over an in-memory FS and
// returns the client, the underlying MapFS (so a test can mutate files
// mid-session and follow up with didChangeWatchedFiles) and the initialize
// result.
func initWorkspaceDiagClient(t *testing.T, files map[string]string) (*lsptestutil.LSPClient, *vfstest.MapFS, lsproto.InitializeResponse) {
	t.Helper()

	prefs := &lsutil.UserPreferences{}
	base := vfstest.FromMap(files, false)
	baseFS := base.(iovfs.FsWithSys).FSys().(*vfstest.MapFS)
	fs := bundled.WrapFS(base)

	onServerRequest := func(_ context.Context, req *lsproto.RequestMessage) *lsproto.ResponseMessage {
		switch req.Method {
		case lsproto.MethodWorkspaceConfiguration:
			return &lsproto.ResponseMessage{
				ID:      req.ID,
				JSONRPC: req.JSONRPC,
				Result:  []any{prefs},
			}
		case lsproto.MethodClientRegisterCapability, lsproto.MethodClientUnregisterCapability, lsproto.MethodWindowWorkDoneProgressCreate:
			return &lsproto.ResponseMessage{
				ID:      req.ID,
				JSONRPC: req.JSONRPC,
				Result:  lsproto.Null{},
			}
		default:
			return nil
		}
	}

	client, closeClient := lsptestutil.NewLSPClient(t, lsp.ServerOptions{
		Err:                io.Discard,
		Cwd:                workspaceDiagRoot,
		FS:                 fs,
		DefaultLibraryPath: bundled.LibPath(),
	}, onServerRequest)
	t.Cleanup(func() { _ = closeClient() })

	initMsg, initResult, ok := lsptestutil.SendRequest(t, client, lsproto.InitializeInfo, &lsproto.InitializeParams{
		Capabilities: &lsproto.ClientCapabilities{},
	})
	assert.Assert(t, ok && initMsg.AsResponse().Error == nil, "Initialize failed")
	lsptestutil.SendNotification(t, client, lsproto.InitializedInfo, &lsproto.InitializedParams{})
	<-client.Server.InitComplete()

	return client, baseFS, initResult
}

func workspaceDiagURI(fileName string) lsproto.DocumentUri {
	return lsconv.FileNameToDocumentURI(fileName)
}

func openWorkspaceDiagFile(t *testing.T, client *lsptestutil.LSPClient, fileName string, content string) {
	t.Helper()
	lsptestutil.SendNotification(t, client, lsproto.TextDocumentDidOpenInfo, &lsproto.DidOpenTextDocumentParams{
		TextDocument: &lsproto.TextDocumentItem{
			Uri:        workspaceDiagURI(fileName),
			LanguageId: "typescript",
			Text:       content,
		},
	})
}

type workspaceDiagReport = lsproto.WorkspaceFullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport

// pullWorkspaceDiagnostics sends one `workspace/diagnostic` request and asserts
// it succeeded.
func pullWorkspaceDiagnostics(t *testing.T, client *lsptestutil.LSPClient, previous []lsproto.PreviousResultId) []workspaceDiagReport {
	t.Helper()
	if previous == nil {
		previous = []lsproto.PreviousResultId{}
	}
	msg, resp, ok := lsptestutil.SendRequest(t, client, lsproto.WorkspaceDiagnosticInfo, &lsproto.WorkspaceDiagnosticParams{
		PreviousResultIds: previous,
	})
	assert.Assert(t, ok, "expected a workspace/diagnostic response")
	assert.Assert(t, msg.AsResponse().Error == nil, "workspace/diagnostic failed: %s", msg.AsResponse().Error)
	assert.Assert(t, resp != nil, "expected a workspace diagnostic report")
	return resp.Items
}

func workspaceDiagReportURI(item workspaceDiagReport) lsproto.DocumentUri {
	if item.FullDocumentDiagnosticReport != nil {
		return item.FullDocumentDiagnosticReport.Uri
	}
	return item.UnchangedDocumentDiagnosticReport.Uri
}

func workspaceDiagReportURIs(items []workspaceDiagReport) []lsproto.DocumentUri {
	uris := make([]lsproto.DocumentUri, 0, len(items))
	for _, item := range items {
		uris = append(uris, workspaceDiagReportURI(item))
	}
	return uris
}

func findWorkspaceDiagReport(items []workspaceDiagReport, uri lsproto.DocumentUri) (workspaceDiagReport, bool) {
	for _, item := range items {
		if workspaceDiagReportURI(item) == uri {
			return item, true
		}
	}
	return workspaceDiagReport{}, false
}

func fullWorkspaceDiagReportFor(t *testing.T, items []workspaceDiagReport, uri lsproto.DocumentUri) *lsproto.WorkspaceFullDocumentDiagnosticReport {
	t.Helper()
	item, ok := findWorkspaceDiagReport(items, uri)
	assert.Assert(t, ok, "expected a report for %s, got %v", uri, workspaceDiagReportURIs(items))
	assert.Assert(t, item.FullDocumentDiagnosticReport != nil, "expected a full report for %s, got an unchanged report", uri)
	return item.FullDocumentDiagnosticReport
}

func workspaceDiagMessage(diagnostic *lsproto.Diagnostic) string {
	if diagnostic.Message.String != nil {
		return *diagnostic.Message.String
	}
	if diagnostic.Message.MarkupContent != nil {
		return diagnostic.Message.MarkupContent.Value
	}
	return ""
}

func workspaceDiagSummary(diagnostics []*lsproto.Diagnostic) string {
	parts := make([]string, 0, len(diagnostics))
	for _, diagnostic := range diagnostics {
		code := ""
		if diagnostic.Code != nil && diagnostic.Code.Integer != nil {
			code = fmt.Sprintf("TS%d: ", *diagnostic.Code.Integer)
		}
		parts = append(parts, code+workspaceDiagMessage(diagnostic))
	}
	return strings.Join(parts, "; ")
}

func workspaceDiagCode(t *testing.T, diagnostic *lsproto.Diagnostic) int32 {
	t.Helper()
	assert.Assert(t, diagnostic.Code != nil && diagnostic.Code.Integer != nil, "expected a numeric diagnostic code")
	return *diagnostic.Code.Integer
}

// assertOnlyWorkspaceDiagFiles guards against lib/bundled files leaking into a
// workspace pull.
func assertOnlyWorkspaceDiagFiles(t *testing.T, items []workspaceDiagReport) {
	t.Helper()
	for _, uri := range workspaceDiagReportURIs(items) {
		assert.Assert(t, strings.HasPrefix(string(uri), "file:///root/"), "unexpected reported URI outside the workspace: %s", uri)
	}
}

func TestWorkspaceDiagnosticCapability(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	_, _, initResult := initWorkspaceDiagClient(t, workspaceDiagFiles())

	assert.Assert(t, initResult != nil)
	assert.Assert(t, initResult.Capabilities != nil)
	provider := initResult.Capabilities.DiagnosticProvider
	assert.Assert(t, provider != nil, "expected a diagnosticProvider capability")
	assert.Assert(t, provider.Options != nil, "expected diagnostic options, not registration options")
	assert.Equal(t, provider.Options.WorkspaceDiagnostics, true)
}

func TestWorkspaceDiagnosticColdPullCoversUnopenedFiles(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	client, _, _ := initWorkspaceDiagClient(t, workspaceDiagFiles())
	openWorkspaceDiagFile(t, client, "/root/a.tlua", workspaceDiagFileA)

	items := pullWorkspaceDiagnostics(t, client, nil)
	assertOnlyWorkspaceDiagFiles(t, items)

	// b is not open, yet its error is reported.
	report := fullWorkspaceDiagReportFor(t, items, workspaceDiagURI("/root/b.tlua"))
	assert.Equal(t, len(report.Items), 1, "b diagnostics: %s", workspaceDiagSummary(report.Items))
	assert.Equal(t, workspaceDiagCode(t, report.Items[0]), int32(2322))
	assert.Equal(t, workspaceDiagMessage(report.Items[0]), "Type 'string' is not assignable to type 'number'.")

	// Never-opened documents are reported against a null version.
	assert.Assert(t, report.Version.Integer == nil, "expected a null version for the unopened b.tlua")

	// The report is cacheable: the client gets an ID to send back next time.
	assert.Assert(t, report.ResultId != nil && *report.ResultId != "", "expected a non-empty resultId for b.tlua")

	// Clean files the client has never been told about are omitted entirely.
	_, hasA := findWorkspaceDiagReport(items, workspaceDiagURI("/root/a.tlua"))
	assert.Assert(t, !hasA, "clean, never-reported a.tlua should be omitted, got %v", workspaceDiagReportURIs(items))
	_, hasC := findWorkspaceDiagReport(items, workspaceDiagURI("/root/c.tlua"))
	assert.Assert(t, !hasC, "clean, never-reported c.tlua should be omitted, got %v", workspaceDiagReportURIs(items))
}

func TestWorkspaceDiagnosticUnchangedOnRepull(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	client, _, _ := initWorkspaceDiagClient(t, workspaceDiagFiles())
	openWorkspaceDiagFile(t, client, "/root/a.tlua", workspaceDiagFileA)

	bURI := workspaceDiagURI("/root/b.tlua")
	first := fullWorkspaceDiagReportFor(t, pullWorkspaceDiagnostics(t, client, nil), bURI)
	assert.Assert(t, first.ResultId != nil)

	items := pullWorkspaceDiagnostics(t, client, []lsproto.PreviousResultId{{Uri: bURI, Value: *first.ResultId}})
	assertOnlyWorkspaceDiagFiles(t, items)

	item, ok := findWorkspaceDiagReport(items, bURI)
	assert.Assert(t, ok, "expected a report for b.tlua, got %v", workspaceDiagReportURIs(items))
	assert.Assert(t, item.UnchangedDocumentDiagnosticReport != nil, "expected an unchanged report for b.tlua")
	assert.Equal(t, item.UnchangedDocumentDiagnosticReport.ResultId, *first.ResultId)
}

func TestWorkspaceDiagnosticEditInvalidatesSelectively(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	client, fs, _ := initWorkspaceDiagClient(t, workspaceDiagFiles())
	openWorkspaceDiagFile(t, client, "/root/a.tlua", workspaceDiagFileA)

	aURI := workspaceDiagURI("/root/a.tlua")
	bURI := workspaceDiagURI("/root/b.tlua")
	cURI := workspaceDiagURI("/root/c.tlua")

	first := fullWorkspaceDiagReportFor(t, pullWorkspaceDiagnostics(t, client, nil), bURI)
	assert.Assert(t, first.ResultId != nil)
	oldResultID := *first.ResultId

	// b stays unopened: the edit arrives on disk, announced by a watcher event.
	assert.NilError(t, fs.WriteFile("root/b.tlua", workspaceDiagFileBEdited, 0o666))
	lsptestutil.SendNotification(t, client, lsproto.WorkspaceDidChangeWatchedFilesInfo, &lsproto.DidChangeWatchedFilesParams{
		Changes: []*lsproto.FileEvent{{Uri: bURI, Type: lsproto.FileChangeTypeChanged}},
	})

	items := pullWorkspaceDiagnostics(t, client, []lsproto.PreviousResultId{{Uri: bURI, Value: oldResultID}})
	assertOnlyWorkspaceDiagFiles(t, items)

	report := fullWorkspaceDiagReportFor(t, items, bURI)
	assert.Equal(t, len(report.Items), 1, "b diagnostics: %s", workspaceDiagSummary(report.Items))
	assert.Equal(t, workspaceDiagCode(t, report.Items[0]), int32(2322))
	assert.Equal(t, workspaceDiagMessage(report.Items[0]), "Type 'number' is not assignable to type 'string'.")
	assert.Assert(t, report.ResultId != nil && *report.ResultId != "", "expected a resultId for the re-reported b.tlua")
	assert.Assert(t, *report.ResultId != oldResultID, "expected a new resultId after b.tlua changed")

	// a imports b, so it may be rechecked and re-reported; it must never come
	// back as anything other than a full report (it is still clean, so being
	// omitted is the expected outcome).
	if item, ok := findWorkspaceDiagReport(items, aURI); ok {
		assert.Assert(t, item.FullDocumentDiagnosticReport != nil, "a.tlua must not be reported unchanged against a stale id")
	}

	// c neither imports b nor was ever reported: it stays out of the response.
	_, hasC := findWorkspaceDiagReport(items, cURI)
	assert.Assert(t, !hasC, "untouched, never-reported c.tlua should stay absent, got %v", workspaceDiagReportURIs(items))
}

func TestWorkspaceDiagnosticDeletedFileIsCleared(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	client, fs, _ := initWorkspaceDiagClient(t, workspaceDiagFiles())
	openWorkspaceDiagFile(t, client, "/root/a.tlua", workspaceDiagFileA)

	aURI := workspaceDiagURI("/root/a.tlua")
	bURI := workspaceDiagURI("/root/b.tlua")

	first := fullWorkspaceDiagReportFor(t, pullWorkspaceDiagnostics(t, client, nil), bURI)
	assert.Assert(t, first.ResultId != nil)

	assert.NilError(t, fs.Remove("root/b.tlua"))
	lsptestutil.SendNotification(t, client, lsproto.WorkspaceDidChangeWatchedFilesInfo, &lsproto.DidChangeWatchedFilesParams{
		Changes: []*lsproto.FileEvent{{Uri: bURI, Type: lsproto.FileChangeTypeDeleted}},
	})

	items := pullWorkspaceDiagnostics(t, client, []lsproto.PreviousResultId{{Uri: bURI, Value: *first.ResultId}})
	assertOnlyWorkspaceDiagFiles(t, items)

	// An empty full report is what clears the client's state for a URI no
	// project reports on any more.
	cleared := fullWorkspaceDiagReportFor(t, items, bURI)
	assert.Equal(t, len(cleared.Items), 0, "expected b.tlua to be cleared, got %s", workspaceDiagSummary(cleared.Items))
	assert.Assert(t, cleared.ResultId == nil, "a clearing report must not hand out a result id")
	assert.Assert(t, cleared.Version.Integer == nil, "a clearing report has no version")

	// a is rechecked because it imported b, but an unresolved dotted `require`
	// falls back to tlua's ambient overloads (`any`) rather than erroring, so a
	// stays clean and is therefore omitted. Only assert that it never comes back
	// as `unchanged`.
	if item, ok := findWorkspaceDiagReport(items, aURI); ok {
		assert.Assert(t, item.FullDocumentDiagnosticReport != nil, "a.tlua must not be reported unchanged after its import disappeared")
	}
}

func TestWorkspaceDiagnosticDeterminism(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	// Two consecutive pulls in one session, with no edits and no previous ids,
	// must produce byte-identical items.
	client, _, _ := initWorkspaceDiagClient(t, workspaceDiagFiles())
	openWorkspaceDiagFile(t, client, "/root/a.tlua", workspaceDiagFileA)

	firstItems := pullWorkspaceDiagnostics(t, client, nil)
	secondItems := pullWorkspaceDiagnostics(t, client, nil)
	assert.DeepEqual(t, workspaceDiagReportURIs(firstItems), workspaceDiagReportURIs(secondItems))

	firstJSON, err := json.Marshal(firstItems)
	assert.NilError(t, err)
	secondJSON, err := json.Marshal(secondItems)
	assert.NilError(t, err)
	assert.Equal(t, string(firstJSON), string(secondJSON))

	// A second, independent session over the same files reports the same URIs.
	// (Result IDs carry a per-process nonce, so only the URI set is comparable.)
	other, _, _ := initWorkspaceDiagClient(t, workspaceDiagFiles())
	openWorkspaceDiagFile(t, other, "/root/a.tlua", workspaceDiagFileA)
	otherItems := pullWorkspaceDiagnostics(t, other, nil)
	assert.DeepEqual(t, workspaceDiagReportURIs(firstItems), workspaceDiagReportURIs(otherItems))
}

func TestWorkspaceDiagnosticCancellation(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	// A project big enough that a cold pull takes long enough to be cancelled
	// mid-flight. This is inherently a race, so the test accepts either outcome
	// and only pins down what a cancelled pull must look like.
	files := map[string]string{
		"/root/tluaconfig.json": `{ "compilerOptions": { "noEmit": true } }`,
	}
	for i := range 300 {
		files[fmt.Sprintf("/root/f%d.tlua", i)] = fmt.Sprintf("local x%d: number = \"oops\"\nreturn { value = x%d }\n", i, i)
	}

	client, _, _ := initWorkspaceDiagClient(t, files)
	openWorkspaceDiagFile(t, client, "/root/f0.tlua", files["/root/f0.tlua"])

	// Wait for the project to be loaded before starting the pull. A
	// cancellation for a request the dispatch loop has not picked up yet is
	// dropped (it is not yet in the server's pending-request table), and while
	// the didOpen is still building the program the pull is exactly that.
	warmupMsg, _, warmupOk := lsptestutil.SendRequest(t, client, lsproto.CustomProjectInfoInfo, &lsproto.ProjectInfoParams{
		TextDocument: lsproto.TextDocumentIdentifier{Uri: workspaceDiagURI("/root/f0.tlua")},
	})
	assert.Assert(t, warmupOk && warmupMsg.AsResponse().Error == nil, "failed to load the project")

	// SendRequestAsync mints the next request ID; burn one to learn what it
	// will be, since the client exposes no other way to name the in-flight
	// request to $/cancelRequest.
	pullID := client.NextID() + 1
	waitForPull := lsptestutil.SendRequestAsync(t, client, lsproto.WorkspaceDiagnosticInfo, &lsproto.WorkspaceDiagnosticParams{
		PreviousResultIds: []lsproto.PreviousResultId{},
	})
	// Land the cancellation inside the checking loop rather than before the
	// handler starts.
	time.Sleep(2 * time.Millisecond)
	lsptestutil.SendNotification(t, client, lsproto.CancelRequestInfo, &lsproto.CancelParams{
		Id: lsproto.IntegerOrString{Integer: &pullID},
	})

	msg, resp, _ := waitForPull()
	assert.Assert(t, msg != nil, "expected a response to the workspace pull")
	responseError := msg.AsResponse().Error
	if responseError == nil {
		// The pull won the race and completed normally.
		t.Log("the pull completed before the cancellation landed")
		assert.Assert(t, resp != nil, "expected a workspace diagnostic report")
		return
	}

	// Cancelled: the code must be ServerCancelled, not RequestCancelled, and it
	// must carry the retrigger flag, or vscode-languageclient gives up on
	// workspace pulls after a handful of cancellations.
	assert.Equal(t, int32(lsproto.ErrorCodeServerCancelled), int32(-32802))
	assert.Equal(t, responseError.Code, int32(lsproto.ErrorCodeServerCancelled))
	assert.Assert(t, responseError.Data != nil, "expected DiagnosticServerCancellationData on a cancelled pull")
	raw, err := json.Marshal(responseError.Data)
	assert.NilError(t, err)
	var data lsproto.DiagnosticServerCancellationData
	assert.NilError(t, json.Unmarshal(raw, &data))
	assert.Equal(t, data.RetriggerRequest, true)
}
