Fs::
//// [/app.tlua]


//// [/tluaconfig.json]
{
			    "compilerOptions": {
				"sourcemap": true,
				"declarationmap": true,
				"nouncheckedindexedaccess": true,
				"exactoptionalpropertytypes": true,
				"verbatimmodulesyntax": true,
				"isolatedmodules": true,
				"nouncheckedsideeffectimports": true,
				"moduledetection": "force",
				"skiplibcheck": true,
				"checkjs": true
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
[91merror[0m[90m TLUA5025: [0mUnknown compiler option 'sourcemap'. Did you mean 'sourceMap'?
[91merror[0m[90m TLUA5025: [0mUnknown compiler option 'declarationmap'. Did you mean 'declarationMap'?
[91merror[0m[90m TLUA5025: [0mUnknown compiler option 'nouncheckedindexedaccess'. Did you mean 'noUncheckedIndexedAccess'?
[91merror[0m[90m TLUA5025: [0mUnknown compiler option 'exactoptionalpropertytypes'. Did you mean 'exactOptionalPropertyTypes'?
[91merror[0m[90m TLUA5023: [0mUnknown compiler option 'verbatimmodulesyntax'.
[91merror[0m[90m TLUA5025: [0mUnknown compiler option 'isolatedmodules'. Did you mean 'isolatedModules'?
[91merror[0m[90m TLUA5025: [0mUnknown compiler option 'nouncheckedsideeffectimports'. Did you mean 'noUncheckedSideEffectImports'?
[91merror[0m[90m TLUA5023: [0mUnknown compiler option 'moduledetection'.
[91merror[0m[90m TLUA5025: [0mUnknown compiler option 'skiplibcheck'. Did you mean 'skipLibCheck'?
[91merror[0m[90m TLUA5023: [0mUnknown compiler option 'checkjs'.
