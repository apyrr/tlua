currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/tslibs/TS/Lib/lib.es2015.iterable.d.tlua] *new* 
interface SymbolConstructor {
    readonly iterator: unique symbol;
}
interface IteratorYieldResult<TYield> {
    done?: false;
    value: TYield;
}
interface IteratorReturnResult<TReturn> {
    done: true;
    value: TReturn;
}
type IteratorResult<T, TReturn = any> = IteratorYieldResult<T> | IteratorReturnResult<TReturn>;
interface Iterator<T, TReturn = any, TNext = any> {
    // NOTE: 'next' is defined using a tuple to ensure we report the correct assignability errors in all places.
    next(...[value]: [] | [TNext]): IteratorResult<T, TReturn>;
    return?(value?: TReturn): IteratorResult<T, TReturn>;
    throw?(e?: any): IteratorResult<T, TReturn>;
}
interface Iterable<T, TReturn = any, TNext = any> {
    [Symbol.iterator](): Iterator<T, TReturn, TNext>;
}
interface IterableIterator<T, TReturn = any, TNext = any> extends Iterator<T, TReturn, TNext> {
    [Symbol.iterator](): IterableIterator<T, TReturn, TNext>;
}
interface IteratorObject<T, TReturn = unknown, TNext = unknown> extends Iterator<T, TReturn, TNext> {
    [Symbol.iterator](): IteratorObject<T, TReturn, TNext>;
}
type BuiltinIteratorReturn = intrinsic;
interface ArrayIterator<T> extends IteratorObject<T, BuiltinIteratorReturn, unknown> {
    [Symbol.iterator](): ArrayIterator<T>;
}
interface Array<T> {
    [Symbol.iterator](): ArrayIterator<T>;
    entries(): ArrayIterator<[number, T]>;
    keys(): ArrayIterator<number>;
    values(): ArrayIterator<T>;
}
//// [/home/src/tslibs/TS/Lib/lib.es2017.full.d.tlua] *new* 
/// <reference lib="es2015.iterable"/>
interface File {
}
interface FileList {
    readonly length: number;
    item(index: number): File | null;
    [index: number]: File;
    [Symbol.iterator](): ArrayIterator<File>;
}/// <reference no-default-lib="true"/>
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
//// [/home/src/workspaces/project/a.tlua] *new* 
local createFileListFromFiles = function(files: File[]): FileList
local fileList: FileList = {
    length: files.length,
    item: function(index: number): File | null return files[index] || null end,
    [Symbol.iterator]: (0 as any) as () => ArrayIterator<File>,
} as unknown as FileList;

return fileList;
end;
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "target": "es2017",
        "strict": true
    }
}

tlua 
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96ma.tlua[0m:[93m3[0m:[93m11[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m     length: files.length,
[7m [0m [91m          ~[0m

[96ma.tlua[0m:[93m4[0m:[93m9[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m4[0m     item: function(index: number): File | null return files[index] || null end,
[7m [0m [91m        ~[0m

[96ma.tlua[0m:[93m5[0m:[93m22[0m - [91merror[0m[90m TLUA1005: [0m'=' expected.

[7m5[0m     [Symbol.iterator]: (0 as any) as () => ArrayIterator<File>,
[7m [0m [91m                     ~[0m


Found 3 errors in the same file, starting at: a.tlua[90m:3[0m

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
//// [/home/src/workspaces/project/a.lua] *new* 
local createFileListFromFiles = function(files)
  local fileList = {
    length, files.length,
    item, function(index: number): File | nil
      return files[index] or nil;
    end,
    [Symbol.iterator] = , (0 as any) as () => ArrayIterator<File>,
  };
  return fileList;
end;




Edit [0]:: no change

tlua 
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96ma.tlua[0m:[93m3[0m:[93m11[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m     length: files.length,
[7m [0m [91m          ~[0m

[96ma.tlua[0m:[93m4[0m:[93m9[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m4[0m     item: function(index: number): File | null return files[index] || null end,
[7m [0m [91m        ~[0m

[96ma.tlua[0m:[93m5[0m:[93m22[0m - [91merror[0m[90m TLUA1005: [0m'=' expected.

[7m5[0m     [Symbol.iterator]: (0 as any) as () => ArrayIterator<File>,
[7m [0m [91m                     ~[0m


Found 3 errors in the same file, starting at: a.tlua[90m:3[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*



Edit [1]:: no change with incremental

tlua --incremental
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96ma.tlua[0m:[93m3[0m:[93m11[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m     length: files.length,
[7m [0m [91m          ~[0m

[96ma.tlua[0m:[93m4[0m:[93m9[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m4[0m     item: function(index: number): File | null return files[index] || null end,
[7m [0m [91m        ~[0m

[96ma.tlua[0m:[93m5[0m:[93m22[0m - [91merror[0m[90m TLUA1005: [0m'=' expected.

[7m5[0m     [Symbol.iterator]: (0 as any) as () => ArrayIterator<File>,
[7m [0m [91m                     ~[0m


Found 3 errors in the same file, starting at: a.tlua[90m:3[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"root":[2],"fileNames":["lib.luajit.d.tlua","./a.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"dd33ec948bae8b9c1de34d88d969a302-local createFileListFromFiles = function(files: File[]): FileList\nlocal fileList: FileList = {\n    length: files.length,\n    item: function(index: number): File | null return files[index] || null end,\n    [Symbol.iterator]: (0 as any) as () => ArrayIterator<File>,\n} as unknown as FileList;\n\nreturn fileList;\nend;"],"options":{"strict":true,"target":4},"semanticDiagnosticsPerFile":[1,2]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./a.tlua"
  ],
  "fileInfos": [
    {
      "fileName": "lib.luajit.d.tlua",
      "version": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
      "signature": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./a.tlua",
      "version": "dd33ec948bae8b9c1de34d88d969a302-local createFileListFromFiles = function(files: File[]): FileList\nlocal fileList: FileList = {\n    length: files.length,\n    item: function(index: number): File | null return files[index] || null end,\n    [Symbol.iterator]: (0 as any) as () => ArrayIterator<File>,\n} as unknown as FileList;\n\nreturn fileList;\nend;",
      "signature": "dd33ec948bae8b9c1de34d88d969a302-local createFileListFromFiles = function(files: File[]): FileList\nlocal fileList: FileList = {\n    length: files.length,\n    item: function(index: number): File | null return files[index] || null end,\n    [Symbol.iterator]: (0 as any) as () => ArrayIterator<File>,\n} as unknown as FileList;\n\nreturn fileList;\nend;",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "strict": true,
    "target": 4
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "./a.tlua"
  ],
  "size": 1348
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
Signatures::


Edit [2]:: no change with incremental that reads buildInfo

tlua --incremental
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96ma.tlua[0m:[93m3[0m:[93m11[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m     length: files.length,
[7m [0m [91m          ~[0m

[96ma.tlua[0m:[93m4[0m:[93m9[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m4[0m     item: function(index: number): File | null return files[index] || null end,
[7m [0m [91m        ~[0m

[96ma.tlua[0m:[93m5[0m:[93m22[0m - [91merror[0m[90m TLUA1005: [0m'=' expected.

[7m5[0m     [Symbol.iterator]: (0 as any) as () => ArrayIterator<File>,
[7m [0m [91m                     ~[0m


Found 3 errors in the same file, starting at: a.tlua[90m:3[0m


tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
Signatures::
