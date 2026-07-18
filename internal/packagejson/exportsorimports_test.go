package packagejson_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/json"
	"github.com/apyrr/tlua/internal/packagejson"
	"gotest.tools/v3/assert"
)

func TestExports(t *testing.T) {
	t.Parallel()

	t.Run("UnmarshalJSONV2", func(t *testing.T) {
		t.Parallel()
		testExports(t, func(in []byte, out any) error { return json.Unmarshal(in, out) })
	})
}

func testExports(t *testing.T, unmarshal func([]byte, any) error) {
	type Exports struct {
		Imports packagejson.ExportsOrImports `json:"imports"`
		Exports packagejson.ExportsOrImports `json:"exports"`
	}

	var e Exports

	jsonString := `{
		"imports": {
			"#foo": {
				"import": "./foo.tlua"
			}
		},
		"exports": {
			".": {
				"import": "./test.tlua",
				"default": "./test.tlua"
			},
			"./test": [
				"./test1.tlua",
				"./test2.tlua",
				null
			],
			"./null": null
		}
	}`

	err := unmarshal([]byte(jsonString), &e)
	assert.NilError(t, err)

	assert.Assert(t, e.Exports.IsSubpaths())
	assert.Equal(t, e.Exports.AsObject().Size(), 3)
	assert.Assert(t, e.Exports.AsObject().GetOrZero(".").IsConditions())
	assert.Assert(t, e.Exports.AsObject().GetOrZero(".").AsObject().GetOrZero("import").Type == packagejson.JSONValueTypeString)
	assert.Equal(t, e.Exports.AsObject().GetOrZero("./test").AsArray()[2].Type, packagejson.JSONValueTypeNull)
	assert.Assert(t, e.Exports.AsObject().GetOrZero("./null").Type == packagejson.JSONValueTypeNull)

	assert.Assert(t, e.Imports.IsImports())
	assert.Equal(t, e.Imports.AsObject().Size(), 1)
	assert.Assert(t, e.Imports.AsObject().GetOrZero("#foo").IsConditions())
	assert.Assert(t, e.Imports.AsObject().GetOrZero("#foo").AsObject().GetOrZero("import").Type == packagejson.JSONValueTypeString)
}
