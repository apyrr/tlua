//// [tests/cases/compiler/tluaLabelDeadBlock.tlua] ////

//// [tluaLabelDeadBlock.tlua]
// A dead block still scopes its labels: an unused label inside one is reported
// exactly like an unused label directly in the dead tail of a live block.
function deadLabelBare(): number
    return 1;
    ::L1::
end

function deadLabelInBlock(): number {
    return 1;
    { ::L2:: }
}

// And a goto inside a dead block still marks its label referenced.
function deadBlockGotoReferences(): number {
    return 1;
    { ::L3:: goto L3 }
}


//// [tluaLabelDeadBlock.lua]
-- A dead block still scopes its labels: an unused label inside one is reported
-- exactly like an unused label directly in the dead tail of a live block.
function deadLabelBare()
    return 1;
    ::L1::
end
function deadLabelInBlock() {
    return 1;
    {
        ::L2::
    }
}
-- And a goto inside a dead block still marks its label referenced.
function deadBlockGotoReferences() {
    return 1;
    {
        ::L3::
        goto L3;
    }
}
