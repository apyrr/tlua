package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestLinkedEditingJsxTag12(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /incomplete.tsx
function Test()
    return <div>
        </*0*/
        <div {...{}}>
        </div>
    </div>
end
// @Filename: /incompleteMismatched.tsx
function Test()
    return <div>
        <T
        <div {...{}}>
        </div>
    </div>
end
// @Filename: /incompleteMismatched2.tsx
function Test()
    return <div>
        <T
        <div {...{}}>
        T</div>
    </div>
end
// @Filename: /incompleteMismatched3.tsx
function Test()
    return <div>
        <div {...{}}>
        </div>
        <T
    </div>
end
// @Filename: /mismatched.tsx
function Test()
    return <div>
        <T>
        <div {...{}}>
        </div>
    </div>
end
// @Filename: /matched.tsx
function Test()
    return <div>

        <div {...{}}>
        </div>
    </div>
end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyLinkedEditing(t, map[string][]lsproto.Range{"0": nil})
	f.VerifyBaselineLinkedEditing(t)
}
