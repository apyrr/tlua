currentDirectory::/home/src/projects/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/projects/project/base.json] *new* 
{
    "files": [1],
}
//// [/home/src/projects/project/main.tlua] *new* 
local x = 1;
//// [/home/src/projects/project/tsconfig.json] *new* 
{
    "extends": "./base.json",
}

tlua -p tsconfig.json --pretty false
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
base.json(2,15): error TS5024: Compiler option 'files' requires a value of type string.

