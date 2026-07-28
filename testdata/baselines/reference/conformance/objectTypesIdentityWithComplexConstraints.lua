//// [tests/cases/conformance/ported/objectTypesIdentityWithComplexConstraints.tlua] ////

//// [objectTypesIdentityWithComplexConstraints.tlua]
-- ported from tests/cases/conformance/types/typeRelationships/typeAndMemberIdentity/objectTypesIdentityWithComplexConstraints.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: C-style function body braces, rewritten to Lua `end`

interface A {
    <T extends {
        <S extends A>(x: T, y: S): void
    }>(x: T, y: T): void
}

interface B {
    <U extends B>(x: U, y: U): void
}

-- ok, not considered identical because the steps of contextual signature instantiation create fresh type parameters
function foo(x: A);
function foo(x: B); -- error after constraints above made illegal
function foo(x: any) end


//// [objectTypesIdentityWithComplexConstraints.lua]
-- ported from tests/cases/conformance/types/typeRelationships/typeAndMemberIdentity/objectTypesIdentityWithComplexConstraints.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: C-style function body braces, rewritten to Lua `end`
function foo(x)
end
