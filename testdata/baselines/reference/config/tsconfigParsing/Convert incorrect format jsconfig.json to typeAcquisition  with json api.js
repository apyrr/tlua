Fs::
//// [/apath/a.tlua]


//// [/apath/b.tlua]


//// [/apath/jsconfig.json]
{
	"typeAcquisition": {
		"enableAutoDiscovy": true,
	},
}


configFileName:: jsconfig.json
CompilerOptions::
{
  "noEmit": true,
  "skipLibCheck": true,
  "configFilePath": "/apath/jsconfig.json"
}

TypeAcquisition::
{
  "enable": true
}

FileNames::
/apath/a.tlua,/apath/b.tlua
Errors::
[91merror[0m[90m TLUA17010: [0mUnknown type acquisition option 'enableAutoDiscovy'.
