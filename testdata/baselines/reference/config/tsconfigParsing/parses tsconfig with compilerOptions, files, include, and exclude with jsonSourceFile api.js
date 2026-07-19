Fs::
//// [/apath/dist/output.lua]


//// [/apath/node_modules/module.tlua]


//// [/apath/src/app.tlua]


//// [/apath/src/index.tlua]


//// [/apath/tluaconfig.json]
{
  "compilerOptions": {
    "outDir": "./dist",
    "strict": true,
    "noImplicitAny": true,
    "target": "ES2017",
    "module": "ESNext",
    "jsx": "react",
	"maxNodeModuleJsDepth": 1
  },
  "files": ["/apath/src/index.tlua", "/apath/src/app.tlua"],
  "include": ["/apath/src/**/*"],
  "exclude": ["/apath/node_modules", "/apath/dist"]
}


configFileName:: /apath/tluaconfig.json
CompilerOptions::
{
  "jsx": 3,
  "module": 99,
  "noImplicitAny": true,
  "outDir": "/apath/dist",
  "strict": true,
  "target": 4,
  "configFilePath": "/apath/tluaconfig.json"
}

TypeAcquisition::
{}

FileNames::
/apath/src/index.tlua,/apath/src/app.tlua
Errors::
[96mtluaconfig.json[0m:[93m9[0m:[93m2[0m - [91merror[0m[90m TLUA5023: [0mUnknown compiler option 'maxNodeModuleJsDepth'.

[7m9[0m  "maxNodeModuleJsDepth": 1
[7m [0m [91m ~~~~~~~~~~~~~~~~~~~~~~[0m

