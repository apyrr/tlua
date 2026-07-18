//// [tests/cases/compiler/tluaGotoForwardDeadEdge.tlua] ////

//// [tluaGotoForwardDeadEdge.tlua]
// A label's entry is the join of the fall-through AND every forward goto. A
// forward edge the checker later proves dead (here: behind a never call) must
// not mask a live fall-through — only the goto itself is unreachable.
declare function fail(): never;
declare function log(n: number): void;
function liveFallThroughDeadGoto(b: boolean): void
    if b then
        fail();
        goto done
    end
    log(1);
    ::done::
    log(2);
end

// The mirror case: several forward gotos, the FIRST-bound one dead, a later one
// live, fall-through dead. The label is still reachable through the live edge.
function liveSecondGoto(b: boolean): number
    if b then
        fail();
        goto out
    end
    goto out
    ::out::
    return 1;
end

// Narrowing still joins all live entries when a dead edge is among them.
function narrowPastDeadEdge(x: string | nil): string
    if x == nil then
        fail();
        goto has
    end
    ::has::
    return x;
end


//// [tluaGotoForwardDeadEdge.lua]
function liveFallThroughDeadGoto(b)
    if b then
        fail();
        goto done;
    end
    log(1);
    ::done::
    log(2);
end
-- The mirror case: several forward gotos, the FIRST-bound one dead, a later one
-- live, fall-through dead. The label is still reachable through the live edge.
function liveSecondGoto(b)
    if b then
        fail();
        goto out;
    end
    goto out;
    ::out::
    return 1;
end
-- Narrowing still joins all live entries when a dead edge is among them.
function narrowPastDeadEdge(x)
    if x == nil then
        fail();
        goto has;
    end
    ::has::
    return x;
end
