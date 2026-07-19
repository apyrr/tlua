currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/src/index.tlua] *new* 
local a = 1;
//// [/home/src/workspaces/project/test/test1.tlua] *new* 
local src = require("src.index");
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "strict": true
    },
    "exclude": [
        "test"
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
        "./src/index.tlua"
    ],
    "exclude": [
        "test"
    ]
}
