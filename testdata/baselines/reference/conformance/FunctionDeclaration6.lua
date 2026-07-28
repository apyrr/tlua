//// [tests/cases/conformance/ported/FunctionDeclaration6.tlua] ////

//// [FunctionDeclaration6.tlua]
-- ported from tests/cases/compiler/FunctionDeclaration6.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: C-style statement block rewritten as Lua `do`/`end`

do
    function foo();
    function bar()
    end
end


//// [FunctionDeclaration6.lua]
-- ported from tests/cases/compiler/FunctionDeclaration6.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: C-style statement block rewritten as Lua `do`/`end`
do
  function bar()
  end
end
