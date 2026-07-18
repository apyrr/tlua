currentDirectory::/home/src/workspace/projects
useCaseSensitiveFileNames::true
Input::
//// [/home/src/tslibs/TS/Lib/lib.dom.d.tlua] *new* 
interface DOMInterface { }
//// [/home/src/tslibs/TS/Lib/lib.webworker.d.tlua] *new* 
interface WebWorkerInterface { }
//// [/home/src/workspace/projects/node_modules/@typescript/unlreated/index.d.tlua] *new* 
declare unrelated: number;
//// [/home/src/workspace/projects/project1/core.d.tlua] *new* 
declare core: number;
//// [/home/src/workspace/projects/project1/file.tlua] *new* 
local file = 10;
//// [/home/src/workspace/projects/project1/file2.tlua] *new* 
/// <reference lib="webworker"/>
/// <reference lib="es5"/>
//// [/home/src/workspace/projects/project1/index.tlua] *new* 
local x = "type1";
//// [/home/src/workspace/projects/project1/tsconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "typeRoots": ["./typeroot1"],
        "lib": ["es5", "dom"],
        "traceResolution": true,
        "libReplacement": false
    }
}
//// [/home/src/workspace/projects/project1/typeroot1/sometype/index.d.tlua] *new* 
type TheNum = "type1";
//// [/home/src/workspace/projects/project1/utils.d.tlua] *new* 
declare y: number;
//// [/home/src/workspace/projects/project2/index.tlua] *new* 
local y = 10
//// [/home/src/workspace/projects/project2/tsconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "lib": ["es5", "dom"],
        "traceResolution": true,
        "libReplacement": false
    }
}
//// [/home/src/workspace/projects/project2/utils.d.tlua] *new* 
declare y: number;
//// [/home/src/workspace/projects/project3/index.tlua] *new* 
local z = 10
//// [/home/src/workspace/projects/project3/tsconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "lib": ["es5", "dom"],
        "traceResolution": true,
        "libReplacement": false
    }
}
//// [/home/src/workspace/projects/project3/utils.d.tlua] *new* 
declare y: number;
//// [/home/src/workspace/projects/project4/index.tlua] *new* 
local z = 10
//// [/home/src/workspace/projects/project4/tsconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "lib": ["esnext", "dom", "webworker"],
        "traceResolution": true,
        "libReplacement": false
    }
}
//// [/home/src/workspace/projects/project4/utils.d.tlua] *new* 
declare y: number;

tlua -b project1 project2 project3 project4 --verbose --explainFiles
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * project1/tsconfig.json
    * project2/tsconfig.json
    * project3/tsconfig.json
    * project4/tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'project1/tsconfig.json' is out of date because output file 'project1/tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'project1/tsconfig.json'...

[91merror[0m[90m TS2318: [0mCannot find global type 'Array'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Boolean'.
[91merror[0m[90m TS2318: [0mCannot find global type 'CallableFunction'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Function'.
[91merror[0m[90m TS2318: [0mCannot find global type 'IArguments'.
[91merror[0m[90m TS2318: [0mCannot find global type 'NewableFunction'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Number'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Object'.
[91merror[0m[90m TS2318: [0mCannot find global type 'RegExp'.
[91merror[0m[90m TS2318: [0mCannot find global type 'String'.
[96mproject1/file.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local file = 10;
[7m [0m [91m~~~~~[0m

[96mproject1/file2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m /// <reference lib="webworker"/>
[7m [0m [91m~[0m

[96mproject1/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = "type1";
[7m [0m [91m~~~~~[0m

[96mproject1/tsconfig.json[0m:[93m5[0m:[93m17[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m5[0m         "lib": ["es5", "dom"],
[7m [0m [91m                ~~~~~[0m

[96mproject1/tsconfig.json[0m:[93m5[0m:[93m24[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m5[0m         "lib": ["es5", "dom"],
[7m [0m [91m                       ~~~~~[0m

project1/core.d.tlua
   Matched by default include pattern '**/*'
project1/file.tlua
   Matched by default include pattern '**/*'
project1/file2.tlua
   Matched by default include pattern '**/*'
project1/index.tlua
   Matched by default include pattern '**/*'
project1/utils.d.tlua
   Matched by default include pattern '**/*'
project1/typeroot1/sometype/index.d.tlua
   Matched by default include pattern '**/*'
[[90mHH:MM:SS AM[0m] Project 'project2/tsconfig.json' is out of date because output file 'project2/tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'project2/tsconfig.json'...

[91merror[0m[90m TS2318: [0mCannot find global type 'Array'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Boolean'.
[91merror[0m[90m TS2318: [0mCannot find global type 'CallableFunction'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Function'.
[91merror[0m[90m TS2318: [0mCannot find global type 'IArguments'.
[91merror[0m[90m TS2318: [0mCannot find global type 'NewableFunction'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Number'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Object'.
[91merror[0m[90m TS2318: [0mCannot find global type 'RegExp'.
[91merror[0m[90m TS2318: [0mCannot find global type 'String'.
[96mproject2/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local y = 10
[7m [0m [91m~~~~~[0m

[96mproject2/tsconfig.json[0m:[93m4[0m:[93m17[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m         "lib": ["es5", "dom"],
[7m [0m [91m                ~~~~~[0m

[96mproject2/tsconfig.json[0m:[93m4[0m:[93m24[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m         "lib": ["es5", "dom"],
[7m [0m [91m                       ~~~~~[0m

project2/index.tlua
   Matched by default include pattern '**/*'
project2/utils.d.tlua
   Matched by default include pattern '**/*'
[[90mHH:MM:SS AM[0m] Project 'project3/tsconfig.json' is out of date because output file 'project3/tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'project3/tsconfig.json'...

[91merror[0m[90m TS2318: [0mCannot find global type 'Array'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Boolean'.
[91merror[0m[90m TS2318: [0mCannot find global type 'CallableFunction'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Function'.
[91merror[0m[90m TS2318: [0mCannot find global type 'IArguments'.
[91merror[0m[90m TS2318: [0mCannot find global type 'NewableFunction'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Number'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Object'.
[91merror[0m[90m TS2318: [0mCannot find global type 'RegExp'.
[91merror[0m[90m TS2318: [0mCannot find global type 'String'.
[96mproject3/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local z = 10
[7m [0m [91m~~~~~[0m

[96mproject3/tsconfig.json[0m:[93m4[0m:[93m17[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m         "lib": ["es5", "dom"],
[7m [0m [91m                ~~~~~[0m

[96mproject3/tsconfig.json[0m:[93m4[0m:[93m24[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m         "lib": ["es5", "dom"],
[7m [0m [91m                       ~~~~~[0m

project3/index.tlua
   Matched by default include pattern '**/*'
project3/utils.d.tlua
   Matched by default include pattern '**/*'
[[90mHH:MM:SS AM[0m] Project 'project4/tsconfig.json' is out of date because output file 'project4/tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'project4/tsconfig.json'...

[91merror[0m[90m TS2318: [0mCannot find global type 'Array'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Boolean'.
[91merror[0m[90m TS2318: [0mCannot find global type 'CallableFunction'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Function'.
[91merror[0m[90m TS2318: [0mCannot find global type 'IArguments'.
[91merror[0m[90m TS2318: [0mCannot find global type 'NewableFunction'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Number'.
[91merror[0m[90m TS2318: [0mCannot find global type 'Object'.
[91merror[0m[90m TS2318: [0mCannot find global type 'RegExp'.
[91merror[0m[90m TS2318: [0mCannot find global type 'String'.
[96mproject4/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local z = 10
[7m [0m [91m~~~~~[0m

[96mproject4/tsconfig.json[0m:[93m4[0m:[93m17[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m         "lib": ["esnext", "dom", "webworker"],
[7m [0m [91m                ~~~~~~~~[0m

[96mproject4/tsconfig.json[0m:[93m4[0m:[93m27[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m         "lib": ["esnext", "dom", "webworker"],
[7m [0m [91m                          ~~~~~[0m

[96mproject4/tsconfig.json[0m:[93m4[0m:[93m34[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m         "lib": ["esnext", "dom", "webworker"],
[7m [0m [91m                                 ~~~~~~~~~~~[0m

project4/index.tlua
   Matched by default include pattern '**/*'
project4/utils.d.tlua
   Matched by default include pattern '**/*'

Found 55 errors in 10 files.

Errors  Files
     1  project1/file.tlua[90m:1[0m
     1  project1/file2.tlua[90m:1[0m
     1  project1/index.tlua[90m:1[0m
     2  project1/tsconfig.json[90m:5[0m
     1  project2/index.tlua[90m:1[0m
     2  project2/tsconfig.json[90m:4[0m
     1  project3/index.tlua[90m:1[0m
     2  project3/tsconfig.json[90m:4[0m
     1  project4/index.tlua[90m:1[0m
     3  project4/tsconfig.json[90m:4[0m

//// [/home/src/workspace/projects/project1/file.lua] *new* 
local file = 10;

//// [/home/src/workspace/projects/project1/file2.lua] *new* 
-- / <reference lib="webworker"/>
-- / <reference lib="es5"/>

//// [/home/src/workspace/projects/project1/index.lua] *new* 
local x = "type1";

//// [/home/src/workspace/projects/project1/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[[1,6]],"fileNames":["./core.d.tlua","./file.tlua","./file2.tlua","./index.tlua","./utils.d.tlua","./typeroot1/sometype/index.d.tlua"],"fileInfos":[{"version":"04f379066c1bb74390b1066587d91241-declare core: number;","affectsGlobalScope":true,"impliedNodeFormat":1},"2c71e802ca277d650ab0cd4cb0d56c21-local file = 10;","8eae5f7ae94ffd8697370957483d1f51-/// <reference lib=\"webworker\"/>\n/// <reference lib=\"es5\"/>","a07894186a824f107b104601266cd424-local x = \"type1\";",{"version":"f6fcbf475027f205d43c7595f33c0601-declare y: number;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"b2f9eca773dbd225c9d8bc83d036316e-type TheNum = \"type1\";","affectsGlobalScope":true,"impliedNodeFormat":1}],"options":{"composite":true},"semanticDiagnosticsPerFile":[1,2,3,4,5,6],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4]}
//// [/home/src/workspace/projects/project1/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./core.d.tlua",
        "./file.tlua",
        "./file2.tlua",
        "./index.tlua",
        "./utils.d.tlua",
        "./typeroot1/sometype/index.d.tlua"
      ],
      "original": [
        1,
        6
      ]
    }
  ],
  "fileNames": [
    "./core.d.tlua",
    "./file.tlua",
    "./file2.tlua",
    "./index.tlua",
    "./utils.d.tlua",
    "./typeroot1/sometype/index.d.tlua"
  ],
  "fileInfos": [
    {
      "fileName": "./core.d.tlua",
      "version": "04f379066c1bb74390b1066587d91241-declare core: number;",
      "signature": "04f379066c1bb74390b1066587d91241-declare core: number;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "04f379066c1bb74390b1066587d91241-declare core: number;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./file.tlua",
      "version": "2c71e802ca277d650ab0cd4cb0d56c21-local file = 10;",
      "signature": "2c71e802ca277d650ab0cd4cb0d56c21-local file = 10;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./file2.tlua",
      "version": "8eae5f7ae94ffd8697370957483d1f51-/// <reference lib=\"webworker\"/>\n/// <reference lib=\"es5\"/>",
      "signature": "8eae5f7ae94ffd8697370957483d1f51-/// <reference lib=\"webworker\"/>\n/// <reference lib=\"es5\"/>",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./index.tlua",
      "version": "a07894186a824f107b104601266cd424-local x = \"type1\";",
      "signature": "a07894186a824f107b104601266cd424-local x = \"type1\";",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./utils.d.tlua",
      "version": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
      "signature": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./typeroot1/sometype/index.d.tlua",
      "version": "b2f9eca773dbd225c9d8bc83d036316e-type TheNum = \"type1\";",
      "signature": "b2f9eca773dbd225c9d8bc83d036316e-type TheNum = \"type1\";",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "b2f9eca773dbd225c9d8bc83d036316e-type TheNum = \"type1\";",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    }
  ],
  "options": {
    "composite": true
  },
  "semanticDiagnosticsPerFile": [
    "./core.d.tlua",
    "./file.tlua",
    "./file2.tlua",
    "./index.tlua",
    "./utils.d.tlua",
    "./typeroot1/sometype/index.d.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./file.tlua",
      [
        {
          "end": 5,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./file2.tlua",
      [
        {
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./index.tlua",
      [
        {
          "end": 5,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./file.tlua",
      "original": 2
    },
    {
      "file": "./file2.tlua",
      "original": 3
    },
    {
      "file": "./index.tlua",
      "original": 4
    }
  ],
  "size": 1218
}
//// [/home/src/workspace/projects/project2/index.lua] *new* 
local y = 10;

//// [/home/src/workspace/projects/project2/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[[1,2]],"fileNames":["./index.tlua","./utils.d.tlua"],"fileInfos":["1079f0bf5852c8c62312e52b8128dfae-local y = 10",{"version":"f6fcbf475027f205d43c7595f33c0601-declare y: number;","affectsGlobalScope":true,"impliedNodeFormat":1}],"options":{"composite":true},"semanticDiagnosticsPerFile":[1,2],"emitDiagnosticsPerFile":[[1,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[1]}
//// [/home/src/workspace/projects/project2/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./index.tlua",
        "./utils.d.tlua"
      ],
      "original": [
        1,
        2
      ]
    }
  ],
  "fileNames": [
    "./index.tlua",
    "./utils.d.tlua"
  ],
  "fileInfos": [
    {
      "fileName": "./index.tlua",
      "version": "1079f0bf5852c8c62312e52b8128dfae-local y = 10",
      "signature": "1079f0bf5852c8c62312e52b8128dfae-local y = 10",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./utils.d.tlua",
      "version": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
      "signature": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    }
  ],
  "options": {
    "composite": true
  },
  "semanticDiagnosticsPerFile": [
    "./index.tlua",
    "./utils.d.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./index.tlua",
      [
        {
          "end": 5,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./index.tlua",
      "original": 1
    }
  ],
  "size": 496
}
//// [/home/src/workspace/projects/project3/index.lua] *new* 
local z = 10;

//// [/home/src/workspace/projects/project3/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[[1,2]],"fileNames":["./index.tlua","./utils.d.tlua"],"fileInfos":["e62fa1c528c515ce549fcb84013e1ae4-local z = 10",{"version":"f6fcbf475027f205d43c7595f33c0601-declare y: number;","affectsGlobalScope":true,"impliedNodeFormat":1}],"options":{"composite":true},"semanticDiagnosticsPerFile":[1,2],"emitDiagnosticsPerFile":[[1,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[1]}
//// [/home/src/workspace/projects/project3/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./index.tlua",
        "./utils.d.tlua"
      ],
      "original": [
        1,
        2
      ]
    }
  ],
  "fileNames": [
    "./index.tlua",
    "./utils.d.tlua"
  ],
  "fileInfos": [
    {
      "fileName": "./index.tlua",
      "version": "e62fa1c528c515ce549fcb84013e1ae4-local z = 10",
      "signature": "e62fa1c528c515ce549fcb84013e1ae4-local z = 10",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./utils.d.tlua",
      "version": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
      "signature": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    }
  ],
  "options": {
    "composite": true
  },
  "semanticDiagnosticsPerFile": [
    "./index.tlua",
    "./utils.d.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./index.tlua",
      [
        {
          "end": 5,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./index.tlua",
      "original": 1
    }
  ],
  "size": 496
}
//// [/home/src/workspace/projects/project4/index.lua] *new* 
local z = 10;

//// [/home/src/workspace/projects/project4/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[[1,2]],"fileNames":["./index.tlua","./utils.d.tlua"],"fileInfos":["e62fa1c528c515ce549fcb84013e1ae4-local z = 10",{"version":"f6fcbf475027f205d43c7595f33c0601-declare y: number;","affectsGlobalScope":true,"impliedNodeFormat":1}],"options":{"composite":true},"semanticDiagnosticsPerFile":[1,2],"emitDiagnosticsPerFile":[[1,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[1]}
//// [/home/src/workspace/projects/project4/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./index.tlua",
        "./utils.d.tlua"
      ],
      "original": [
        1,
        2
      ]
    }
  ],
  "fileNames": [
    "./index.tlua",
    "./utils.d.tlua"
  ],
  "fileInfos": [
    {
      "fileName": "./index.tlua",
      "version": "e62fa1c528c515ce549fcb84013e1ae4-local z = 10",
      "signature": "e62fa1c528c515ce549fcb84013e1ae4-local z = 10",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./utils.d.tlua",
      "version": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
      "signature": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    }
  ],
  "options": {
    "composite": true
  },
  "semanticDiagnosticsPerFile": [
    "./index.tlua",
    "./utils.d.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./index.tlua",
      [
        {
          "end": 5,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./index.tlua",
      "original": 1
    }
  ],
  "size": 496
}

project1/tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/workspace/projects/project1/core.d.tlua
*not cached* /home/src/workspace/projects/project1/file.tlua
*not cached* /home/src/workspace/projects/project1/file2.tlua
*not cached* /home/src/workspace/projects/project1/index.tlua
*not cached* /home/src/workspace/projects/project1/utils.d.tlua
*not cached* /home/src/workspace/projects/project1/typeroot1/sometype/index.d.tlua
Signatures::

project2/tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/workspace/projects/project2/index.tlua
*not cached* /home/src/workspace/projects/project2/utils.d.tlua
Signatures::

project3/tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/workspace/projects/project3/index.tlua
*not cached* /home/src/workspace/projects/project3/utils.d.tlua
Signatures::

project4/tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/workspace/projects/project4/index.tlua
*not cached* /home/src/workspace/projects/project4/utils.d.tlua
Signatures::
