package ast

import (
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/apyrr/tlua/internal/jsnum"
)

// Symbol

type Symbol struct {
	Flags            SymbolFlags
	CheckFlags       CheckFlags // Non-zero only in transient symbols created by Checker
	Name             string
	Declarations     []*Node
	ValueDeclaration *Node
	Members          SymbolTable
	Exports          SymbolTable
	id               atomic.Uint64
	Parent           *Symbol
	ExportSymbol     *Symbol
}

func (s *Symbol) IsExternalModule() bool {
	return s.Flags&SymbolFlagsModule != 0 && len(s.Name) > 0 && s.Name[0] == '"'
}

func (s *Symbol) IsStatic() bool {
	if s.ValueDeclaration == nil {
		return false
	}
	modifierFlags := s.ValueDeclaration.ModifierFlags()
	return modifierFlags&ModifierFlagsStatic != 0
}

// See comment on `declareModuleMember` in `binder.go`.
func (s *Symbol) CombinedLocalAndExportSymbolFlags() SymbolFlags {
	if s.ExportSymbol != nil {
		return s.Flags | s.ExportSymbol.Flags
	}
	return s.Flags
}

// SymbolTable

type SymbolTable map[string]*Symbol

const InternalSymbolNamePrefix = "\xFE" // Invalid UTF8 sequence, will never occur as IdentifierName

const (
	InternalSymbolNameCall                    = InternalSymbolNamePrefix + "call"                    // Call signatures
	InternalSymbolNameConstructor             = InternalSymbolNamePrefix + "constructor"             // Constructor implementations
	InternalSymbolNameNew                     = InternalSymbolNamePrefix + "new"                     // Constructor signatures
	InternalSymbolNameIndex                   = InternalSymbolNamePrefix + "index"                   // Index signatures
	InternalSymbolNameExportStar              = InternalSymbolNamePrefix + "export"                  // Module export * declarations
	InternalSymbolNameGlobal                  = InternalSymbolNamePrefix + "global"                  // Global self-reference
	InternalSymbolNameMissing                 = InternalSymbolNamePrefix + "missing"                 // Indicates missing symbol
	InternalSymbolNameType                    = InternalSymbolNamePrefix + "type"                    // Anonymous type literal symbol
	InternalSymbolNameObject                  = InternalSymbolNamePrefix + "object"                  // Anonymous object literal declaration
	InternalSymbolNameJSXAttributes           = InternalSymbolNamePrefix + "jsxAttributes"           // Anonymous JSX attributes object literal declaration
	InternalSymbolNameClass                   = InternalSymbolNamePrefix + "class"                   // Unnamed class expression
	InternalSymbolNameFunction                = InternalSymbolNamePrefix + "function"                // Unnamed function expression
	InternalSymbolNameComputed                = InternalSymbolNamePrefix + "computed"                // Computed property name declaration with dynamic name
	InternalSymbolNameAssignmentDeclaration   = InternalSymbolNamePrefix + "assignment"              // Assignment declarations
	InternalSymbolNameInstantiationExpression = InternalSymbolNamePrefix + "instantiationExpression" // Instantiation expressions
	InternalSymbolNameExportEquals            = "export="                                            // Export assignment symbol
	InternalSymbolNameDefault                 = "default"                                            // Default export symbol (technically not wholly internal, but included here for usability)
	InternalSymbolNameThis                    = "this"
	InternalSymbolNameModuleExports           = "module.exports"
)

// numberKeyPrefix marks symbol-table names that denote Lua number keys,
// following the "\xFE@" known-symbol / "\xFE#" private-name precedent.
// Lua distinguishes number keys from string keys (t[1] != t["1"]), so
// number-keyed properties live in their own name namespace.
const numberKeyPrefix = InternalSymbolNamePrefix + "n:"

// NumberKeyName returns the symbol-table name for the number key value.
func NumberKeyName(value jsnum.Number) string {
	return numberKeyPrefix + value.String()
}

// numberKeyNameSmall caches the names for the small non-negative indexes that
// dominate tuple/array element access, avoiding a per-lookup allocation. 33
// entries cover keys 0..32, so a 32-element tuple's last key (1-based) still
// hits the cache.
var numberKeyNameSmall = func() [33]string {
	var names [33]string
	for i := range names {
		names[i] = numberKeyPrefix + strconv.Itoa(i)
	}
	return names
}()

// NumberKeyNameFromInt returns the number-key name for a non-negative integer
// index. It is byte-identical to NumberKeyName(jsnum.Number(i)) for indexes in
// the safe-integer range (which all tuple/array positions are) but skips the
// float round-trip.
func NumberKeyNameFromInt(i int) string {
	if i >= 0 && i < len(numberKeyNameSmall) {
		return numberKeyNameSmall[i]
	}
	return numberKeyPrefix + strconv.Itoa(i)
}

// NumberKeyNameFromPosition returns the number-key name for the 0-based element
// position pos: Lua sequences are 1-based, so position pos occupies key pos+1.
// Every conversion from an element/parameter/argument position to a number key
// must go through here (or state its own +1 explicitly); a bare
// NumberKeyNameFromInt(pos) at such a site is an off-by-one bug.
func NumberKeyNameFromPosition(pos int) string {
	return NumberKeyNameFromInt(pos + 1)
}

// NumberKeyNameFromText canonicalizes numeric-literal source text through ES
// ToNumber so every spelling of a value ("1.0", "0x10", "+16", "1e1") names
// the same key as "16".
func NumberKeyNameFromText(text string) string {
	// Fast path: canonical decimal text (which is what the scanner produces for
	// numeric-literal tokens) is already in ToNumber-normal form, so it names
	// its key directly without the FromString parse and float format.
	if isCanonicalNumberKeyText(text) {
		return numberKeyPrefix + text
	}
	return NumberKeyName(jsnum.FromString(text))
}

// isCanonicalNumberKeyText reports whether text is the ToNumber-canonical
// spelling of a non-negative integer within the safe-integer range: all ASCII
// digits, no redundant leading zero, and short enough to round-trip exactly.
func isCanonicalNumberKeyText(text string) bool {
	if len(text) == 0 || len(text) > 15 {
		return false
	}
	if text[0] == '0' {
		return len(text) == 1 // only "0" itself; "01" is not canonical
	}
	for i := range len(text) {
		if text[i] < '0' || text[i] > '9' {
			return false
		}
	}
	return true
}

// IsNumberKeyName reports whether name denotes a number key.
func IsNumberKeyName(name string) bool {
	return strings.HasPrefix(name, numberKeyPrefix)
}

// NumberKeyValue returns the canonical decimal text ("16", "-1", "1.5") of a
// number-key name.
func NumberKeyValue(name string) (string, bool) {
	return strings.CutPrefix(name, numberKeyPrefix)
}

// NumberKeyDisplayName returns the user-visible text of a symbol name: a
// number-key name ("\xFEn:1") decodes to its decimal value ("1"); any other
// name passes through unchanged. Use this wherever a symbol name flows into a
// diagnostic message or other user-facing text, so the "\xFEn:" sentinel never
// leaks.
func NumberKeyDisplayName(name string) string {
	if value, ok := NumberKeyValue(name); ok {
		return value
	}
	return name
}

// NumberKeyNumber returns the numeric value of a number-key name.
func NumberKeyNumber(name string) (jsnum.Number, bool) {
	if v, ok := NumberKeyValue(name); ok {
		return jsnum.FromString(v), true
	}
	return 0, false
}

func SymbolName(symbol *Symbol) string {
	return symbol.Name
}

// EscapeAllInternalSymbolNames replaces internal symbol name markers ("\xFE") with "__".
func EscapeAllInternalSymbolNames(name string) string {
	return strings.ReplaceAll(name, InternalSymbolNamePrefix, "__")
}

func EscapeInternalSymbolName(name string) string {
	if rest, ok := strings.CutPrefix(name, InternalSymbolNamePrefix); ok {
		return "__" + rest
	}
	return name
}

// EscapeSymbolName converts a binder symbol name into its escaped "__String"
// form. Internal names (prefixed with the "\xFE" sentinel) become "__"-prefixed,
// and user names that already begin with "__" gain an extra leading underscore
// so they can be distinguished from internal names.
func EscapeSymbolName(name string) string {
	if rest, ok := strings.CutPrefix(name, InternalSymbolNamePrefix); ok {
		return "__" + rest
	}
	if len(name) >= 2 && name[0] == '_' && name[1] == '_' {
		return "_" + name
	}
	return name
}
