package ast_test

import (
	"strconv"
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/jsnum"
	"gotest.tools/v3/assert"
)

// The fast paths in NumberKeyNameFromInt and NumberKeyNameFromText must produce
// exactly the same name as the canonical NumberKeyName(jsnum.Number) form, or
// number keys built on different paths would silently split into two symbols.

func TestNumberKeyNameFromIntMatchesCanonical(t *testing.T) {
	t.Parallel()
	// Cover the cached range (0..31), just past it, and larger indexes.
	for _, i := range []int{0, 1, 2, 5, 31, 32, 33, 100, 65536, 1_000_000} {
		assert.Equal(t, ast.NumberKeyNameFromInt(i), ast.NumberKeyName(jsnum.Number(i)))
	}
}

func TestNumberKeyNameFromTextFastPathMatchesCanonical(t *testing.T) {
	t.Parallel()
	// Canonical decimal spellings take the fast path; all must match the value
	// they name via ToNumber.
	for _, text := range []string{"0", "1", "9", "10", "42", "100", "999999999999999"} {
		assert.Equal(t, ast.NumberKeyNameFromText(text), ast.NumberKeyName(jsnum.FromString(text)))
	}
	// Non-canonical spellings take the slow path but must still canonicalize to
	// the same key as their value.
	for _, text := range []string{"01", "1.0", "0x10", "+16", "1e1", "-1", "1.5", "9007199254740993"} {
		assert.Equal(t, ast.NumberKeyNameFromText(text), ast.NumberKeyName(jsnum.FromString(text)))
	}
}

// isCanonicalNumberKeyText is unexported; assert the property it guards via the
// public function: a fast-path spelling names its own text.
func TestNumberKeyNameFromTextCanonicalIsIdentity(t *testing.T) {
	t.Parallel()
	for _, text := range []string{"0", "7", "128", "999999999999999"} {
		value, ok := ast.NumberKeyValue(ast.NumberKeyNameFromText(text))
		assert.Assert(t, ok)
		assert.Equal(t, value, text)
	}
	// Leading zero and 16+ digit spellings are not canonical, so they are
	// re-normalized rather than passed through verbatim.
	name := ast.NumberKeyNameFromText("007")
	value, _ := ast.NumberKeyValue(name)
	assert.Equal(t, value, strconv.Itoa(7))
}
