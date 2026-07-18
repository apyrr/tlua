package checker

import (
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
)

// A Lua table accepts any non-nil value as a key: a boolean, another table, a
// function, a coroutine. TypeScript's index signatures accept only the types that can
// name a *property* -- string, number, symbol -- and that narrower domain is still the
// right one wherever the checker turns a key back into a property name, mapped types
// most of all. The two predicates below keep the domains apart.

// isPropertyNameIndexKeyType reports whether t can name a property. This is
// TypeScript's index-key domain, and it deliberately does not widen to the Lua one.
func (c *Checker) isPropertyNameIndexKeyType(t *Type) bool {
	return t.flags&(TypeFlagsString|TypeFlagsNumber|TypeFlagsESSymbol) != 0 ||
		c.isPatternLiteralType(t) ||
		t.flags&TypeFlagsIntersection != 0 && !c.isGenericType(t) && core.Some(t.Types(), c.isPropertyNameIndexKeyType)
}

// isValidIndexKeyType reports whether t may be declared as the key of an index
// signature. Every type keys a Lua table except nil, so the rule is stated negatively:
// a key must not admit nil, which rules out nil itself, void, and the top types. A
// bare type parameter is accepted *without* consulting its constraint -- Luau's rule --
// and an EXPLICIT argument it later receives is re-validated by the gate in
// checkTypeArgumentConstraints (100048), which checks the *argument* domain below (so
// `any` passes there); the `Table<unknown, unknown>` a failed inference falls back to
// still escapes.
//
// Like the property-name predicate it replaces, this does not distribute over unions;
// callers wrap it in everyType or forEachIndexKeyType.
func (c *Checker) isValidIndexKeyType(t *Type) bool {
	return t.flags&(TypeFlagsNil|TypeFlagsVoid|TypeFlagsAnyOrUnknown|TypeFlagsNever) == 0
}

// isValidIndexArgumentType reports whether t may be the key in an element access. The
// declaration domain, plus `any` and `never`, which are unchecked here as everywhere.
func (c *Checker) isValidIndexArgumentType(t *Type) bool {
	if t.flags&(TypeFlagsAny|TypeFlagsNever) != 0 {
		return true
	}
	return everyType(t, c.isValidIndexKeyType)
}

// isIndexKeyTypeParameter reports whether the parameter keys an index signature
// of its declaring interface or type alias, which subjects the arguments it
// receives to the index-key domain. The answer comes from the declaration AST
// alone: resolving the declared type here would run inside another type's
// constraint check and perturb the depth limiters.
func isIndexKeyTypeParameter(typeParameter *Type) bool {
	symbol := typeParameter.symbol
	if symbol == nil {
		return false
	}
	name := symbol.Name
	for _, declaration := range symbol.Declarations {
		parent := declaration.Parent
		if parent == nil {
			continue
		}
		var members []*ast.Node
		switch {
		case ast.IsInterfaceDeclaration(parent):
			members = parent.Members()
		case ast.IsTypeOrJSTypeAliasDeclaration(parent):
			if target := parent.Type(); target != nil && target.Kind == ast.KindTypeLiteral {
				members = target.Members()
			}
		}
		for _, member := range members {
			if member.Kind != ast.KindIndexSignature {
				continue
			}
			parameters := member.Parameters()
			if len(parameters) != 0 && parameters[0].Type() != nil && typeNodeNamesTypeParameter(parameters[0].Type(), name) {
				return true
			}
		}
	}
	return false
}

// typeNodeNamesTypeParameter reports whether the key type node mentions the
// parameter by name -- directly or as a union constituent. Inside the
// declaring container the bare name can only mean the parameter.
func typeNodeNamesTypeParameter(node *ast.Node, name string) bool {
	switch node.Kind {
	case ast.KindTypeReference:
		typeName := node.AsTypeReferenceNode().TypeName
		return ast.IsIdentifier(typeName) && typeName.Text() == name
	case ast.KindUnionType:
		for _, constituent := range node.AsUnionTypeNode().Types.Nodes {
			if typeNodeNamesTypeParameter(constituent, name) {
				return true
			}
		}
	}
	return false
}

// forEachIndexKeyType calls f once for every key type a declared index signature
// contributes. This is forEachType with one correction: `boolean` is interned as the
// union `false | true`, so forEachType would shred a boolean key into two *literal*
// keys -- which the literal-key rule then rejects, leaving the signature with no index
// at all. When both boolean literals are present -- fresh or regular -- they fold into
// one `boolean`, in the position the first occupies.
func (c *Checker) forEachIndexKeyType(t *Type, f func(keyType *Type)) {
	if t.flags&TypeFlagsUnion == 0 {
		f(t)
		return
	}
	types := t.Types()
	hasFalse := false
	hasTrue := false
	for _, u := range types {
		if u.flags&TypeFlagsBooleanLiteral != 0 {
			if c.getRegularTypeOfLiteralType(u) == c.regularFalseType {
				hasFalse = true
			} else {
				hasTrue = true
			}
		}
	}
	isBoolean := hasFalse && hasTrue
	booleanDone := false
	for _, u := range types {
		if isBoolean && u.flags&TypeFlagsBooleanLiteral != 0 {
			if !booleanDone {
				booleanDone = true
				f(c.booleanType)
			}
			continue
		}
		f(u)
	}
}
