//// [tests/cases/conformance/ported/declarationEmitTypeAliasWithTypeParameters3.tlua] ////

//// [declarationEmitTypeAliasWithTypeParameters3.tlua]
-- ported from tests/cases/compiler/declarationEmitTypeAliasWithTypeParameters3.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @declaration: true -- tlua has declaration-emit machinery, but it reports
--   TLUA100054 for every .tlua source, so restoring the directive would replace this test's
--   subject with that one diagnostic. tluaModuleDeclarationEmitUnsupported.tlua covers it.


type Foo<T> = {
    foo<U>(): Foo<U>
}

function bar()
    return {} as Foo<number>
end


//// [declarationEmitTypeAliasWithTypeParameters3.lua]
-- ported from tests/cases/compiler/declarationEmitTypeAliasWithTypeParameters3.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @declaration: true -- tlua has declaration-emit machinery, but it reports
--   TLUA100054 for every .tlua source, so restoring the directive would replace this test's
--   subject with that one diagnostic. tluaModuleDeclarationEmitUnsupported.tlua covers it.
function bar()
  return {};
end
