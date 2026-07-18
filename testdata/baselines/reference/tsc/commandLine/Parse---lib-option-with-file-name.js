currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/first.tlua] *new* 
local Key = Symbol()

tlua --lib es6  first.tlua
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

