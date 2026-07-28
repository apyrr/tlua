//// [tests/cases/conformance/ported/functionOverloads10.tlua] ////

//// [functionOverloads10.tlua]
-- ported from tests/cases/compiler/functionOverloads10.ts
-- dropped: @target: es2015 and @strict: false directives (tlua uses latest target and strict checking)

function foo(foo: string, bar: number): any;
function foo(foo: string): any;
function foo(foo: any): any
end


//// [functionOverloads10.lua]
-- ported from tests/cases/compiler/functionOverloads10.ts
-- dropped: @target: es2015 and @strict: false directives (tlua uses latest target and strict checking)
function foo(foo)
end
