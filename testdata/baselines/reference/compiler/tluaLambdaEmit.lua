//// [tests/cases/compiler/tluaLambdaEmit.tlua] ////

//// [tluaLambdaEmit.tlua]
// A lambda compiles to the Lua function constructor. An expression body
// becomes a single-line `function(x) return x + 1 end`; a braced body inlines
// its statements before `end`.
local function two(): (number, string)
  return 1, "s"
end
local add = (x: number, y: number) => x + y
local block = (x: number) => { return x * 2; }
local nested = (x: number) => (y: number) => x + y
local noParams = () => 42

// A concise body is typed as one value, but a direct call in Lua tail
// position would propagate its whole pack -- the emit parenthesizes it so the
// runtime truncates to the value the type names.
local truncated = () => two()
local one: number = truncated()

// A parenthesized table-literal body keeps its parens; `return { ... } end`
// would also be valid Lua, but the source spelling round-trips.
local table = () => ({ k = 1 })
print(add(1, 2), block(2), nested(1)(2), noParams(), table().k)


//// [tluaLambdaEmit.lua]
-- A lambda compiles to the Lua function constructor. An expression body
-- becomes a single-line `function(x) return x + 1 end`; a braced body inlines
-- its statements before `end`.
local function two()
    return 1, "s";
end
local add = function(x, y) return x + y end;
local block = function(x)
    return x * 2;
end;
local nested = function(x) return function(y) return x + y end end;
local noParams = function() return 42 end;
-- A concise body is typed as one value, but a direct call in Lua tail
-- position would propagate its whole pack -- the emit parenthesizes it so the
-- runtime truncates to the value the type names.
local truncated = function() return (two()) end;
local one = truncated();
-- A parenthesized table-literal body keeps its parens; `return { ... } end`
-- would also be valid Lua, but the source spelling round-trips.
local table = function() return ({ k = 1 }) end;
print(add(1, 2), block(2), nested(1)(2), noParams(), table().k);
