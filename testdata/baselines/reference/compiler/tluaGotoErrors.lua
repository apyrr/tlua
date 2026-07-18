//// [tests/cases/compiler/tluaGotoErrors.tlua] ////

//// [tluaGotoErrors.tlua]
// Error: no such label anywhere.
function missing(): void
    goto nowhere
end

// Error: a label in a nested block is not visible outside it.
function nestedNotVisible(): void
    goto inner
    do
        ::inner::
    end
end

// Error: a label never crosses a function boundary, so it is simply not visible.
function crossFunction(): void
    ::here::
    local function inner(): void
        goto here
    end
    inner();
end

// Error: two labels of the same name in one block.
function duplicate(): void
    ::a::
    ::a::
    goto a
end

// Accepted: a nested block may shadow an enclosing label.
function shadowIsFine(): void
    ::a::
    do
        ::a::
        goto a
    end
    goto a
end


//// [tluaGotoErrors.lua]
-- Error: no such label anywhere.
function missing()
    goto nowhere;
end
-- Error: a label in a nested block is not visible outside it.
function nestedNotVisible()
    goto inner;
    do
        ::inner::
    end
end
-- Error: a label never crosses a function boundary, so it is simply not visible.
function crossFunction()
    ::here::
    local function inner()
        goto here;
    end
    inner();
end
-- Error: two labels of the same name in one block.
function duplicate()
    ::a::
    ::a::
    goto a;
end
-- Accepted: a nested block may shadow an enclosing label.
function shadowIsFine()
    ::a::
    do
        ::a::
        goto a;
    end
    goto a;
end
