//// [tests/cases/conformance/ported/functionCall13.tlua] ////

//// [functionCall13.tlua]
-- ported from tests/cases/compiler/functionCall13.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(a: string, ...: number)
end

foo('foo', 1)
foo('foo')
foo()
foo(1, 'bar')
foo('foo', 1, 3)


//// [functionCall13.lua]
-- ported from tests/cases/compiler/functionCall13.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(a, ...)
end
foo('foo', 1);
foo('foo');
foo();
foo(1, 'bar');
foo('foo', 1, 3);
