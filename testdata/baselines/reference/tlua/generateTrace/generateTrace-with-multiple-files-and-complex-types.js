currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/main.tlua] *new* 
local c: Container<number> = { value: 42, map: (fn) => ({ value: fn(42), map: c.map }) };
local n: Nullable<string> = "hello";
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "strict": true,
        "noEmit": true
    }
}
//// [/home/src/workspaces/project/types.tlua] *new* 
interface Container<T> {
    value: T;
    map<U>(fn: (x: T) => U): Container<U>;
}
type Nullable<T> = T | null | undefined;

tlua --generateTrace /home/src/workspaces/project/trace --singleThreaded
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96mmain.tlua[0m:[93m1[0m:[93m37[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local c: Container<number> = { value: 42, map: (fn) => ({ value: fn(42), map: c.map }) };
[7m [0m [91m                                    ~[0m

[96mmain.tlua[0m:[93m1[0m:[93m46[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local c: Container<number> = { value: 42, map: (fn) => ({ value: fn(42), map: c.map }) };
[7m [0m [91m                                             ~[0m

[96mmain.tlua[0m:[93m1[0m:[93m77[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local c: Container<number> = { value: 42, map: (fn) => ({ value: fn(42), map: c.map }) };
[7m [0m [91m                                                                            ~[0m


Found 3 errors in the same file, starting at: main.tlua[90m:1[0m

//// [/home/src/tslibs/TS/Lib/lib.luajit.d.tlua] *Lib*
/// <reference no-default-lib="true"/>
interface Boolean {}
interface Function {}
interface CallableFunction {}
interface NewableFunction {}
interface IArguments {}
interface Number { toExponential: any; }
interface Object {}
interface RegExp {}
interface String { charAt: any; }
interface Array<T> { length: number; [n: number]: T; }
interface ReadonlyArray<T> {}
interface SymbolConstructor {
    (desc?: string | number): symbol;
    for(name: string): symbol;
    readonly toStringTag: symbol;
}
declare Symbol: SymbolConstructor;
interface Symbol {
    readonly [Symbol.toStringTag]: string;
}
declare console: { log(msg: any): void; };
declare function require(module: string): any;
//// [/home/src/workspaces/project/trace/legend.json] *new* 
[
  {
    "configFilePath": "/home/src/workspaces/project/tluaconfig.json",
    "tracePath": "/home/src/workspaces/project/trace/trace.json",
    "typesPath": "/home/src/workspaces/project/trace/types_0.json",
    "checkerId": 0
  }
]
//// [/home/src/workspaces/project/trace/trace.json] *new* 
[
{"pid":1,"tid":1,"ph":"M","cat":"__metadata","ts":1,"name":"process_name","args":{"name":"tlua"}},
{"pid":1,"tid":1,"ph":"M","cat":"__metadata","ts":1,"name":"thread_name","args":{"name":"Main"}},
{"pid":1,"tid":1,"ph":"M","cat":"disabled-by-default-devtools.timeline","ts":1,"name":"TracingStartedInBrowser"},
{"pid":1,"tid":1,"ph":"B","cat":"program","ts":2,"name":"createProgram","args":{"configFilePath":"/home/src/workspaces/project/tluaconfig.json"}},
{"pid":1,"tid":119112257,"ph":"M","cat":"__metadata","ts":1,"name":"thread_name","args":{"name":"file:/home/src/tslibs/TS/Lib/lib.luajit.d.tlua"}},
{"pid":1,"tid":119112257,"ph":"B","cat":"parse","ts":3,"name":"createSourceFile","args":{"path":"/home/src/tslibs/TS/Lib/lib.luajit.d.tlua"}},
{"pid":1,"tid":119112257,"ph":"E","cat":"parse","ts":4,"name":"createSourceFile","args":{"path":"/home/src/tslibs/TS/Lib/lib.luajit.d.tlua"}},
{"pid":1,"tid":544349038,"ph":"M","cat":"__metadata","ts":1,"name":"thread_name","args":{"name":"file:/home/src/workspaces/project/types.tlua"}},
{"pid":1,"tid":544349038,"ph":"B","cat":"parse","ts":5,"name":"createSourceFile","args":{"path":"/home/src/workspaces/project/types.tlua"}},
{"pid":1,"tid":544349038,"ph":"E","cat":"parse","ts":6,"name":"createSourceFile","args":{"path":"/home/src/workspaces/project/types.tlua"}},
{"pid":1,"tid":423323186,"ph":"M","cat":"__metadata","ts":1,"name":"thread_name","args":{"name":"file:/home/src/workspaces/project/main.tlua"}},
{"pid":1,"tid":423323186,"ph":"B","cat":"parse","ts":7,"name":"createSourceFile","args":{"path":"/home/src/workspaces/project/main.tlua"}},
{"pid":1,"tid":423323186,"ph":"E","cat":"parse","ts":8,"name":"createSourceFile","args":{"path":"/home/src/workspaces/project/main.tlua"}},
{"pid":1,"tid":1,"ph":"E","cat":"program","ts":9,"name":"createProgram","args":{"configFilePath":"/home/src/workspaces/project/tluaconfig.json"}},
{"pid":1,"tid":1,"ph":"B","cat":"emit","ts":10,"name":"emit"},
{"pid":1,"tid":544349038,"ph":"B","cat":"bind","ts":11,"name":"bindSourceFile","args":{"path":"/home/src/workspaces/project/types.tlua"}},
{"pid":1,"tid":544349038,"ph":"E","cat":"bind","ts":12,"name":"bindSourceFile","args":{"path":"/home/src/workspaces/project/types.tlua"}},
{"pid":1,"tid":423323186,"ph":"B","cat":"bind","ts":13,"name":"bindSourceFile","args":{"path":"/home/src/workspaces/project/main.tlua"}},
{"pid":1,"tid":423323186,"ph":"E","cat":"bind","ts":14,"name":"bindSourceFile","args":{"path":"/home/src/workspaces/project/main.tlua"}},
{"pid":1,"tid":119112257,"ph":"B","cat":"bind","ts":15,"name":"bindSourceFile","args":{"path":"/home/src/tslibs/TS/Lib/lib.luajit.d.tlua"}},
{"pid":1,"tid":119112257,"ph":"E","cat":"bind","ts":16,"name":"bindSourceFile","args":{"path":"/home/src/tslibs/TS/Lib/lib.luajit.d.tlua"}},
{"pid":1,"tid":544349038,"ph":"B","cat":"emit","ts":17,"name":"emit","args":{"path":"/home/src/workspaces/project/types.tlua"}},
{"pid":1,"tid":544349038,"ph":"E","cat":"emit","ts":18,"name":"emit","args":{"path":"/home/src/workspaces/project/types.tlua"}},
{"pid":1,"tid":423323186,"ph":"B","cat":"emit","ts":19,"name":"emit","args":{"path":"/home/src/workspaces/project/main.tlua"}},
{"pid":1,"tid":423323186,"ph":"E","cat":"emit","ts":20,"name":"emit","args":{"path":"/home/src/workspaces/project/main.tlua"}},
{"pid":1,"tid":1,"ph":"E","cat":"emit","ts":21,"name":"emit"}
]

//// [/home/src/workspaces/project/trace/types_0.json] *new* 
[{"id":1,"intrinsicName":"any","recursionId":0,"flags":["Any"]},
{"id":2,"intrinsicName":"any","recursionId":1,"flags":["Any"]},
{"id":3,"intrinsicName":"any","recursionId":2,"flags":["Any"]},
{"id":4,"intrinsicName":"any","recursionId":3,"flags":["Any"]},
{"id":5,"intrinsicName":"error","recursionId":4,"flags":["Any"]},
{"id":6,"intrinsicName":"unresolved","recursionId":5,"flags":["Any"]},
{"id":7,"intrinsicName":"any","recursionId":6,"flags":["Any"]},
{"id":8,"intrinsicName":"intrinsic","recursionId":7,"flags":["Any"]},
{"id":9,"intrinsicName":"unknown","recursionId":8,"flags":["Unknown"]},
{"id":10,"intrinsicName":"nil","recursionId":9,"flags":["Undefined"]},
{"id":11,"intrinsicName":"nil","recursionId":10,"flags":["Undefined"]},
{"id":12,"intrinsicName":"nil","recursionId":11,"flags":["Undefined"]},
{"id":13,"intrinsicName":"string","recursionId":12,"flags":["String"]},
{"id":14,"intrinsicName":"number","recursionId":13,"flags":["Number"]},
{"id":15,"recursionId":14,"flags":["BooleanLiteral"],"display":"false"},
{"id":16,"recursionId":15,"flags":["BooleanLiteral"],"display":"false"},
{"id":17,"recursionId":16,"flags":["BooleanLiteral"],"display":"true"},
{"id":18,"recursionId":17,"flags":["BooleanLiteral"],"display":"true"},
{"id":19,"recursionId":18,"unionTypes":[15,17],"flags":["Boolean","Union"],"display":"boolean"},
{"id":20,"intrinsicName":"symbol","recursionId":19,"flags":["ESSymbol"]},
{"id":21,"intrinsicName":"void","recursionId":20,"flags":["Void"]},
{"id":22,"intrinsicName":"never","recursionId":21,"flags":["Never"]},
{"id":23,"intrinsicName":"never","recursionId":22,"flags":["Never"]},
{"id":24,"intrinsicName":"never","recursionId":23,"flags":["Never"]},
{"id":25,"intrinsicName":"never","recursionId":24,"flags":["Never"]},
{"id":26,"intrinsicName":"table","recursionId":25,"flags":["NonPrimitive"]},
{"id":27,"recursionId":26,"unionTypes":[13,14,20],"flags":["Union"],"display":"string | number | symbol"},
{"id":28,"recursionId":27,"unionTypes":[10,13],"flags":["Union"],"display":"string | nil"},
{"id":29,"recursionId":28,"unionTypes":[10,14],"flags":["Union"],"display":"number | nil"},
{"id":30,"recursionId":29,"unionTypes":[10,13,14,15,17],"flags":["Union"],"display":"string | number | boolean | nil"},
{"id":31,"intrinsicName":"never","recursionId":30,"flags":["Never"]},
{"id":32,"recursionId":31,"flags":["Object"],"display":"{}"},
{"id":33,"recursionId":32,"flags":["Object"],"display":"{}"},
{"id":34,"recursionId":33,"flags":["Object"],"display":"{}"},
{"id":35,"symbolName":"__type","recursionId":34,"flags":["Object"],"display":"{}"},
{"id":36,"recursionId":35,"flags":["Object"],"display":"{}"},
{"id":37,"recursionId":36,"unionTypes":[10,36],"flags":["Union"],"display":"{} | nil"},
{"id":38,"recursionId":37,"flags":["Object"],"display":"{}"},
{"id":39,"recursionId":38,"flags":["Object"],"display":"{}"},
{"id":40,"recursionId":39,"flags":["Object"],"display":"{}"},
{"id":41,"recursionId":40,"flags":["Object"],"display":"{}"},
{"id":42,"recursionId":41,"flags":["Object"],"display":"{}"},
{"id":43,"recursionId":42,"flags":["TypeParameter"]},
{"id":44,"recursionId":43,"flags":["TypeParameter"]},
{"id":45,"recursionId":44,"flags":["TypeParameter"]},
{"id":46,"recursionId":45,"flags":["TypeParameter"]},
{"id":47,"recursionId":46,"flags":["TypeParameter"]},
{"id":48,"symbolName":"IArguments","recursionId":47,"firstDeclaration":{"path":"/home/src/tslibs/ts/lib/lib.luajit.d.tlua","start":{"line":6,"character":1},"end":{"line":6,"character":24}},"flags":["Object"]},
{"id":49,"symbolName":"_G","recursionId":48,"flags":["Object"],"display":"typeof _G"},
{"id":50,"symbolName":"Array","recursionId":49,"instantiatedType":50,"typeArguments":[51],"firstDeclaration":{"path":"/home/src/tslibs/ts/lib/lib.luajit.d.tlua","start":{"line":11,"character":1},"end":{"line":11,"character":55}},"flags":["Object"]},
{"id":51,"symbolName":"T","recursionId":50,"firstDeclaration":{"path":"/home/src/tslibs/ts/lib/lib.luajit.d.tlua","start":{"line":11,"character":17},"end":{"line":11,"character":18}},"flags":["TypeParameter"]},
{"id":52,"symbolName":"Array","recursionId":49,"firstDeclaration":{"path":"/home/src/tslibs/ts/lib/lib.luajit.d.tlua","start":{"line":11,"character":1},"end":{"line":11,"character":55}},"flags":["TypeParameter"]},
{"id":53,"recursionId":51,"flags":["Object"],"display":"{}"},
{"id":54,"recursionId":52,"flags":["Object"],"display":"{}"},
{"id":55,"symbolName":"String","recursionId":53,"firstDeclaration":{"path":"/home/src/tslibs/ts/lib/lib.luajit.d.tlua","start":{"line":10,"character":1},"end":{"line":10,"character":34}},"flags":["Object"]},
{"id":56,"symbolName":"Number","recursionId":54,"firstDeclaration":{"path":"/home/src/tslibs/ts/lib/lib.luajit.d.tlua","start":{"line":7,"character":1},"end":{"line":7,"character":41}},"flags":["Object"]},
{"id":57,"symbolName":"Boolean","recursionId":55,"firstDeclaration":{"path":"/home/src/tslibs/ts/lib/lib.luajit.d.tlua","start":{"line":2,"character":1},"end":{"line":2,"character":21}},"flags":["Object"]},
{"id":58,"symbolName":"RegExp","recursionId":56,"firstDeclaration":{"path":"/home/src/tslibs/ts/lib/lib.luajit.d.tlua","start":{"line":9,"character":1},"end":{"line":9,"character":20}},"flags":["Object"]},
{"id":59,"recursionId":57,"instantiatedType":59,"typeArguments":[60],"flags":["Object"]},
{"id":60,"recursionId":58,"flags":["TypeParameter"]},
{"id":61,"recursionId":59,"flags":["TypeParameter"]},
{"id":62,"recursionId":60,"instantiatedType":62,"typeArguments":[63],"flags":["Object"]},
{"id":63,"recursionId":61,"flags":["TypeParameter"]},
{"id":64,"recursionId":62,"flags":["TypeParameter"]},
{"id":65,"recursionId":63,"instantiatedType":59,"typeArguments":[1],"flags":["Object"]},
{"id":66,"recursionId":64,"flags":["TypeParameter"]},
{"id":67,"recursionId":65,"isTuple":true,"instantiatedType":67,"typeArguments":[66],"flags":["Object"]},
{"id":68,"recursionId":66,"flags":["TypeParameter"]},
{"id":69,"recursionId":65,"instantiatedType":67,"typeArguments":[1],"flags":["Object"]},
{"id":70,"recursionId":67,"instantiatedType":59,"typeArguments":[2],"flags":["Object"]},
{"id":71,"symbolName":"ReadonlyArray","recursionId":68,"instantiatedType":71,"typeArguments":[72],"firstDeclaration":{"path":"/home/src/tslibs/ts/lib/lib.luajit.d.tlua","start":{"line":12,"character":1},"end":{"line":12,"character":30}},"flags":["Object"]},
{"id":72,"symbolName":"T","recursionId":69,"firstDeclaration":{"path":"/home/src/tslibs/ts/lib/lib.luajit.d.tlua","start":{"line":12,"character":25},"end":{"line":12,"character":26}},"flags":["TypeParameter"]},
{"id":73,"symbolName":"ReadonlyArray","recursionId":68,"firstDeclaration":{"path":"/home/src/tslibs/ts/lib/lib.luajit.d.tlua","start":{"line":12,"character":1},"end":{"line":12,"character":30}},"flags":["TypeParameter"]},
{"id":74,"recursionId":70,"instantiatedType":62,"typeArguments":[1],"flags":["Object"]}]


