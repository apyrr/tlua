//// [tests/cases/compiler/reachabilityChecksIgnored.tlua] ////

//// [reachabilityChecksIgnored.tlua]
function a()
    throw "";

    // @ts-ignore
    console.log("unreachable");
end

function b()
    throw "";

    // @ts-expect-error
    console.log("unreachable");
end


//// [reachabilityChecksIgnored.lua]
function a()
  throw "";
  -- @ts-ignore
  console.log("unreachable");
end
function b()
  throw "";
  -- @ts-expect-error
  console.log("unreachable");
end
