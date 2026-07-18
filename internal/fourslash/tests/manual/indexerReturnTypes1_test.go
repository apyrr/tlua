package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestIndexerReturnTypes1(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Numeric {
    [x: number]: Date;
}
}
interface Stringy {
    [x: string]: RegExp;
}
}
interface NumericPlus {
    [x: number]: Date;
    foo(): Date;
}
}
interface StringyPlus {
    [x: string]: RegExp;
    foo(): RegExp;
}
}
interface NumericG<T> {
    [x: number]: T;
}
}
interface StringyG<T> {
    [x: string]: T;
}
}
interface Ty<T> {
    [x: number]: Ty<T>;
}
interface Ty2<T> {
    [x: number]: { [x: number]: T };
}


}
local numeric: Numeric;
local stringy: Stringy;
local numericPlus: NumericPlus;
local stringPlus: StringyPlus;
local numericG: NumericG<Date>;
local stringyG: StringyG<Date>;
local ty: Ty<Date>;
local ty2: Ty2<Date>;

local /*1*/r1 = numeric[1];
local /*2*/r2 = numeric['1'];
local /*3*/r3 = stringy[1];
local /*4*/r4 = stringy['1'];
local /*5*/r5 = numericPlus[1];
local /*6*/r6 = numericPlus['1'];
local /*7*/r7 = stringPlus[1];
local /*8*/r8 = stringPlus['1'];
local /*9*/r9 = numericG[1];
local /*10*/r10 = numericG['1'];
local /*11*/r11 = stringyG[1];
local /*12*/r12 = stringyG['1'];
local /*13*/r13 = ty[1];
local /*14*/r14 = ty['1'];
local /*15*/r15 = ty2[1];
local /*16*/r16 = ty2['1'];`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	// tlua: number and string keys are disjoint, so a string-key access on a
	// number-only index (and vice versa) is not covered and resolves to `any`.
	f.VerifyQuickInfoAt(t, "1", "local r1: Date", "")
	f.VerifyQuickInfoAt(t, "2", "local r2: any", "")
	f.VerifyQuickInfoAt(t, "3", "local r3: any", "")
	f.VerifyQuickInfoAt(t, "4", "local r4: RegExp", "")
	f.VerifyQuickInfoAt(t, "5", "local r5: Date", "")
	f.VerifyQuickInfoAt(t, "6", "local r6: any", "")
	f.VerifyQuickInfoAt(t, "7", "local r7: any", "")
	f.VerifyQuickInfoAt(t, "8", "local r8: RegExp", "")
	f.VerifyQuickInfoAt(t, "9", "local r9: Date", "")
	f.VerifyQuickInfoAt(t, "10", "local r10: any", "")
	f.VerifyQuickInfoAt(t, "11", "local r11: any", "")
	f.VerifyQuickInfoAt(t, "12", "local r12: Date", "")
	f.VerifyQuickInfoAt(t, "13", "local r13: Ty<Date>", "")
	f.VerifyQuickInfoAt(t, "14", "local r14: any", "")
	f.VerifyQuickInfoAt(t, "15", "local r15: {\n    [x: number]: Date;\n}", "")
	f.VerifyQuickInfoAt(t, "16", "local r16: any", "")
}
