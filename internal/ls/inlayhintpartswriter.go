package ls

import (
	"strings"
	"unicode/utf8"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/printer"
	"github.com/apyrr/tlua/internal/stringutil"
)

var _ printer.EmitTextWriter = &inlayHintPartsWriter{}

// inlayHintPartsWriter captures printed text as inlay hint label parts. A name the printer
// writes with its symbol becomes a part that links to the symbol's declaration; all other
// text merges into plain parts. Hints are one line, so a line break prints as a space --
// one, however many breaks and spaces meet (writeLineBreak).
type inlayHintPartsWriter struct {
	state       *inlayHintState
	parts       []*lsproto.InlayHintLabelPart
	pending     strings.Builder
	length      int
	lastWritten string
}

func (w *inlayHintPartsWriter) write(text string) {
	if text == "" {
		return
	}
	w.pending.WriteString(text)
	w.length += len(text)
	w.lastWritten = text
}

func (w *inlayHintPartsWriter) flush() {
	if w.pending.Len() != 0 {
		w.parts = append(w.parts, &lsproto.InlayHintLabelPart{Value: w.pending.String()})
		w.pending.Reset()
	}
}

func (w *inlayHintPartsWriter) labelParts() []*lsproto.InlayHintLabelPart {
	w.flush()
	return w.parts
}

func (w *inlayHintPartsWriter) WriteSymbol(text string, symbol *ast.Symbol) {
	if symbol != nil && len(symbol.Declarations) != 0 {
		if name := ast.GetNameOfDeclaration(symbol.Declarations[0]); name != nil {
			w.flush()
			w.parts = append(w.parts, w.state.getNodeDisplayPart(text, name))
			w.length += len(text)
			w.lastWritten = text
			return
		}
	}
	w.write(text)
}

func (w *inlayHintPartsWriter) writeLineBreak() {
	if w.length != 0 && !w.HasTrailingWhitespace() {
		w.write(" ")
	}
}

func (w *inlayHintPartsWriter) String() string {
	var b strings.Builder
	for _, part := range w.parts {
		b.WriteString(part.Value)
	}
	b.WriteString(w.pending.String())
	return b.String()
}

func (w *inlayHintPartsWriter) Clear() {
	w.parts = nil
	w.pending.Reset()
	w.length = 0
	w.lastWritten = ""
}

func (w *inlayHintPartsWriter) HasTrailingWhitespace() bool {
	ch, _ := utf8.DecodeLastRuneInString(w.lastWritten)
	return ch != utf8.RuneError && stringutil.IsWhiteSpaceLike(ch)
}

func (w *inlayHintPartsWriter) Write(s string)                     { w.write(s) }
func (w *inlayHintPartsWriter) WriteTrailingSemicolon(text string) { w.write(text) }
func (w *inlayHintPartsWriter) WriteComment(text string)           { w.write(text) }
func (w *inlayHintPartsWriter) WriteKeyword(text string)           { w.write(text) }
func (w *inlayHintPartsWriter) WriteOperator(text string)          { w.write(text) }
func (w *inlayHintPartsWriter) WritePunctuation(text string)       { w.write(text) }
func (w *inlayHintPartsWriter) WriteSpace(text string)             { w.write(text) }
func (w *inlayHintPartsWriter) WriteStringLiteral(text string)     { w.write(text) }
func (w *inlayHintPartsWriter) WriteParameter(text string)         { w.write(text) }
func (w *inlayHintPartsWriter) WriteProperty(text string)          { w.write(text) }
func (w *inlayHintPartsWriter) WriteLiteral(s string)              { w.write(s) }
func (w *inlayHintPartsWriter) RawWrite(s string)                  { w.write(s) }
func (w *inlayHintPartsWriter) WriteLine()                         { w.writeLineBreak() }
func (w *inlayHintPartsWriter) WriteLineForce(force bool)          { w.writeLineBreak() }
func (w *inlayHintPartsWriter) IncreaseIndent()                    {}
func (w *inlayHintPartsWriter) DecreaseIndent()                    {}
func (w *inlayHintPartsWriter) GetTextPos() int                    { return w.length }
func (w *inlayHintPartsWriter) GetLine() int                       { return 0 }
func (w *inlayHintPartsWriter) GetColumn() core.UTF16Offset        { return 0 }
func (w *inlayHintPartsWriter) GetIndent() int                     { return 0 }
func (w *inlayHintPartsWriter) IsAtStartOfLine() bool              { return false }
func (w *inlayHintPartsWriter) HasTrailingComment() bool           { return false }
