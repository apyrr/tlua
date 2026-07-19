currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/primary/a.tlua] *new* 
local primaryA = 0;
//// [/home/src/workspaces/project/primary/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": false,
        "outDir": "bin",
    }
}
//// [/home/src/workspaces/project/reference/b.tlua] *new* 
local mod_1 = require("primary.a");
//// [/home/src/workspaces/project/reference/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "outDir": "bin",
        "rootDir": "..",
    },
    "files": [ ],
    "references": [{
        "path": "../primary"
    }]
}

tlua --p reference/tluaconfig.json
ExitStatus:: Success
Output::
//// [/home/src/workspaces/project/reference/bin/reference/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","fileInfos":[],"options":{"composite":true,"outDir":"..","rootDir":"../../.."}}
//// [/home/src/workspaces/project/reference/bin/reference/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "fileInfos": [],
  "options": {
    "composite": true,
    "outDir": "..",
    "rootDir": "../../.."
  },
  "size": 106
}

reference/tluaconfig.json::
SemanticDiagnostics::
Signatures::
