//// [tests/cases/conformance/ported/infinitelyGenerativeInheritance1.tlua] ////

//// [infinitelyGenerativeInheritance1.tlua]
-- ported from tests/cases/compiler/infinitelyGenerativeInheritance1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Stack<T> {
    pop(): T
    zip<S>(a: Stack<S>): Stack<{ x: T; y: S }>
}

interface MyStack<T> extends Stack<T> {
    zip<S>(a: Stack<S>): Stack<{ x: T; y: S }>
}


//// [infinitelyGenerativeInheritance1.lua]
-- ported from tests/cases/compiler/infinitelyGenerativeInheritance1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
