currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/c.tlua] *new* 
local c = 10;
//// [/home/src/workspaces/project/src/a.tlua] *new* 
local a = 10;
//// [/home/src/workspaces/project/src/b.tlua] *new* 
local b = 10;
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "include": ["src"],
}

tlua src/a.tlua
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[91merror[0m[90m TLUA5112: [0mtluaconfig.json is present but will not be loaded if files are specified on commandline. Use '--ignoreConfig' to skip this error.

