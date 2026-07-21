# tlua

`tlua` is a typed Lua compiler, language service, and tooling stack. It brings
TypeScript's type system to Lua's syntax and semantics: you write Lua
(`function … end`, `local`, `{ x = 1 }` tables, `and`/`or`/`not`, colon methods,
`goto`, `require`), and it is statically type-checked and compiled to Lua.

tlua is a hard fork of
[microsoft/typescript-go](https://github.com/microsoft/typescript-go): it keeps
the checker and language-service architecture and replaces the language.

## At a glance

```lua
-- main.tlua
type Vec2 = { x: number, y: number }

local function length(v: Vec2): number
  return math.sqrt(v.x * v.x + v.y * v.y)
end

local origin: Vec2 = { x = 3, y = 4 }
print(length(origin))
```

compiles to plain Lua - the types simply disappear:

```lua
-- main.lua
local function length(v)
    return math.sqrt(v.x * v.x + v.y * v.y);
end
local origin = { x = 3, y = 4 };
print(length(origin));
```

## Install

```sh
npm install --save-dev @tlua/cli
npx tlua --help
```

## Your first project

Create a `tluaconfig.json`:

```json
{
  "compilerOptions": {
    "strict": true,
    "rootDir": "src",
    "outDir": "out"
  },
  "include": ["src"]
}
```

Then compile:

```sh
npx tlua              # compile the project in the current directory
npx tlua -p ./path    # compile the project at a given path
npx tlua --watch      # recompile on change
npx tlua --noEmit     # type-check only
npx tlua app.tlua     # ignore tluaconfig.json, compile these files
```

Sources are `.tlua`; declaration files are `.d.tlua`; output is `.lua`.

## A short tour

**Type errors point at Lua source, the way you'd expect:**

```lua
local function area(w: number, h: number): number
  return w * h
end

print(area(3, "4"))
```

```text
oops.tlua(5,15): error TLUA2345: Argument of type 'string' is not assignable to parameter of type 'number'.
```

**Unions, narrowing, and generics.** `nil` is tlua's absent value (there is no
`null`/`undefined` split), and `if`/`elseif` chains narrow discriminated unions:

```lua
type Shape =
  | { kind: "circle", radius: number }
  | { kind: "rect", w: number, h: number }

local function describe(s: Shape): string
  if s.kind == "circle" then
    return "circle r=" .. s.radius
  else
    return "rect " .. s.w .. "x" .. s.h
  end
end

local function first<T>(items: T[]): T | nil
  for _, item in ipairs(items) do
    return item
  end
  return nil
end

local shapes: Shape[] = { { kind = "circle", radius = 2 } }
local head = first(shapes)
if head ~= nil and head.kind == "circle" then
  print(describe(head))
end
```

**Modules are Lua modules.** There is no `import`/`export`: a module returns a
table, and `require` is typed against it.

```lua
-- src/counter.tlua
interface Counter {
  count: number,
  bump: (self: Counter, by: number) => number,
}

local function create(): Counter
  return {
    count = 0,
    bump = function(self: Counter, by: number): number
      self.count = self.count + by
      return self.count
    end,
  }
end

return { create = create }
```

```lua
-- src/main.tlua
local counter = require("counter")

local c = counter.create()
c:bump(5)          -- colon calls type the implicit `self`
print(c.count)
```

**Ambient declarations** describe globals your host environment injects, in a
`.d.tlua` file:

```lua
-- globals.d.tlua
declare version: string
declare limits: { min: number, max: number }
```

**Optional chaining and string templates** lower to ordinary Lua. `?.`
short-circuits on `nil` using `and` guards, hoisting a temporary when a
subexpression would otherwise be evaluated twice, and `` `...${}` `` templates
become `..` concatenation with `tostring`:

```lua
type User = { name: string, address: { city: string } | nil }

local function cityOf(u: User | nil): string
  local city = u?.address?.city
  return city or "unknown"
end

local u: User = { name = "ada", address = { city = "london" } }
print(`${u.name} lives in ${cityOf(u)}`)
```

```lua
local function cityOf(u)
    local _a = u and u.address;
    local city = _a and _a.city;
    return city or "unknown";
end
local u = { name = "ada", address = { city = "london" } };
print((tostring(u.name) .. " lives in " .. tostring(cityOf(u))));
```

The checker tracks the `nil` through the chain, so `city` is typed
`string | nil` and the `or` default is what makes `cityOf` return a plain
`string`.

## Coming from TypeScript?

tlua's goal is typed Lua, *not* TypeScript compatibility. The type system is
familiar - interfaces, unions, generics, narrowing, mapped and conditional
types - but the language is Lua:

- `nil` replaces both `null` and `undefined`.
- `local` is the only declaration form; no `let`/`const`/`var`.
- Modules use `require` and a top-level `return`, not `import`/`export`.
- `and`/`or`/`not` are the logical operators, with Lua truthiness (`0` and `""`
  are truthy). There is no ternary or conditional expression.
- `?.` is supported, but `??` and `??=` are not - use `or` for defaults.
- Tables use `=` keys (`{ x = 1 }`) and are 1-based.
- No `switch` and no `bigint`. Use `if`/`elseif` chains for exhaustive
  dispatch.
- `..` concatenates; `#` measures length.

## VS Code extension

The extension in `_extension` provides language support for `.tlua` and
`.d.tlua` files, including diagnostics, completions, navigation, and hovers.

Build a VSIX for the current platform with:

```sh
npx hereby tlua:pack-extensions
```

Artifacts are written to `built/vsix`.

## Development

```sh
npm ci
npx hereby test
```

Run `npx hereby --tasks` to see all available build, test, and lint tasks.

## License

Apache-2.0. tlua is a fork of
[microsoft/typescript-go](https://github.com/microsoft/typescript-go); see
[LICENSE](LICENSE) and [NOTICE.txt](NOTICE.txt) for the retained upstream
copyright and attribution.
