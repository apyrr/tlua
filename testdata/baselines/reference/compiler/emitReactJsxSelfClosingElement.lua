//// [tests/cases/compiler/emitReactJsxSelfClosingElement.tsx] ////

//// [a.tsx]
local app = <App />;


//// [a.lua]
import { jsx as _jsx } from "react/jsx-runtime";
local app = _jsx(App, {});
