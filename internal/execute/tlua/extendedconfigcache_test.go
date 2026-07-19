package tlua_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/execute/tlua"
	"github.com/apyrr/tlua/internal/tsoptions"
	"github.com/apyrr/tlua/internal/vfs"
	"github.com/apyrr/tlua/internal/vfs/vfstest"
)

type testParseConfigHost struct {
	fs  vfs.FS
	cwd string
}

func (h *testParseConfigHost) FS() vfs.FS                  { return h.fs }
func (h *testParseConfigHost) GetCurrentDirectory() string { return h.cwd }

func TestExtendedConfigCacheExtendsCircularity(t *testing.T) {
	t.Parallel()

	t.Run("self-referencing extends", func(t *testing.T) {
		t.Parallel()

		// Regression test: a tsconfig extends cycle should produce an error,
		// not a deadlock when using the tlua ExtendedConfigCache.
		files := map[string]any{
			"/project/tluaconfig.json": `{"extends": "./base.json"}`,
			"/project/base.json":       `{"extends": "./base.json"}`,
			"/project/main.tlua":       `// Hello World!`,
		}

		fs := vfstest.FromMap(files, false /*useCaseSensitiveFileNames*/)
		host := &testParseConfigHost{fs: fs, cwd: "/project"}
		cache := &tlua.ExtendedConfigCache{}

		cmd, _ := tsoptions.GetParsedCommandLineOfConfigFile("/project/tluaconfig.json", nil, nil, host, cache)
		if cmd == nil {
			t.Fatal("expected non-nil ParsedCommandLine")
		}
		assertHasCircularityDiagnostic(t, cmd)
	})

	t.Run("mutual extends cycle", func(t *testing.T) {
		t.Parallel()

		// Two config files that extend each other.
		files := map[string]any{
			"/project/tluaconfig.json": `{"extends": "./other.json"}`,
			"/project/other.json":      `{"extends": "./tluaconfig.json"}`,
			"/project/main.tlua":       `// Hello World!`,
		}

		fs := vfstest.FromMap(files, false /*useCaseSensitiveFileNames*/)
		host := &testParseConfigHost{fs: fs, cwd: "/project"}
		cache := &tlua.ExtendedConfigCache{}

		cmd, _ := tsoptions.GetParsedCommandLineOfConfigFile("/project/tluaconfig.json", nil, nil, host, cache)
		if cmd == nil {
			t.Fatal("expected non-nil ParsedCommandLine")
		}
		assertHasCircularityDiagnostic(t, cmd)
	})

	t.Run("case-insensitive self-referencing extends", func(t *testing.T) {
		t.Parallel()

		// On a case-insensitive FS, ./Base.json and ./base.json resolve to the same
		// cache entry. The cycle check must use canonical paths to avoid deadlock.
		files := map[string]any{
			"/project/tluaconfig.json": `{"extends": "./Base.json"}`,
			"/project/base.json":       `{"extends": "./base.json"}`,
			"/project/main.tlua":       `// Hello World!`,
		}

		fs := vfstest.FromMap(files, false /*useCaseSensitiveFileNames*/)
		host := &testParseConfigHost{fs: fs, cwd: "/project"}
		cache := &tlua.ExtendedConfigCache{}

		cmd, _ := tsoptions.GetParsedCommandLineOfConfigFile("/project/tluaconfig.json", nil, nil, host, cache)
		if cmd == nil {
			t.Fatal("expected non-nil ParsedCommandLine")
		}
		assertHasCircularityDiagnostic(t, cmd)
	})
}

func TestExtendedConfigCacheNullExtendsDoesNotPanic(t *testing.T) {
	t.Parallel()

	files := map[string]any{
		"/project/tluaconfig.json": `{"extends": null}`,
		"/project/main.tlua":       `// Hello World!`,
	}

	fs := vfstest.FromMap(files, false /*useCaseSensitiveFileNames*/)
	host := &testParseConfigHost{fs: fs, cwd: "/project"}
	cache := &tlua.ExtendedConfigCache{}

	cmd, _ := tsoptions.GetParsedCommandLineOfConfigFile("/project/tluaconfig.json", nil, nil, host, cache)
	if cmd == nil {
		t.Fatal("expected non-nil ParsedCommandLine")
	}
	if len(cmd.Errors) == 0 {
		t.Fatal("expected diagnostics for invalid null extends")
	}
}

func assertHasCircularityDiagnostic(t *testing.T, cmd *tsoptions.ParsedCommandLine) {
	t.Helper()
	for _, d := range cmd.Errors {
		if d != nil && d.Code() == 18000 {
			return
		}
	}
	t.Error("expected circularity diagnostic (code 18000), but none was found")
}
