Fs::
//// [/apath/a.tlua]


//// [/apath/b.tlua]


//// [/apath/jsconfig.json]
{
	"typeAcquisition": {
		"enable": false,
		"include": ["0.d.tlua"],
		"exclude": ["0.lua"],
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
  "enable": false,
  "include": [
    "0.d.tlua"
  ],
  "exclude": [
    "0.lua"
  ]
}

FileNames::
/apath/a.tlua,/apath/b.tlua
Errors::

