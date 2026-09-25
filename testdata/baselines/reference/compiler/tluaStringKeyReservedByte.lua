//// [tests/cases/compiler/tluaStringKeyReservedByte.tlua] ////

//// [tluaStringKeyReservedByte.tlua]
// A Lua string key may begin with any byte, including 0xFE, the byte that begins
// the checker's internal names ("\xFEcall", "\xFEindex", "\xFEn:1", ...). Such a
// key is an ordinary string key: it names neither a number key nor a call or
// index signature.
local t = { ["\xfen:1"] = "s" }
local s: string = t["\xfen:1"]
local n: number = t[1]

interface Keys { "\xfecall": number; "\xfeindex": string; "\xfefoo": boolean }
local missing: Keys = { ["\xfecall"] = 1, ["\xfeindex"] = "x" }
local k: Keys = { ["\xfecall"] = 1, ["\xfeindex"] = "x", ["\xfefoo"] = true }
local c: number = k["\xfecall"]
local notCallable = k()

type KeyNames = keyof Keys
local key: KeyNames = "\xfefoo"
local notAKey: KeyNames = "foo"

// Diagnostics spell a byte that is not valid UTF-8 as `\xHH`, so two different
// bytes never look alike.
local ab = { a = 1 }
local ff = ab["\xff"]
local fe = ab["\xfe"]

// Inferring an object type from a key keeps the key a string key.
declare function fromKey<T>(k: keyof T): T
local numberLike = fromKey("\xfen:1")
local callLike = fromKey("\xfecall")


//// [tluaStringKeyReservedByte.lua]
-- A Lua string key may begin with any byte, including 0xFE, the byte that begins
-- the checker's internal names ("\xFEcall", "\xFEindex", "\xFEn:1", ...). Such a
-- key is an ordinary string key: it names neither a number key nor a call or
-- index signature.
local t = { ["\xfen:1"] = "s" };
local s = t["\xfen:1"];
local n = t[1];
local missing = { ["\xfecall"] = 1, ["\xfeindex"] = "x" };
local k = { ["\xfecall"] = 1, ["\xfeindex"] = "x", ["\xfefoo"] = true };
local c = k["\xfecall"];
local notCallable = k();
local key = "\xfefoo";
local notAKey = "foo";
-- Diagnostics spell a byte that is not valid UTF-8 as `\xHH`, so two different
-- bytes never look alike.
local ab = { a = 1 };
local ff = ab["\xff"];
local fe = ab["\xfe"];
local numberLike = fromKey("\xfen:1");
local callLike = fromKey("\xfecall");
