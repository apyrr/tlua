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
[96mtluaconfig.json[0m:[93m3[0m:[93m3[0m - [91merror[0m[90m TLUA17010: [0mUnknown type acquisition option 'enableAutoDiscovy'.

[7m3[0m   "enableAutoDiscovy": true,
[7m [0m [91m  ~~~~~~~~~~~~~~~~~~~[0m

