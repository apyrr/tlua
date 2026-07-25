package checker

import (
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
)

var luaIONames = []string{"read", "lines"}

type luaIOCall struct {
	name        string
	formatIndex int
}

func (c *Checker) getLuaIOCall(node *ast.Node) *luaIOCall {
	call := c.resolveLuaBuiltinCall(node, luaIONames, c.getLuaIOGlobalSymbol, c.getLuaFileType)
	if call == nil {
		return nil
	}
	// Where the format sits in the signature, before any colon shift: `io.read`
	// takes it first, and everything else reserves that slot -- `io.lines` for the
	// filename, the file members for the handle.
	declaredIndex := 1
	if call.namespaceForm && call.name == "read" {
		declaredIndex = 0
	}
	formatIndex := ast.LuaWrittenArgumentIndex(node, declaredIndex)
	if formatIndex < 0 {
		// `io:read("*a")` hands the io table itself to the format parameter, so
		// there is no written format to read and nothing to refine. Without this
		// the read loop below indexes the argument list at -1.
		return nil
	}
	return &luaIOCall{name: call.name, formatIndex: formatIndex}
}

func (c *Checker) checkLuaIOCall(node *ast.Node, checkMode CheckMode) *Type {
	call := c.getLuaIOCall(node)
	if call == nil {
		return nil
	}
	args := node.Arguments()
	types := make([]*Type, 0, max(len(args)-call.formatIndex, 1))
	openTail := false
	for i := call.formatIndex; i < len(args) && !openTail; i++ {
		formatType := c.checkExpressionCachedEx(args[i], checkMode)
		switch {
		case formatType.flags&TypeFlagsNumberLiteral != 0:
			types = append(types, c.stringOrNilType)
		case formatType.flags&TypeFlagsStringLiteral != 0:
			switch getStringLiteralValue(formatType) {
			case "*n":
				types = append(types, c.numberOrNilType)
			case "*a":
				// "*a" reads "" at end of file, never nil -- but a FAILED
				// earlier format stops the read, so only the first format can
				// promise a value.
				types = append(types, core.IfElse(i == call.formatIndex, c.stringType, c.stringOrNilType))
			case "*l":
				types = append(types, c.stringOrNilType)
			default:
				// A literal outside LuaReadFormat: the declared union already
				// reports it; no refinement.
				return nil
			}
		default:
			// A format only known at runtime ends the analysis, but the
			// literal prefix keeps its precision under an open tail.
			if len(types) == 0 {
				return nil
			}
			openTail = true
		}
	}
	if len(types) == 0 {
		types = append(types, c.stringOrNilType) // no formats read one line
	}
	resultPack := c.createLuaValuePack(types, openTail)
	if call.name == "read" {
		return resultPack
	}
	return c.createLuaIteratorType(resultPack)
}
