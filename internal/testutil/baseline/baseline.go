package baseline

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/apyrr/tlua/internal/repo"
	"github.com/apyrr/tlua/internal/stringutil"
	"github.com/peter-evans/patience"
)

type Options struct {
	Subfolder string
}

const NoContent = "<no content>"

func Run(t *testing.T, fileName string, actual string, opts Options) {
	localPath := filepath.Join(localRoot, opts.Subfolder, fileName)
	referencePath := filepath.Join(referenceRoot, opts.Subfolder, fileName)

	// Record this baseline for tracking unused baselines
	recordBaseline(t, filepath.Join(opts.Subfolder, fileName))

	writeComparison(t, actual, localPath, referencePath, false)
}

func DiffText(oldName string, newName string, expected string, actual string) string {
	lines := patience.Diff(stringutil.SplitLines(expected), stringutil.SplitLines(actual))
	return patience.UnifiedDiffTextWithOptions(lines, patience.UnifiedDiffOptions{
		Precontext:  3,
		Postcontext: 3,
		SrcHeader:   oldName,
		DstHeader:   newName,
	})
}

func RunAgainstSubmodule(t *testing.T, fileName string, actual string, opts Options) {
	// Record this baseline for tracking unused baselines
	recordBaseline(t, filepath.Join(opts.Subfolder, fileName))

	local := filepath.Join(localRoot, opts.Subfolder, fileName)
	reference := filepath.Join(submoduleReferenceRoot, opts.Subfolder, fileName)
	writeComparison(t, actual, local, reference, true)
}

func writeComparison(t *testing.T, actualContent string, local, reference string, comparingAgainstSubmodule bool) {
	if actualContent == "" {
		panic("the generated content was \"\". Return 'baseline.NoContent' if no baselining is required.")
	}

	if err := os.MkdirAll(filepath.Dir(local), 0o755); err != nil {
		t.Error(fmt.Errorf("failed to create directories for the local baseline file %s: %w", local, err))
		return
	}

	if _, err := os.Stat(local); err == nil {
		if err := os.Remove(local); err != nil {
			t.Error(fmt.Errorf("failed to remove the local baseline file %s: %w", local, err))
			return
		}
	}

	expected := NoContent
	foundExpected := false
	if content, err := os.ReadFile(reference); err == nil {
		expected = string(content)
		foundExpected = true
	}

	if expected != actualContent || actualContent == NoContent && foundExpected {
		if actualContent == NoContent {
			if err := os.WriteFile(local+".delete", []byte{}, 0o644); err != nil {
				t.Error(fmt.Errorf("failed to write the local baseline file %s: %w", local+".delete", err))
				return
			}
		} else {
			if err := os.WriteFile(local, []byte(actualContent), 0o644); err != nil {
				t.Error(fmt.Errorf("failed to write the local baseline file %s: %w", local, err))
				return
			}
		}

		if _, err := os.Stat(reference); err != nil {
			if comparingAgainstSubmodule {
				t.Errorf("the baseline file %s does not exist in the TypeScript submodule", reference)
			} else {
				t.Errorf("new baseline created at %s.", local)
			}
		} else if comparingAgainstSubmodule {
			t.Errorf("the baseline file %s does not match the reference in the TypeScript submodule", reference)
		} else {
			t.Errorf("the baseline file %s has changed. (Run `hereby baseline-accept` if the new baseline is correct.)", reference)
		}
	}
}

var (
	localRoot              = filepath.Join(repo.TestDataPath(), "baselines", "local")
	referenceRoot          = filepath.Join(repo.TestDataPath(), "baselines", "reference")
	submoduleReferenceRoot = filepath.Join(repo.TypeScriptSubmodulePath(), "tests", "baselines", "reference")
)
