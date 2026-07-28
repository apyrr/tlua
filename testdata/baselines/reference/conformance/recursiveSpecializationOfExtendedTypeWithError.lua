//// [tests/cases/conformance/ported/recursiveSpecializationOfExtendedTypeWithError.tlua] ////

//// [recursiveSpecializationOfExtendedTypeWithError.tlua]
-- ported from tests/cases/compiler/recursiveSpecializationOfExtendedTypeWithError.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @lib: es5 directive (tlua uses its built-in Lua library)

interface HTMLSelectElement {
    options: HTMLSelectElement,
    <A>(name: A): any,
}


//// [recursiveSpecializationOfExtendedTypeWithError.lua]
-- ported from tests/cases/compiler/recursiveSpecializationOfExtendedTypeWithError.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @lib: es5 directive (tlua uses its built-in Lua library)
