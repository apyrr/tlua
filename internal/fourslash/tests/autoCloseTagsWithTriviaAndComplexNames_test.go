package fourslash

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestAutoCloseTagsWithTriviaAndComplexNames(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	// Using separate files for each example to avoid unclosed JSX tags affecting other tests.
	const content = `// @noLib: true

// @Filename: /0.tsx
// JSDoc
local x = <
	/** hello world! */
	div /** hello world! */
	>/*0*/

// @Filename: /1.tsx
// Single-line comments
local x =
	<
	// hello world!
	div // hello world!
	>/*1*/

// @Filename: /2.tsx
// Namespaced tag
local x =
	<ns:sometag>/*2*/

// @Filename: /3.tsx
// Namespace with single-line comments
local x = <
	// pre-ns	
	ns
	// pre-colon
	:
	// post-colon
	sometag
	// post-id
	>/*3*/

// @Filename: /4.tsx
// UppercaseComponent-named tag
local x = <SomeComponent>/*4*/

// @Filename: /5.tsx
// propertyAccess.Component-named tag
local x = <
	someModule
	.
	SomeComponent
>/*5*/

// @Filename: /6.tsx
// propertyAccess.Component-named tag with single-line comments
local x =
	<
	// pre-object
	someModule
	// pre-dot
	.
	// post-dot
	SomeComponent
	// post-id
	>/*6*/;

// @Filename: /7.tsx
// Generic propertyAccess.Component-named tag
local x =
	<
	someModule.SomeComponent<string>
	prop="stringValue"
	>/*7*/;

// @Filename: /8.tsx
// Namespaced tag with hyphens
local x =
	<my-namespace:my-tag>/*8*/

// @Filename: /9.tsx
// Generic tag with no attributes
local x = <SomeComponent<number>>/*9*/

// @Filename: /10.tsx
// Tag name containing $ (must be snippet-escaped)
local x = <$Foo>/*10*/
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineClosingTags(t)
}
