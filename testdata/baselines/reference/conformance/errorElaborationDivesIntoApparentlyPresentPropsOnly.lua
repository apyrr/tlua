//// [tests/cases/conformance/ported/errorElaborationDivesIntoApparentlyPresentPropsOnly.tlua] ////

//// [errorElaborationDivesIntoApparentlyPresentPropsOnly.tlua]
-- ported from tests/cases/compiler/errorElaborationDivesIntoApparentlyPresentPropsOnly.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @strict: false directive (tlua strictness is fixed by the harness)

function foo<T extends { a: string }>(x: T)
    x = { a = "abc", b = 20, c = 30 }
end

function bar<T extends { a: string }>(x: T)
    x = { a = 20 }
end

function baz<T extends { a: string }>(x: T)
    x = { a = "not ok" }
end


//// [errorElaborationDivesIntoApparentlyPresentPropsOnly.lua]
-- ported from tests/cases/compiler/errorElaborationDivesIntoApparentlyPresentPropsOnly.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @strict: false directive (tlua strictness is fixed by the harness)
function foo(x)
  x = { a = "abc", b = 20, c = 30 };
end
function bar(x)
  x = { a = 20 };
end
function baz(x)
  x = { a = "not ok" };
end
