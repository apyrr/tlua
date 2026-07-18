currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/tsconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "noEmit": true,
    },
    "typeAcquisition": {
        "enable": true,
        "include": ["0.d.tlua", "1.d.tlua"],
        "exclude": ["0.lua", "1.lua"],
        "disableFilenameBasedTypeAcquisition": true,
    },
}

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[91merror[0m[90m TS18003: [0mNo inputs were found in config file '/home/src/workspaces/project/tsconfig.json'. Specified 'include' paths were '["**/*"]' and 'exclude' paths were '[]'.

Found 1 error.

//// [/home/src/workspaces/project/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"fileInfos":[],"options":{"composite":true}}
//// [/home/src/workspaces/project/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "fileInfos": [],
  "options": {
    "composite": true
  },
  "size": 85
}

tsconfig.json::
SemanticDiagnostics::
Signatures::
