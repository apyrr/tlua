package ast

type FunctionFlags uint32

const (
	FunctionFlagsNormal FunctionFlags = 0
	// 1 << 0 was FunctionFlagsGenerator; generators are removed in tlua.
	FunctionFlagsAsync   FunctionFlags = 1 << 1
	FunctionFlagsInvalid FunctionFlags = 1 << 2
)

func GetFunctionFlags(node *Node) FunctionFlags {
	if node == nil {
		return FunctionFlagsInvalid
	}
	data := node.BodyData()
	if data == nil {
		return FunctionFlagsInvalid
	}
	flags := FunctionFlagsNormal
	switch node.Kind {
	case KindFunctionDeclaration, KindFunctionExpression, KindArrowFunction:
		if HasSyntacticModifier(node, ModifierFlagsAsync) {
			flags |= FunctionFlagsAsync
		}
	}
	if data.Body == nil {
		flags |= FunctionFlagsInvalid
	}
	return flags
}
