//// [tests/cases/compiler/tluaGotoFlow.tlua] ////

//// [tluaGotoFlow.tlua]
// The narrowed type travels along the goto edge into the label.
function narrowThroughGoto(x: string | nil): string
    if x ~= nil then goto has end
    return "";
    ::has::
    return x;
end

// The label is a join: its type is the union of every edge that reaches it.
function joinAtLabel(x: string | number): string
    if type(x) == "string" then goto stringify end
    x = "n";
    ::stringify::
    return x;
end

// Code between an unconditional goto and its label is unreachable.
function unreachableAfterGoto(n: number): number
    goto skip
    n = n + 1;
    ::skip::
    return n;
end

// A label reached only by a goto keeps the code after it reachable, even when
// the fall-through into the label is dead.
function reachableViaGotoOnly(flag: boolean): number
    if flag then goto tail end
    return 1;
    ::tail::
    return 2;
end

// A label nothing jumps to and nothing falls into is unreachable.
function deadLabel(): number
    return 1;
    ::never::
    return 2;
end


//// [tluaGotoFlow.lua]
-- The narrowed type travels along the goto edge into the label.
function narrowThroughGoto(x)
    if x ~= nil then
        goto has;
    end
    return "";
    ::has::
    return x;
end
-- The label is a join: its type is the union of every edge that reaches it.
function joinAtLabel(x)
    if type(x) == "string" then
        goto stringify;
    end
    x = "n";
    ::stringify::
    return x;
end
-- Code between an unconditional goto and its label is unreachable.
function unreachableAfterGoto(n)
    goto skip;
    n = n + 1;
    ::skip::
    return n;
end
-- A label reached only by a goto keeps the code after it reachable, even when
-- the fall-through into the label is dead.
function reachableViaGotoOnly(flag)
    if flag then
        goto tail;
    end
    return 1;
    ::tail::
    return 2;
end
-- A label nothing jumps to and nothing falls into is unreachable.
function deadLabel()
    return 1;
    ::never::
    return 2;
end
