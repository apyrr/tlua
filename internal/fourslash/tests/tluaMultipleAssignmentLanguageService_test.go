package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestTluaMultipleAssignmentGoToDefinition(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `Root = {};
Root.x, side, Root["named"], Root[1] = 1, 2, 3, 4;
Repeated, Repeated = 5, "overwritten";
AssertedRoot = {};
AssertedRoot!.value, assertedMemberSide = 7, 0;
ReceiverRoot = {};
(ReceiverRoot).value, receiverMemberSide = "receiver", 0;
local member = Root./*memberUse*/x;
local global = /*globalUse*/side;
local named = Root[/*namedUse*/"named"];
local numeric = Root[/*numericUse*/1];
local repeated = /*repeatedUse*/Repeated;
local assertedMember = AssertedRoot./*assertedMemberUse*/value;
local receiverMember = ReceiverRoot./*receiverMemberUse*/value;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineGoToDefinition(t, false /*includeOriginalSelectionRange*/, "memberUse", "globalUse", "namedUse", "numericUse", "repeatedUse", "assertedMemberUse", "receiverMemberUse")
}

func TestTluaMultipleAssignmentFindAllReferences(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `Root = {};
Root.[|x|], [|side|] = 1, 2;
[|Repeated|], [|Repeated|] = 5, "overwritten";
local member = Root.[|x|];
local global = [|side|];
local repeated = [|Repeated|];`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineFindAllReferences(t)
}

func TestTluaMultipleAssignmentFindAllReferencesWithoutDeclarations(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `Root = {};
Root.x, side = 1, 2;
Repeated, Repeated = 5, "overwritten";
local member = Root./*memberUse*/x;
local global = /*globalUse*/side;
local repeated = /*repeatedUse*/Repeated;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineFindAllReferencesWithoutDeclarations(t, "memberUse", "globalUse", "repeatedUse")
}

func TestTluaMultipleAssignmentQuickInfoDocumentation(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/** Global target docs. */
DocumentedGlobal, documentedSide = 1, 0;
DocumentedRoot = {};
/** Member target docs. */
DocumentedRoot.value, documentedMemberSide = "value", 0;
local globalUse = DocumentedGlobal/*globalUse*/;
local memberUse = DocumentedRoot.value/*memberUse*/;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "globalUse", "local DocumentedGlobal: number", "Global target docs.")
	f.VerifyQuickInfoAt(t, "memberUse", "(property) value: string", "Member target docs.")
}
