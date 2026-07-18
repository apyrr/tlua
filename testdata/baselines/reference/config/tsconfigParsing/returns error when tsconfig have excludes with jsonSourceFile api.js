Fs::
//// [/apath/foge.tlua]


//// [/apath/test.tlua]


//// [/apath/tsconfig.json]
{
                    "compilerOptions": {
                        "lib": ["es5"]
                    },
                    "excludes": [
                        "foge.tlua"
                    ]
                }


configFileName:: tsconfig.json
CompilerOptions::
{
  "lib": [],
  "configFilePath": "/apath/tsconfig.json"
}

TypeAcquisition::
{}

FileNames::
/apath/foge.tlua,/apath/test.tlua
Errors::
[96mtsconfig.json[0m:[93m3[0m:[93m33[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m3[0m                         "lib": ["es5"]
[7m [0m [91m                                ~~~~~[0m

[96mtsconfig.json[0m:[93m5[0m:[93m21[0m - [91merror[0m[90m TS6114: [0mUnknown option 'excludes'. Did you mean 'exclude'?

[7m5[0m                     "excludes": [
[7m [0m [91m                    ~~~~~~~~~~[0m

