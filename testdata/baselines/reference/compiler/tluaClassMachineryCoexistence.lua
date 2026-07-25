//// [tests/cases/compiler/tluaClassMachineryCoexistence.tlua] ////

//// [tluaClassMachineryCoexistence.tlua]
// Class *declarations* are gone, and so is `new` — the expression, `new.target`,
// construct signatures and constructor types all went with it. What classes
// relied on that MUST keep working is the shared type machinery interfaces, the
// DOM libs and Lua tables still use: `extends` heritage, `instanceof`, and
// `typeof`. (Object-literal `super` died with table-literal methods in the table
// slice. Polymorphic `this` died with the `this` slice — see tluaNoThis.tlua.)

interface Base {
  b: number;
}

// interface `extends` heritage
interface Derived extends Base {
  d: number;
}

// With classes and construct signatures gone, the only right-hand side
// `instanceof` still accepts is a value of `function` type.
declare K: function;
declare inst: Derived;

local isK = inst instanceof K; // `instanceof`
type TK = typeof K; // `typeof`

// A non-callable right-hand side is still an error: the check runs through
// isTypeDerivedFrom, not a structural subtype test the empty Function sentinel
// would make vacuous.
declare notCallable: { area: number };
local bad1 = inst instanceof notCallable;
local bad2 = inst instanceof 3;

// `InstanceType` and `ConstructorParameters` are both gone: the first was
// defined over construct signatures, which no longer exist, and the second was
// built on `...args: infer P`, which needs a rest parameter whose type IS the
// parameter tuple. See the vararg slice.


//// [tluaClassMachineryCoexistence.lua]
-- Class *declarations* are gone, and so is `new` — the expression, `new.target`,
-- construct signatures and constructor types all went with it. What classes
-- relied on that MUST keep working is the shared type machinery interfaces, the
-- DOM libs and Lua tables still use: `extends` heritage, `instanceof`, and
-- `typeof`. (Object-literal `super` died with table-literal methods in the table
-- slice. Polymorphic `this` died with the `this` slice — see tluaNoThis.tlua.)
local isK = inst instanceof K; -- `instanceof`
local bad1 = inst instanceof notCallable;
local bad2 = inst instanceof 3;
-- `InstanceType` and `ConstructorParameters` are both gone: the first was
-- defined over construct signatures, which no longer exist, and the second was
-- built on `...args: infer P`, which needs a rest parameter whose type IS the
-- parameter tuple. See the vararg slice.
