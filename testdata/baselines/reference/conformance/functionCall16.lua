//// [tests/cases/conformance/ported/functionCall16.tlua] ////

//// [functionCall16.tlua]
-- ported from tests/cases/compiler/functionCall16.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: named rest parameter `...c`; tlua uses the unnamed vararg `...` form.

function foo(a: string, b?: string, ...: number)
end

foo('foo', 1)
foo('foo')
foo('foo', 'bar')
foo()
foo(1, 'bar')
foo('foo', 'bar', 3)


//// [functionCall16.lua]
-- ported from tests/cases/compiler/functionCall16.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: named rest parameter `...c`; tlua uses the unnamed vararg `...` form.
function foo(a, b, ...)
end
foo('foo', 1);
foo('foo');
foo('foo', 'bar');
foo();
foo(1, 'bar');
foo('foo', 'bar', 3);
