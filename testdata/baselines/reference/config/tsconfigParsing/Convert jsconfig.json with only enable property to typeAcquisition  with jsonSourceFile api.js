Fs::
//// [/apath/a.tlua]


//// [/apath/b.tlua]


//// [/apath/jsconfig.json]
{
	"typeAcquisition": {
		"enable": false,
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
  "enable": false
}

FileNames::
/apath/a.tlua,/apath/b.tlua
Errors::

