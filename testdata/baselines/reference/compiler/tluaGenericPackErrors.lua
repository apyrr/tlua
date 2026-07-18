//// [tests/cases/compiler/tluaGenericPackErrors.tlua] ////

//// [tluaGenericPackErrors.tlua]
// The grammar and usage rules that keep a generic pack parameter in a pack
// position.

// A pack parameter must trail the plain ones: a plain parameter cannot follow it.
declare function ordering<...A, T>(...: A): T;

// A pack parameter names several values, so it cannot stand in a single-value
// position: as a plain annotation, a tuple/return element, or a value type.
declare function asValue<...A>(x: A): void;
declare function asReturn<...A>(): A;
declare function inList<...A>(): (number, A);

// A pack parameter is allowed on an alias, but not an interface or class.
type Alias<...A> = number;
interface Iface<...A> {}

// Any function-shaped signature carries one, though: an interface METHOD or call
// signature is a function, so neither of these errors.
interface WithMethods {
    m<...A>(...: A): (...A);
    <...R>(...: R): (...R);
}

// A written constraint or default on a pack parameter is rejected: the implicit
// constraint is the open `(...any)` pack, and silently ignoring the syntax would
// enforce nothing.
declare function constrained<...A extends string>(...: A): void;
declare function defaulted<...A = string>(...: A): void;

// `...T` where T is an ordinary (non-pack) type parameter stays a homogeneous
// pack of element T -- this is the Stage 1 form and is NOT an error.
declare function homogeneous<T>(...: T): (...T);
local okA, okB = homogeneous(1, 2);

// A well-formed pack function, so the errors above are isolated.
declare function fine<...A, ...R>(f: (...: A) => (...R), ...: A): (...R);

// Explicit scalar arguments are lifted even when a failed call uses its
// candidate signature for error recovery.
declare function id<...A>(...: A): (...A);
id<number>("s");


//// [tluaGenericPackErrors.lua]
-- The grammar and usage rules that keep a generic pack parameter in a pack
-- position.
local okA, okB = homogeneous(1, 2);
id("s");
