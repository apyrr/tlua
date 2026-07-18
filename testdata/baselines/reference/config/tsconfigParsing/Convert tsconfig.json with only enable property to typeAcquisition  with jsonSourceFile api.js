Fs::
//// [/apath/a.tlua]


//// [/apath/b.tlua]


//// [/apath/tsconfig.json]
{
	"typeAcquisition": {
		"enable": true,
	},
}


configFileName:: tsconfig.json
CompilerOptions::
{
  "configFilePath": "/apath/tsconfig.json"
}

TypeAcquisition::
{
  "enable": true
}

FileNames::
/apath/a.tlua,/apath/b.tlua
Errors::

