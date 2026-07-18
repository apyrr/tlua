currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/extra.tlua] *new* 
local c = 3;
//// [/home/src/workspaces/project/src/main.tlua] *new* 
local a = 1;
//// [/home/src/workspaces/project/src/util.tlua] *new* 
local b = 2;
//// [/home/src/workspaces/project/tsconfig.json] *new* 
{
    "compilerOptions": {
        "strict": true
    },
    "include": [
        "src/**/*"
    ]
}

tlua -p tsconfig.json --showConfig
ExitStatus:: Success
Output::
{
    "compilerOptions": {
        "strict": true
    },
    "files": [
        "./src/main.tlua",
        "./src/util.tlua"
    ],
    "include": [
        "src/**/*"
    ]
}
