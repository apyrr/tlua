package parser

// ParseFlags

type ParseFlags uint32

const (
	ParseFlagsNone ParseFlags = 0
	// 1 << 0 was ParseFlagsYield and 1 << 1 was ParseFlagsAwait; generators,
	// await expressions, and their parsing contexts are removed in tlua.
	ParseFlagsType                   ParseFlags = 1 << 2
	ParseFlagsIgnoreMissingOpenBrace ParseFlags = 1 << 4
	ParseFlagsJSDoc                  ParseFlags = 1 << 5
)
