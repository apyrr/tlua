package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestDocumentSymbolPrivateName(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: first.tlua
class A {
  #foo() {
    class B {
      #bar() {   
         function baz () {
         }
      }
    }
  }
}

class B {
	constructor(private prop: string) {}
}

// @Filename: second.tlua
class Foo {
	#privateProp: string;
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentSymbol(t)
	f.GoToFile(t, "second.tlua")
	f.VerifyBaselineDocumentSymbol(t)
}
