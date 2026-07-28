//// [tests/cases/conformance/ported/functionCall9.tlua] ////

//// [functionCall9.tlua]
-- ported from tests/cases/compiler/functionCall9.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(a?: string, b?: number)
end

foo('foo', 1)
foo('foo')
foo('foo', 'bar')
foo('foo', 1, 'bar')
foo()


//// [functionCall9.lua]
-- ported from tests/cases/compiler/functionCall9.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(a, b)
end
foo('foo', 1);
foo('foo');
foo('foo', 'bar');
foo('foo', 1, 'bar');
foo();
