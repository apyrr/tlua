//// [tests/cases/conformance/ported/propertiesAndIndexers2.tlua] ////

//// [propertiesAndIndexers2.tlua]
-- ported from tests/cases/compiler/propertiesAndIndexers2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface A {
    [n: number]: string
    [s: string]: number
}

-- All of these should fail.
interface B extends A {
    c: string
    3: string
    Infinity: string
    "-Infinity": string
    NaN: string
    "-NaN": string
    6(): string
}


//// [propertiesAndIndexers2.lua]
-- ported from tests/cases/compiler/propertiesAndIndexers2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
