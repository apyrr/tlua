currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/extra.tlua] *new* 
local c = 3;
//// [/home/src/workspaces/project/src/main.tlua] *new* 
local a = 1;
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "strict": true
    },
    "files": [
        "extra.tlua"
    ],
    "include": [
        "src/**/*"
    ]
}

tlua -p tluaconfig.json --showConfig
ExitStatus:: Success
Output::
{
    "compilerOptions": {
        "strict": true
    },
    "files": [
        "./extra.tlua",
        "./src/main.tlua"
    ],
    "include": [
        "src/**/*"
    ]
}
