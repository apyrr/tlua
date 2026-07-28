//// [tests/cases/conformance/ported/functionCall10.tlua] ////

//// [functionCall10.tlua]
-- ported from tests/cases/compiler/functionCall10.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(...: number)
end
foo(0, 1)
foo('foo')
foo()
foo(1, 'bar')


//// [functionCall10.lua]
-- ported from tests/cases/compiler/functionCall10.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(...)
end
foo(0, 1);
foo('foo');
foo();
foo(1, 'bar');
