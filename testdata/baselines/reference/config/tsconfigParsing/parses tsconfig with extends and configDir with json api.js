Fs::
//// [/dist/output.lua]


//// [/node_modules/module.tlua]


//// [/src/app.tlua]


//// [/src/index.tlua]


//// [/tluaconfig.base.json]
{
  "compilerOptions": {
    "outFile": "${configDir}/outFile",
    "outDir": "${configDir}/outDir",
    "rootDir": "${configDir}/rootDir",
    "tsBuildInfoFile": "${configDir}/tsBuildInfoFile",
    "declarationDir": "${configDir}/declarationDir",
  }
}

//// [/tluaconfig.json]
{
				"extends": "./tluaconfig.base.json"
			}


configFileName:: tluaconfig.json
CompilerOptions::
{
  "declarationDir": "/declarationDir",
  "outDir": "/outDir",
  "rootDir": "/rootDir",
  "tsBuildInfoFile": "/tsBuildInfoFile",
  "outFile": "/outFile",
  "configFilePath": "/tluaconfig.json"
}

TypeAcquisition::
{}

FileNames::
/src/app.tlua,/src/index.tlua
Errors::

