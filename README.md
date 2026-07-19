# tlua

`tlua` is a typed Lua compiler, language service, and tooling stack. It grafts
TypeScript's type system onto Lua's syntax and semantics: you write Lua
(`function … end`, `local`, `{ x = 1 }` tables, `and`/`or`/`not`, colon methods,
`goto`, `require`), and it is statically type-checked and compiled to Lua.

tlua is a hard fork of
[microsoft/typescript-go](https://github.com/microsoft/typescript-go): it keeps
the checker and language-service architecture and replaces the language.

## Install

```sh
npm install --save-dev @tlua/cli
npx tlua --help
```

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
