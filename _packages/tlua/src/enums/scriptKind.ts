export var ScriptKind: any;
(function (ScriptKind) {
    ScriptKind[ScriptKind["Unknown"] = 0] = "Unknown";
    ScriptKind[ScriptKind["JS"] = 1] = "JS";
    ScriptKind[ScriptKind["TS"] = 2] = "TS";
    ScriptKind[ScriptKind["External"] = 3] = "External";
    ScriptKind[ScriptKind["JSON"] = 4] = "JSON";
    ScriptKind[ScriptKind["Deferred"] = 5] = "Deferred";
})(ScriptKind || (ScriptKind = {}));
