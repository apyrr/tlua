//// [tests/cases/conformance/ported/restParameters.tlua] ////

//// [restParameters.tlua]
-- ported from tests/cases/compiler/restParameters.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function f18(a?: string, ...: number)
end

function f19(a?: string, b?: number, ...: number)
end

function f20(a: string, b?: string, ...: number)
end

function f21(a: string, b?: string, c?: number, ...: number)
end


//// [restParameters.lua]
-- ported from tests/cases/compiler/restParameters.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function f18(a, ...)
end
function f19(a, b, ...)
end
function f20(a, b, ...)
end
function f21(a, b, c, ...)
end
