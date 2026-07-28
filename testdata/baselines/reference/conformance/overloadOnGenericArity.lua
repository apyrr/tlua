//// [tests/cases/conformance/ported/overloadOnGenericArity.tlua] ////

//// [overloadOnGenericArity.tlua]
-- ported from tests/cases/compiler/overloadOnGenericArity.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target); upstream has no other directives
-- rewritten: the second overload's `Date` return type became `number`, since tlua has no JS `Date`; any type distinct from `string` preserves the subject, which is a generic and a non-generic overload sharing a value-parameter list but differing in return type
-- note: upstream carries a stale inline "Error: Overloads cannot differ only by return type" comment but has no errors baseline; differing only by return type is accepted here too, so no diagnostic is expected.

interface Test {
    then<U>(p: string): string
    then(p: string): number
}


//// [overloadOnGenericArity.lua]
-- ported from tests/cases/compiler/overloadOnGenericArity.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target); upstream has no other directives
-- rewritten: the second overload's `Date` return type became `number`, since tlua has no JS `Date`; any type distinct from `string` preserves the subject, which is a generic and a non-generic overload sharing a value-parameter list but differing in return type
-- note: upstream carries a stale inline "Error: Overloads cannot differ only by return type" comment but has no errors baseline; differing only by return type is accepted here too, so no diagnostic is expected.
