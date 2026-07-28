//// [tests/cases/conformance/ported/declarationEmitTypeAliasWithTypeParameters4.tlua] ////

//// [declarationEmitTypeAliasWithTypeParameters4.tlua]
-- ported from tests/cases/compiler/declarationEmitTypeAliasWithTypeParameters4.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: declaration emit is unsupported for Lua modules (TLUA100054); kept generic type-alias, method-signature, alias-instantiation, and assertion checker coverage.

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
-- dropped: declaration emit is unsupported for Lua modules (TLUA100054); kept generic type-alias, method-signature, alias-instantiation, and assertion checker coverage.
function foo()
  return {};
end
