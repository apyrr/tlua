//// [tests/cases/conformance/ported/recursiveTypes1.tlua] ////

//// [recursiveTypes1.tlua]
-- ported from tests/cases/compiler/recursiveTypes1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Entity<T extends Entity<T>> {
    X: T
    Y: T
}

interface Person<U extends Person<U>> extends Entity<U> {
    n: number
}

interface Customer extends Person<Customer> {
    s: string
}


//// [recursiveTypes1.lua]
-- ported from tests/cases/compiler/recursiveTypes1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
