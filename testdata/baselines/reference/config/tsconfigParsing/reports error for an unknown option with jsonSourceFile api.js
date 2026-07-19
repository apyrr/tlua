Fs::
//// [/app.tlua]


//// [/tluaconfig.json]
{
			    "compilerOptions": {
				"unknown": true
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
[96mtluaconfig.json[0m:[93m3[0m:[93m5[0m - [91merror[0m[90m TLUA5023: [0mUnknown compiler option 'unknown'.

[7m3[0m     "unknown": true
[7m [0m [91m    ~~~~~~~~~[0m

