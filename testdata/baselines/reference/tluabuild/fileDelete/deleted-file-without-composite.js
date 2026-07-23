currentDirectory::/home/src/workspaces/solution
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/solution/child/child.tlua] *new* 
local child2 = require("child2");
function child()
    child2.child2();
end
return { child = child };
//// [/home/src/workspaces/solution/child/child2.tlua] *new* 
function child2() end
return { child2 = child2 };
//// [/home/src/workspaces/solution/child/tluaconfig.json] *new* 
{
    "compilerOptions": { }
}

tlua --b child/tluaconfig.json -v --traceResolution --explainFiles
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * child/tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'child/tluaconfig.json' is out of date because output file 'child/tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'child/tluaconfig.json'...

======== Resolving module 'child2' from '/home/src/workspaces/solution/child/child.tlua'. ========
File '/home/src/workspaces/solution/child/child2.tlua' exists - use it as a name resolution result.
======== Module name 'child2' was successfully resolved to '/home/src/workspaces/solution/child/child2.tlua'. ========
../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
child/child2.tlua
   Imported via "child2" from file 'child/child.tlua'
   Matched by default include pattern '**/*'
child/child.tlua
   Matched by default include pattern '**/*'
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
//// [/home/src/workspaces/solution/child/child.lua] *new* 
local child2 = require("child2");
function child()
  child2.child2();
end
return { child = child };

//// [/home/src/workspaces/solution/child/child2.lua] *new* 
function child2()
end
return { child2 = child2 };

//// [/home/src/workspaces/solution/child/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":["./child.tlua","./child2.tlua"]}
//// [/home/src/workspaces/solution/child/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./child.tlua"
      ],
      "original": "./child.tlua"
    },
    {
      "files": [
        "./child2.tlua"
      ],
      "original": "./child2.tlua"
    }
  ],
  "size": 67
}

child/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/solution/child/child2.tlua
*refresh*    /home/src/workspaces/solution/child/child.tlua
Signatures::


Edit [0]:: delete child2 file
//// [/home/src/workspaces/solution/child/child2.lua] *deleted*
//// [/home/src/workspaces/solution/child/child2.tlua] *deleted*

tlua --b child/tluaconfig.json -v --traceResolution --explainFiles
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * child/tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'child/tluaconfig.json' is out of date because buildinfo file 'child/tluaconfig.tluabuildinfo' indicates that file 'child/child2.tlua' was root file of compilation but not any more.

[[90mHH:MM:SS AM[0m] Building project 'child/tluaconfig.json'...

======== Resolving module 'child2' from '/home/src/workspaces/solution/child/child.tlua'. ========
File '/home/src/workspaces/solution/child/child2.tlua' does not exist.
File '/home/src/workspaces/solution/child/child2.tsx' does not exist.
File '/home/src/workspaces/solution/child/child2/init.tlua' does not exist.
File '/home/src/workspaces/solution/child/child2/init.tsx' does not exist.
File '/home/src/workspaces/solution/child/child2.lua' does not exist.
File '/home/src/workspaces/solution/child/child2/init.lua' does not exist.
======== Module name 'child2' was not resolved. ========
../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
child/child.tlua
   Matched by default include pattern '**/*'
//// [/home/src/workspaces/solution/child/child.lua] *rewrite with same content*
//// [/home/src/workspaces/solution/child/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":["./child.tlua"]}
//// [/home/src/workspaces/solution/child/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./child.tlua"
      ],
      "original": "./child.tlua"
    }
  ],
  "size": 51
}

child/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/solution/child/child.tlua
Signatures::
