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
[96mtluaconfig.json[0m:[93m3[0m:[93m5[0m - [91merror[0m[90m TLUA5024: [0mCompiler option 'include' requires a value of type string.

[7m3[0m     [
[7m [0m [91m    ~[0m
[7m4[0m       "./**/*.tlua"
[7m [0m [91m~~~~~~~~~~~~~~~~~~~[0m
[7m5[0m     ]
[7m [0m [91m~~~~~[0m

[91merror[0m[90m TLUA18003: [0mNo inputs were found in config file '/apath/tluaconfig.json'. Specified 'include' paths were '[["./**/*.tlua"]]' and 'exclude' paths were '[]'.
