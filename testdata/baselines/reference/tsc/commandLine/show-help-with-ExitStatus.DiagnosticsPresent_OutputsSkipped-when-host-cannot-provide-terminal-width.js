currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
Version FakeTSVersion
tsc: The TypeScript Compiler - Version FakeTSVersion

[1mCOMMON COMMANDS[22m

  [94mtsc[39m
  Compiles the current project (tsconfig.json in the working directory.)

  [94mtsc app.ts util.ts[39m
  Ignoring tsconfig.json, compiles the specified files with default compiler options.

  [94mtsc -b[39m
  Build a composite project in the working directory.

  [94mtsc --init[39m
  Creates a tsconfig.json with the recommended settings in the working directory.

  [94mtsc -p ./path/to/tsconfig.json[39m
  Compiles the TypeScript project located at the specified path.

  [94mtsc --help --all[39m
  An expanded version of this information, showing all possible compiler options

  [94mtsc --noEmit[39m
  [94mtsc --target esnext[39m
  Compiles the current project, with additional settings.

[1mCOMMAND LINE FLAGS[22m

[94m--help, -h[39m
Print this message.

[94m--watch, -w[39m
Watch input files.

[94m--all[39m
Show all compiler options.

[94m--version, -v[39m
Print the compiler's version.

[94m--init[39m
Initializes a TypeScript project and creates a tsconfig.json file.

[94m--project, -p[39m
Compile the project given the path to its configuration file, or to a folder with a 'tsconfig.json'.

[94m--showConfig[39m
Print the final configuration instead of building.

[94m--ignoreConfig[39m
Ignore the tsconfig found and build with commandline options and files.

[94m--build, -b[39m
Build one or more projects and their dependencies, if out of date

[1mCOMMON COMPILER OPTIONS[22m

[94m--pretty[39m
Enable color and formatting in TypeScript's output to make compiler errors easier to read.
type: boolean
default: true

[94m--declaration, -d[39m
Generate .d.ts files from TypeScript and JavaScript files in your project.
type: boolean
default: `false`, unless `composite` is set

[94m--declarationMap[39m
Create sourcemaps for d.ts files.
type: boolean
default: false

[94m--emitDeclarationOnly[39m
Only output d.ts files and not JavaScript files.
type: boolean
default: false

[94m--sourceMap[39m
Create source map files for emitted JavaScript files.
type: boolean
default: false

[94m--noEmit[39m
Disable emitting files from a compilation.
type: boolean
default: false

[94m--target, -t[39m
Set the JavaScript language version for emitted JavaScript and include compatible library declarations.
one of: es6/es2015, es2016, es2017, es2018, es2019, es2020, es2021, es2022, es2023, es2024, es2025, esnext
default: es2025

[94m--module, -m[39m
Specify what module code is generated.
one of: commonjs, es6/es2015, es2020, es2022, esnext, node16, node18, node20, nodenext, preserve
default: undefined

[94m--lib[39m
Specify a set of bundled library declaration files that describe the target runtime environment.
one or more: luajit
default: undefined

[94m--jsx[39m
Specify what JSX code is generated.
one of: preserve, react-native, react-jsx, react-jsxdev, react
default: undefined

[94m--outFile[39m
Specify a file that bundles all outputs into one JavaScript file. If 'declaration' is true, also designates a file that bundles all .d.ts output.

[94m--outDir[39m
Specify an output folder for all emitted files.

[94m--removeComments[39m
Disable emitting comments.
type: boolean
default: false

[94m--strict[39m
Enable all strict type-checking options.
type: boolean
default: true

[94m--types[39m
Specify type package names to be included without being referenced in a source file.

You can learn about all of the compiler options at https://aka.ms/tsc


