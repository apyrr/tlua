//// [tests/cases/conformance/ported/functionCall17.tlua] ////

//// [functionCall17.tlua]
-- ported from tests/cases/compiler/functionCall17.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: named rest parameter `...d` and its array annotation; rewritten as tlua's unnamed scalar vararg `...: number`

function foo(a: string, b?: string, c?: number, ...: number)
end

foo("foo", 1)
foo("foo")
foo()
foo(1, "bar")
foo("foo", 1, 3)
foo("foo", "bar", 3, 4)


//// [functionCall17.lua]
-- ported from tests/cases/compiler/functionCall17.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: named rest parameter `...d` and its array annotation; rewritten as tlua's unnamed scalar vararg `...: number`
function foo(a, b, c, ...)
end
foo("foo", 1);
foo("foo");
foo();
foo(1, "bar");
foo("foo", 1, 3);
foo("foo", "bar", 3, 4);
