This repository is **tlua**: a typed Lua compiler, language service, and
tooling stack.

tlua is a hard fork of TypeScript-Go (`microsoft/typescript-go`): it keeps the
checker and language-service architecture and replaces the language.

Goal: typed Lua based on TypeScript.
Non-goal: typed Lua with TypeScript compatibility.

## Repository Layout

- `internal` contains most compiler and language service code.
- `_extension` contains the preview VS Code extension.
- Most behavior should be tested with compiler tests.

## Working Rules

- Add short comments that clarify non-obvious code.
- Do not add or change dependencies unless asked.
- Do not remove debug assertions or panic calls just because they look
  strict.
- Do not use `timeout` around tests or commands unless specifically
  debugging a hang.
- Do not change ts protocol version.
- Run tests less often. For example if task is large, run tests after
  full implementation.

## Common Commands

Run `npx hereby --tasks` to see all available tasks.

```sh
npx hereby build  # Build the tlua binary; not required for normal tests
npx hereby test   # Run the main test suite
npx hereby lint   # Run linters
npx hereby format # Format the repo
```

## Testing

CI installs dependencies and then runs these test tasks:

```sh
npm ci
npx hereby test
npx hereby test:benchmarks
npx hereby test:tools
npx hereby test:api
```

For day-to-day local testing, use `npx hereby test` (or equivalently `npm
test`). For the full CI test-job equivalent, use:

```sh
npx hereby test:all
```

`hereby test` runs Go tests over `./...`, clears
`testdata/baselines/local` first, and supports a filter:

```sh
npx hereby test --tests 'TestLocal/myTest'
```

Filtered `hereby test -t ...` runs can still be slow because they invoke
`go test -run='<filter>' ./...`; most packages may report `[no tests to run]`,
but they still have to load or compile, and central changes invalidate cache.
Use direct `go test ./internal/testrunner` for quick compiler-test debugging,
then rerun the filtered hereby command before finalizing because it clears
`testdata/baselines/local` and matches repo behavior.

For direct debugging of one compiler test:

```sh
go test -run='TestLocal/<test name>' ./internal/testrunner -v
```

Prereqs: `npm ci`, Go, Node, and the TypeScript submodule for full
coverage. CI uses Node LTS and the repo's Microsoft Go setup action,
which defaults to `go1.26`.

## Compiler Tests And Baselines

New compiler tests live under `testdata/tests/cases/compiler/`.
This repo primarily uses snapshot/baseline/golden tests rather than
small unit tests.

When tests run, they write local output under:

```text
testdata/baselines/local
```

Compare that output with:

```text
testdata/baselines/reference
```

If baseline output changes are expected, accept it with:

```sh
npx hereby baseline-accept
git diff
```

For a focused compiler test:

```sh
go test -run='TestLocal/<test name>' ./internal/testrunner
```

`TestLocal` runs the tests in `testdata/tests/cases`. (The vendored
upstream corpus and its `TestSubmodule` runner have been retired.)

For a detailed guide to writing compiler and fourslash tests (file
formats, test directives, fourslash markers, and baselines), see
[.github/skills/compiler-and-fourslash-tests/SKILL.md](.github/skills/compiler-and-fourslash-tests/SKILL.md).
