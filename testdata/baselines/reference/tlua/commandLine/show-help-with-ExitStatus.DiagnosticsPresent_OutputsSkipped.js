currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
Version FakeTSVersion
tlua: The Lua Compiler - Version FakeTSVersion                                                                         [44m     [39;49m
                                                                                                                   [44m[97m Lua [39m[39;49m
[1mCOMMON COMMANDS[22m

  [94mtlua[39m
  Compiles the current project (tluaconfig.json in the working directory.)

  [94mtlua app.tlua util.tlua[39m
  Ignoring tluaconfig.json, compiles the specified files with default compiler options.

  [94mtlua -b[39m
  Build a composite project in the working directory.

  [94mtlua --init[39m
  Creates a tluaconfig.json with the recommended settings in the working directory.

  [94mtlua -p ./path/to/tluaconfig.json[39m
  Compiles the Lua project located at the specified path.

  [94mtlua --help --all[39m
  An expanded version of this information, showing all possible compiler options

  [94mtlua --noEmit[39m
  [94mtlua --target esnext[39m
  Compiles the current project, with additional settings.

[1mCOMMAND LINE FLAGS[22m

[94m      --help, -h  [39mPrint this message.


[94m     --watch, -w  [39mWatch input files.


[94m           --all  [39mShow all compiler options.


[94m   --version, -v  [39mPrint the compiler's version.


[94m          --init  [39mInitializes a tlua project and creates a tluaconfig.json file.


[94m   --project, -p  [39mCompile the project given the path to its configuration file, or to a folder with a 'tluaconfig.json'.


[94m    --showConfig  [39mPrint the final configuration instead of building.


[94m  --ignoreConfig  [39mIgnore the tluaconfig found and build with commandline options and files.


[94m     --build, -b  [39mBuild one or more projects and their dependencies, if out of date


[1mCOMMON COMPILER OPTIONS[22m

[94m               --pretty  [39mEnable color and formatting in tlua's output to make compiler errors easier to read.

                  type:  boolean

               default:  true


[94m      --declaration, -d  [39mGenerate .d.ts files from Lua and JavaScript files in your project.

                  type:  boolean

               default:  `false`, unless `composite` is set


[94m       --declarationMap  [39mCreate sourcemaps for d.ts files.

                  type:  boolean

               default:  false


[94m  --emitDeclarationOnly  [39mOnly output d.ts files and not JavaScript files.

                  type:  boolean

               default:  false


[94m            --sourceMap  [39mCreate source map files for emitted JavaScript files.

                  type:  boolean

               default:  false


[94m               --noEmit  [39mDisable emitting files from a compilation.

                  type:  boolean

               default:  false


[94m           --target, -t  [39mSet the JavaScript language version for emitted JavaScript and include compatible library decla
                         rations.

                one of:  es6/es2015, es2016, es2017, es2018, es2019, es2020, es2021, es2022, es2023, es2024, es2025, esn
                         ext

               default:  es2025


[94m           --module, -m  [39mSpecify what module code is generated.

                one of:  commonjs, es6/es2015, es2020, es2022, esnext, node16, node18, node20, nodenext, preserve

               default:  undefined


[94m                  --lib  [39mSpecify a set of bundled library declaration files that describe the target runtime environment
                         .

           one or more:  luajit

               default:  undefined


[94m                  --jsx  [39mSpecify what JSX code is generated.

                one of:  preserve, react-native, react-jsx, react-jsxdev, react

               default:  undefined


[94m              --outFile  [39mSpecify a file that bundles all outputs into one JavaScript file. If 'declaration' is true, als
                         o designates a file that bundles all .d.ts output.


[94m               --outDir  [39mSpecify an output folder for all emitted files.


[94m       --removeComments  [39mDisable emitting comments.

                  type:  boolean

               default:  false


[94m               --strict  [39mEnable all strict type-checking options.

                  type:  boolean

               default:  true


[94m                --types  [39mSpecify type package names to be included without being referenced in a source file.


You can learn about all of the compiler options at https://aka.ms/tsc


