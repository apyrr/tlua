//// [tests/cases/compiler/tluaTableTagNarrowing.tlua] ////

//// [tluaTableTagNarrowing.tlua]
// `type(x) == "table"` says a value is a table and nothing about its keys, so a type
// that states nothing about the value (unknown, a bare type parameter) narrows to
// Table<{}, unknown>: any key reads unknown through its index signature, and
// pairs visits it. A declared `table` states no key and stays closed; code that reads
// arbitrary keys declares Table<K, V>.

local function isMarked(value: unknown): boolean
    return type(value) == "table" and value._marked == true
end

-- Flow facts narrow the read, in either spelling.
local function readCode(err: unknown): string | nil
    if type(err) == "table" and type(err.code) == "string" then
        return err.code
    end
    return nil
end
local function readSlot(err: unknown): string | nil
    if type(err) == "table" and type(err["name"]) == "string" then
        return err["name"]
    end
    return nil
end

-- Any key reads: strings, numbers, booleans, tables.
local function readKeys(value: unknown, key: string, other: table): (unknown, unknown, unknown, unknown)
    if type(value) ~= "table" then
        return nil, nil, nil, nil
    end
    return value[key], value[1], value[true], value[other]
end

-- The narrowed type is the declared Table<{}, unknown>, so it iterates, and it
-- is not a table of anything narrower.
local function keysOf(value: unknown): unknown[]
    local keys: unknown[] = {}
    if type(value) == "table" then
        for key, v in pairs(value) do
            keys[#keys + 1] = key
            local asString: string = key
        end
        local anyTable: Table<{}, unknown> = value
        local asTable: table = value
        local numbers: Table<string, number> = value
        local strings: Table<string, unknown> = value
        local cast = value as Table<string, unknown>
    end
    return keys
end

-- A generic narrows to an intersection with it. It reads any key, but a generic
-- is indexed for reading only, so writes need a stated type.
local function readGeneric<T>(value: T): unknown
    if type(value) == "table" then
        value.field = 1
        value["slot"] = 2
        return value.field
    end
    return nil
end

-- A stated type keeps its shape: the tag selects the table members of a union.
interface Named { name: string; }
local function readUnion(value: string | Named | nil): string | nil
    if type(value) == "table" then
        return value.name
    end
    return nil
end
local function readList(value: string | number[]): number | nil
    if type(value) == "table" then
        return value[1]
    end
    return nil
end
local function notTable(value: string | Named): string
    if type(value) ~= "table" then
        return value
    end
    return value.name
end

-- A declared `table` states no key: it stays `table` under the tag, a read is an
-- error, and it does not iterate. Declare the keys it holds instead.
local function readDeclared(t: table): unknown
    if type(t) == "table" then
        return t.name
    end
    return nil
end
-- pairs and next visit it with unknown keys and values, as Luau's do; ipairs and
-- the table library want a stated array.
local function iterateDeclared(t: table, list: Table<number, string>)
    for key, value in pairs(t) do
        local k: {} = key
        local n: number = key
        local v: unknown = value
        local s: string = value
    end
    local firstKey, firstValue = next(t)
    local k2: {} | nil = firstKey
    local v2: {} = firstValue
    for key, value in next, t do
        local v3: {} = value
    end
    for index, value in pairs(list) do
        local i: number = index
        local s: string = value
    end
    for index, value in ipairs(t) do end
    table.insert(t, 1)
end
local function readStated(t: Table<string, unknown>): unknown
    for key, value in pairs(t) do
        local asString: string = key
    end
    return t.name
end
local function readConstrained<T extends table>(t: T): unknown
    if type(t) == "table" then
        return t.field
    end
    return nil
end

-- Writes go through the index signature too.
local function mark(value: unknown)
    if type(value) == "table" then
        value._marked = true
        value[1] = "x"
    end
end

-- Nested reads go through unknown, which states nothing.
local function nested(value: unknown): unknown
    if type(value) == "table" then
        return value.inner.value
    end
    return nil
end

-- A key no table takes is an error on the narrowed table too.
local function readBadKeys(value: unknown, k: string | nil, u: unknown)
    if type(value) == "table" then
        local a = value[k]
        local b = value[u]
    end
end

-- The narrowed table is a table, stays one under a second check, and takes a
-- metatable.
local function useAsTable(value: unknown)
    if type(value) == "table" then
        if type(value) == "table" then
            local again: Table<{}, unknown> = value
        end
        setmetatable(value, nil)
        local n: number = #value
    end
end

-- A type guard to a stated shape still narrows the table to it.
local function isRecord(value: unknown): value is Table<string, unknown>
    return type(value) == "table"
end
local function guarded(value: unknown)
    if type(value) == "table" and isRecord(value) then
        local record: Table<string, unknown> = value
    end
end

-- A utility over any table is generic, as the lib's pairs is; a table stating
-- members passes as one whose keys and values are those members'. Its values
-- narrow like any unknown.
local function deepCopy<K, V>(t: Table<K, V>): Table<K, V>
    local copy: Table<K, V> = {}
    for key, value in pairs(t) do
        if type(value) == "table" then
            local inner: unknown = value.anything
        end
        rawset(copy, key, value)
    end
    return copy
end
local named: Named = { name = "x" }
local copied: Table<"name", string> = deepCopy(named)
local list: number[] = deepCopy({ 1, 2 })
local record: Table<string, number> = deepCopy({} as Table<string, number>)
for key, value in pairs(named) do
    local asString: string = value
end

-- Table<{}, unknown> is a type of its own, not the type of every table: K is
-- invariant, so a Table of narrower keys is not one, nor is a `table`, and it is not
-- a Table of narrower keys. A record passes because its members stand in for the
-- index signature, as they do for every Table.
local function readAny(t: Table<{}, unknown>): unknown
    return t.anything
end
readAny(named)
readAny({} as Table<string, number>)
readAny({} as table)
local function toStated(t: Table<{}, unknown>): Table<string, unknown>
    return t
end

-- pairs and next keep one signature, so a caller that reads it still correlates
-- the table's keys and values.
local numbers: Table<string, number> = {}
local ok, firstName, firstNumber = pcall(next, numbers)
local n: number | nil = firstNumber
local okPairs, iterate = pcall(pairs, numbers)
local typedIterator: LuaPairsIterator<string, number> = iterate
for key, value in pairs(42) do end

-- A union with an opaque `table` iterates with the keys and values that side may
-- hold, not only the keyed side's.
local function iterateMixed(x: Table<string, number> | table, y: Named | table)
    for key, value in pairs(x) do
        local s: string = key
        local n: number = value
    end
    for key, value in pairs(y) do
        local s: string = key
        local n: string = value
    end
    local firstKey, firstValue = next(x)
    local k: string | nil = firstKey
    local v: number = firstValue
    local afterKey = next(x, "a")
    local k2: string | nil = afterKey
end
-- The same holds for any signature of that shape: a stated key, an intersection.
interface Tagged { tag?: string; }
local function valuesOf<V>(t: Table<string, V> | table): V
    return nil as any
end
local function taggedValues<K, V>(t: (Table<K, V> & Tagged) | table): V
    return nil as any
end
local function useMixed(x: Table<string, number> | table)
    local a: number = valuesOf(x)
    local b: number = taggedValues(x as (Table<string, number> & Tagged) | table)
end


//// [tluaTableTagNarrowing.lua]
-- `type(x) == "table"` says a value is a table and nothing about its keys, so a type
-- that states nothing about the value (unknown, a bare type parameter) narrows to
-- Table<{}, unknown>: any key reads unknown through its index signature, and
-- pairs visits it. A declared `table` states no key and stays closed; code that reads
-- arbitrary keys declares Table<K, V>.
local function isMarked(value)
  return type(value) == "table" and value._marked == true;
end
-- Flow facts narrow the read, in either spelling.
local function readCode(err)
  if type(err) == "table" and type(err.code) == "string" then
    return err.code;
  end
  return nil;
end
local function readSlot(err)
  if type(err) == "table" and type(err["name"]) == "string" then
    return err["name"];
  end
  return nil;
end
-- Any key reads: strings, numbers, booleans, tables.
local function readKeys(value, key, other)
  if type(value) ~= "table" then
    return nil, nil, nil, nil;
  end
  return value[key], value[1], value[true], value[other];
end
-- The narrowed type is the declared Table<{}, unknown>, so it iterates, and it
-- is not a table of anything narrower.
local function keysOf(value)
  local keys = {};
  if type(value) == "table" then
    for key, v in pairs(value) do
      keys[#keys + 1] = key;
      local asString = key;
    end
    local anyTable = value;
    local asTable = value;
    local numbers = value;
    local strings = value;
    local cast = value;
  end
  return keys;
end
-- A generic narrows to an intersection with it. It reads any key, but a generic
-- is indexed for reading only, so writes need a stated type.
local function readGeneric(value)
  if type(value) == "table" then
    value.field = 1;
    value["slot"] = 2;
    return value.field;
  end
  return nil;
end
local function readUnion(value)
  if type(value) == "table" then
    return value.name;
  end
  return nil;
end
local function readList(value)
  if type(value) == "table" then
    return value[1];
  end
  return nil;
end
local function notTable(value)
  if type(value) ~= "table" then
    return value;
  end
  return value.name;
end
-- A declared `table` states no key: it stays `table` under the tag, a read is an
-- error, and it does not iterate. Declare the keys it holds instead.
local function readDeclared(t)
  if type(t) == "table" then
    return t.name;
  end
  return nil;
end
-- pairs and next visit it with unknown keys and values, as Luau's do; ipairs and
-- the table library want a stated array.
local function iterateDeclared(t, list)
  for key, value in pairs(t) do
    local k = key;
    local n = key;
    local v = value;
    local s = value;
  end
  local firstKey, firstValue = next(t);
  local k2 = firstKey;
  local v2 = firstValue;
  for key, value in next, t do
    local v3 = value;
  end
  for index, value in pairs(list) do
    local i = index;
    local s = value;
  end
  for index, value in ipairs(t) do
  end
  table.insert(t, 1);
end
local function readStated(t)
  for key, value in pairs(t) do
    local asString = key;
  end
  return t.name;
end
local function readConstrained(t)
  if type(t) == "table" then
    return t.field;
  end
  return nil;
end
-- Writes go through the index signature too.
local function mark(value)
  if type(value) == "table" then
    value._marked = true;
    value[1] = "x";
  end
end
-- Nested reads go through unknown, which states nothing.
local function nested(value)
  if type(value) == "table" then
    return value.inner.value;
  end
  return nil;
end
-- A key no table takes is an error on the narrowed table too.
local function readBadKeys(value, k, u)
  if type(value) == "table" then
    local a = value[k];
    local b = value[u];
  end
end
-- The narrowed table is a table, stays one under a second check, and takes a
-- metatable.
local function useAsTable(value)
  if type(value) == "table" then
    if type(value) == "table" then
      local again = value;
    end
    setmetatable(value, nil);
    local n = #value;
  end
end
-- A type guard to a stated shape still narrows the table to it.
local function isRecord(value)
  return type(value) == "table";
end
local function guarded(value)
  if type(value) == "table" and isRecord(value) then
    local record = value;
  end
end
-- A utility over any table is generic, as the lib's pairs is; a table stating
-- members passes as one whose keys and values are those members'. Its values
-- narrow like any unknown.
local function deepCopy(t)
  local copy = {};
  for key, value in pairs(t) do
    if type(value) == "table" then
      local inner = value.anything;
    end
    rawset(copy, key, value);
  end
  return copy;
end
local named = { name = "x" };
local copied = deepCopy(named);
local list = deepCopy({ 1, 2 });
local record = deepCopy({});
for key, value in pairs(named) do
  local asString = value;
end
-- Table<{}, unknown> is a type of its own, not the type of every table: K is
-- invariant, so a Table of narrower keys is not one, nor is a `table`, and it is not
-- a Table of narrower keys. A record passes because its members stand in for the
-- index signature, as they do for every Table.
local function readAny(t)
  return t.anything;
end
readAny(named);
readAny({});
readAny({});
local function toStated(t)
  return t;
end
-- pairs and next keep one signature, so a caller that reads it still correlates
-- the table's keys and values.
local numbers = {};
local ok, firstName, firstNumber = pcall(next, numbers);
local n = firstNumber;
local okPairs, iterate = pcall(pairs, numbers);
local typedIterator = iterate;
for key, value in pairs(42) do
end
-- A union with an opaque `table` iterates with the keys and values that side may
-- hold, not only the keyed side's.
local function iterateMixed(x, y)
  for key, value in pairs(x) do
    local s = key;
    local n = value;
  end
  for key, value in pairs(y) do
    local s = key;
    local n = value;
  end
  local firstKey, firstValue = next(x);
  local k = firstKey;
  local v = firstValue;
  local afterKey = next(x, "a");
  local k2 = afterKey;
end
local function valuesOf(t)
  return nil;
end
local function taggedValues(t)
  return nil;
end
local function useMixed(x)
  local a = valuesOf(x);
  local b = taggedValues(x);
end
