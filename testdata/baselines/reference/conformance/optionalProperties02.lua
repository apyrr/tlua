//// [tests/cases/conformance/ported/optionalProperties02.tlua] ////

//// [optionalProperties02.tlua]
-- ported from tests/cases/conformance/types/typeRelationships/comparable/optionalProperties02.ts
-- dropped: @target: es2015, @strictNullChecks: true, and @declaration directives (tlua defaults to the corresponding strict settings and this test does not need declaration output)

interface Foo {
    a?: string
    b: string
}

local _ = ({ a = nil } as Foo)


//// [optionalProperties02.lua]
-- ported from tests/cases/conformance/types/typeRelationships/comparable/optionalProperties02.ts
-- dropped: @target: es2015, @strictNullChecks: true, and @declaration directives (tlua defaults to the corresponding strict settings and this test does not need declaration output)
local _ = { a = nil };
