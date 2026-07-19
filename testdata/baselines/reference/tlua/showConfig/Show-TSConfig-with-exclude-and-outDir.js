currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/src/bin/tool.tlua] *new* 
local b = 2;
//// [/home/src/workspaces/project/src/index.tlua] *new* 
local a = 1;
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "strict": true,
        "outDir": "./build"
    },
    "exclude": [
        "build"
    ]
}

tlua -p tluaconfig.json --showConfig
ExitStatus:: Success
Output::
{
    "compilerOptions": {
        "outDir": "./build",
        "strict": true
    },
    "files": [
        "./src/index.tlua",
        "./src/bin/tool.tlua"
    ],
    "exclude": [
        "build"
    ]
}
