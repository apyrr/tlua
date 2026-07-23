currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/src/index.tlua] *new* 
local a = 1;
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "outDir": "./lib",
        "module": "commonjs",
        "target": "ES2017",
        "sourceMap": true,
        "experimentalDecorators": true,
        "emitDecoratorMetadata": true
    },
    "include": [
        "./src/**/*"
    ]
}

tlua -p tluaconfig.json --showConfig
ExitStatus:: Success
Output::
{
    "compilerOptions": {
        "module": "commonjs",
        "outDir": "./lib",
        "sourceMap": true,
        "target": "es2017"
    },
    "files": [
        "./src/index.tlua"
    ],
    "include": [
        "./src/**/*"
    ],
    "exclude": [
        "/home/src/workspaces/project/lib"
    ]
}
