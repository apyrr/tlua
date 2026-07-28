//// [tests/cases/conformance/ported/unionTypeWithIndexAndMethodSignature.tlua] ////

//// [unionTypeWithIndexAndMethodSignature.tlua]
-- ported from tests/cases/compiler/unionTypeWithIndexAndMethodSignature.ts
-- dropped: @target: es2015 and @strict: true directives (tlua defaults to esnext and strict checking)

interface Options {
    m(x: number): void
    [key: string]: unknown
}
declare function f(options: number | Options): void
f({
    m = function(x) end,
})


//// [unionTypeWithIndexAndMethodSignature.lua]
-- ported from tests/cases/compiler/unionTypeWithIndexAndMethodSignature.ts
-- dropped: @target: es2015 and @strict: true directives (tlua defaults to esnext and strict checking)
f({
  m = function(x)
  end,
});
