//// [tests/cases/conformance/ported/typeParameterConstraints1.tlua] ////

//// [typeParameterConstraints1.tlua]
-- ported from tests/cases/compiler/typeParameterConstraints1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: Date and RegExp constraints (JavaScript library types unavailable; rewritten as table)
-- dropped: Object constraint (JavaScript library type unavailable; rewritten as table)
-- dropped: undefined constraint (rewritten as nil)

function foo1<T extends any>(test: T) end
function foo2<T extends number>(test: T) end
function foo3<T extends string>(test: T) end
function foo4<T extends table>(test: T) end -- valid
function foo5<T extends table>(test: T) end -- valid
function foo6<T extends hm>(test: T) end
function foo7<T extends table>(test: T) end -- valid
function foo8<T extends "">(test: T) end
function foo9<T extends 1>(test: T) end
function foo10<T extends (1)>(test: T) end
function foo11<T extends nil>(test: T) end
function foo12<T extends nil>(test: T) end
function foo13<T extends void>(test: T) end


//// [typeParameterConstraints1.lua]
-- ported from tests/cases/compiler/typeParameterConstraints1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: Date and RegExp constraints (JavaScript library types unavailable; rewritten as table)
-- dropped: Object constraint (JavaScript library type unavailable; rewritten as table)
-- dropped: undefined constraint (rewritten as nil)
function foo1(test)
end
function foo2(test)
end
function foo3(test)
end
function foo4(test)
end -- valid
function foo5(test)
end -- valid
function foo6(test)
end
function foo7(test)
end -- valid
function foo8(test)
end
function foo9(test)
end
function foo10(test)
end
function foo11(test)
end
function foo12(test)
end
function foo13(test)
end
