//// [tests/cases/compiler/tluaLabelUnused.tlua] ////

//// [tluaLabelUnused.tlua]
// A label no goto targets is reported.
function unused(): void
    ::dead::
end

// A label a goto targets is silent.
function used(): void
    ::live::
    goto live
end


//// [tluaLabelUnused.lua]
-- A label no goto targets is reported.
function unused()
  ::dead::
end
-- A label a goto targets is silent.
function used()
  ::live::
  goto live;
end
