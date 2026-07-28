//// [tests/cases/conformance/ported/inheritedMembersAndIndexSignaturesFromDifferentBases2.tlua] ////

//// [inheritedMembersAndIndexSignaturesFromDifferentBases2.tlua]
-- ported from tests/cases/compiler/inheritedMembersAndIndexSignaturesFromDifferentBases2.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to latest target and strict checking)

interface A<T> {
    [n: number]: T;
}

interface B {
    foo: number;
}

interface C extends B, A<string> { } -- Should succeed


//// [inheritedMembersAndIndexSignaturesFromDifferentBases2.lua]
-- ported from tests/cases/compiler/inheritedMembersAndIndexSignaturesFromDifferentBases2.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to latest target and strict checking)
