//// [tests/cases/conformance/ported/typeParameterConstModifiers.tlua] ////

//// [typeParameterConstModifiers.tlua]
-- ported from tests/cases/conformance/types/typeParameters/typeParameterLists/typeParameterConstModifiers.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @strict: true directive (strict checking is enabled by default)
-- dropped: f6/fb/fc/fn/fn1/factory_55033 rest-parameter cases (rest-tuple
--   correlation is replaced by generic packs, and const packs are rejected)
-- dropped: spread arguments (operand spread no longer parses)
-- dropped: C1/C2 class cases (classes removed)
-- dropped: T4/T5 construct-signature types (`new` removed)
-- converted: arrow functions to function expressions; fx2's trailing-comma
--   variant collapses into fx1 (no arrow ambiguity to disambiguate)

declare f1: <const T>(x: T) => T;

local x11 = f1('a')
local x12 = f1({ 'a', { 'b', 'c' } })
local x13 = f1({ a = 1, b = "c", d = { "e", 2, true, { f = "g" } } })

declare f2: <const T, U>(x: T | nil) => T;

local x21 = f2('a')
local x22 = f2({ 'a', { 'b', 'c' } })
local x23 = f2({ a = 1, b = "c", d = { "e", 2, true, { f = "g" } } })

declare f3: <const T>(x: T) => T[];

local x31 = f3("hello")
local x32 = f3("hello")

declare f4: <const T>(obj: [T, T]) => T;

local x41 = f4({ { 1, 'x' }, { 2, 'y' } })
local x42 = f4({ { a = 1, b = 'x' }, { a = 2, b = 'y' } })

declare f5: <const T>(obj: { x: T, y: T }) => T;

local x51 = f5({ x = { 1, 'x' }, y = { 2, 'y' } })
local x52 = f5({ x = { a = 1, b = 'x' }, y = { a = 2, b = 'y' } })

local fx1 = function<const T>(x: T): T
    return x
end

interface I1<const T> { x: T }  -- Error

interface I2 {
    f<const T>(x: T): T;
}

type T1<const T> = T;  -- Error

type T2 = <const T>(x: T) => T;
type T3 = { <const T>(x: T): T };

-- Corrected repro from #51745

type Obj = { a: { b: { c: "123" } } };

type GetPath<T, P> =
    P extends readonly [] ? T :
    P extends readonly [infer A extends keyof T, ...infer Rest] ? GetPath<T[A], Rest> :
    never;

function set<T, const P extends readonly string[]>(obj: T, path: P, value: GetPath<T, P>)
end

declare obj: Obj;
declare value: "123";

set(obj, { 'a', 'b', 'c' }, value)

-- Repro from #52007

declare inners2: <const T extends readonly any[]>(args: readonly [unknown, ...T, unknown]) => T;

local test2 = inners2({ 1, 2, 3, 4, 5 })

-- Repro from #53307

type NotEmpty<T extends Table<string, any>> = keyof T extends never ? never : T;

local thing = function<const O extends Table<string, any>>(o: NotEmpty<O>)
    return o
end

local t = thing({ foo = '' })  -- readonly { foo: "" }

type NotEmptyMapped<T extends Table<string, any>> = keyof T extends never ? never : { [K in keyof T]: T[K] };

local thingMapped = function<const O extends Table<string, any>>(o: NotEmptyMapped<O>)
    return o
end

local tMapped = thingMapped({ foo = '' })  -- { foo: "" }

-- More examples of non-readonly constraints

declare fa1: <const T extends unknown[]>(args: T) => T;
declare fa2: <const T extends readonly unknown[]>(args: T) => T;

fa1({ "hello", 42 })
fa2({ "hello", 42 })

declare fd1: <const T extends string[] | number[]>(args: T) => T;
declare fd2: <const T extends string[] | readonly number[]>(args: T) => T;
declare fd3: <const T extends readonly string[] | readonly number[]>(args: T) => T;

fd1({ "hello", "world" })
fd1({ 1, 2, 3 })
fd2({ "hello", "world" })
fd2({ 1, 2, 3 })
fd3({ "hello", "world" })
fd3({ 1, 2, 3 })


//// [typeParameterConstModifiers.lua]
-- ported from tests/cases/conformance/types/typeParameters/typeParameterLists/typeParameterConstModifiers.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @strict: true directive (strict checking is enabled by default)
-- dropped: f6/fb/fc/fn/fn1/factory_55033 rest-parameter cases (rest-tuple
--   correlation is replaced by generic packs, and const packs are rejected)
-- dropped: spread arguments (operand spread no longer parses)
-- dropped: C1/C2 class cases (classes removed)
-- dropped: T4/T5 construct-signature types (`new` removed)
-- converted: arrow functions to function expressions; fx2's trailing-comma
--   variant collapses into fx1 (no arrow ambiguity to disambiguate)
local x11 = f1('a');
local x12 = f1({ 'a', { 'b', 'c' } });
local x13 = f1({ a = 1, b = "c", d = { "e", 2, true, { f = "g" } } });
local x21 = f2('a');
local x22 = f2({ 'a', { 'b', 'c' } });
local x23 = f2({ a = 1, b = "c", d = { "e", 2, true, { f = "g" } } });
local x31 = f3("hello");
local x32 = f3("hello");
local x41 = f4({ { 1, 'x' }, { 2, 'y' } });
local x42 = f4({ { a = 1, b = 'x' }, { a = 2, b = 'y' } });
local x51 = f5({ x = { 1, 'x' }, y = { 2, 'y' } });
local x52 = f5({ x = { a = 1, b = 'x' }, y = { a = 2, b = 'y' } });
local fx1 = function(x)
  return x;
end;
function set(obj, path, value)
end
set(obj, { 'a', 'b', 'c' }, value);
local test2 = inners2({ 1, 2, 3, 4, 5 });
local thing = function(o)
  return o;
end;
local t = thing({ foo = '' }); -- readonly { foo: "" }
local thingMapped = function(o)
  return o;
end;
local tMapped = thingMapped({ foo = '' }); -- { foo: "" }
fa1({ "hello", 42 });
fa2({ "hello", 42 });
fd1({ "hello", "world" });
fd1({ 1, 2, 3 });
fd2({ "hello", "world" });
fd2({ 1, 2, 3 });
fd3({ "hello", "world" });
fd3({ 1, 2, 3 });
