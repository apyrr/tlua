currentDirectory::/home/src/projects/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/projects/project/base.json] *new* 
{
    "include": [1],
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
error TS18003: No inputs were found in config file '/home/src/projects/project/tsconfig.json'. Specified 'include' paths were '[1]' and 'exclude' paths were '[]'.
base.json(2,17): error TS5024: Compiler option 'include' requires a value of type string.

