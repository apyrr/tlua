package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestCallHierarchyIncomingCallsNoCrashArrayPush(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function splitNames(name: string)
  return (name || "").split(",").filter(Boolean)
end

function trim(packageNames: string[])
  local nameOrPkgs = packageNames.filter(Boolean)
  local names = {}
  for _, nameOrPkg in ipairs(nameOrPkgs) do
    names./*push*/push(nameOrPkg)
  end
  return names
end
	`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "push")
	f.VerifyBaselineCallHierarchy(t)
}
