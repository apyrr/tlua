//// [tests/cases/compiler/suspendJsxArgumentsAttributeName.tlua] ////

//// [test.tsx]
  interface JsxIntrinsicElements { div: any; }

suspend function f()
  return <div arguments={42} />;
end


//// [test.jsx]
function f()
  return <div arguments={42}/>;
end
