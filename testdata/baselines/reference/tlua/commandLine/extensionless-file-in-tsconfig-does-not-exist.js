currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
                    "files": ["./src/doesNotExist"]
                    }

tlua -p ./tluaconfig.json
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[91merror[0m[90m TLUA6231: [0mCould not resolve the path '/home/src/workspaces/project/src/doesNotExist' with the extensions: '.tlua', '.tsx', '.d.tlua'.
  The file is in the program because:
    Part of 'files' list in tluaconfig.json
  [96mtluaconfig.json[0m:[93m2[0m:[93m31[0m - File is matched by 'files' list specified here.
    [7m2[0m                     "files": ["./src/doesNotExist"]
    [7m [0m [96m                              ~~~~~~~~~~~~~~~~~~~~[0m


Found 1 error.

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

