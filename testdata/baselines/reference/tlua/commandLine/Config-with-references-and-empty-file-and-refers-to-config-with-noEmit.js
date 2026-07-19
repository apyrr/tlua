currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/packages/pkg1/index.tlua] *new* 
local a = 1;
//// [/home/src/workspaces/project/packages/pkg1/tluaconfig.json] *new* 
{
                    "compilerOptions": {
                        "composite": true,
                        "noEmit": true
                    },
                    "files": [
                        "./index.tlua",
                    ],
                }
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
                    "files": [],
                    "references": [
                        {
                            "path": "./packages/pkg1"
                        },
                    ],
                }

tlua -p .
ExitStatus:: Success
Output::

