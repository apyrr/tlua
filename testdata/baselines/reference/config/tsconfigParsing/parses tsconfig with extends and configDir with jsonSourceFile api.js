Fs::
//// [/dist/output.lua]


//// [/node_modules/module.tlua]


//// [/src/app.tlua]


//// [/src/index.tlua]


//// [/tsconfig.base.json]
{
  "compilerOptions": {
    "outFile": "${configDir}/outFile",
    "outDir": "${configDir}/outDir",
    "rootDir": "${configDir}/rootDir",
    "tsBuildInfoFile": "${configDir}/tsBuildInfoFile",
    "declarationDir": "${configDir}/declarationDir",
  }
}

//// [/tsconfig.json]
{
				"extends": "./tsconfig.base.json"
			}


configFileName:: tsconfig.json
CompilerOptions::
{
  "declarationDir": "/declarationDir",
  "outDir": "/outDir",
  "rootDir": "/rootDir",
  "tsBuildInfoFile": "/tsBuildInfoFile",
  "outFile": "/outFile",
  "configFilePath": "/tsconfig.json"
}

TypeAcquisition::
{}

FileNames::
/src/app.tlua,/src/index.tlua
Errors::

