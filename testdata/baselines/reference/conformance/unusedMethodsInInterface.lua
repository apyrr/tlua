//// [tests/cases/conformance/ported/unusedMethodsInInterface.tlua] ////

//// [unusedMethodsInInterface.tlua]
-- ported from tests/cases/compiler/unusedMethodsInInterface.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- compiler gap: tlua reports TLUA6196 on the top-level interface name `I1`; upstream reports
--   nothing, because a top-level interface is a global rather than a local. tlua resolves such
--   declarations globally but still unused-checks them. Same gap as the unusedTypeParameter* ports.


interface I1 {
    f1()
    f2(x: number, y: string)
}


//// [unusedMethodsInInterface.lua]
-- ported from tests/cases/compiler/unusedMethodsInInterface.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- compiler gap: tlua reports TLUA6196 on the top-level interface name `I1`; upstream reports
--   nothing, because a top-level interface is a global rather than a local. tlua resolves such
--   declarations globally but still unused-checks them. Same gap as the unusedTypeParameter* ports.
