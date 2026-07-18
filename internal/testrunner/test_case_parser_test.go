package testrunner

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"gotest.tools/v3/assert"
)

func TestMakeUnitsFromTest(t *testing.T) {
	t.Parallel()
	code := `// @strict: true
// @noEmit: true
// @filename: firstFile.tlua
function foo() { return "a"; }
// normal comment
// @filename: secondFile.tlua
// some other comment
function bar() { return "b"; }`
	testUnit1 := &testUnit{
		content: `function foo() { return "a"; }
// normal comment`,
		name: "firstFile.tlua",
	}
	testUnit2 := &testUnit{
		content: `// some other comment
function bar() { return "b"; }`,
		name: "secondFile.tlua",
	}
	testContent := testCaseContent{
		testUnitData:         []*testUnit{testUnit1, testUnit2},
		tsConfig:             nil,
		tsConfigFileUnitData: nil,
		symlinks:             make(map[string]string),
	}
	assert.DeepEqual(
		t,
		makeUnitsFromTest(code, "simpleTest.tlua"),
		testContent,
		cmp.AllowUnexported(testCaseContent{}, testUnit{}),
	)
}
