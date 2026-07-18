Fs::
//// [/apath/main.tlua]


//// [/apath/tsconfig.json]
{
                "include": ["**/../*.tlua"]
            }


configFileName:: /apath/tsconfig.json
CompilerOptions::
{
  "configFilePath": "/apath/tsconfig.json"
}

TypeAcquisition::
{}

FileNames::

Errors::
[96mtsconfig.json[0m:[93m2[0m:[93m29[0m - [91merror[0m[90m TS5065: [0mFile specification cannot contain a parent directory ('..') that appears after a recursive directory wildcard ('**'): '**/../*.tlua'.

[7m2[0m                 "include": ["**/../*.tlua"]
[7m [0m [91m                            ~~~~~~~~~~~~~~[0m

[91merror[0m[90m TS18003: [0mNo inputs were found in config file '/apath/tsconfig.json'. Specified 'include' paths were '["**/../*.tlua"]' and 'exclude' paths were '[]'.
