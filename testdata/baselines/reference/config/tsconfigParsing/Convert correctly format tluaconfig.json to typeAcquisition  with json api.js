Fs::
//// [/apath/a.tlua]


//// [/apath/b.tlua]


//// [/apath/tluaconfig.json]
{
	"typeAcquisition": {
		"enable": true,
		"include": ["0.d.tlua", "1.d.tlua"],
		"exclude": ["0.lua", "1.lua"],
	},
}


configFileName:: tluaconfig.json
CompilerOptions::
{
  "configFilePath": "/apath/tluaconfig.json"
}

TypeAcquisition::
{
  "enable": true,
  "include": [
    "0.d.tlua",
    "1.d.tlua"
  ],
  "exclude": [
    "0.lua",
    "1.lua"
  ]
}

FileNames::
/apath/a.tlua,/apath/b.tlua
Errors::

