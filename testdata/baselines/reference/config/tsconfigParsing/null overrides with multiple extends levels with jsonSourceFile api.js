Fs::
//// [/app.tlua]


//// [/tluaconfig-base.json]
{
  "compilerOptions": {
    "types": ["node"],
    "lib": ["es2020"],
    "outDir": "./dist",
    "strict": true
  }
}

//// [/tluaconfig-middle.json]
{
  "extends": "./tluaconfig-base.json",
  "compilerOptions": {
    "types": ["jest"],
    "outDir": "./build"
  }
}

//// [/tluaconfig.json]
{
  "extends": "./tluaconfig-middle.json",
  "compilerOptions": {
    "types": null,
    "lib": null
  }
}


configFileName:: tluaconfig.json
CompilerOptions::
{
  "outDir": "/build",
  "strict": true,
  "configFilePath": "/tluaconfig.json"
}

TypeAcquisition::
{}

FileNames::
/app.tlua
Errors::
[96mtluaconfig-base.json[0m:[93m4[0m:[93m13[0m - [91merror[0m[90m TLUA6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m     "lib": ["es2020"],
[7m [0m [91m            ~~~~~~~~[0m

