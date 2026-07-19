currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/c.tlua] *new* 
local c = 10;
//// [/home/src/workspaces/project/src/a.tlua] *new* 
local a = 10;
//// [/home/src/workspaces/project/src/b.tlua] *new* 
local b = 10;

tlua -p .
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[91merror[0m[90m TLUA5081: [0mCannot find a tluaconfig.json file at the current directory: /home/src/workspaces/project/tluaconfig.json.

