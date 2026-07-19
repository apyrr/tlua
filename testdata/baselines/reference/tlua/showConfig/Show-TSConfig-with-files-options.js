currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::

tlua --showConfig file0.tlua file1.tlua file2.tlua
ExitStatus:: Success
Output::
{
    "compilerOptions": {},
    "files": [
        "./file0.tlua",
        "./file1.tlua",
        "./file2.tlua"
    ]
}
