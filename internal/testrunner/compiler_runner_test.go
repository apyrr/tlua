package testrunner

import (
	"testing"

	"github.com/apyrr/tlua/internal/bundled"
	"github.com/apyrr/tlua/internal/collections"
	"github.com/apyrr/tlua/internal/tspath"
	"gotest.tools/v3/assert"
)

// Runs the compiler tests and produces baselines (e.g. `test1.symbols`).
func TestLocal(t *testing.T) { runCompilerTests(t) } //nolint:paralleltest

func runCompilerTests(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		// Without embedding, we'd need to read all of the lib files out from disk into the MapFS.
		// Just skip this for now.
		t.Skip("bundled files are not embedded")
	}

	runners := []*CompilerBaselineRunner{
		NewCompilerBaselineRunner(TestTypeRegression),
		NewCompilerBaselineRunner(TestTypeConformance),
	}

	var seenTests collections.Set[string]
	for _, runner := range runners {
		for _, test := range runner.EnumerateTestFiles() {
			test = tspath.GetBaseFileName(test)
			assert.Assert(t, !seenTests.Has(test), "Duplicate test file: %s", test)
			seenTests.Add(test)
		}
	}

	for _, runner := range runners {
		runner.RunTests(t)
	}
}
