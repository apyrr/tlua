Fs::
//// [/app.tlua]


//// [/tsconfig-base.json]
{
  "compilerOptions": {
    "types": ["node"],
    "lib": ["es2020"],
    "outDir": "./base",
    "strict": true
  }
}

//// [/tsconfig-middle.json]
{
  "extends": "./tsconfig-base.json",
  "compilerOptions": {
    "types": null,
    "lib": null,
    "outDir": "./middle"
  }
}

//// [/tsconfig.json]
{
  "extends": "./tsconfig-middle.json",
  "compilerOptions": {
    "outDir": "./final"
  }
}


configFileName:: tsconfig.json
CompilerOptions::
{
  "outDir": "/final",
  "strict": true,
  "configFilePath": "/tsconfig.json"
}

TypeAcquisition::
{}

FileNames::
/app.tlua
Errors::
[96mtsconfig-base.json[0m:[93m4[0m:[93m13[0m - [91merror[0m[90m TS6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m4[0m     "lib": ["es2020"],
[7m [0m [91m            ~~~~~~~~[0m

