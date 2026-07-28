//// [tests/cases/conformance/ported/functionCall6.tlua] ////

//// [functionCall6.tlua]
-- ported from tests/cases/compiler/functionCall6.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(a: string)
end
foo('bar')
foo(2)
foo('foo', 'bar')
foo()


//// [functionCall6.lua]
-- ported from tests/cases/compiler/functionCall6.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(a)
end
foo('bar');
foo(2);
foo('foo', 'bar');
foo();
