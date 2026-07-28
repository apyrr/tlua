//// [tests/cases/conformance/ported/mappedTypeAndIndexSignatureRelation.tlua] ////

//// [mappedTypeAndIndexSignatureRelation.tlua]
-- ported from tests/cases/compiler/mappedTypeAndIndexSignatureRelation.ts
-- dropped: @target: es2015 and @strict directives (tlua uses the latest target and strict checking by default)
-- adapted: PropertyKey is unavailable in tlua; string | number is the supported table-key constraint

type Same<T> = { [P in keyof T]: T[P] }

type T1<T extends Record<string | number, number>> = T
type T2<U extends Record<string | number, number>> = T1<Same<U>>

-- Repro from #38235

type Foo<IdentifierT extends Record<string | number, string | number>> =
    IdentifierT

type Bar<IdentifierT extends Record<string | number, string | number>, T> =
    {
        [k in keyof T]: Foo<IdentifierT & { k: k }>
    }

type Merge2<T> = { [k in keyof T]: T[k] }
type Bar2<IdentifierT extends Record<string | number, string | number>, T> =
    {
        [k in keyof T]: Foo<Merge2<IdentifierT & { k: k }>>
    }

type Identity<T> = T
type Merge3<T> = Identity<{ [k in keyof T]: T[k] }>
type Bar3<IdentifierT extends Record<string | number, string | number>, T> =
    {
        [k in keyof T]: Foo<Merge3<IdentifierT & { k: k }>>
    }


//// [mappedTypeAndIndexSignatureRelation.lua]
-- ported from tests/cases/compiler/mappedTypeAndIndexSignatureRelation.ts
-- dropped: @target: es2015 and @strict directives (tlua uses the latest target and strict checking by default)
-- adapted: PropertyKey is unavailable in tlua; string | number is the supported table-key constraint
