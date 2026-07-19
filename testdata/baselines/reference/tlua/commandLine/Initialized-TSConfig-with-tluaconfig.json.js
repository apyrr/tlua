currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/first.tlua] *new* 
local a = 1;
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "strict": true,
        "noEmit": true
    }
}

tlua --init
ExitStatus:: Success
Output::
[91merror[0m[90m TLUA5054: [0mA 'tluaconfig.json' file is already defined at: '/home/src/workspaces/project/tluaconfig.json'.

