//// [tests/cases/conformance/ported/parserComputedPropertyName20.tlua] ////

//// [parserComputedPropertyName20.tlua]
-- ported from tests/cases/conformance/parser/ecmascript6/ComputedPropertyNames/parserComputedPropertyName20.ts
-- dropped: @target: ES6 directive (tlua defaults to latest target)

interface I {
    [e](): number
}


//// [parserComputedPropertyName20.lua]
-- ported from tests/cases/conformance/parser/ecmascript6/ComputedPropertyNames/parserComputedPropertyName20.ts
-- dropped: @target: ES6 directive (tlua defaults to latest target)
