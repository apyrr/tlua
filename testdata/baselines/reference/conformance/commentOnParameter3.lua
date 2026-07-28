//// [tests/cases/conformance/ported/commentOnParameter3.tlua] ////

//// [commentOnParameter3.tlua]
-- ported from tests/cases/compiler/commentOnParameter3.ts
-- dropped: @strict: false; explicit any annotations avoid unrelated strict-mode diagnostics.
-- note: upstream emission also drops the final standalone parameter comment; retained as printer parity coverage.


function commentedParameters(
    a: any --[[ parameter a ]],
    b: any --[[ parameter b ]],
    --[[ extra comment ]]
)
end


//// [commentOnParameter3.lua]
-- ported from tests/cases/compiler/commentOnParameter3.ts
-- dropped: @strict: false; explicit any annotations avoid unrelated strict-mode diagnostics.
-- note: upstream emission also drops the final standalone parameter comment; retained as printer parity coverage.
function commentedParameters(a --[[ parameter a ]], b --[[ parameter b ]])
end
