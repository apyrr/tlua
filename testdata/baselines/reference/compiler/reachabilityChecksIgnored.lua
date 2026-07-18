//// [tests/cases/compiler/reachabilityChecksIgnored.tlua] ////

//// [reachabilityChecksIgnored.tlua]
function a() {
    throw new Error("");

    // @ts-ignore
    console.log("unreachable");
}

function b() {
    throw new Error("");

    // @ts-expect-error
    console.log("unreachable");
}


//// [reachabilityChecksIgnored.lua]
function a() {
    throw new Error("");
    -- @ts-ignore
    console.log("unreachable");
}
function b() {
    throw new Error("");
    -- @ts-expect-error
    console.log("unreachable");
}
