package ast

type TokenFlags int32

const (
	TokenFlagsNone                           TokenFlags = 0
	TokenFlagsPrecedingLineBreak             TokenFlags = 1 << 0
	TokenFlagsPrecedingJSDocComment          TokenFlags = 1 << 1
	TokenFlagsUnterminated                   TokenFlags = 1 << 2
	TokenFlagsExtendedUnicodeEscape          TokenFlags = 1 << 3  // e.g. `\u{10ffff}`
	TokenFlagsScientific                     TokenFlags = 1 << 4  // e.g. `10e2`
	TokenFlagsHexSpecifier                   TokenFlags = 1 << 6  // e.g. `0x00000000`
	TokenFlagsUnicodeEscape                  TokenFlags = 1 << 10 // e.g. `\u00a0`
	TokenFlagsContainsInvalidEscape          TokenFlags = 1 << 11 // e.g. `\uhello`
	TokenFlagsHexEscape                      TokenFlags = 1 << 12 // e.g. `\xa0`
	TokenFlagsContainsLeadingZero            TokenFlags = 1 << 13 // e.g. `0888`
	TokenFlagsPrecedingJSDocLeadingAsterisks TokenFlags = 1 << 15
	TokenFlagsSingleQuote                    TokenFlags = 1 << 16 // e.g. `'abc'`
	TokenFlagsPrecedingJSDocWithDeprecated   TokenFlags = 1 << 17 // Preceding JSDoc comment contains @deprecated
	TokenFlagsPrecedingJSDocWithSeeOrLink    TokenFlags = 1 << 18 // Preceding JSDoc comment contains @see or @link
	TokenFlagsWordOperator                   TokenFlags = 1 << 19 // tlua: an operator spelled as a word (`and`, `or`, `not`) rather than as punctuation
	// Lua has hexadecimal literals but no binary or octal form, and no numeric
	// separator, so `0x` is the only numeric specifier tlua scans.
	TokenFlagsWithSpecifier                 TokenFlags = TokenFlagsHexSpecifier
	TokenFlagsStringLiteralFlags            TokenFlags = TokenFlagsUnterminated | TokenFlagsHexEscape | TokenFlagsUnicodeEscape | TokenFlagsExtendedUnicodeEscape | TokenFlagsContainsInvalidEscape | TokenFlagsSingleQuote
	TokenFlagsNumericLiteralFlags           TokenFlags = TokenFlagsScientific | TokenFlagsContainsLeadingZero | TokenFlagsWithSpecifier
	TokenFlagsTemplateLiteralLikeFlags      TokenFlags = TokenFlagsUnterminated | TokenFlagsHexEscape | TokenFlagsUnicodeEscape | TokenFlagsExtendedUnicodeEscape | TokenFlagsContainsInvalidEscape
	TokenFlagsRegularExpressionLiteralFlags TokenFlags = TokenFlagsUnterminated
	TokenFlagsIsInvalid                     TokenFlags = TokenFlagsContainsLeadingZero | TokenFlagsContainsInvalidEscape
)
