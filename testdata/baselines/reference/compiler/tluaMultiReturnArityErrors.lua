//// [tests/cases/compiler/tluaMultiReturnArityErrors.tlua] ////

//// [tluaMultiReturnArityErrors.tlua]
function need2(): (number, string)
  return 1;
end

function allows2(): (number, string)
  return 1, "a", true;
end

function wrongTypes(): (number, string)
  return "a", 1;
end

function single(): number
  return 1, "a";
end

// A bare return provides zero values against a pack that requires two.
function bareReturn(): (number, string)
  return;
end

// An all-optional pack allows returning nothing at all.
function allOptional(): (number?, string?)
end

// A pack with required values must return.
function mustReturn(): (number, string?)
end

// A type predicate cannot start a multiple return type.
function badPredicate(x: unknown): (x is string, number) {
  return type(x) == "string", 1;
}


//// [tluaMultiReturnArityErrors.lua]
function need2()
    return 1;
end
function allows2()
    return 1, "a", true;
end
function wrongTypes()
    return "a", 1;
end
function single()
    return 1, "a";
end
-- A bare return provides zero values against a pack that requires two.
function bareReturn()
    return;
end
-- An all-optional pack allows returning nothing at all.
function allOptional()
end
-- A pack with required values must return.
function mustReturn()
end
-- A type predicate cannot start a multiple return type.
function badPredicate(x) {
    return type(x) == "string", 1;
}
