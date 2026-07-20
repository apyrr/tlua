package checker

import (
	"slices"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/diagnostics"
)

// The typed metatable protocol. A Lua table's failed reads fall through to its metatable's
// __index, so setmetatable(t, mt) does not return the table it was given: it returns a table
// that also carries __index's shape. The result is a MetatableType, the pair of the two
// operands, and getmetatable reads the metatable back out of it.
//
// The pairing is only as precise as the metatable's type. A metatable whose type says no
// more than `LuaMetatable<T> | nil` -- which is exactly what the parameter declares -- cannot
// say which fallback applies, so such a call returns the plain table rather than a wrong
// augmentation. Precision comes from the metatable literal a call passes, or from a variable
// that has one as its type.

// MetatableTypeKey interns a setmetatable result. The pair of operands determines the result:
// everything else the type stores is derived from the metatable.
type MetatableTypeKey struct {
	tableID     TypeId
	metatableID TypeId
}

type luaMetatableCallKind int

const (
	luaMetatableCallNone luaMetatableCallKind = iota
	luaMetatableCallSet
	luaMetatableCallGet
	luaMetatableCallDebugSet
	luaMetatableCallDebugGet
)

func (k luaMetatableCallKind) isSet() bool {
	return k == luaMetatableCallSet || k == luaMetatableCallDebugSet
}

func (k luaMetatableCallKind) isGet() bool {
	return k == luaMetatableCallGet || k == luaMetatableCallDebugGet
}

// getLuaMetatableCall reports which metatable global a call invokes. Only the real globals carry
// the protocol: a shadowing local, or a call through an alias, resolves to a different symbol and
// keeps whatever it declares, as it does for the type() narrowing builtins. The debug library's
// setmetatable and getmetatable are recognized as member calls, the way io.type is: the callee's
// base must be the debug global.
func (c *Checker) getLuaMetatableCall(node *ast.Node) luaMetatableCallKind {
	if !ast.IsCallExpression(node) {
		return luaMetatableCallNone
	}
	callee := ast.SkipParentheses(node.Expression())
	if ast.IsIdentifier(callee) {
		switch callee.Text() {
		case "setmetatable":
			if c.isLuaGlobalReference(callee, c.getLuaSetmetatableGlobalSymbol()) {
				return luaMetatableCallSet
			}
		case "getmetatable":
			if c.isLuaGlobalReference(callee, c.getLuaGetmetatableGlobalSymbol()) {
				return luaMetatableCallGet
			}
		}
		return luaMetatableCallNone
	}
	if ast.IsPropertyAccessExpression(callee) && c.isLuaGlobalReference(ast.SkipParentheses(callee.Expression()), c.getLuaDebugGlobalSymbol()) {
		switch callee.Name().Text() {
		case "setmetatable":
			return luaMetatableCallDebugSet
		case "getmetatable":
			return luaMetatableCallDebugGet
		}
	}
	return luaMetatableCallNone
}

// isLuaMetatableCall reports whether the metatable protocol interprets this call. The quick
// path of getQuickTypeOfExpression must skip such a call: it reads the signature's return type
// without running checkCallExpression's special cases.
func (c *Checker) isLuaMetatableCall(node *ast.Node) bool {
	return c.getLuaMetatableCall(node) != luaMetatableCallNone
}

// checkLuaMetatableCall returns the type of a setmetatable or getmetatable call, or nil when
// the call's declared return type stands.
func (c *Checker) checkLuaMetatableCall(node *ast.Node, returnType *Type) *Type {
	args := node.Arguments()
	kind := c.getLuaMetatableCall(node)
	switch {
	case kind.isSet():
		if len(args) < 2 {
			return nil
		}
		// A protected table refuses a new metatable: the runtime raises, so the old pairing
		// stands unchanged -- the one place setmetatable does not replace. debug.setmetatable
		// bypasses the protection.
		if kind == luaMetatableCallSet && someType(returnType, c.isProtectedMetatableType) {
			c.error(node, diagnostics.Cannot_change_a_protected_metatable)
			return returnType
		}
		// The table is the signature's return type: setmetatable is declared to return its
		// first parameter, so inference has already folded and widened a literal argument. The
		// metatable is read from the argument itself -- its declared parameter type is the
		// imprecise union by construction -- and widened, so a fresh literal never reaches a
		// type that outlives the expression.
		return c.getSetmetatableResultType(returnType, c.getLuaMetatableArgumentType(args[1]))
	case kind.isGet():
		if len(args) < 1 {
			return nil
		}
		return c.getGetmetatableResultType(c.checkExpressionCached(args[0]), kind == luaMetatableCallDebugGet)
	}
	return nil
}

// getLuaMetatableArgumentType is the metatable operand as the pairing sees it: read from the
// argument -- the declared parameter type is the imprecise union by construction -- and widened,
// so a fresh literal never reaches a type that outlives the expression.
func (c *Checker) getLuaMetatableArgumentType(arg *ast.Node) *Type {
	return c.getWidenedType(c.getRegularTypeOfObjectLiteral(c.checkExpressionCached(arg)))
}

// getGetmetatableResultType returns the metatable paired with t, or nil when t carries no
// pairing the checker knows about and the declared return type stands. A protected pairing shows
// its __metatable sentinel instead of the real metatable, unless the caller bypasses it as
// debug.getmetatable does.
func (c *Checker) getGetmetatableResultType(t *Type, bypassProtection bool) *Type {
	if t.flags&TypeFlagsUnion != 0 {
		metatables := make([]*Type, 0, len(t.Types()))
		for _, u := range t.Types() {
			metatable := c.getGetmetatableResultType(u, bypassProtection)
			if metatable == nil {
				// One arm with no known metatable makes the whole union unknown: the value may
				// be that arm, and no answer covers both.
				return nil
			}
			metatables = append(metatables, metatable)
		}
		return c.getUnionType(metatables)
	}
	if isMetatableType(t) {
		d := t.AsMetatableType()
		if !bypassProtection && d.protectedType != nil {
			// __metatable is what a protected table shows in place of its metatable.
			return d.protectedType
		}
		return d.metatableType
	}
	return nil
}

func isMetatableType(t *Type) bool {
	return t.flags&TypeFlagsObject != 0 && t.objectFlags&ObjectFlagsMetatable != 0
}

func (c *Checker) isProtectedMetatableType(t *Type) bool {
	return isMetatableType(t) && t.AsMetatableType().protectedType != nil
}

// getUnpairedTableType strips a pairing from a table: what setmetatable returned before is the
// table it was given, and the metatable it was given is gone.
func (c *Checker) getUnpairedTableType(t *Type) *Type {
	return c.mapType(t, func(u *Type) *Type {
		for isMetatableType(u) {
			u = u.AsMetatableType().tableType
		}
		return u
	})
}

// luaOperatorMetamethods are the operator metamethods the checker dispatches. __tostring, __gc
// and __mode never change a value's type, so none of the three is worth a pairing on its own.
var luaOperatorMetamethods = []string{
	"__eq", "__lt", "__le", "__unm", "__len",
	"__add", "__sub", "__mul", "__div", "__mod", "__pow", "__concat",
}

// getMetatableHandlerType returns the committed, nil-stripped type of a metamethod, or nil. A
// declared-optional member commits to nothing: LuaMetatable<T> declares every metamethod
// optionally, so reading through its optionality would make every LuaMetatable<T>-typed argument
// callable and protected. A metatable literal's members are not optional, so the idiomatic path
// is unaffected. __index is the exception -- its optional form types the defaults idiom -- and
// reads through optionality on its own path.
func (c *Checker) getMetatableHandlerType(metatableType *Type, name string) *Type {
	// Metamethods never live on the global Object/Function augmentations, so the lookup skips
	// them -- and a merged `interface Object { __add: ... }` cannot inject an operator everywhere.
	prop := c.getPropertyOfTypeEx(metatableType, name, true /*skipObjectFunctionPropertyAugment*/, false /*includeTypeOnlyMembers*/)
	if prop == nil || prop.Flags&ast.SymbolFlagsOptional != 0 {
		return nil
	}
	t := c.removeMissingOrUndefinedType(c.getTypeOfSymbol(prop))
	if c.isErrorType(t) {
		return nil
	}
	return t
}

// getMetatableSource returns a metamethod that is a table or a function -- __index and
// __newindex both take that shape -- along with which one it is, or nil when the metamethod is
// absent, ambiguous, or not table-or-function shaped. A callable table is dropped as ambiguous:
// which half of it answers a lookup is not worth guessing at. __index reads through optionality
// (the defaults idiom); every other metamethod does not.
func (c *Checker) getMetatableSource(metatableType *Type, name string, allowOptional bool) (*Type, bool) {
	prop := c.getPropertyOfTypeEx(metatableType, name, true /*skipObjectFunctionPropertyAugment*/, false /*includeTypeOnlyMembers*/)
	if prop == nil || !allowOptional && prop.Flags&ast.SymbolFlagsOptional != 0 {
		return nil, false
	}
	source := c.removeMissingOrUndefinedType(c.getTypeOfSymbol(prop))
	if c.isErrorType(source) || source.flags&(TypeFlagsObject|TypeFlagsIntersection|TypeFlagsInstantiableNonPrimitive) == 0 {
		return nil, false
	}
	isFunction := len(c.getSignaturesOfType(source, SignatureKindCall)) != 0
	if isFunction && (len(c.getPropertiesOfType(source)) != 0 || len(c.getIndexInfosOfType(source)) != 0) {
		return nil, false
	}
	return source, isFunction
}

// getMetatableCallSource returns __call when it commits to a callable, or nil. Its signatures
// make the pairing callable.
func (c *Checker) getMetatableCallSource(metatableType *Type) *Type {
	handler := c.getMetatableHandlerType(metatableType, "__call")
	if handler == nil || len(c.getSignaturesOfType(handler, SignatureKindCall)) == 0 {
		return nil
	}
	return handler
}

// metatableHasOperatorMetamethod reports whether the metatable commits to any operator
// metamethod: a committed, callable handler is what an operator can dispatch to.
func (c *Checker) metatableHasOperatorMetamethod(metatableType *Type) bool {
	return core.Some(luaOperatorMetamethods, func(name string) bool {
		handler := c.getMetatableHandlerType(metatableType, name)
		return handler != nil && len(c.getSignaturesOfType(handler, SignatureKindCall)) != 0
	})
}

// getSetmetatableResultType pairs a table with the metatable it is given. setmetatable
// *replaces* a metatable, so the table is unpaired first: whatever its old metamethods answered,
// they do not answer now. The new pairing forms only when the metatable commits to a metamethod
// the checker interprets; otherwise the unpaired table stands, which is also how
// `setmetatable(t, nil)` detaches one. Returns nil only when the call is already in error and its
// declared type should stand.
func (c *Checker) getSetmetatableResultType(tableType *Type, metatableType *Type) *Type {
	if c.isErrorType(tableType) {
		return nil
	}
	tableType = c.getUnpairedTableType(tableType)
	if c.isErrorType(metatableType) || !c.canCarryMetatableMembers(metatableType) {
		return tableType
	}
	// __index is declared optional, so a nil arm says only that there may be no fallback at all.
	// What remains has to name a single fallback for the read augmentation to mean anything.
	indexSource, indexIsFunction := c.getMetatableSource(metatableType, "__index", true /*allowOptional*/)
	if c.isGenericObjectType(tableType) || indexSource != nil && c.isGenericObjectType(indexSource) {
		// A pairing resolves its members eagerly, so it can only hold concrete types. A generic
		// one falls back to the intersection of table and __index, as a generic spread does:
		// less precise about collisions, blind to every other metamethod, but it instantiates.
		// The imprecision cuts the other way too: as a plain intersection it exposes the __index
		// source's metamethod-named members to ambient operator dispatch, which a real pairing
		// refuses (Lua does not inherit metamethods through __index). Accepted with the rest.
		if indexSource == nil || indexIsFunction {
			return tableType
		}
		return c.getIntersectionType([]*Type{tableType, indexSource})
	}
	if !c.canCarryMetatableMembers(tableType) {
		return tableType
	}
	newindexSource, newindexIsFunction := c.getMetatableSource(metatableType, "__newindex", false /*allowOptional*/)
	callSource := c.getMetatableCallSource(metatableType)
	protectedType := c.getMetatableHandlerType(metatableType, "__metatable")
	if indexSource == nil && newindexSource == nil && callSource == nil && protectedType == nil &&
		!c.metatableHasOperatorMetamethod(metatableType) {
		// The metatable commits to nothing the checker interprets, so the plain table stands.
		return tableType
	}
	key := MetatableTypeKey{tableID: tableType.id, metatableID: metatableType.id}
	if cached := c.metatableTypes[key]; cached != nil {
		return cached
	}
	t := c.newObjectType(ObjectFlagsMetatable, nil /*symbol*/)
	// The merged members come from both operands, so a non-inferrable one makes the pair
	// non-inferrable too.
	t.objectFlags |= (tableType.objectFlags | metatableType.objectFlags) & ObjectFlagsNonInferrableType
	d := t.AsMetatableType()
	d.tableType = tableType
	d.metatableType = metatableType
	d.indexSource = indexSource
	d.indexIsFunction = indexIsFunction
	d.newindexSource = newindexSource
	d.newindexIsFunction = newindexIsFunction
	d.callSource = callSource
	d.protectedType = protectedType
	c.metatableTypes[key] = t
	return t
}

// canCarryMetatableMembers reports whether a pairing can read members out of t. An intersection
// can: it resolves properties and index signatures like the object types it is made of. A union
// cannot -- which of its arms answers a read is exactly what it does not say.
func (c *Checker) canCarryMetatableMembers(t *Type) bool {
	return t.flags&(TypeFlagsObject|TypeFlagsIntersection) != 0
}

// canCarryAmbientMetamethods reports whether an unpaired operand type can carry ambient
// metamethod members: the same structured shapes getMetatableSource reads from. Primitives never
// carry one -- the number/string operator fast paths stay free of property walks -- and a union
// answers nothing here as everywhere (canCarryMetatableMembers): which arm would dispatch is
// exactly what it does not say.
func canCarryAmbientMetamethods(t *Type) bool {
	return t.flags&(TypeFlagsObject|TypeFlagsIntersection|TypeFlagsInstantiableNonPrimitive) != 0
}

// instantiateMetatableType instantiates a pairing's operands and re-pairs them. The pairing is
// not structural -- its members are already merged -- so it cannot be instantiated in place; and
// if the instantiated metatable no longer says what the fallback is, the augmentation goes with
// it and the table stands alone.
func (c *Checker) instantiateMetatableType(t *Type, m *TypeMapper) *Type {
	d := t.AsMetatableType()
	tableType := c.instantiateType(d.tableType, m)
	metatableType := c.instantiateType(d.metatableType, m)
	if tableType == d.tableType && metatableType == d.metatableType {
		return t
	}
	if result := c.getSetmetatableResultType(tableType, metatableType); result != nil {
		return result
	}
	return tableType
}

// resolveMetatableTypeMembers merges a table with the metamethods of its metatable. A read the
// table misses falls through to __index; the table wins a collision, because a read hits the
// table first -- but only when the table answers with a value: a key that is absent or nil is
// exactly what makes Lua consult the metatable, so a member that admits nil reads as its own
// non-nil half or the fallback. __call makes the pairing callable. __index never contributes call
// signatures: a callable __index table does not make the table itself callable, which is __call's
// job.
func (c *Checker) resolveMetatableTypeMembers(t *Type) {
	d := t.AsMetatableType()
	// The table's own shape. Its property symbols are shared rather than cloned, exactly as an
	// interface shares the symbols of the members it inherits.
	members := createSymbolTable(c.getPropertiesOfType(d.tableType))
	if members == nil {
		members = make(ast.SymbolTable)
	}
	callSignatures := c.getSignaturesOfType(d.tableType, SignatureKindCall)
	constructSignatures := c.getSignaturesOfType(d.tableType, SignatureKindConstruct)
	indexInfos := slices.Clone(c.getIndexInfosOfType(d.tableType))
	// Publish the table's shape before reading the metatable, so a lookup that re-enters this
	// type while a fallback resolves sees a resolved, if incomplete, type instead of recurring.
	c.setStructuredTypeMembers(t, members, callSignatures, constructSignatures, indexInfos)
	t.objectFlags |= ObjectFlagsUnresolvedMembers
	if d.indexSource != nil {
		fallbackInfos := c.getMetatableFallbackIndexInfos(d)
		for name, prop := range members {
			if merged := c.getMetatableFallthroughSymbol(d, fallbackInfos, prop); merged != nil {
				members[name] = merged
			}
		}
		if !d.indexIsFunction {
			// The fallback's own members fill the names the table does not have at all.
			members = c.addInheritedMembers(members, c.getPropertiesOfType(d.indexSource))
		}
		for _, info := range fallbackInfos {
			if findIndexInfo(indexInfos, info.keyType) == nil {
				indexInfos = append(indexInfos, info)
			}
		}
	}
	if d.callSource != nil {
		// __call makes the pairing callable: a call passes the table itself as the metamethod's
		// first argument, so a call site's arguments begin at its second parameter.
		callSignatures = slices.Clone(callSignatures)
		for _, signature := range c.getSignaturesOfType(d.callSource, SignatureKindCall) {
			callSignatures = append(callSignatures, c.getLuaCallMetamethodSignature(signature))
		}
	}
	t.objectFlags &^= ObjectFlagsUnresolvedMembers
	c.setStructuredTypeMembers(t, members, callSignatures, constructSignatures, indexInfos)
}

// getLuaCallMetamethodSignature drops __call's receiver parameter: Lua passes the table itself as
// the metamethod's first argument, so a call site supplies the rest. A signature whose position 0
// is already its rest parameter keeps it -- the rest absorbed the receiver and absorbs the
// arguments the same way.
func (c *Checker) getLuaCallMetamethodSignature(sig *Signature) *Signature {
	if len(sig.parameters) == 0 || signatureHasRestParameter(sig) && len(sig.parameters) == 1 {
		return sig
	}
	result := c.newSignature(sig.flags&SignatureFlagsPropagatingFlags, sig.declaration, sig.typeParameters,
		sig.thisParameter, sig.parameters[1:], c.getReturnTypeOfSignature(sig), nil,
		max(int(sig.minArgumentCount)-1, 0))
	result.target = sig.target
	result.mapper = sig.mapper
	return result
}

// getMetatableFallbackIndexInfos returns the index signatures __index contributes. A table
// contributes the ones it declares. A function contributes one per key domain its parameter
// admits: it answers a read of *any* key there, which is an index signature and not a property,
// since nothing names the keys it will be asked for. The domains are the ones a computed table
// key synthesizes -- `any` keys both strings and numbers, and a literal key widens, because an
// index may not be keyed by one.
func (c *Checker) getMetatableFallbackIndexInfos(d *MetatableType) []*IndexInfo {
	if !d.indexIsFunction {
		return c.getIndexInfosOfType(d.indexSource)
	}
	var infos []*IndexInfo
	for _, signature := range c.getSignaturesOfType(d.indexSource, SignatureKindCall) {
		keyType := c.getTypeAtPosition(signature, 1)
		if !c.isValidIndexArgumentType(keyType) {
			continue
		}
		valueType := c.adjustMultiReturn(c.getReturnTypeOfSignature(signature))
		c.forEachObjectLiteralIndexKeyType(keyType, func(indexKeyType *Type) {
			// Overloads that answer the same key domain both apply: a read of such a key gets
			// whichever one Lua dispatches to. The infos are ours until they are published, so
			// widening one in place is safe.
			if existing := findIndexInfo(infos, indexKeyType); existing != nil {
				existing.valueType = c.getUnionType([]*Type{existing.valueType, valueType})
				return
			}
			infos = append(infos, c.newIndexInfo(indexKeyType, valueType, false /*isReadonly*/, nil /*declaration*/, nil /*components*/))
		})
	}
	return infos
}

// getMetatableFallthroughSymbol merges a table member that may be nil with what __index answers
// for the same key, or returns nil when the member stands as it is. A raw read that comes back
// nil is what makes Lua run the metatable, so such a member is not the whole answer: it reads as
// its own non-nil half or the fallback's value. A member that cannot be nil never falls through.
func (c *Checker) getMetatableFallthroughSymbol(d *MetatableType, fallbackInfos []*IndexInfo, prop *ast.Symbol) *ast.Symbol {
	propType := c.getTypeOfSymbol(prop)
	nonNilType := c.GetNonNullableType(propType)
	if nonNilType == propType && prop.Flags&ast.SymbolFlagsOptional == 0 {
		return nil
	}
	fallbackType := c.getMetatableFallbackType(d, fallbackInfos, prop.Name)
	if fallbackType == nil {
		return nil
	}
	merged := c.createSymbolWithType(prop, c.getUnionType([]*Type{nonNilType, fallbackType}))
	// The key now always answers, whether the table has it or the metatable does.
	merged.Flags &^= ast.SymbolFlagsOptional
	return merged
}

// getMetatableFallbackType returns what __index answers for a key named name: the fallback
// table's own member, or the index signature that covers the key's domain.
func (c *Checker) getMetatableFallbackType(d *MetatableType, fallbackInfos []*IndexInfo, name string) *Type {
	if !d.indexIsFunction {
		if prop := c.getPropertyOfType(d.indexSource, name); prop != nil {
			return c.getTypeOfSymbol(prop)
		}
	}
	keyType := c.esSymbolType
	if !isLateBoundName(name) {
		keyType = c.keyTypeForPropertyName(name)
	}
	if info := c.findApplicableIndexInfo(fallbackInfos, keyType); info != nil {
		return info.valueType
	}
	return nil
}

// luaMetamethodForOperator maps a binary operator token to the metamethod Lua runs for it and
// whether the operands swap first: a > b runs __lt(b, a) and a >= b runs __le(b, a). A compound
// assignment runs the same metamethod its operator does.
func luaMetamethodForOperator(operator ast.Kind) (string, bool) {
	switch operator {
	case ast.KindPlusToken, ast.KindPlusEqualsToken:
		return "__add", false
	case ast.KindMinusToken, ast.KindMinusEqualsToken:
		return "__sub", false
	case ast.KindAsteriskToken, ast.KindAsteriskEqualsToken:
		return "__mul", false
	case ast.KindSlashToken, ast.KindSlashEqualsToken:
		return "__div", false
	case ast.KindPercentToken, ast.KindPercentEqualsToken:
		return "__mod", false
	case ast.KindAsteriskAsteriskToken:
		return "__pow", false
	case ast.KindDotDotToken:
		return "__concat", false
	case ast.KindLessThanToken:
		return "__lt", false
	case ast.KindGreaterThanToken:
		return "__lt", true
	case ast.KindLessThanEqualsToken:
		return "__le", false
	case ast.KindGreaterThanEqualsToken:
		return "__le", true
	}
	return "", false
}

// getLuaMetamethodSignatures returns the call signatures of the metamethod an operand's type
// commits to. A pairing reads the handler off its metatable and off nothing else -- setmetatable
// told the checker exactly where dispatch goes. Any other type that can carry ambient
// metamethods reads a metamethod-NAMED member off the type itself: the ambient-operator
// convention, by which a declared userdata type (GMod's Vector) states its metatable behavior as
// `__add`-style members that no visible setmetatable ever attaches. Optionality commits to
// nothing on either path, so a LuaMetatable<T>-typed value stays inert. Only OPERATOR
// metamethods dispatch ambiently -- __index, __newindex and __call keep meaning only what a
// real pairing's metatable says.
func (c *Checker) getLuaMetamethodSignatures(operandType *Type, name string) []*Signature {
	carrier := operandType
	if isMetatableType(operandType) {
		carrier = operandType.AsMetatableType().metatableType
	} else if !canCarryAmbientMetamethods(operandType) {
		return nil
	}
	handler := c.getMetatableHandlerType(carrier, name)
	if handler == nil {
		return nil
	}
	return c.getSignaturesOfType(handler, SignatureKindCall)
}

// hasLuaEqualityMetamethod reports whether either operand carries an __eq handler against a
// peer that could dispatch it. Lua only consults __eq between two tables (or two userdata), so a
// handler suppresses the no-overlap error only when the other operand could hold a table --
// `span == "hello"` stays flagged. A union peer counts when some arm could dispatch. The
// equality paths already yield boolean, so presence is all that matters -- no parameter check.
func (c *Checker) hasLuaEqualityMetamethod(leftType *Type, rightType *Type) bool {
	return len(c.getLuaMetamethodSignatures(leftType, "__eq")) != 0 && someType(rightType, c.typeMayBeLuaTable) ||
		len(c.getLuaMetamethodSignatures(rightType, "__eq")) != 0 && someType(leftType, c.typeMayBeLuaTable)
}

// typeMayBeLuaTable reports whether a value of this type could be a table (or userdata) at
// runtime -- the shapes an __eq dispatch needs on BOTH sides. A type variable answers by its
// base constraint, not its own non-primitive flag: `T extends string` never holds a table, while
// an unconstrained T might.
func (c *Checker) typeMayBeLuaTable(t *Type) bool {
	if t.flags&TypeFlagsInstantiableNonPrimitive != 0 {
		if constraint := c.getBaseConstraintOfType(t); constraint != nil {
			return someType(constraint, c.typeMayBeLuaTable)
		}
		return true
	}
	return t.flags&(TypeFlagsObject|TypeFlagsIntersection) != 0
}

// checkLuaBinaryMetamethod dispatches a binary operator to a paired operand's metamethod, or
// returns nil when neither operand carries one and ordinary checking stands. Lua consults the
// left operand's handler first, then the right's, and always passes the operands in dispatch
// order -- so a handler found on the right still receives the left operand first, and its
// parameter types say whether that is allowed. The first signature that admits both operands
// wins; when none does, the first names the mismatch. The result is the handler's first return
// value, as at any call boundary.
func (c *Checker) checkLuaBinaryMetamethod(operator ast.Kind, left *ast.Node, right *ast.Node, leftType *Type, rightType *Type) *Type {
	name, swapped := luaMetamethodForOperator(operator)
	if name == "" {
		return nil
	}
	if swapped {
		left, right = right, left
		leftType, rightType = rightType, leftType
	}
	signatures := c.getLuaMetamethodSignatures(leftType, name)
	if len(signatures) == 0 {
		signatures = c.getLuaMetamethodSignatures(rightType, name)
	}
	// A handler receives both operands whatever it declares, but a signature that does not name
	// an operand's position cannot check it: an under-arity handler commits to nothing (the
	// out-of-range position would read as `any` and admit every operand vacuously), and ordinary
	// operand checking stands. Deliberately AFTER the left-then-right selection: Lua would still
	// run the left operand's handler at runtime, so an under-arity left refuses dispatch outright
	// rather than falling through to a right handler that would never be consulted.
	signatures = core.Filter(signatures, func(s *Signature) bool {
		return c.tryGetTypeAtPosition(s, 0) != nil && c.tryGetTypeAtPosition(s, 1) != nil
	})
	if len(signatures) == 0 {
		return nil
	}
	signature := core.Find(signatures, func(s *Signature) bool {
		return c.isTypeAssignableTo(leftType, c.getTypeAtPosition(s, 0)) &&
			c.isTypeAssignableTo(rightType, c.getTypeAtPosition(s, 1))
	})
	if signature == nil {
		signature = signatures[0]
		c.checkTypeAssignableTo(leftType, c.getTypeAtPosition(signature, 0), left, nil)
		c.checkTypeAssignableTo(rightType, c.getTypeAtPosition(signature, 1), right, nil)
	}
	return c.adjustMultiReturn(c.getReturnTypeOfSignature(signature))
}

// checkLuaUnaryMetamethod dispatches a unary operator to an operand's metamethod. An
// under-arity handler commits to nothing, as at the binary sites.
func (c *Checker) checkLuaUnaryMetamethod(operandType *Type, name string) *Type {
	signatures := core.Filter(c.getLuaMetamethodSignatures(operandType, name), func(s *Signature) bool {
		return c.tryGetTypeAtPosition(s, 0) != nil
	})
	if len(signatures) == 0 {
		return nil
	}
	return c.adjustMultiReturn(c.getReturnTypeOfSignature(signatures[0]))
}

// getMetatableNewindexWriteType returns the type __newindex accepts for a write of a key the
// merged shape misses, or nil when the pairing commits to no __newindex or the handler does not
// answer this key. The table form receives the write, so its own member or index signature types
// it; the function form is called as (table, key, value), so its value parameter does. name is
// the written property's name, or "" for a computed key.
func (c *Checker) getMetatableNewindexWriteType(t *Type, keyType *Type, name string) *Type {
	if !isMetatableType(t) {
		return nil
	}
	d := t.AsMetatableType()
	if d.newindexSource == nil {
		return nil
	}
	if d.newindexIsFunction {
		if !c.isValidIndexArgumentType(keyType) {
			return nil
		}
		for _, signature := range c.getSignaturesOfType(d.newindexSource, SignatureKindCall) {
			if c.isTypeAssignableTo(keyType, c.getTypeAtPosition(signature, 1)) {
				return c.getTypeAtPosition(signature, 2)
			}
		}
		return nil
	}
	if name != "" {
		if prop := c.getPropertyOfType(d.newindexSource, name); prop != nil {
			return c.getWriteTypeOfSymbol(prop)
		}
	}
	if info := c.getApplicableIndexInfo(d.newindexSource, keyType); info != nil && !info.isReadonly {
		return info.valueType
	}
	return nil
}

// getMetatableNewindexWriteTypeForName is the named-member entry point for the write fallback.
func (c *Checker) getMetatableNewindexWriteTypeForName(t *Type, name string) *Type {
	keyType := c.esSymbolType
	if !isLateBoundName(name) {
		keyType = c.keyTypeForPropertyName(name)
	}
	return c.getMetatableNewindexWriteType(t, keyType, name)
}
