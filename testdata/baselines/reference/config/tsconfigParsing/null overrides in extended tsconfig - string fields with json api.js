Fs::
//// [/app.tlua]


//// [/tluaconfig-base.json]
{
  "compilerOptions": {
    "outDir": "./dist",
    "rootDir": "./src"
  }
}

//// [/tluaconfig.json]
{
  "extends": "./tluaconfig-base.json",
  "compilerOptions": {
    "outDir": null,
    "rootDir": null
  }
}


configFileName:: tluaconfig.json
CompilerOptions::
{
  "configFilePath": "/tluaconfig.json"
}

TypeAcquisition::
{}

FileNames::
/app.tlua
Errors::

