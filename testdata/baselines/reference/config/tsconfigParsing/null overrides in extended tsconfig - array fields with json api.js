Fs::
//// [/app.tlua]


//// [/tsconfig-base.json]
{
  "compilerOptions": {
    "types": ["node", "@types/jest"],
    "lib": ["es2020", "dom"],
    "typeRoots": ["./types", "./node_modules/@types"]
  }
}

//// [/tsconfig.json]
{
  "extends": "./tsconfig-base.json",
  "compilerOptions": {
    "types": null,
    "lib": null,
    "typeRoots": null
  }
}


configFileName:: tsconfig.json
CompilerOptions::
{
  "configFilePath": "/tsconfig.json"
}

TypeAcquisition::
{}

FileNames::
/app.tlua
Errors::
[96mtsconfig-base.json[0m:[93m4[0m:[93m13[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m     "lib": ["es2020", "dom"],
[7m [0m [91m            ~~~~~~~~[0m

[96mtsconfig-base.json[0m:[93m4[0m:[93m23[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m     "lib": ["es2020", "dom"],
[7m [0m [91m                      ~~~~~[0m

