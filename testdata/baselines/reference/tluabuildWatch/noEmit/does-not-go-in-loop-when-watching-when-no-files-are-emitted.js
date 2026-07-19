currentDirectory::/user/username/projects/myproject
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/myproject/a.tlua] *new* 

//// [/user/username/projects/myproject/b.tlua] *new* 

//// [/user/username/projects/myproject/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "noEmit": true,
    },
}

tlua -b -w -verbose
ExitStatus:: Success
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output file 'tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

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
//// [/user/username/projects/myproject/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":["./a.tlua","./b.tlua"]}
//// [/user/username/projects/myproject/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    },
    {
      "files": [
        "./b.tlua"
      ],
      "original": "./b.tlua"
    }
  ],
  "size": 58
}

Watch Registrations::
Directory watches::
  /user/username/projects/myproject (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/myproject/a.tlua
*refresh*    /user/username/projects/myproject/b.tlua
Signatures::


Edit [0]:: No change
//// [/user/username/projects/myproject/a.tlua] *mTime changed*


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output 'tluaconfig.tluabuildinfo' is older than input 'a.tlua'

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[[90mHH:MM:SS AM[0m] Found 0 errors. Watching for file changes.

//// [/user/username/projects/myproject/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/user/username/projects/myproject/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

Watch Registrations::
Directory watches::
  /user/username/projects/myproject (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/myproject/a.tlua
*refresh*    /user/username/projects/myproject/b.tlua
Signatures::


Edit [1]:: change
//// [/user/username/projects/myproject/a.tlua] *modified* 
local x = 10;


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output 'tluaconfig.tluabuildinfo' is older than input 'a.tlua'

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[[90mHH:MM:SS AM[0m] Found 0 errors. Watching for file changes.

//// [/user/username/projects/myproject/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/user/username/projects/myproject/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

Watch Registrations::
Directory watches::
  /user/username/projects/myproject (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/myproject/a.tlua
*refresh*    /user/username/projects/myproject/b.tlua
Signatures::
