//// [tests/cases/conformance/ported/contextualTyping33.tlua] ////

//// [contextualTyping33.tlua]
-- ported from tests/cases/compiler/contextualTyping33.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- The second function intentionally returns string and should produce a checker error.

function foo(param: { (): number; (i: number): number }[])
end

foo({ function() return 1 end, function() return "foo" end })


//// [contextualTyping33.lua]
-- ported from tests/cases/compiler/contextualTyping33.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- The second function intentionally returns string and should produce a checker error.
function foo(param)
end
foo({ function()
    return 1;
  end, function()
    return "foo";
  end });
