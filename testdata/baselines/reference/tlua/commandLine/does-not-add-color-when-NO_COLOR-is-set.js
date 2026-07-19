currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
Version FakeTSVersion
tlua: The Lua Compiler - Version FakeTSVersion

COMMON COMMANDS

  tlua
  Compiles the current project (tluaconfig.json in the working directory.)

  tlua app.tlua util.tlua
  Ignoring tluaconfig.json, compiles the specified files with default compiler options.

  tlua -b
  Build a composite project in the working directory.

  tlua --init
  Creates a tluaconfig.json with the recommended settings in the working directory.

  tlua -p ./path/to/tluaconfig.json
  Compiles the Lua project located at the specified path.

  tlua --help --all
  An expanded version of this information, showing all possible compiler options

  tlua --noEmit
  tlua --target esnext
  Compiles the current project, with additional settings.

COMMAND LINE FLAGS

--help, -h
Print this message.

--watch, -w
Watch input files.

--all
Show all compiler options.

--version, -v
Print the compiler's version.

--init
Initializes a tlua project and creates a tluaconfig.json file.

--project, -p
Compile the project given the path to its configuration file, or to a folder with a 'tluaconfig.json'.

--showConfig
Print the final configuration instead of building.

--ignoreConfig
Ignore the tluaconfig found and build with commandline options and files.

--build, -b
Build one or more projects and their dependencies, if out of date

COMMON COMPILER OPTIONS

--pretty
Enable color and formatting in tlua's output to make compiler errors easier to read.
type: boolean
default: true

--declaration, -d
Generate .d.ts files from Lua and JavaScript files in your project.
type: boolean
default: `false`, unless `composite` is set

--declarationMap
Create sourcemaps for d.ts files.
type: boolean
default: false

--emitDeclarationOnly
Only output d.ts files and not JavaScript files.
type: boolean
default: false

--sourceMap
Create source map files for emitted JavaScript files.
type: boolean
default: false

--noEmit
Disable emitting files from a compilation.
type: boolean
default: false

--target, -t
Set the JavaScript language version for emitted JavaScript and include compatible library declarations.
one of: es6/es2015, es2016, es2017, es2018, es2019, es2020, es2021, es2022, es2023, es2024, es2025, esnext
default: es2025

--module, -m
Specify what module code is generated.
one of: commonjs, es6/es2015, es2020, es2022, esnext, node16, node18, node20, nodenext, preserve
default: undefined

--lib
Specify a set of bundled library declaration files that describe the target runtime environment.
one or more: luajit
default: undefined

--jsx
Specify what JSX code is generated.
one of: preserve, react-native, react-jsx, react-jsxdev, react
default: undefined

--outFile
Specify a file that bundles all outputs into one JavaScript file. If 'declaration' is true, also designates a file that bundles all .d.ts output.

--outDir
Specify an output folder for all emitted files.

--removeComments
Disable emitting comments.
type: boolean
default: false

--strict
Enable all strict type-checking options.
type: boolean
default: true

--types
Specify type package names to be included without being referenced in a source file.

You can learn about all of the compiler options at https://aka.ms/tsc


