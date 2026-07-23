//// [tests/cases/compiler/tluaGotoIntoLetScope.tlua] ////

//// [tluaGotoIntoLetScope.tlua]
// The jump-into-local-scope rule covers the transitional let/const/local forms
// too: they declare names that compile to Lua locals.
declare function log(n: number): void;
function intoLetScope(): void
    goto done
    local x = 1;
    ::done::
    log(x);
end

function intoConstScope(): void
    goto done
    local c = 1;
    ::done::
    log(c);
end

function intoVarScope(): void
    goto done
    local v = 1;
    ::done::
    log(v);
end


//// [tluaGotoIntoLetScope.lua]
function intoLetScope()
  goto done;
  local x = 1;
  ::done::
  log(x);
end
function intoConstScope()
  goto done;
  local c = 1;
  ::done::
  log(c);
end
function intoVarScope()
  goto done;
  local v = 1;
  ::done::
  log(v);
end
