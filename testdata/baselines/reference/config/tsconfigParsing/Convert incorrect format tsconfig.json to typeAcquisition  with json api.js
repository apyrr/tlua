Fs::
//// [/apath/a.tlua]


//// [/apath/b.tlua]


//// [/apath/tsconfig.json]
{
	"typeAcquisition": {
		"enableAutoDiscovy": true,
	}
}


configFileName:: tsconfig.json
CompilerOptions::
{
  "configFilePath": "/apath/tsconfig.json"
}

TypeAcquisition::
{}

FileNames::
/apath/a.tlua,/apath/b.tlua
Errors::
[91merror[0m[90m TS17010: [0mUnknown type acquisition option 'enableAutoDiscovy'.
