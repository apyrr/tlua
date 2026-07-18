package fixtures

import (
	"path/filepath"

	"github.com/apyrr/tlua/internal/repo"
	"github.com/apyrr/tlua/internal/testutil/filefixture"
)

var BenchFixtures = []filefixture.Fixture{
	filefixture.FromString("empty.tlua", "empty.tlua", ""),
	filefixture.FromFile("checker.tlua", filepath.Join(repo.TypeScriptSubmodulePath(), "src/compiler/checker.ts")),
	filefixture.FromFile("dom.generated.d.tlua", filepath.Join(repo.TypeScriptSubmodulePath(), "src/lib/dom.generated.d.ts")),
	// The upstream Herebyfile.mjs JS-parse fixture is gone: tlua has no
	// JavaScript inputs (.mjs is not a recognized extension).
	filefixture.FromFile("jsxComplexSignatureHasApplicabilityError.tsx", filepath.Join(repo.TypeScriptSubmodulePath(), "tests/cases/compiler/jsxComplexSignatureHasApplicabilityError.tsx")),
}
