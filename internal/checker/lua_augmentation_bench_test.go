package checker_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/bundled"
	"github.com/apyrr/tlua/internal/compiler"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/tsoptions"
	"github.com/apyrr/tlua/internal/vfs/vfstest"
)

// luaAugmentationSource builds a program shaped like the two paths suspected of
// dominating augmentation cost: many constructor arms reached through aliases,
// and many narrowed member reads that drive isMatchingReference.
func luaAugmentationSource(roots int, membersPerRoot int, readsPerRoot int) string {
	var b strings.Builder
	for r := range roots {
		fmt.Fprintf(&b, "Root%d = {};\n", r)
		fmt.Fprintf(&b, "Root%d.nested = {};\n", r)
		for m := range membersPerRoot {
			fmt.Fprintf(&b, "Root%d.m%d = %d;\n", r, m, m)
			fmt.Fprintf(&b, "Root%d.nested.n%d = \"v%d\";\n", r, m, m)
		}
		// A stable alias so arm resolution has to canonicalize a root.
		fmt.Fprintf(&b, "local alias%d = Root%d.nested;\n", r, r)
		fmt.Fprintf(&b, "alias%d.viaAlias = true;\n", r)
		// A defaulted guard, the other constructor-discovery shape.
		fmt.Fprintf(&b, "Root%d.guard = Root%d.guard or {};\n", r, r)
		fmt.Fprintf(&b, "Root%d.guard.inner = %d;\n", r, r)
		for k := range readsPerRoot {
			m := k % membersPerRoot
			fmt.Fprintf(&b, "local function read%d_%d()\n", r, k)
			fmt.Fprintf(&b, "  if type(Root%d.m%d) == \"number\" then\n", r, m)
			fmt.Fprintf(&b, "    local a: number = Root%d.m%d;\n", r, m)
			fmt.Fprintf(&b, "    local c: string = Root%d.nested.n%d;\n", r, m)
			fmt.Fprintf(&b, "    local d: boolean = alias%d.viaAlias;\n", r)
			b.WriteString("  end\n")
			b.WriteString("end\n")
		}
	}
	return b.String()
}

func benchmarkLuaAugmentation(b *testing.B, roots, members, reads int) {
	b.Helper()
	source := luaAugmentationSource(roots, members, reads)
	fs := vfstest.FromMap(map[string]string{
		"/bench.tlua": source,
		"/tluaconfig.json": `{
			"compilerOptions": { "strict": true },
			"files": ["bench.tlua"]
		}`,
	}, false /*useCaseSensitiveFileNames*/)
	fs = bundled.WrapFS(fs)
	host := compiler.NewCompilerHost("/", fs, bundled.LibPath(), nil, nil)
	parsed, errors := tsoptions.GetParsedCommandLineOfConfigFile("/tluaconfig.json", &core.CompilerOptions{}, nil, host, nil)
	if len(errors) != 0 {
		b.Fatalf("config errors: %v", errors)
	}

	b.ReportAllocs()
	for b.Loop() {
		p := compiler.NewProgram(compiler.ProgramOptions{Config: parsed, Host: host})
		p.BindSourceFiles()
		file := p.GetSourceFile("/bench.tlua")
		p.GetSemanticDiagnostics(b.Context(), file)
	}
}

func BenchmarkLuaAugmentationWide(b *testing.B) {
	benchmarkLuaAugmentation(b, 40 /*roots*/, 8 /*members*/, 4 /*reads*/)
}

func BenchmarkLuaAugmentationDeepReads(b *testing.B) {
	benchmarkLuaAugmentation(b, 8 /*roots*/, 6 /*members*/, 30 /*reads*/)
}
