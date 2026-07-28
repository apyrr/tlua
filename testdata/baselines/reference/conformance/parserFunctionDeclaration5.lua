//// [tests/cases/conformance/ported/parserFunctionDeclaration5.tlua] ////

//// [parserFunctionDeclaration5.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/FunctionDeclarations/parserFunctionDeclaration5.ts
-- dropped: @target: es2015 (tlua targets latest; the ES target is not a tlua concept)
-- dropped: C-style function body, rewritten with Lua's `end` terminator

function foo();
function foo()
end


//// [parserFunctionDeclaration5.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/FunctionDeclarations/parserFunctionDeclaration5.ts
-- dropped: @target: es2015 (tlua targets latest; the ES target is not a tlua concept)
-- dropped: C-style function body, rewritten with Lua's `end` terminator
function foo()
end
