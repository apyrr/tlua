//// [tests/cases/conformance/ported/parserFunctionDeclaration5.tlua] ////

//// [parserFunctionDeclaration5.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/FunctionDeclarations/parserFunctionDeclaration5.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict)
-- dropped: C-style function body, rewritten with Lua's `end` terminator

function foo();
function foo()
end


//// [parserFunctionDeclaration5.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/FunctionDeclarations/parserFunctionDeclaration5.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict)
-- dropped: C-style function body, rewritten with Lua's `end` terminator
function foo()
end
