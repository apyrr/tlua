//// [tests/cases/conformance/ported/overloadOnConstDuplicateOverloads1.tlua] ////

//// [overloadOnConstDuplicateOverloads1.tlua]
-- ported from tests/cases/compiler/overloadOnConstDuplicateOverloads1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(a: 'hi', x: string);
function foo(a: 'hi', x: string);
function foo(a: any, x: any)
end

function foo2(a: 'hi', x: string);
function foo2(a: 'hi', x: string);
function foo2(a: string, x: string);
function foo2(a: any, x: any)
end


//// [overloadOnConstDuplicateOverloads1.lua]
-- ported from tests/cases/compiler/overloadOnConstDuplicateOverloads1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(a, x)
end
function foo2(a, x)
end
