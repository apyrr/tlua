currentDirectory::/user/username/projects/demo
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/demo/animals/animal.tlua] *new* 
type Size = "small" | "medium" | "large";
interface Animal {
    size: Size;
}
//// [/user/username/projects/demo/animals/dog.tlua] *new* 
local utilities = require('core.utilities');

interface Dog extends Animal {
    woof(): void;
    name: string;
}

function createDog(): Dog
    return ({
        size: "medium",
        woof: function()
            console.log("Woof!");
        end,
        name: utilities.makeRandomName()
    });
end

return { createDog = createDog };
//// [/user/username/projects/demo/animals/index.tlua] *new* 
local dog = require('animals.dog');

return { createDog = dog.createDog };
//// [/user/username/projects/demo/animals/tluaconfig.json] *new* 
{
    "extends": "../tluaconfig-base.json",
    "compilerOptions": {
        "outDir": "../lib/animals",
        "rootDir": ".."
    },
    "references": [
        { "path": "../core" }
    ]
}
//// [/user/username/projects/demo/core/tluaconfig.json] *new* 
{
    "extends": "../tluaconfig-base.json",
    "compilerOptions": {
        "outDir": "../lib/core",
        "rootDir": ".."
    },
    "references": [
        {
            "path": "../zoo",
        }
    ]
}
//// [/user/username/projects/demo/core/utilities.tlua] *new* 
function makeRandomName()
    return "Bob!?! ";
end

function lastElementOf<T>(arr: T[]): T | undefined
    if arr.length == 0 then return undefined end
    return arr[arr.length - 1];
end

return { makeRandomName = makeRandomName, lastElementOf = lastElementOf };
//// [/user/username/projects/demo/tluaconfig-base.json] *new* 
{
    "compilerOptions": {
        "declaration": true,
        "target": "es5",
        "module": "commonjs",
        "strict": true,
        "noUnusedLocals": true,
        "noUnusedParameters": true,
        "noImplicitReturns": true,
        "composite": true,
    },
}
//// [/user/username/projects/demo/tluaconfig.json] *new* 
{
    "files": [],
    "references": [
        {
            "path": "./core"
        },
        {
            "path": "./animals",
        },
        {
            "path": "./zoo",
        },
    ],
}
//// [/user/username/projects/demo/zoo/tluaconfig.json] *new* 
{
    "extends": "../tluaconfig-base.json",
    "compilerOptions": {
        "outDir": "../lib/zoo",
        "rootDir": ".."
    },
    "references": [
        {
            "path": "../animals"
        }
    ]
}
//// [/user/username/projects/demo/zoo/zoo.tlua] *new* 
local animals = require('animals.index');

function createZoo(): Array<Dog>
    return {
        animals.createDog()
    };
end

return { createZoo = createZoo };

tlua --b --verbose
ExitStatus:: ProjectReferenceCycle_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * animals/tluaconfig.json
    * zoo/tluaconfig.json
    * core/tluaconfig.json
    * tluaconfig.json

[91merror[0m[90m TLUA6202: [0mProject references may not form a circular graph. Cycle detected: /user/username/projects/demo/tluaconfig.json
/user/username/projects/demo/core/tluaconfig.json
/user/username/projects/demo/zoo/tluaconfig.json
/user/username/projects/demo/animals/tluaconfig.json

Found 1 error.


