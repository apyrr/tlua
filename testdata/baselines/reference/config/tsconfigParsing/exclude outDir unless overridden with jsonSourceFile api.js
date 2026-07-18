Fs::
//// [/b.tlua]


//// [/bin/a.tlua]


//// [/tsconfig.json]
{
                "compilerOptions": {
                    "outDir": "bin"
                }
            }


configFileName:: tsconfig.json
CompilerOptions::
{
  "outDir": "/bin",
  "configFilePath": "/tsconfig.json"
}

TypeAcquisition::
{}

FileNames::
/b.tlua
Errors::


Fs::
//// [/b.tlua]


//// [/bin/a.tlua]


//// [/tsconfig.json]
{
                "compilerOptions": {
                    "outDir": "bin"
                },
                "exclude": [ "obj" ]
            }


configFileName:: tsconfig.json
CompilerOptions::
{
  "outDir": "/bin",
  "configFilePath": "/tsconfig.json"
}

TypeAcquisition::
{}

FileNames::
/b.tlua,/bin/a.tlua
Errors::

