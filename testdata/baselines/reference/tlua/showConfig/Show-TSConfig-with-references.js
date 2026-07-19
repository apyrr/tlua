currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/src/index.tlua] *new* 
local a = 1;
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "strict": true
    },
    "references": [
        { "path": "./packages/a" },
        { "path": "./packages/b" }
    ]
}

tlua -p tluaconfig.json --showConfig
ExitStatus:: Success
Output::
{
    "compilerOptions": {
        "composite": true,
        "strict": true,
        "declaration": true,
        "incremental": true
    },
    "references": [
        {
            "path": "./packages/a"
        },
        {
            "path": "./packages/b"
        }
    ],
    "files": [
        "./src/index.tlua"
    ]
}
