Fs::
//// [/b.tlua]


//// [/bin/a.tlua]


//// [/tluaconfig.json]
{
                "compilerOptions": {
                    "outDir": "bin"
                }
            }


configFileName:: tluaconfig.json
CompilerOptions::
{
  "outDir": "/bin",
  "configFilePath": "/tluaconfig.json"
}

TypeAcquisition::
{}

FileNames::
/b.tlua
Errors::


Fs::
//// [/b.tlua]


//// [/bin/a.tlua]


//// [/tluaconfig.json]
{
                "compilerOptions": {
                    "outDir": "bin"
                },
                "exclude": [ "obj" ]
            }


configFileName:: tluaconfig.json
CompilerOptions::
{
  "outDir": "/bin",
  "configFilePath": "/tluaconfig.json"
}

TypeAcquisition::
{}

FileNames::
/b.tlua,/bin/a.tlua
Errors::

