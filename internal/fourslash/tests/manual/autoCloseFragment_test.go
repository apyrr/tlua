package fourslash

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestAutoCloseFragment(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	// Using separate files for each example to avoid unclosed JSX tags affecting other tests.
	const content = `// @noLib: true
// @Filename: /0.tsx
local x = <>/*0*/;

// @Filename: /1.tsx
local x = <> foo/*1*/ </>;

// @Filename: /2.tsx
local x = <></>/*2*/;

// @Filename: /3.tsx
local x = </>/*3*/;

// @Filename: /4.tsx
local x = <div>
    <>/*4*/
    </div>
</>;

// @Filename: /5.tsx
local x = <> text /*5*/;

// @Filename: /6.tsx
local x = <>
    <>/*6*/
</>;

// @Filename: /7.tsx
local x = <div>
    <>/*7*/
</div>;

// @Filename: /8.tsx
local x = <div>
    <>/*8*/</>
</div>;

// @Filename: /9.tsx
local x = <p>
    <>
        <>/*9*/
    </>
</p>`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyJsxClosingTag(t, map[string]*string{
		"0": new("</>"),
		"1": nil,
		"2": nil,
		"3": nil,
		"4": new("</>"),
		"5": new("</>"),
		"6": new("</>"),
		"7": new("</>"),
		"8": nil,
		"9": new("</>"),
	})
}
