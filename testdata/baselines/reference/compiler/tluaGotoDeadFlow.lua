//// [tests/cases/compiler/tluaGotoDeadFlow.tlua] ////

//// [tluaGotoDeadFlow.tlua]
// A label reached only through dead fall-through, then a backward goto to it,
// used to hang the flow analysis: the label's loop node became its own only
// antecedent, a cycle with no base. The label is now correctly unreachable, and
// the compiler terminates.
function selfGotoInDeadCode(): number
    return 1;
    ::again::
    goto again
end

// A label whose only goto sits in dead code is NOT unused: the goto still names
// it, even though it wires no flow edge.
function deadGotoStillReferences(): number
    ::x::
    return 1;
    goto x
end

// A label nothing references at all is unused.
function trulyUnused(): number
    ::dead::
    return 1;
end

// A live forward goto keeps a label — and the code after it — reachable, even
// when the fall-through into the label is dead.
function reachableViaGotoOnly(flag: boolean): number
    if flag then goto tail end
    return 1;
    ::tail::
    return 2;
end


//// [tluaGotoDeadFlow.lua]
-- A label reached only through dead fall-through, then a backward goto to it,
-- used to hang the flow analysis: the label's loop node became its own only
-- antecedent, a cycle with no base. The label is now correctly unreachable, and
-- the compiler terminates.
function selfGotoInDeadCode()
  return 1;
  ::again::
  goto again;
end
-- A label whose only goto sits in dead code is NOT unused: the goto still names
-- it, even though it wires no flow edge.
function deadGotoStillReferences()
  ::x::
  return 1;
  goto x;
end
-- A label nothing references at all is unused.
function trulyUnused()
  ::dead::
  return 1;
end
-- A live forward goto keeps a label — and the code after it — reachable, even
-- when the fall-through into the label is dead.
function reachableViaGotoOnly(flag)
  if flag then
    goto tail;
  end
  return 1;
  ::tail::
  return 2;
end
