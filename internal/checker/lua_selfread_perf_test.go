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

// benchmarkLuaSelfReadGuards checks n files that each open with the defaulted
// guard of one shared global — the cross-file addon idiom whose self-reads all
// enter the ordered snapshot path.
func benchmarkLuaSelfReadGuards(b *testing.B, n int) {
	b.Helper()
	files := map[string]string{}
	names := make([]string, 0, n)
	for i := range n {
		name := fmt.Sprintf("g%03d.tlua", i)
		files["/"+name] = fmt.Sprintf("M = M or {};\nM.f%d = %d;\n", i, i)
		names = append(names, name)
	}
	quoted := make([]string, 0, n)
	for _, name := range names {
		quoted = append(quoted, fmt.Sprintf("%q", name))
	}
	files["/tluaconfig.json"] = fmt.Sprintf(`{"compilerOptions": {"strict": true}, "files": [%s]}`, strings.Join(quoted, ","))
	fs := vfstest.FromMap(files, false /*useCaseSensitiveFileNames*/)
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
		for _, name := range names {
			file := p.GetSourceFile("/" + name)
			p.GetSemanticDiagnostics(b.Context(), file)
		}
	}
}

func BenchmarkLuaSelfReadGuards64(b *testing.B)  { benchmarkLuaSelfReadGuards(b, 64) }
func BenchmarkLuaSelfReadGuards256(b *testing.B) { benchmarkLuaSelfReadGuards(b, 256) }
