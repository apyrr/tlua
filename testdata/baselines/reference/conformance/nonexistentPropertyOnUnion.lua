//// [tests/cases/conformance/ported/nonexistentPropertyOnUnion.tlua] ////

//// [nonexistentPropertyOnUnion.tlua]
-- ported from tests/cases/compiler/nonexistentPropertyOnUnion.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: Promise union constituent (Promise is deleted; replaced with a structurally distinct interface)

interface Other {
    value: string
}

function f(x: string | Other)
    x:lower()
end


//// [nonexistentPropertyOnUnion.lua]
-- ported from tests/cases/compiler/nonexistentPropertyOnUnion.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: Promise union constituent (Promise is deleted; replaced with a structurally distinct interface)
function f(x)
  x:lower();
end
