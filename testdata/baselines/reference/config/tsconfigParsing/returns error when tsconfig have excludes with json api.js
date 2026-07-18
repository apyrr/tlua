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
[91merror[0m[90m TS6114: [0mUnknown option 'excludes'. Did you mean 'exclude'?
[91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.
