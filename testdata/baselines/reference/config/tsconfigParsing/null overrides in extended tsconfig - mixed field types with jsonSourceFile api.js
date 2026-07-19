Fs::
//// [/app.tlua]


//// [/tluaconfig-base.json]
{
  "compilerOptions": {
    "types": ["node"],
    "lib": ["es2020", "dom"],
    "outDir": "./dist",
    "strict": true,
    "allowJs": true,
    "target": "es2020"
  }
}

//// [/tluaconfig.json]
{
  "extends": "./tluaconfig-base.json",
  "compilerOptions": {
    "types": null,
    "outDir": null,
    "strict": false,
    "lib": ["es2022"],
    "allowJs": null
  }
}


configFileName:: tluaconfig.json
CompilerOptions::
{
  "lib": [],
  "strict": false,
  "target": 7,
  "configFilePath": "/tluaconfig.json"
}

TypeAcquisition::
{}

FileNames::
/app.tlua
Errors::
[96mtluaconfig.json[0m:[93m7[0m:[93m13[0m - [91merror[0m[90m TLUA6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m7[0m     "lib": ["es2022"],
[7m [0m [91m            ~~~~~~~~[0m

[96mtluaconfig-base.json[0m:[93m4[0m:[93m13[0m - [91merror[0m[90m TLUA6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m     "lib": ["es2020", "dom"],
[7m [0m [91m            ~~~~~~~~[0m

[96mtluaconfig-base.json[0m:[93m4[0m:[93m23[0m - [91merror[0m[90m TLUA6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m     "lib": ["es2020", "dom"],
[7m [0m [91m                      ~~~~~[0m

[96mtluaconfig-base.json[0m:[93m7[0m:[93m5[0m - [91merror[0m[90m TLUA5023: [0mUnknown compiler option 'allowJs'.

[7m7[0m     "allowJs": true,
[7m [0m [91m    ~~~~~~~~~[0m

