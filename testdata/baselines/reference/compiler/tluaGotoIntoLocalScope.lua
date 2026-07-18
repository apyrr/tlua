//// [tests/cases/compiler/tluaGotoIntoLocalScope.tlua] ////

//// [tluaGotoIntoLocalScope.tlua]
// Error: `x` is still in scope at ::done::, so the jump enters its scope.
function intoScope(): number
    goto done
    local x = 1;
    ::done::
    return x;
end

// Error: a `local function` declares a local too.
function intoFunctionScope(): void
    goto done
    local function helper(): void
    end
    ::done::
    helper();
end

// Accepted: the label sits in the block's trailing void region, where no local
// is in scope. This is the relaxation that makes the skip-to-end idiom legal.
function trailingLabel(): void
    goto done
    local x = 1;
    ::done::
end

// Accepted: a local inside a nested block died with that block.
function subBlockLocal(): void
    goto done
    do
        local x = 1;
    end
    ::done::
    return;
end

// Error: a multi-name local declares every binding at once, so jumping past it
// enters their scope. The message names the first binding.
function intoMultiNameScope(): number
    goto done
    local a, b = 1, 2;
    ::done::
    return b;
end

// Accepted: a backward jump leaves scopes, it never enters one.
function backwardPastLocal(): void
    ::top::
    local x = 1;
    if x < 5 then goto top end
end

// Accepted: the loop body's trailing label, jumped to past the body's local.
function skipToEnd(n: number): number
    local total = 0;
    while n > 0 do
        n = n - 1;
        if n == 3 then goto next end
        local step = n * 2;
        total = total + step;
        ::next::
    end
    return total;
end


//// [tluaGotoIntoLocalScope.lua]
-- Error: `x` is still in scope at ::done::, so the jump enters its scope.
function intoScope()
    goto done;
    local x = 1;
    ::done::
    return x;
end
-- Error: a `local function` declares a local too.
function intoFunctionScope()
    goto done;
    local function helper()
    end
    ::done::
    helper();
end
-- Accepted: the label sits in the block's trailing void region, where no local
-- is in scope. This is the relaxation that makes the skip-to-end idiom legal.
function trailingLabel()
    goto done;
    local x = 1;
    ::done::
end
-- Accepted: a local inside a nested block died with that block.
function subBlockLocal()
    goto done;
    do
        local x = 1;
    end
    ::done::
    return;
end
-- Error: a multi-name local declares every binding at once, so jumping past it
-- enters their scope. The message names the first binding.
function intoMultiNameScope()
    goto done;
    local a, b = 1, 2;
    ::done::
    return b;
end
-- Accepted: a backward jump leaves scopes, it never enters one.
function backwardPastLocal()
    ::top::
    local x = 1;
    if x < 5 then
        goto top;
    end
end
-- Accepted: the loop body's trailing label, jumped to past the body's local.
function skipToEnd(n)
    local total = 0;
    while n > 0 do
        n = n - 1;
        if n == 3 then
            goto next;
        end
        local step = n * 2;
        total = total + step;
        ::next::
    end
    return total;
end
