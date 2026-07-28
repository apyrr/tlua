//// [tests/cases/conformance/ported/commentOnParameter1.tlua] ////

//// [commentOnParameter1.tlua]
-- ported from tests/cases/compiler/commentOnParameter1.ts
-- dropped: @strict: false; tlua test defaults keep strict mode enabled, so parameters are explicitly any to avoid unrelated errors.


function commentedParameters(
--[[ Parameter a ]]
a: any
--[[ End of parameter a ]]
--[[ Parameter b ]]
,
b: any
--[[ End of parameter b ]]
)
end


//// [commentOnParameter1.lua]
-- ported from tests/cases/compiler/commentOnParameter1.ts
-- dropped: @strict: false; tlua test defaults keep strict mode enabled, so parameters are explicitly any to avoid unrelated errors.
function commentedParameters(
--[[ Parameter a ]]
a
--[[ End of parameter a ]]
--[[ Parameter b ]]
, b
--[[ End of parameter b ]]
)
end
