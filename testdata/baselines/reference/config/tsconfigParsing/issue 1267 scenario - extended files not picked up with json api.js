Fs::
//// [/src/main.tlua]
export {}

//// [/src/utils.tlua]
export {}

//// [/tluaconfig-base/backend.json]
{
  "$schema": "https://json.schemastore.org/tsconfig",
  "display": "Backend",
  "compilerOptions": {
    "allowJs": true,
    "module": "nodenext",
    "removeComments": true,
    "emitDecoratorMetadata": true,
    "experimentalDecorators": true,
    "target": "esnext",
    "lib": ["ESNext"],
    "incremental": false,
    "noImplicitAny": true,
    "types": ["node", "vitest/globals"],
    "sourceMap": true,
    "strictPropertyInitialization": false
  },
  "files": [
    "types/ical2json.d.tlua",
    "types/express.d.tlua",
    "types/multer.d.tlua",
    "types/reset.d.tlua",
    "types/stripe-custom-typings.d.tlua",
    "types/nestjs-modules.d.tlua",
    "types/luxon.d.tlua",
    "types/nestjs-pino.d.tlua"
  ],
  "ts-node": {
    "files": true
  }
}

//// [/tluaconfig-base/types/express.d.tlua]
export {}

//// [/tluaconfig-base/types/ical2json.d.tlua]
export {}

//// [/tluaconfig-base/types/luxon.d.tlua]
declare module 'luxon' {
  interface TSSettings {
    throwOnInvalid: true
  }
}
export {}

//// [/tluaconfig-base/types/multer.d.tlua]
export {}

//// [/tluaconfig-base/types/nestjs-modules.d.tlua]
export {}

//// [/tluaconfig-base/types/nestjs-pino.d.tlua]
export {}

//// [/tluaconfig-base/types/reset.d.tlua]
export {}

//// [/tluaconfig-base/types/stripe-custom-typings.d.tlua]
export {}

//// [/tluaconfig.json]
{
  "extends": "./tluaconfig-base/backend.json",
  "compilerOptions": {
    "outDir": "dist",
    "rootDir": "src"
  },
  "exclude": ["node_modules", "dist"],
  "include": ["src/**/*"]
}


configFileName:: tluaconfig.json
CompilerOptions::
{
  "incremental": false,
  "lib": [],
  "module": 199,
  "noImplicitAny": true,
  "outDir": "/dist",
  "removeComments": true,
  "rootDir": "/src",
  "strictPropertyInitialization": false,
  "sourceMap": true,
  "target": 99,
  "types": [
    "node",
    "vitest/globals"
  ],
  "configFilePath": "/tluaconfig.json"
}

TypeAcquisition::
{}

FileNames::
/tluaconfig-base/types/ical2json.d.tlua,/tluaconfig-base/types/express.d.tlua,/tluaconfig-base/types/multer.d.tlua,/tluaconfig-base/types/reset.d.tlua,/tluaconfig-base/types/stripe-custom-typings.d.tlua,/tluaconfig-base/types/nestjs-modules.d.tlua,/tluaconfig-base/types/luxon.d.tlua,/tluaconfig-base/types/nestjs-pino.d.tlua,/src/main.tlua,/src/utils.tlua
Errors::
[96mtluaconfig-base/backend.json[0m:[93m5[0m:[93m5[0m - [91merror[0m[90m TLUA5023: [0mUnknown compiler option 'allowJs'.

[7m5[0m     "allowJs": true,
[7m [0m [91m    ~~~~~~~~~[0m

[96mtluaconfig-base/backend.json[0m:[93m8[0m:[93m5[0m - [91merror[0m[90m TLUA5023: [0mUnknown compiler option 'emitDecoratorMetadata'.

[7m8[0m     "emitDecoratorMetadata": true,
[7m [0m [91m    ~~~~~~~~~~~~~~~~~~~~~~~[0m

[96mtluaconfig-base/backend.json[0m:[93m9[0m:[93m5[0m - [91merror[0m[90m TLUA5023: [0mUnknown compiler option 'experimentalDecorators'.

[7m9[0m     "experimentalDecorators": true,
[7m [0m [91m    ~~~~~~~~~~~~~~~~~~~~~~~~[0m

[96mtluaconfig-base/backend.json[0m:[93m11[0m:[93m13[0m - [91merror[0m[90m TLUA6046: [0mArgument for '--lib' option must be: 'luajit'.

[7m11[0m     "lib": ["ESNext"],
[7m  [0m [91m            ~~~~~~~~[0m

