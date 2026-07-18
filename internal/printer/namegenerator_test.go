package printer_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/binder"
	"github.com/apyrr/tlua/internal/printer"
	"github.com/apyrr/tlua/internal/testutil/parsetestutil"
	"gotest.tools/v3/assert"
)

func TestTempVariable1(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewTempVariable()
	name2 := ec.Factory.NewTempVariable()

	g := &printer.NameGenerator{Context: ec}
	text1 := g.GenerateName(name1)
	text2 := g.GenerateName(name2)

	assert.Equal(t, "_a", text1)
	assert.Equal(t, "_b", text2)
}

func TestTempVariable2(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewTempVariableEx(printer.AutoGenerateOptions{
		Prefix: "A",
		Suffix: "B",
	})
	name2 := ec.Factory.NewTempVariableEx(printer.AutoGenerateOptions{
		Prefix: "A",
		Suffix: "B",
	})

	g := &printer.NameGenerator{Context: ec}
	text1 := g.GenerateName(name1)
	text2 := g.GenerateName(name2)

	assert.Equal(t, "A_aB", text1)
	assert.Equal(t, "A_bB", text2)
}

func TestTempVariable3(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewTempVariable()

	g := &printer.NameGenerator{Context: ec}
	text1 := g.GenerateName(name1)
	text2 := g.GenerateName(name1)

	assert.Equal(t, "_a", text1)
	assert.Equal(t, "_a", text2)
}

func TestTempVariableScoped(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewTempVariable()
	name2 := ec.Factory.NewTempVariable()

	g := &printer.NameGenerator{Context: ec}
	text1 := g.GenerateName(name1)
	g.PushScope(false)
	text2 := g.GenerateName(name2)
	g.PopScope(false)

	assert.Equal(t, "_a", text1)
	assert.Equal(t, "_a", text2)
}

func TestTempVariableScopedReserved(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewTempVariableEx(printer.AutoGenerateOptions{Flags: printer.GeneratedIdentifierFlagsReservedInNestedScopes})
	name2 := ec.Factory.NewTempVariable()

	g := &printer.NameGenerator{Context: ec}
	text1 := g.GenerateName(name1)
	g.PushScope(false)
	text2 := g.GenerateName(name2)
	g.PopScope(false)

	assert.Equal(t, "_a", text1)
	assert.Equal(t, "_b", text2)
}

func TestLoopVariable1(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewLoopVariable()
	name2 := ec.Factory.NewLoopVariable()

	g := &printer.NameGenerator{Context: ec}
	text1 := g.GenerateName(name1)
	text2 := g.GenerateName(name2)

	assert.Equal(t, "_i", text1)
	assert.Equal(t, "_a", text2)
}

func TestLoopVariable2(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewLoopVariableEx(printer.AutoGenerateOptions{
		Prefix: "A",
		Suffix: "B",
	})
	name2 := ec.Factory.NewLoopVariableEx(printer.AutoGenerateOptions{
		Prefix: "A",
		Suffix: "B",
	})

	g := &printer.NameGenerator{Context: ec}
	text1 := g.GenerateName(name1)
	text2 := g.GenerateName(name2)

	assert.Equal(t, "A_iB", text1)
	assert.Equal(t, "A_aB", text2)
}

func TestLoopVariable3(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewLoopVariable()

	g := &printer.NameGenerator{Context: ec}
	text1 := g.GenerateName(name1)
	text2 := g.GenerateName(name1)

	assert.Equal(t, "_i", text1)
	assert.Equal(t, "_i", text2)
}

func TestLoopVariableScoped(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewLoopVariable()
	name2 := ec.Factory.NewLoopVariable()

	g := &printer.NameGenerator{Context: ec}
	text1 := g.GenerateName(name1)
	g.PushScope(false)
	text2 := g.GenerateName(name2)
	g.PopScope(false)

	assert.Equal(t, "_i", text1)
	assert.Equal(t, "_i", text2)
}

func TestUniqueName1(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewUniqueName("foo")
	name2 := ec.Factory.NewUniqueName("foo")

	g := &printer.NameGenerator{Context: ec}
	text1 := g.GenerateName(name1)
	text2 := g.GenerateName(name2)

	assert.Equal(t, "foo_1", text1)
	assert.Equal(t, "foo_2", text2)
}

func TestUniqueName2(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewUniqueName("foo")

	g := &printer.NameGenerator{Context: ec}
	text1 := g.GenerateName(name1)
	text2 := g.GenerateName(name1)

	assert.Equal(t, "foo_1", text1)
	// Expected to be same because GenerateName goes off object identity
	assert.Equal(t, "foo_1", text2)
}

func TestUniqueNameScoped(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewUniqueName("foo")
	name2 := ec.Factory.NewUniqueName("foo")

	g := &printer.NameGenerator{Context: ec}
	assert.Equal(t, "foo_1", g.GenerateName(name1))

	g.PushScope(false)
	assert.Equal(t, "foo_2", g.GenerateName(name2)) // Matches Strada, but is incorrect
	// assert.Equal(t, "foo_1", g.GenerateName(name2)) // TODO: Fix after Strada port is complete.
	g.PopScope(false)
}

func TestUniquePrivateName1(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewUniquePrivateName("#foo")
	name2 := ec.Factory.NewUniquePrivateName("#foo")

	g := &printer.NameGenerator{Context: ec}
	text1 := g.GenerateName(name1)
	text2 := g.GenerateName(name2)

	assert.Equal(t, "#foo_1", text1)
	assert.Equal(t, "#foo_2", text2)
}

func TestUniquePrivateName2(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewUniquePrivateName("#foo")

	g := &printer.NameGenerator{Context: ec}
	text1 := g.GenerateName(name1)
	text2 := g.GenerateName(name1)

	assert.Equal(t, "#foo_1", text1)
	assert.Equal(t, "#foo_1", text2)
}

func TestUniquePrivateNameScoped(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()
	name1 := ec.Factory.NewUniquePrivateName("#foo")
	name2 := ec.Factory.NewUniquePrivateName("#foo")

	g := &printer.NameGenerator{Context: ec}
	assert.Equal(t, "#foo_1", g.GenerateName(name1))

	g.PushScope(false) // private names are always reserved in nested scopes
	assert.Equal(t, "#foo_2", g.GenerateName(name2))
	g.PopScope(false)
}

func TestGeneratedNameForIdentifier1(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()

	file := parsetestutil.ParseTypeScript("function f() {}", false /*jsx*/)
	binder.BindSourceFile(file)

	n := file.Statements.Nodes[0].Name()
	name1 := ec.Factory.NewGeneratedNameForNode(n)

	g := &printer.NameGenerator{Context: ec, GetTextOfNode: (*ast.Node).Text}
	text1 := g.GenerateName(name1)

	assert.Equal(t, "f_1", text1)
}

func TestGeneratedNameForIdentifier2(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()

	file := parsetestutil.ParseTypeScript("function f() {}", false /*jsx*/)
	binder.BindSourceFile(file)

	n := file.Statements.Nodes[0].Name()
	name1 := ec.Factory.NewGeneratedNameForNodeEx(n, printer.AutoGenerateOptions{
		Prefix: "a",
		Suffix: "b",
	})

	g := &printer.NameGenerator{Context: ec, GetTextOfNode: (*ast.Node).Text}
	text1 := g.GenerateName(name1)

	assert.Equal(t, "afb", text1)
}

func TestGeneratedNameForIdentifier3(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()

	file := parsetestutil.ParseTypeScript("function f() {}", false /*jsx*/)
	binder.BindSourceFile(file)

	n := file.Statements.Nodes[0].Name()
	name1 := ec.Factory.NewGeneratedNameForNodeEx(n, printer.AutoGenerateOptions{
		Prefix: "a",
		Suffix: "b",
	})
	name2 := ec.Factory.NewGeneratedNameForNode(name1)

	g := &printer.NameGenerator{Context: ec, GetTextOfNode: (*ast.Node).Text}
	text1 := g.GenerateName(name2)

	assert.Equal(t, "afb_1", text1)
}

// namespace reuses name if it does not collide with locals
func TestGeneratedNameForNodeCached(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()

	file := parsetestutil.ParseTypeScript("function foo() { local foo; }", false /*jsx*/)
	binder.BindSourceFile(file)

	ns1 := file.Statements.Nodes[0]
	name1 := ec.Factory.NewGeneratedNameForNode(ns1)
	name2 := ec.Factory.NewGeneratedNameForNode(ns1)

	g := &printer.NameGenerator{Context: ec, GetTextOfNode: (*ast.Node).Text}
	text1 := g.GenerateName(name1)
	text2 := g.GenerateName(name2)

	assert.Equal(t, "foo_1", text1)
	assert.Equal(t, "foo_1", text2)
}

// newSyntheticMethod builds a MethodDeclaration via the factory: object-literal
// methods no longer parse in tlua, but transforms still synthesize
// method/accessor members, so the generator path stays covered.
func newSyntheticMethod(ec *printer.EmitContext, name string) *ast.Node {
	return ec.Factory.NewMethodDeclaration(
		nil, /*modifiers*/
		ec.Factory.NewIdentifier(name),
		nil, /*postfixToken*/
		nil, /*typeParameters*/
		ec.Factory.NewNodeList(nil),
		nil, /*returnType*/
		nil, /*fullSignature*/
		ec.Factory.NewBlock(ec.Factory.NewNodeList(nil), false /*multiline*/),
	)
}

func TestGeneratedNameForMethod1(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()

	n := newSyntheticMethod(ec, "m")
	name1 := ec.Factory.NewGeneratedNameForNode(n)

	g := &printer.NameGenerator{Context: ec, GetTextOfNode: (*ast.Node).Text}
	text1 := g.GenerateName(name1)

	assert.Equal(t, "m_1", text1)
}

func TestGeneratedNameForMethod2(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()

	// method with a non-identifier (numeric) name falls back to a temp name
	n := ec.Factory.NewMethodDeclaration(
		nil, /*modifiers*/
		ec.Factory.NewNumericLiteral("0", ast.TokenFlagsNone),
		nil, /*postfixToken*/
		nil, /*typeParameters*/
		ec.Factory.NewNodeList(nil),
		nil, /*returnType*/
		nil, /*fullSignature*/
		ec.Factory.NewBlock(ec.Factory.NewNodeList(nil), false /*multiline*/),
	)
	name1 := ec.Factory.NewGeneratedNameForNode(n)

	g := &printer.NameGenerator{Context: ec, GetTextOfNode: (*ast.Node).Text}
	text1 := g.GenerateName(name1)

	assert.Equal(t, "_a", text1)
}

func TestGeneratedPrivateNameForMethod(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()

	n := newSyntheticMethod(ec, "m")
	name1 := ec.Factory.NewGeneratedPrivateNameForNode(n)

	g := &printer.NameGenerator{Context: ec, GetTextOfNode: (*ast.Node).Text}
	text1 := g.GenerateName(name1)

	assert.Equal(t, "#m_1", text1)
}

func TestGeneratedNameForComputedPropertyName(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()

	// object-literal computed property name (KindComputedPropertyName)
	file := parsetestutil.ParseTypeScript("({ [x]: 0 })", false /*jsx*/)
	binder.BindSourceFile(file)

	n := file.Statements.Nodes[0].Expression().Expression().AsObjectLiteralExpression().Properties.Nodes[0].Name()
	name1 := ec.Factory.NewGeneratedNameForNode(n)

	g := &printer.NameGenerator{Context: ec, GetTextOfNode: (*ast.Node).Text}
	text1 := g.GenerateName(name1)

	assert.Equal(t, "_a", text1)
}

func TestGeneratedNameForOther(t *testing.T) {
	t.Parallel()

	ec := printer.NewEmitContext()

	file := parsetestutil.ParseTypeScript("interface C {}", false /*jsx*/)
	binder.BindSourceFile(file)

	n := ec.Factory.NewObjectLiteralExpression(
		ec.Factory.NewNodeList([]*ast.Node{}),
		false, /*multiLine*/
	)
	name1 := ec.Factory.NewGeneratedNameForNode(n)

	g := &printer.NameGenerator{Context: ec, GetTextOfNode: (*ast.Node).Text}
	text1 := g.GenerateName(name1)

	assert.Equal(t, "_a", text1)
}
