//// [tests/cases/compiler/jsxSpreadWithAssertion.tsx] ////

//// [jsxSpreadWithAssertion.tsx]
declare EntryTextDialog: any;

local c = <EntryTextDialog
    {...{
        first: 0,
        foo: 1,
        bar: 2 as any,
        baz: 3,
        last: 4
    }}
/>


//// [jsxSpreadWithAssertion.lua]
import { jsx as _jsx } from "react/jsx-runtime";
local c = _jsx(EntryTextDialog, {});
{
    ...;
    {
        first;
        0,
            foo;
        1,
            bar;
        2,
            baz;
        3,
            last;
        4;
    }
}
/>;
