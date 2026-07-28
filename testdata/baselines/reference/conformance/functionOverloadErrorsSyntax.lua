//// [tests/cases/conformance/ported/functionOverloadErrorsSyntax.tlua] ////

//// [functionOverloadErrorsSyntax.tlua]
-- ported from tests/cases/conformance/functions/functionOverloadErrorsSyntax.ts
-- dropped: nothing essential; TS's bodyless overload-signature list survives as
-- tlua's semicolon-terminated overload signatures, which (unlike `declare
-- function`) still require a final implementation, so a trivial empty body was
-- added to each overload group to keep the parse legal.


-- Function overload signature with optional parameter followed by non-optional parameter
function fn4a(x?: number, y: string);
function fn4a()
end

function fn4b(n: string, x?: number, y: string);
function fn4b()
end

-- Function overload signature with rest param followed by non-optional parameter
function fn5(x: string, ...: any, z: string);
function fn5()
end


//// [functionOverloadErrorsSyntax.lua]
-- ported from tests/cases/conformance/functions/functionOverloadErrorsSyntax.ts
-- dropped: nothing essential; TS's bodyless overload-signature list survives as
-- tlua's semicolon-terminated overload signatures, which (unlike `declare
-- function`) still require a final implementation, so a trivial empty body was
-- added to each overload group to keep the parse legal.
function fn4a()
end
function fn4b()
end
function fn5()
end
