//// [tests/cases/conformance/ported/parserAmbiguity1.tlua] ////

//// [parserAmbiguity1.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/Generics/parserAmbiguity1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

f(g<A, B>(7))


//// [parserAmbiguity1.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/Generics/parserAmbiguity1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
f(g(7));
