Fs::
//// [/app.tlua]


//// [/tluaconfig.json]
{
			    "compilerOptions": {
				"target": "invalid value",
				"removeComments": "should be a boolean",
				"jsx": "invalid value"
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
[91merror[0m[90m TLUA5024: [0mCompiler option 'removeComments' requires a value of type boolean.
