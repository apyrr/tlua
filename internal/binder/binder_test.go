package binder

import (
	"runtime"
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/parser"
	"github.com/apyrr/tlua/internal/testutil/fixtures"
	"github.com/apyrr/tlua/internal/tspath"
	"github.com/apyrr/tlua/internal/vfs/osvfs"
)

func TestLuaAssignmentListSymbolsStayTargetLocal(t *testing.T) {
	t.Parallel()

	parseOptions := ast.SourceFileParseOptions{
		FileName: "/assignment.tlua",
		Path:     "/assignment.tlua",
	}
	file := parser.ParseSourceFile(parseOptions, "Scalar = 1; First, Second = 2, 3;", core.ScriptKindTS)
	BindSourceFile(file)

	scalar := file.Statements.Nodes[0].Expression()
	listAssignment := file.Statements.Nodes[1].Expression()
	if scalar.Symbol() == nil {
		t.Fatal("scalar assignment should retain its declaration symbol")
	}
	if listAssignment.Symbol() != nil {
		t.Fatal("assignment list must not expose an arbitrary target symbol")
	}
	candidates := file.LuaWriteCandidates
	if len(candidates) != 3 {
		t.Fatalf("got %d candidates, want 3", len(candidates))
	}
	if candidates[1].Symbol == candidates[2].Symbol {
		t.Fatal("assignment list targets must have distinct declaration symbols")
	}
	// ValueIndex is what maps each target to its positional value.
	if candidates[0].ValueIndex != 0 || candidates[1].ValueIndex != 0 || candidates[2].ValueIndex != 1 {
		t.Fatalf("got value indexes %d/%d/%d, want 0/0/1",
			candidates[0].ValueIndex, candidates[1].ValueIndex, candidates[2].ValueIndex)
	}
	for _, candidate := range candidates {
		if candidate.Target == nil {
			t.Fatal("every assignment candidate should carry its target")
		}
		if candidate.Symbol.ValueDeclaration != candidate.Target {
			t.Fatal("assignment target should be its own value declaration")
		}
		if len(candidate.Symbol.Declarations) != 1 || candidate.Symbol.Declarations[0] != candidate.Target {
			t.Fatal("assignment target should be its own declaration span")
		}
	}
}

func BenchmarkBind(b *testing.B) {
	for _, f := range fixtures.BenchFixtures {
		b.Run(f.Name(), func(b *testing.B) {
			f.SkipIfNotExist(b)

			fileName := tspath.GetNormalizedAbsolutePath(f.Path(), "/")
			path := tspath.ToPath(fileName, "/", osvfs.FS().UseCaseSensitiveFileNames())
			sourceText := f.ReadFile(b)

			parseOptions := ast.SourceFileParseOptions{
				FileName: fileName,
				Path:     path,
			}
			// The fixture name carries the tlua extension; the on-disk path may
			// be an upstream .ts corpus file used purely as parse input.
			scriptKind := core.GetScriptKindFromFileName(f.Name())

			sourceFiles := make([]*ast.SourceFile, b.N)
			for i := range b.N {
				sourceFiles[i] = parser.ParseSourceFile(parseOptions, sourceText, scriptKind)
			}

			// The above parses do a lot of work; ensure GC is finished before we start collecting performance data.
			// GC must be called twice to allow things to settle.
			runtime.GC()
			runtime.GC()

			b.ResetTimer()
			for i := range b.N {
				BindSourceFile(sourceFiles[i])
			}
		})
	}
}
