//// [tests/cases/conformance/ported/commentOnBlock1.tlua] ////

//// [commentOnBlock1.tlua]
-- ported from tests/cases/compiler/commentOnBlock1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

-- note: upstream /*asdf*/ block-attached comment is attached to the Lua do block.
-- compiler gap: inline block comments are emitted as line comments; exact inline placement is not preserved.
function f()
    -- asdf
    do -- asdf
    end
end


//// [commentOnBlock1.lua]
-- ported from tests/cases/compiler/commentOnBlock1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- note: upstream /*asdf*/ block-attached comment is attached to the Lua do block.
-- compiler gap: inline block comments are emitted as line comments; exact inline placement is not preserved.
function f()
  -- asdf
  do
  end
end
