Fs::
//// [/app.tlua]


//// [/tluaconfig-base.json]
{
  "compilerOptions": {
    "types": ["node", "@types/jest"],
    "lib": ["es2020", "dom"],
    "typeRoots": ["./types", "./node_modules/@types"]
  }
}

//// [/tluaconfig.json]
{
  "extends": "./tluaconfig-base.json",
  "compilerOptions": {
    "types": null,
    "lib": null,
    "typeRoots": null
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
[96mtluaconfig-base.json[0m:[93m4[0m:[93m13[0m - [91merror[0m[90m TLUA6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m     "lib": ["es2020", "dom"],
[7m [0m [91m            ~~~~~~~~[0m

[96mtluaconfig-base.json[0m:[93m4[0m:[93m23[0m - [91merror[0m[90m TLUA6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m     "lib": ["es2020", "dom"],
[7m [0m [91m                      ~~~~~[0m

