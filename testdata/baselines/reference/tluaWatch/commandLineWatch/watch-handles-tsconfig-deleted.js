currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/index.tlua] *new* 
local x = 1;
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{}

tlua --watch
ExitStatus:: Success
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[[90mHH:MM:SS AM[0m] Found 0 errors. Watching for file changes.

//// [/home/src/tslibs/TS/Lib/lib.luajit.d.tlua] *Lib*
/// <reference no-default-lib="true"/>
interface Boolean {}
interface Function {}
interface CallableFunction {}
interface NewableFunction {}
interface IArguments {}
interface Number { toExponential: any; }
interface Object {}
interface RegExp {}
interface String { charAt: any; }
interface Array<T> { length: number; [n: number]: T; }
interface ReadonlyArray<T> {}
interface SymbolConstructor {
    (desc?: string | number): symbol;
    for(name: string): symbol;
    readonly toStringTag: symbol;
}
declare Symbol: SymbolConstructor;
interface Symbol {
    readonly [Symbol.toStringTag]: string;
}
declare console: { log(msg: any): void; };
declare function require(module: string): any;
//// [/home/src/workspaces/project/index.lua] *new* 
local x = 1;


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/index.tlua
Signatures::


Edit [0]:: delete tsconfig
//// [/home/src/workspaces/project/tluaconfig.json] *deleted*


Output::
[91merror[0m[90m TLUA5083: [0mCannot read file '/home/src/workspaces/project/tluaconfig.json'.
[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)


Diff:: incremental reports config read error while clean build without tsconfig prints usage help
--- nonIncremental.output.txt
+++ incremental.output.txt
@@ -1,129 +1,1 @@
-Version FakeTSVersion
-tlua: The Lua Compiler - Version FakeTSVersion
-
-[1mCOMMON COMMANDS[22m
-
-  [94mtlua[39m
-  Compiles the current project (tluaconfig.json in the working directory.)
-
-  [94mtlua app.tlua util.tlua[39m
-  Ignoring tluaconfig.json, compiles the specified files with default compiler options.
-
-  [94mtlua -b[39m
-  Build a composite project in the working directory.
-
-  [94mtlua --init[39m
-  Creates a tluaconfig.json with the recommended settings in the working directory.
-
-  [94mtlua -p ./path/to/tluaconfig.json[39m
-  Compiles the Lua project located at the specified path.
-
-  [94mtlua --help --all[39m
-  An expanded version of this information, showing all possible compiler options
-
-  [94mtlua --noEmit[39m
-  [94mtlua --target esnext[39m
-  Compiles the current project, with additional settings.
-
-[1mCOMMAND LINE FLAGS[22m
-
-[94m--help, -h[39m
-Print this message.
-
-[94m--watch, -w[39m
-Watch input files.
-
-[94m--all[39m
-Show all compiler options.
-
-[94m--version, -v[39m
-Print the compiler's version.
-
-[94m--init[39m
-Initializes a tlua project and creates a tluaconfig.json file.
-
-[94m--project, -p[39m
-Compile the project given the path to its configuration file, or to a folder with a 'tluaconfig.json'.
-
-[94m--showConfig[39m
-Print the final configuration instead of building.
-
-[94m--ignoreConfig[39m
-Ignore the tluaconfig found and build with commandline options and files.
-
-[94m--build, -b[39m
-Build one or more projects and their dependencies, if out of date
-
-[1mCOMMON COMPILER OPTIONS[22m
-
-[94m--pretty[39m
-Enable color and formatting in tlua's output to make compiler errors easier to read.
-type: boolean
-default: true
-
-[94m--declaration, -d[39m
-Generate .d.ts files from Lua and JavaScript files in your project.
-type: boolean
-default: `false`, unless `composite` is set
-
-[94m--declarationMap[39m
-Create sourcemaps for d.ts files.
-type: boolean
-default: false
-
-[94m--emitDeclarationOnly[39m
-Only output d.ts files and not JavaScript files.
-type: boolean
-default: false
-
-[94m--sourceMap[39m
-Create source map files for emitted JavaScript files.
-type: boolean
-default: false
-
-[94m--noEmit[39m
-Disable emitting files from a compilation.
-type: boolean
-default: false
-
-[94m--target, -t[39m
-Set the JavaScript language version for emitted JavaScript and include compatible library declarations.
-one of: es6/es2015, es2016, es2017, es2018, es2019, es2020, es2021, es2022, es2023, es2024, es2025, esnext
-default: es2025
-
-[94m--module, -m[39m
-Specify what module code is generated.
-one of: commonjs, es6/es2015, es2020, es2022, esnext, node16, node18, node20, nodenext, preserve
-default: undefined
-
-[94m--lib[39m
-Specify a set of bundled library declaration files that describe the target runtime environment.
-one or more: luajit
-default: undefined
-
-[94m--jsx[39m
-Specify what JSX code is generated.
-one of: preserve, react-native, react-jsx, react-jsxdev, react
-default: undefined
-
-[94m--outFile[39m
-Specify a file that bundles all outputs into one JavaScript file. If 'declaration' is true, also designates a file that bundles all .d.ts output.
-
-[94m--outDir[39m
-Specify an output folder for all emitted files.
-
-[94m--removeComments[39m
-Disable emitting comments.
-type: boolean
-default: false
-
-[94m--strict[39m
-Enable all strict type-checking options.
-type: boolean
-default: true
-
-[94m--types[39m
-Specify type package names to be included without being referenced in a source file.
-
-You can learn about all of the compiler options at https://aka.ms/tsc
-
+[91merror[0m[90m TLUA5083: [0mCannot read file '/home/src/workspaces/project/tluaconfig.json'.