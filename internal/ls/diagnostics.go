package ls

import (
	"context"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/compiler"
	"github.com/apyrr/tlua/internal/ls/lsconv"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
)

// getAllDiagnosticsCore collects all diagnostics for a file: syntactic,
// semantic, suggestion (when requested), and (when declarations are emitted)
// declaration diagnostics.
func getAllDiagnosticsCore(ctx context.Context, program *compiler.Program, file *ast.SourceFile, includeSuggestions bool) []*ast.Diagnostic {
	var diags []*ast.Diagnostic
	diags = append(diags, program.GetSyntacticDiagnostics(ctx, file)...)
	diags = append(diags, program.GetSemanticDiagnostics(ctx, file)...)
	if includeSuggestions {
		diags = append(diags, program.GetSuggestionDiagnostics(ctx, file)...)
	}
	if program.Options().GetEmitDeclarations() {
		diags = append(diags, program.GetDeclarationDiagnostics(ctx, file)...)
	}
	return diags
}

// getAllDiagnostics collects all diagnostics for a file, including suggestions.
func getAllDiagnostics(ctx context.Context, program *compiler.Program, file *ast.SourceFile) []*ast.Diagnostic {
	return getAllDiagnosticsCore(ctx, program, file, true /*includeSuggestions*/)
}

func (l *LanguageService) ProvideDiagnostics(ctx context.Context, uri lsproto.DocumentUri) (lsproto.DocumentDiagnosticResponse, error) {
	program, file := l.getProgramAndFile(uri)

	if l.UserPreferences().EnableValidation.IsFalse() {
		diagnostics := []*lsproto.Diagnostic{}
		return lsproto.RelatedFullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport{
			FullDocumentDiagnosticReport: &lsproto.RelatedFullDocumentDiagnosticReport{
				Items: diagnostics,
			},
		}, nil
	}

	diagnostics := getAllDiagnostics(ctx, program, file)

	return lsproto.RelatedFullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport{
		FullDocumentDiagnosticReport: &lsproto.RelatedFullDocumentDiagnosticReport{
			Items: l.toLSPDiagnostics(ctx, diagnostics),
		},
	}, nil
}

// ProvideFileDiagnostics computes LSP diagnostics for a single file of this
// language service's program. Unlike ProvideDiagnostics it takes the file
// directly (no URI lookup / panic path) and lets the caller choose whether
// suggestion diagnostics are included (workspace pulls skip them: the client
// favors document-pull results for open documents, so suggestions stay
// document-pull-only).
func (l *LanguageService) ProvideFileDiagnostics(ctx context.Context, file *ast.SourceFile, includeSuggestions bool) []*lsproto.Diagnostic {
	program := l.GetProgram()
	if program.GetSourceFileByPath(file.Path()) != file {
		// A file from another program would check against the wrong graph and
		// return silently wrong results; fail like getProgramAndFile does.
		panic("ProvideFileDiagnostics: file does not belong to this language service's program")
	}
	return l.toLSPDiagnostics(ctx, getAllDiagnosticsCore(ctx, program, file, includeSuggestions))
}

func (l *LanguageService) toLSPDiagnostics(ctx context.Context, diagnostics ...[]*ast.Diagnostic) []*lsproto.Diagnostic {
	size := 0
	for _, diagSlice := range diagnostics {
		size += len(diagSlice)
	}
	lspDiagnostics := make([]*lsproto.Diagnostic, 0, size)
	for _, diagSlice := range diagnostics {
		for _, diag := range diagSlice {
			lspDiagnostics = append(lspDiagnostics, lsconv.DiagnosticToLSPPull(ctx, l.converters, diag, l.UserPreferences().ReportStyleChecksAsWarnings.IsTrue()))
		}
	}
	return lspDiagnostics
}
