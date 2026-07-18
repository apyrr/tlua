Fs::
//// [/a.tlua]


//// [/declarations/a.d.tlua]


//// [/tsconfig.json]
{
                "compilerOptions": {
                    "declarationDir": "declarations"
                }
            }


configFileName:: tsconfig.json
CompilerOptions::
{
  "declarationDir": "/declarations",
  "configFilePath": "/tsconfig.json"
}

TypeAcquisition::
{}

FileNames::
/a.tlua
Errors::


Fs::
//// [/a.tlua]


//// [/declarations/a.d.tlua]


//// [/tsconfig.json]
{
                "compilerOptions": {
                    "declarationDir": "declarations"
                },
                "exclude": [ "types" ]
            }


configFileName:: tsconfig.json
CompilerOptions::
{
  "declarationDir": "/declarations",
  "configFilePath": "/tsconfig.json"
}

TypeAcquisition::
{}

FileNames::
/a.tlua,/declarations/a.d.tlua
Errors::

