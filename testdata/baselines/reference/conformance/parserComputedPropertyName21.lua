//// [tests/cases/conformance/ported/parserComputedPropertyName21.tlua] ////

//// [parserComputedPropertyName21.tlua]
-- ported from tests/cases/conformance/parser/ecmascript6/ComputedPropertyNames/parserComputedPropertyName21.ts
-- dropped: @target: ES6 directive (tlua defaults to latest target)

interface I {
    [e]: number
}


//// [parserComputedPropertyName21.lua]
-- ported from tests/cases/conformance/parser/ecmascript6/ComputedPropertyNames/parserComputedPropertyName21.ts
-- dropped: @target: ES6 directive (tlua defaults to latest target)
