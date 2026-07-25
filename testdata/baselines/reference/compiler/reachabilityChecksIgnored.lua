//// [tests/cases/compiler/reachabilityChecksIgnored.tlua] ////

//// [reachabilityChecksIgnored.tlua]
function a()
    error("");

    // @ts-ignore
    console.log("unreachable");
end

function b()
    error("");

    // @ts-expect-error
    console.log("unreachable");
end


//// [reachabilityChecksIgnored.lua]
function a()
  error("");
  -- @ts-ignore
  console.log("unreachable");
end
function b()
  error("");
  -- @ts-expect-error
  console.log("unreachable");
end
