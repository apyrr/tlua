Fs::
//// [/b.tlua]


//// [/bin/a.tlua]


//// [/tluaconfig.json]
{
  "extends": "./tsconfigWithoutConfigDir.json"
}

//// [/tsconfigWithoutConfigDir.json]
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
  "extends": "./tsconfigWithConfigDir.json"
}

//// [/tsconfigWithConfigDir.json]
{
  "compilerOptions": {
    "outDir": "${configDir}/bin"
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

