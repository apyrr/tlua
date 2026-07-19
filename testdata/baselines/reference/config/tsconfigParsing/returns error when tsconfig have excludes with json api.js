Fs::
//// [/apath/foge.tlua]


//// [/apath/test.tlua]


//// [/apath/tluaconfig.json]
{
                    "compilerOptions": {
                        "lib": ["es5"]
                    },
                    "excludes": [
                        "foge.tlua"
                    ]
                }


configFileName:: tluaconfig.json
CompilerOptions::
{
  "lib": [],
  "configFilePath": "/apath/tluaconfig.json"
}

TypeAcquisition::
{}

FileNames::
/apath/foge.tlua,/apath/test.tlua
Errors::
[91merror[0m[90m TLUA6114: [0mUnknown option 'excludes'. Did you mean 'exclude'?
[91merror[0m[90m TLUA6046: [0mArgument for '--lib' option must be: 'luajit'.
