currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/src/index.tlua] *new* 
local a = 1;
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "target": "es5",
        "module": "commonjs",
        "strict": true
    },
    "compileOnSave": true,
    "exclude": [
        "dist"
    ],
    "files": [],
    "include": [
        "src/*"
    ],
    "references": [
        { "path": "./test" }
    ]
}

tlua -p tluaconfig.json --showConfig
ExitStatus:: Success
Output::
{
    "compilerOptions": {
        "module": "commonjs",
        "strict": true,
        "target": "es5"
    },
    "references": [
        {
            "path": "./test"
        }
    ],
    "files": [
        "./src/index.tlua"
    ],
    "include": [
        "src/*"
    ],
    "exclude": [
        "dist"
    ]
}
