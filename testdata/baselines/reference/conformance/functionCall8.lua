//// [tests/cases/conformance/ported/functionCall8.tlua] ////

//// [functionCall8.tlua]
-- ported from tests/cases/compiler/functionCall8.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(a?: string)
end

foo('foo')
foo('foo', 'bar')
foo(4)
foo()


//// [functionCall8.lua]
-- ported from tests/cases/compiler/functionCall8.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(a)
end
foo('foo');
foo('foo', 'bar');
foo(4);
foo();
