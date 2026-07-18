# tlua

`tlua` is a typed Lua compiler, language service, and tooling stack based on the
TypeScript compiler architecture.

## Usage

Install the package and run the compiler:

```sh
npm install --save-dev tlua
npx tlua --help
```

The package selects the native binary for the current operating system and
architecture through an optional `@tlua/tlua-<platform>` dependency.

## JavaScript API

The experimental JavaScript API is available through the `tlua/unstable/*`
exports. These APIs may change between releases.

## Issues and feedback

Report problems on the [issue tracker](https://github.com/apyrr/tlua/issues).
