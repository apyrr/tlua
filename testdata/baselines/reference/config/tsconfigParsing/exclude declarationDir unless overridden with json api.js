Fs::
//// [/a.tlua]


//// [/declarations/a.d.tlua]


//// [/tluaconfig.json]
{
                "compilerOptions": {
                    "declarationDir": "declarations"
                }
            }


configFileName:: tluaconfig.json
CompilerOptions::
{
  "declarationDir": "/declarations",
  "configFilePath": "/tluaconfig.json"
}

TypeAcquisition::
{}

FileNames::
/a.tlua
Errors::


Fs::
//// [/a.tlua]


//// [/declarations/a.d.tlua]


//// [/tluaconfig.json]
{
                "compilerOptions": {
                    "declarationDir": "declarations"
                },
                "exclude": [ "types" ]
            }


configFileName:: tluaconfig.json
CompilerOptions::
{
  "declarationDir": "/declarations",
  "configFilePath": "/tluaconfig.json"
}

TypeAcquisition::
{}

FileNames::
/a.tlua,/declarations/a.d.tlua
Errors::

