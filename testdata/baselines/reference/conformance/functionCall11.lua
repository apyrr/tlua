//// [tests/cases/conformance/ported/functionCall11.tlua] ////

//// [functionCall11.tlua]
-- ported from tests/cases/compiler/functionCall11.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(a: string, b?: number)
end

foo('foo', 1)
foo('foo')
foo()
foo(1, 'bar')
foo('foo', 1, 'bar')


//// [functionCall11.lua]
-- ported from tests/cases/compiler/functionCall11.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(a, b)
end
foo('foo', 1);
foo('foo');
foo();
foo(1, 'bar');
foo('foo', 1, 'bar');
