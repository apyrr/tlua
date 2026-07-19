currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/tluaconfig.first.json] *new* 
{
    "extends": "./foobar.json",
    "compilerOptions": {
        "composite": true
    }
}
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true
    },
    "references": [
        { "path": "./tluaconfig.first.json" },
        { "path": "./tluaconfig.second.json" }
    ]
}
//// [/home/src/workspaces/project/tluaconfig.second.json] *new* 
{
    "extends": "./foobar.json",
    "compilerOptions": {
        "composite": true
    }
}

tlua --b
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[91merror[0m[90m TLUA5083: [0mCannot read file '/home/src/workspaces/project/foobar.json'.
[91merror[0m[90m TLUA18003: [0mNo inputs were found in config file '/home/src/workspaces/project/tluaconfig.first.json'. Specified 'include' paths were '["**/*"]' and 'exclude' paths were '[]'.
[91merror[0m[90m TLUA5083: [0mCannot read file '/home/src/workspaces/project/foobar.json'.
[91merror[0m[90m TLUA18003: [0mNo inputs were found in config file '/home/src/workspaces/project/tluaconfig.second.json'. Specified 'include' paths were '["**/*"]' and 'exclude' paths were '[]'.

Found 4 errors.

//// [/home/src/workspaces/project/tluaconfig.first.tluabuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"fileInfos":[],"options":{"composite":true}}
//// [/home/src/workspaces/project/tluaconfig.first.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "fileInfos": [],
  "options": {
    "composite": true
  },
  "size": 85
}
//// [/home/src/workspaces/project/tluaconfig.second.tluabuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"fileInfos":[],"options":{"composite":true}}
//// [/home/src/workspaces/project/tluaconfig.second.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "fileInfos": [],
  "options": {
    "composite": true
  },
  "size": 85
}

tluaconfig.first.json::
SemanticDiagnostics::
Signatures::

tluaconfig.second.json::
SemanticDiagnostics::
Signatures::
