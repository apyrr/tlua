Fs::
//// [/app.tlua]


//// [/tsconfig-base.json]
{
  "compilerOptions": {
    "outDir": "./dist",
    "rootDir": "./src"
  }
}

//// [/tsconfig.json]
{
  "extends": "./tsconfig-base.json",
  "compilerOptions": {
    "outDir": null,
    "rootDir": null
  }
}


configFileName:: tsconfig.json
CompilerOptions::
{
  "configFilePath": "/tsconfig.json"
}

TypeAcquisition::
{}

FileNames::
/app.tlua
Errors::

