//// [tests/cases/conformance/ported/functionOverloadsOnGenericArity2.tlua] ////

//// [functionOverloadsOnGenericArity2.tlua]
-- ported from tests/cases/compiler/functionOverloadsOnGenericArity2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: JavaScript-library Date return type; table preserves overload arity coverage.

interface I {
    then(p: string): string
    then<U>(p: string): string
    then<U, T>(p: string): table
}


//// [functionOverloadsOnGenericArity2.lua]
-- ported from tests/cases/compiler/functionOverloadsOnGenericArity2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: JavaScript-library Date return type; table preserves overload arity coverage.
