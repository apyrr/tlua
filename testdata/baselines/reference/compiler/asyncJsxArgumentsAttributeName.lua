//// [tests/cases/compiler/asyncJsxArgumentsAttributeName.tlua] ////

//// [test.tsx]
  interface JsxIntrinsicElements { div: any; }

async function f()
  return <div arguments={42} />;
end


//// [test.jsx]
async function f()
    return <div arguments={42}/>;
end
