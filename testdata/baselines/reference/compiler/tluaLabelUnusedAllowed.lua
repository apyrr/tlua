//// [tests/cases/compiler/tluaLabelUnusedAllowed.tlua] ////

//// [tluaLabelUnusedAllowed.tlua]
// allowUnusedLabels suppresses the report for a label no goto targets.
function unused(): void
    ::dead::
end


//// [tluaLabelUnusedAllowed.lua]
-- allowUnusedLabels suppresses the report for a label no goto targets.
function unused()
    ::dead::
end
