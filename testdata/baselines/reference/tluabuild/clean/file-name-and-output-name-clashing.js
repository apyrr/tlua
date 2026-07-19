currentDirectory::/home/src/workspaces/solution
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/solution/bar.tlua] *new* 

//// [/home/src/workspaces/solution/index.lua] *new* 

//// [/home/src/workspaces/solution/tluaconfig.json] *new* 
{
    "compilerOptions": {}
}

tlua --b --clean
ExitStatus:: Success
Output::

