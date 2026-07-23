package ast

type NodeFlags uint32

const (
	NodeFlagsNone NodeFlags = 0
	// Let and Const survive the deletion of the `let`/`const` surface syntax as
	// internal scope markers: the ambient-global form `declare name: Type`
	// desugars to a Const list (read-only global), and Let marks the residual
	// synthesized block-scoped lists. A plain `local` list carries only
	// NodeFlagsLuaLocal — its block scoping comes from the Lua-local binder
	// path, not the Let bit. Bit 2 (Using) is free.
	NodeFlagsLet                             NodeFlags = 1 << 0  // Variable declaration
	NodeFlagsConst                           NodeFlags = 1 << 1  // Variable declaration
	NodeFlagsReparsed                        NodeFlags = 1 << 3  // Node was synthesized during parsing
	NodeFlagsSynthesized                     NodeFlags = 1 << 4  // Node was synthesized during transformation
	NodeFlagsOptionalChain                   NodeFlags = 1 << 5  // Chained MemberExpression rooted to a pseudo-OptionalExpression
	NodeFlagsExportContext                   NodeFlags = 1 << 6  // Export context (initialized by binding)
	NodeFlagsContainsSelf                    NodeFlags = 1 << 7  // Interface contains references to the polymorphic `self` type
	NodeFlagsHasImplicitReturn               NodeFlags = 1 << 8  // If function implicitly returns on one of codepaths (initialized by binding)
	NodeFlagsHasExplicitReturn               NodeFlags = 1 << 9  // If function has explicit reachable return on one of codepaths (initialized by binding)
	NodeFlagsDisallowInContext               NodeFlags = 1 << 10 // If node was parsed in a context where 'in-expressions' are not allowed
	NodeFlagsYieldContext                    NodeFlags = 1 << 11 // If node was parsed in the 'yield' context created when parsing a generator
	NodeFlagsAwaitContext                    NodeFlags = 1 << 13 // If node was parsed in the 'await' context created when parsing an async function
	NodeFlagsDisallowConditionalTypesContext NodeFlags = 1 << 14 // If node was parsed in a context where conditional types are not allowed
	NodeFlagsThisNodeHasError                NodeFlags = 1 << 15 // If the parser encountered an error when parsing the code that created this node
	NodeFlagsJavaScriptFile                  NodeFlags = 1 << 16 // If node was parsed in a JavaScript
	NodeFlagsThisNodeOrAnySubNodesHasError   NodeFlags = 1 << 17 // If this node or any of its children had an error
	NodeFlagsHasAsyncFunctions               NodeFlags = 1 << 18 // If the file has async functions (initialized by binding)
	// NodeFlagsHasAggregatedChildData is deprecated. Use `subtreeFacts` instead.

	// These flags will be set when the parser encounters a dynamic import expression or 'import.meta' to avoid
	// walking the tree if the flags are not set. However, these flags are just a approximation
	// (hence why it's named "PossiblyContainsDynamicImport") because once set, the flags never get cleared.
	// During editing, if a dynamic import is removed, incremental parsing will *NOT* clear this flag.
	// This means that the tree will always be traversed during module resolution, or when looking for external module indicators.
	// However, the removal operation should not occur often and in the case of the
	// removal, it is likely that users will add the import anyway.
	// The advantage of this approach is its simplicity. For the case of batch compilation,
	// we guarantee that users won't have to pay the price of walking the tree if a dynamic import isn't used.
	NodeFlagsPossiblyContainsDynamicImport NodeFlags = 1 << 19
	NodeFlagsPossiblyContainsImportMeta    NodeFlags = 1 << 20

	NodeFlagsHasJSDoc                      NodeFlags = 1 << 21 // If node has preceding JSDoc comment(s)
	NodeFlagsJSDoc                         NodeFlags = 1 << 22 // If node was parsed inside jsdoc
	NodeFlagsAmbient                       NodeFlags = 1 << 23 // If node was inside an ambient context -- a declaration file, or inside something with the `declare` modifier.
	NodeFlagsJsonFile                      NodeFlags = 1 << 25 // If node was parsed in a Json
	NodeFlagsPossiblyContainsDeprecatedTag NodeFlags = 1 << 26 // Set during parse if comment text contains '@deprecated'; must confirm via JSDoc lookup
	NodeFlagsUnreachable                   NodeFlags = 1 << 27 // If node is unreachable according to the binder
	NodeFlagsReparserTransformedLiteral    NodeFlags = 1 << 28 // If node was transformed during parsing, making its' naive text source not match the AST
	NodeFlagsLuaLocal                      NodeFlags = 1 << 29 // Lua local variable declaration
	NodeFlagsLuaBlock                      NodeFlags = 1 << 30 // On a Block: a keyword-delimited Lua block (`then`/`do`/`else` ... `end`). Used by source formatting.
	NodeFlagsLuaTableField                 NodeFlags = 1 << 31 // On a PropertyAssignment: parsed with Lua `[k] = v` syntax rather than `[k]: v`. Set only by the parser; drives `=` printing. NOTE: bits 2 and 24 are free; after them, future Lua flags must repurpose per-kind bits (like NodeFlagsNestedNamespace below).

	NodeFlagsBlockScoped = NodeFlagsLet | NodeFlagsConst

	NodeFlagsReachabilityCheckFlags   = NodeFlagsHasImplicitReturn | NodeFlagsHasExplicitReturn
	NodeFlagsReachabilityAndEmitFlags = NodeFlagsReachabilityCheckFlags | NodeFlagsHasAsyncFunctions

	// Parsing context flags
	NodeFlagsContextFlags NodeFlags = NodeFlagsDisallowInContext | NodeFlagsDisallowConditionalTypesContext | NodeFlagsYieldContext | NodeFlagsAwaitContext | NodeFlagsJavaScriptFile | NodeFlagsAmbient

	// Exclude these flags when parsing a Type
	NodeFlagsTypeExcludesFlags NodeFlags = NodeFlagsYieldContext | NodeFlagsAwaitContext

	// Represents all flags that are potentially set once and
	// never cleared on SourceFiles which get re-used in between incremental parses.
	// See the comment above on `PossiblyContainsDynamicImport` and `PossiblyContainsImportMeta`.
	NodeFlagsPermanentlySetIncrementalFlags NodeFlags = NodeFlagsPossiblyContainsDynamicImport | NodeFlagsPossiblyContainsImportMeta

	// The following flags repurpose other NodeFlags as different meanings for Identifier nodes
	NodeFlagsIdentifierHasExtendedUnicodeEscape NodeFlags = NodeFlagsContainsSelf      // Indicates whether the identifier contains an extended unicode escape sequence
	NodeFlagsIdentifierIsInJSDocNamespace       NodeFlags = NodeFlagsHasAsyncFunctions // Indicates the identifier is the innermost name of a JSDoc namespace declaration

	// The following flag repurposes other NodeFlags for ModuleDeclaration nodes
	NodeFlagsNestedNamespace NodeFlags = NodeFlagsOptionalChain // If ModuleDeclaration is a nested namespace (e.g. inner part of A.B.C)
)
