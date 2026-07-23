currentDirectory::/user/username/projects/sample1
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/sample1/core/anotherModule.tlua] *new* 
local World = "hello";
return { World = World };
//// [/user/username/projects/sample1/core/index.tlua] *new* 
local someString: string = "HELLO WORLD";
function leftPad(s: string, n: number) return s + n; end
function multiply(a: number, b: number) return a * b; end
return { someString = someString, leftPad = leftPad, multiply = multiply };
//// [/user/username/projects/sample1/core/some_decl.d.tlua] *new* 
declare dts: any;
//// [/user/username/projects/sample1/core/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "declaration": true,
        "declarationMap": true,
        "skipDefaultLibCheck": true,
    },
}
//// [/user/username/projects/sample1/logic/index.tlua] *new* 
local c = require('core.index');
function getSecondsInDay()
    return c.multiply(10, 15);
end
local mod = require('core.anotherModule');
return { getSecondsInDay = getSecondsInDay, m = mod };
//// [/user/username/projects/sample1/logic/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "declaration": true,
        "sourceMap": true,
        "skipDefaultLibCheck": true,
        "rootDir": "..",
    },
    "references": [
        { "path": "../core" },
    ],
}
//// [/user/username/projects/sample1/tests/index.tlua] *new* 
local c = require('core.index');
local logic = require('logic.index');

c.leftPad("", 10);
logic.getSecondsInDay();

local mod = require('core.anotherModule');
return { m = mod };
//// [/user/username/projects/sample1/tests/tluaconfig.json] *new* 
{
    "references": [
        { "path": "../core" },
        { "path": "../logic" },
    ],
    "files": ["index.tlua"],
    "compilerOptions": {
        "composite": true,
        "declaration": true,
        "skipDefaultLibCheck": true,
        "rootDir": "..",
        "rootDir": "..",
    },
}

tlua --b logic2 --clean
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[91merror[0m[90m TLUA6053: [0mFile '/user/username/projects/sample1/logic2/tluaconfig.json' not found.

Found 1 error.


