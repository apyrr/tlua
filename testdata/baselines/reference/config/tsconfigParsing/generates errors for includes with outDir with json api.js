Fs::
//// [/apath/a.tlua]


//// [/apath/tluaconfig.json]
{
                "compilerOptions": {
                    "outDir": "./"
                },
                "include": ["**/*"]
            }


configFileName:: /apath/tluaconfig.json
CompilerOptions::
{
  "outDir": "/apath",
  "configFilePath": "/apath/tluaconfig.json"
}

TypeAcquisition::
{}

FileNames::

Errors::
[91merror[0m[90m TLUA18003: [0mNo inputs were found in config file '/apath/tluaconfig.json'. Specified 'include' paths were '["**/*"]' and 'exclude' paths were '["/apath"]'.
