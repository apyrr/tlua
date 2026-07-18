currentDirectory::/home/src/projects/myproject
useCaseSensitiveFileNames::true
Input::
//// [/home/src/projects/configs/first/tsconfig.json] *new* 
{
    "extends": "../second/tsconfig.json",
    "include": ["${configDir}/src"],
    "compilerOptions": {
        "typeRoots": ["root1", "${configDir}/root2", "root3"],
        "types": [],
    },
}
//// [/home/src/projects/configs/second/tsconfig.json] *new* 
{
    "files": ["${configDir}/main.tlua"],
    "compilerOptions": {
        "declarationDir": "${configDir}/decls",
    },
    "watchOptions": {
        "excludeFiles": ["${configDir}/main.tlua"],
    },
}
//// [/home/src/projects/myproject/main.tlua] *new* 
// some comment
local y = 10;
//// [/home/src/projects/myproject/tsconfig.json] *new* 
{
    "extends": "../configs/first/tsconfig.json",
    "compilerOptions": {
        "declaration": true,
        "outDir": "outDir",
        "traceResolution": true,
    },
}

tlua --showConfig
ExitStatus:: Success
Output::
{
    "compilerOptions": {
        "declaration": true,
        "declarationDir": "./decls",
        "outDir": "./outDir",
        "traceResolution": true,
        "typeRoots": [
            "../configs/first/root1",
            "./root2",
            "../configs/first/root3"
        ],
        "types": []
    },
    "files": [
        "./main.tlua"
    ],
    "include": [
        "/home/src/projects/myproject/src"
    ],
    "exclude": [
        "/home/src/projects/myproject/outDir",
        "/home/src/projects/myproject/decls"
    ]
}
