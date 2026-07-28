//// [tests/cases/conformance/ported/numberToString.tlua] ////

//// [numberToString.tlua]
-- ported from tests/cases/compiler/numberToString.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function f1(n: number): string
    return n -- error return type mismatch
end

function f2(s: string): void
end

f1(3)
f2(3) -- error no coercion to string
f2(3 .. "") -- ok string concatenation


//// [numberToString.lua]
-- ported from tests/cases/compiler/numberToString.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function f1(n)
  return n; -- error return type mismatch
end
function f2(s)
end
f1(3);
f2(3); -- error no coercion to string
f2(3 .. ""); -- ok string concatenation
