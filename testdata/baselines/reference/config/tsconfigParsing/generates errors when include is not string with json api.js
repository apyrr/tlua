Fs::
//// [/apath/a.tlua]


//// [/apath/tluaconfig.json]
{
  "include": [
    [
      "./**/*.tlua"
    ]
  ]
}


configFileName:: /apath/tluaconfig.json
CompilerOptions::
{
  "configFilePath": "/apath/tluaconfig.json"
}

TypeAcquisition::
{}

FileNames::

Errors::
[91merror[0m[90m TLUA5024: [0mCompiler option 'include' requires a value of type string.
[91merror[0m[90m TLUA18003: [0mNo inputs were found in config file '/apath/tluaconfig.json'. Specified 'include' paths were '[["./**/*.tlua"]]' and 'exclude' paths were '[]'.
