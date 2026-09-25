package scanner

import (
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/diagnostics"
	"gotest.tools/v3/assert"
)

func TestScanStringLuaEscapes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		text    string
		value   string
		invalid bool
	}{
		{text: `"\xff\255\0\65"`, value: "\xff\xff\x00A"},
		{text: `"\u{48}\u{e9}\u{1F980}"`, value: "Hé🦀"},
		{text: `"\a\b\f\n\r\t\v\\\"\'"`, value: "\a\b\f\n\r\t\v\\\"'"},
		{text: "\"a\\z \n\t b\"", value: "ab"},
		{text: "\"a\\\r\nb\"", value: "a\nb"},
		// LuaJIT's \u{...} encodes Unicode scalar values only.
		{text: `"\u{D83D}"`, value: `\u{D83D}`, invalid: true},
		{text: `"\256"`, value: `\256`, invalid: true},
		{text: `"\q"`, value: `\q`, invalid: true},
		{text: "\"\\`\"", value: "\\`", invalid: true},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			t.Parallel()
			s := NewScanner()
			errors := 0
			s.SetOnError(func(*diagnostics.Message, int, int, ...any) { errors++ })
			s.SetText(tt.text)
			assert.Equal(t, s.Scan(), ast.KindStringLiteral)
			assert.Equal(t, s.TokenValue(), tt.value)
			assert.Equal(t, errors != 0, tt.invalid)
		})
	}
}

func TestScanLocalKeyword(t *testing.T) {
	t.Parallel()
	s := NewScanner()
	s.SetText("local")
	assert.Equal(t, s.Scan(), ast.KindLocalKeyword)
}

func TestScanEndKeyword(t *testing.T) {
	t.Parallel()
	s := NewScanner()
	s.SetText("end")
	assert.Equal(t, s.Scan(), ast.KindEndKeyword)
}

func TestScanSelfKeyword(t *testing.T) {
	t.Parallel()
	s := NewScanner()
	s.SetText("self")
	assert.Equal(t, s.Scan(), ast.KindSelfKeyword)
}
