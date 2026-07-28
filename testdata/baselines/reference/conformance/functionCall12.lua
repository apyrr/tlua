//// [tests/cases/conformance/ported/functionCall12.tlua] ////

//// [functionCall12.tlua]
-- ported from tests/cases/compiler/functionCall12.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(a: string, b?: number, c?: string)
end

foo('foo', 1)
foo('foo')
foo()
foo(1, 'bar')
foo('foo', 1, 'bar')
foo('foo', 1, 3)


//// [functionCall12.lua]
-- ported from tests/cases/compiler/functionCall12.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(a, b, c)
end
foo('foo', 1);
foo('foo');
foo();
foo(1, 'bar');
foo('foo', 1, 'bar');
foo('foo', 1, 3);
