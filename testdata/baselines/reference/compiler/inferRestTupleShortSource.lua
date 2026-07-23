//// [tests/cases/compiler/inferRestTupleShortSource.tlua] ////

//// [inferRestTupleShortSource.tlua]
// Regression test: tlua panics when inferring [...rest, ...T] from a tuple shorter than fixed-arity constraint

function f<T extends [string]>(args: [...string[], ...T])
  // ...
end

f([])


//// [inferRestTupleShortSource.lua]
-- Regression test: tlua panics when inferring [...rest, ...T] from a tuple shorter than fixed-arity constraint
function f(args)
    -- ...
end
f();
