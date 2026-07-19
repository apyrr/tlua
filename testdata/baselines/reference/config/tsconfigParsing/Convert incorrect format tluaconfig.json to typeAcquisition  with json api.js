Fs::
//// [/apath/a.tlua]


//// [/apath/b.tlua]


//// [/apath/tluaconfig.json]
{
	"typeAcquisition": {
		"enableAutoDiscovy": true,
	}
}


configFileName:: tluaconfig.json
CompilerOptions::
{
  "configFilePath": "/apath/tluaconfig.json"
}

TypeAcquisition::
{}

FileNames::
/apath/a.tlua,/apath/b.tlua
Errors::
[91merror[0m[90m TLUA17010: [0mUnknown type acquisition option 'enableAutoDiscovy'.
