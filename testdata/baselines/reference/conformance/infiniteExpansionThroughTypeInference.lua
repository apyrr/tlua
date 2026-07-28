//// [tests/cases/conformance/ported/infiniteExpansionThroughTypeInference.tlua] ////

//// [infiniteExpansionThroughTypeInference.tlua]
-- ported from tests/cases/conformance/types/typeRelationships/recursiveTypes/infiniteExpansionThroughTypeInference.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface G<T> {
    x: G<G<T>>
    y: T
}

function ff<T>(g: G<T>): void
    ff(g)
end


//// [infiniteExpansionThroughTypeInference.lua]
-- ported from tests/cases/conformance/types/typeRelationships/recursiveTypes/infiniteExpansionThroughTypeInference.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function ff(g)
  ff(g);
end
