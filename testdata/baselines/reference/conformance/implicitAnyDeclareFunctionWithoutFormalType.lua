//// [tests/cases/conformance/ported/implicitAnyDeclareFunctionWithoutFormalType.tlua] ////

//// [implicitAnyDeclareFunctionWithoutFormalType.tlua]
-- ported from tests/cases/compiler/implicitAnyDeclareFunctionWithoutFormalType.ts
-- dropped: @target: es2015 and @noImplicitAny: true directives (tlua defaults to latest target and strict checking)
-- dropped: named rest parameter `...args`; rewritten as tlua's unnamed vararg `...`, which has no
--   name for a diagnostic to point at, so upstream's implicit-any error on `args` has no tlua analogue
-- note: upstream's blanket "these should be errors" heading is stale. The upstream baseline reports
--   no diagnostic for the null/undefined-defaulted parameters either, and tlua likewise reports none
--   once they are normalized to nil.


-- these are implicit-any errors
function foo(x)
end

function bar(x: number, y)
end -- error at "y"; no error at "x"

function func2(a, b, c)
end -- error at "a,b,c"

-- these are not errors
function func3(...)
end

function func4(z = nil, w = nil)
end

function noError1(x = 3, y = 2)
end

function noError2(x: number, y: string)
end


//// [implicitAnyDeclareFunctionWithoutFormalType.lua]
-- ported from tests/cases/compiler/implicitAnyDeclareFunctionWithoutFormalType.ts
-- dropped: @target: es2015 and @noImplicitAny: true directives (tlua defaults to latest target and strict checking)
-- dropped: named rest parameter `...args`; rewritten as tlua's unnamed vararg `...`, which has no
--   name for a diagnostic to point at, so upstream's implicit-any error on `args` has no tlua analogue
-- note: upstream's blanket "these should be errors" heading is stale. The upstream baseline reports
--   no diagnostic for the null/undefined-defaulted parameters either, and tlua likewise reports none
--   once they are normalized to nil.
-- these are implicit-any errors
function foo(x)
end
function bar(x, y)
end -- error at "y"; no error at "x"
function func2(a, b, c)
end -- error at "a,b,c"
-- these are not errors
function func3(...)
end
function func4(z, w)
  if z == nil then
    z = nil;
  end
  if w == nil then
    w = nil;
  end
end
function noError1(x, y)
  if x == nil then
    x = 3;
  end
  if y == nil then
    y = 2;
  end
end
function noError2(x, y)
end
