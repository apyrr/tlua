# tlua

This extension provides language support for **tlua** — a typed Lua dialect based on
TypeScript — via the native tlua language server. It provides features like
go-to-definition, completions, errors and diagnostics, quick info/tooltip hovers, and
more.

## Usage

1. Install the extension.
2. Open a tlua file (`.tlua` or `.d.tlua`) in your editor. The language server starts
   automatically.

## Configuration

Settings live under the `tlua.*` namespace. A few useful ones:

```jsonc
{
    // Point at a directory containing a local tlua binary instead of the
    // bundled one (e.g. a local build).
    "tlua.tsdk.path": "./built/local",

    // Trace language server communication ("off" | "messages" | "verbose").
    "tlua.trace.server": "off"
}
```

## Feedback

If you encounter any issues or have suggestions for improvement, please open an issue on
the [GitHub repository](https://github.com/apyrr/tlua).
