Fs::
//// [/dist/output.lua]


//// [/node_modules/module.tlua]


//// [/src/app.tlua]


//// [/src/index.tlua]


//// [/tsconfig.json]
{
				"extends": "./tsconfigWithExtends.json",
				"compilerOptions": {
				    "outDir": "./dist",
    				"strict": true,
    				"noImplicitAny": true,
				},
			}

//// [/tsconfigWithExtends.json]
{
  "files": ["/src/index.tlua", "/src/app.tlua"],
  "include": ["/src/**/*"],
  "exclude": [],
  "ts-node": {
    "compilerOptions": {
      "module": "commonjs"
    },
    "transpileOnly": true
  }
}


configFileName:: tsconfig.json
CompilerOptions::
{
  "noImplicitAny": true,
  "outDir": "/dist",
  "strict": true,
  "configFilePath": "/tsconfig.json"
}

TypeAcquisition::
{}

FileNames::
/src/index.tlua,/src/app.tlua
Errors::

