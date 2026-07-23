//// [tests/cases/compiler/reachabilityChecksIgnored.tlua] ////

//// [reachabilityChecksIgnored.tlua]
function a()
    throw new Error("");

    // @ts-ignore
    console.log("unreachable");
end

function b()
    throw new Error("");

    // @ts-expect-error
    console.log("unreachable");
end


//// [reachabilityChecksIgnored.lua]
function a()
    throw new Error("");
    -- @ts-ignore
    console.log("unreachable");
end
function b()
    throw new Error("");
    -- @ts-expect-error
    console.log("unreachable");
end
