//// [tests/cases/conformance/ported/functionCall14.tlua] ////

//// [functionCall14.tlua]
-- ported from tests/cases/compiler/functionCall14.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: named rest parameter `...b` and its array annotation; rewritten as tlua's unnamed scalar vararg `...: number`

function foo(a?: string, ...: number)
end

foo("foo", 1)
foo("foo")
foo()
foo(1, "bar")
foo("foo", 1, 3)


//// [functionCall14.lua]
-- ported from tests/cases/compiler/functionCall14.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: named rest parameter `...b` and its array annotation; rewritten as tlua's unnamed scalar vararg `...: number`
function foo(a, ...)
end
foo("foo", 1);
foo("foo");
foo();
foo(1, "bar");
foo("foo", 1, 3);
