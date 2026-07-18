currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/a.d.tlua] *new* 
declare a: "hello";
//// [/home/src/workspaces/project/b.tlua] *new* 
local b = 10;
//// [/home/src/workspaces/project/tsconfig.json] *new* 
{
    "compilerOptions": {
        "declaration": true,
        "incremental": true,
        "lib": ["es6"],
    },
}

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
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
[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mtsconfig.json[0m:[93m5[0m:[93m17[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m5[0m         "lib": ["es6"],
[7m [0m [91m                ~~~~~[0m


Found 12 errors in 2 files.

Errors  Files
     1  b.tlua[90m:1[0m
     1  tsconfig.json[90m:5[0m

//// [/home/src/workspaces/project/b.lua] *new* 
local b = 10;

//// [/home/src/workspaces/project/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[[1,2]],"fileNames":["./a.d.tlua","./b.tlua"],"fileInfos":[{"version":"3339861bee6cf69ba81681a9309e0361-declare a: \"hello\";","affectsGlobalScope":true,"impliedNodeFormat":1},"f264fbb7445fd7350e2974971e7a3290-local b = 10;"],"options":{"declaration":true},"semanticDiagnosticsPerFile":[1,2],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
//// [/home/src/workspaces/project/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.d.tlua",
        "./b.tlua"
      ],
      "original": [
        1,
        2
      ]
    }
  ],
  "fileNames": [
    "./a.d.tlua",
    "./b.tlua"
  ],
  "fileInfos": [
    {
      "fileName": "./a.d.tlua",
      "version": "3339861bee6cf69ba81681a9309e0361-declare a: \"hello\";",
      "signature": "3339861bee6cf69ba81681a9309e0361-declare a: \"hello\";",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "3339861bee6cf69ba81681a9309e0361-declare a: \"hello\";",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./b.tlua",
      "version": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "signature": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": true
  },
  "semanticDiagnosticsPerFile": [
    "./a.d.tlua",
    "./b.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./b.tlua",
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
  "size": 473
}

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/workspaces/project/a.d.tlua
*not cached* /home/src/workspaces/project/b.tlua
Signatures::


Edit [0]:: with --noLib

tlua --noLib
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
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
[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mtsconfig.json[0m:[93m5[0m:[93m9[0m - [91merror[0m[90m TS5053: [0mOption 'lib' cannot be specified with option 'noLib'.

[7m5[0m         "lib": ["es6"],
[7m [0m [91m        ~~~~~[0m

[96mtsconfig.json[0m:[93m5[0m:[93m17[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m5[0m         "lib": ["es6"],
[7m [0m [91m                ~~~~~[0m


Found 13 errors in 2 files.

Errors  Files
     1  b.tlua[90m:1[0m
     2  tsconfig.json[90m:5[0m


tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/workspaces/project/a.d.tlua
*not cached* /home/src/workspaces/project/b.tlua
Signatures::
