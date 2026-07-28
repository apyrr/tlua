//// [tests/cases/conformance/ported/typePredicatesInUnion_noMatch.tlua] ////

//// [typePredicatesInUnion_noMatch.tlua]
-- ported from tests/cases/compiler/typePredicatesInUnion_noMatch.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface A {
    pred(x: {}, y: {}): x is boolean
}
interface B {
    pred(x: {}, y: {}): y is string
}

type Or = A | B

function f(o: Or, x: {}, y: {})
    if o.pred(x, y) then
        local _ = x
        local _ = y
    end
end


//// [typePredicatesInUnion_noMatch.lua]
-- ported from tests/cases/compiler/typePredicatesInUnion_noMatch.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function f(o, x, y)
  if o.pred(x, y) then
    local _ = x;
    local _ = y;
  end
end
