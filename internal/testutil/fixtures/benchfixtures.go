package fixtures

import (
	"path/filepath"

	"github.com/apyrr/tlua/internal/repo"
	"github.com/apyrr/tlua/internal/testutil/filefixture"
)

// BenchFixtures is the in-repo parse/bind benchmark corpus. The upstream
// checker.ts/dom.generated.d.ts fixtures went away with the TypeScript
// submodule; the largest bundled lib and the largest in-repo test case
// replace them so the benchmarks keep running.
var BenchFixtures = []filefixture.Fixture{
	filefixture.FromString("empty.tlua", "empty.tlua", ""),
	filefixture.FromFile("lib.luajit.d.tlua", filepath.Join(repo.RootPath(), "internal", "bundled", "libs", "lib.luajit.d.tlua")),
	// The upstream Herebyfile.mjs JS-parse fixture is gone: tlua has no
	// JavaScript inputs (.mjs is not a recognized extension).
	filefixture.FromFile("tluaMultipleAssignment.tlua", filepath.Join(repo.TestDataPath(), "tests", "cases", "compiler", "tluaMultipleAssignment.tlua")),
}
