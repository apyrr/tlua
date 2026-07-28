//// [tests/cases/conformance/ported/declarationEmitTypeAliasWithTypeParameters4.tlua] ////

//// [declarationEmitTypeAliasWithTypeParameters4.tlua]
-- ported from tests/cases/compiler/declarationEmitTypeAliasWithTypeParameters4.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @declaration: true -- tlua has declaration-emit machinery, but it reports
--   TLUA100054 for every .tlua source, so restoring the directive would replace this test's
--   subject with that one diagnostic. tluaModuleDeclarationEmitUnsupported.tlua covers it.


type Foo<T, Y> = {
    foo<U, J>(): Foo<U, J>
}
type SubFoo<R> = Foo<string, R>

function foo()
    return {} as SubFoo<number>
end


//// [declarationEmitTypeAliasWithTypeParameters4.lua]
-- ported from tests/cases/compiler/declarationEmitTypeAliasWithTypeParameters4.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @declaration: true -- tlua has declaration-emit machinery, but it reports
--   TLUA100054 for every .tlua source, so restoring the directive would replace this test's
--   subject with that one diagnostic. tluaModuleDeclarationEmitUnsupported.tlua covers it.
function foo()
  return {};
end
