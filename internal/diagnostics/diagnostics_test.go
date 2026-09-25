package diagnostics

import (
	"testing"

	"github.com/apyrr/tlua/internal/locale"
	"golang.org/x/text/language"
	"gotest.tools/v3/assert"
)

func TestLocalize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		message  *Message
		locale   locale.Locale
		args     []any
		expected string
	}{
		{
			name:     "english default",
			message:  Identifier_expected,
			locale:   locale.Locale(language.English),
			expected: "Identifier expected.",
		},
		{
			name:     "undefined locale uses english",
			message:  Identifier_expected,
			locale:   locale.Locale(language.Und),
			expected: "Identifier expected.",
		},
		{
			name:     "with single argument",
			message:  X_0_expected,
			locale:   locale.Locale(language.English),
			args:     []any{")"},
			expected: "')' expected.",
		},
		{
			name:     "with multiple arguments",
			message:  The_parser_expected_to_find_a_1_to_match_the_0_token_here,
			locale:   locale.Locale(language.English),
			args:     []any{"{", "}"},
			expected: "The parser expected to find a '}' to match the '{' token here.",
		},
		{
			name:     "fallback to english for unknown locale",
			message:  Identifier_expected,
			locale:   locale.Locale(language.MustParse("af-ZA")),
			expected: "Identifier expected.",
		},
		{
			name:     "german",
			message:  Identifier_expected,
			locale:   locale.Locale(language.MustParse("de-DE")),
			expected: "Es wurde ein Bezeichner erwartet.",
		},
		{
			name:     "french",
			message:  Identifier_expected,
			locale:   locale.Locale(language.MustParse("fr-FR")),
			expected: "Identificateur attendu.",
		},
		{
			name:     "spanish",
			message:  Identifier_expected,
			locale:   locale.Locale(language.MustParse("es-ES")),
			expected: "Se esperaba un identificador.",
		},
		{
			name:     "japanese",
			message:  Identifier_expected,
			locale:   locale.Locale(language.MustParse("ja-JP")),
			expected: "識別子が必要です。",
		},
		{
			name:     "chinese simplified",
			message:  Identifier_expected,
			locale:   locale.Locale(language.MustParse("zh-CN")),
			expected: "应为标识符。",
		},
		{
			name:     "korean",
			message:  Identifier_expected,
			locale:   locale.Locale(language.MustParse("ko-KR")),
			expected: "식별자가 필요합니다.",
		},
		{
			name:     "russian",
			message:  Identifier_expected,
			locale:   locale.Locale(language.MustParse("ru-RU")),
			expected: "Ожидался идентификатор.",
		},
		{
			// The key keeps only the first 100 characters of the text, so it survived
			// the rewording that changed 'undefined' to 'nil'; the archived German
			// translation still says 'undefined' and must not be shown.
			name:     "reworded message falls back to english",
			message:  Argument_of_type_0_is_not_assignable_to_parameter_of_type_1_with_exactOptionalPropertyTypes_Colon_true_Consider_adding_nil_to_the_types_of_the_target_s_properties,
			locale:   locale.Locale(language.MustParse("de-DE")),
			args:     []any{"A", "B"},
			expected: "Argument of type 'A' is not assignable to parameter of type 'B' with 'exactOptionalPropertyTypes: true'. Consider adding 'nil' to the types of the target's properties.",
		},
		{
			name:     "german with args",
			message:  X_0_expected,
			locale:   locale.Locale(language.MustParse("de-DE")),
			args:     []any{")"},
			expected: "\")\" wurde erwartet.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := tt.message.Localize(tt.locale, tt.args...)
			assert.Equal(t, result, tt.expected)
		})
	}
}

func TestLocalize_ByKey(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		key      Key
		locale   locale.Locale
		args     []string
		expected string
	}{
		{
			name:     "by key without args",
			key:      "Identifier_expected_1003",
			locale:   locale.Locale(language.English),
			expected: "Identifier expected.",
		},
		{
			name:     "by key with args",
			key:      "_0_expected_1005",
			locale:   locale.Locale(language.English),
			args:     []string{")"},
			expected: "')' expected.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := Localize(tt.locale, nil, tt.key, tt.args...)
			assert.Equal(t, result, tt.expected)
		})
	}
}
