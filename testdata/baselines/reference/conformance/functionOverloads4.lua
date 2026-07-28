//// [tests/cases/conformance/ported/functionOverloads4.tlua] ////

//// [functionOverloads4.tlua]
-- ported from tests/cases/compiler/functionOverloads4.ts
-- dropped: @target: es2015 directive (tlua defaults to esnext)

function foo(): number;
function foo(): string
    return "a"
end


//// [functionOverloads4.lua]
-- ported from tests/cases/compiler/functionOverloads4.ts
-- dropped: @target: es2015 directive (tlua defaults to esnext)
function foo()
  return "a";
end
