//// [tests/cases/conformance/ported/interfacePropertiesWithSameName1.tlua] ////

//// [interfacePropertiesWithSameName1.tlua]
-- ported from tests/cases/compiler/interfacePropertiesWithSameName1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Mover {
    move(): void
    getStatus(): { speed: number }
}
interface Shaker {
    shake(): void
    getStatus(): { frequency: number }
}

interface MoverShaker extends Mover, Shaker {
    getStatus(): { speed: number, frequency: number }
}


//// [interfacePropertiesWithSameName1.lua]
-- ported from tests/cases/compiler/interfacePropertiesWithSameName1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
