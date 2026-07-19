Fs::
//// [/apath/a.tlua]


//// [/apath/b.tlua]


//// [/apath/tluaconfig.json]
{
	"typeAcquisition": {
		"enable": true,
	},
}


configFileName:: tluaconfig.json
CompilerOptions::
{
  "configFilePath": "/apath/tluaconfig.json"
}

TypeAcquisition::
{
  "enable": true
}

FileNames::
/apath/a.tlua,/apath/b.tlua
Errors::

