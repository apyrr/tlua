package scanner

import (
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/json"
	"github.com/apyrr/tlua/internal/repo"
	"gotest.tools/v3/assert"
)

// grammarWordPattern matches one word alternation guarded the way the grammar
// guards a whole word: `(?<![_$[:alnum:]])(?:(?<=\.\.\.)|(?<!\.))(a|b|c)(?![_$[:alnum:]])`.
var grammarWordPattern = regexp.MustCompile(`\(\?<!\[_\$\[:alnum:\]\]\)\(\?:\(\?<=\\\.\\\.\\\.\)\|\(\?<!\\\.\)\)\(?([A-Za-z][A-Za-z|]*)\)?\(\?!\[_\$\[:alnum:\]\]\)`)

// The editor grammar is a hand-maintained copy of what the scanner knows, and it
// drifts when a word leaves the language: `null`, `undefined`, `Infinity` and
// `NaN` were still coloured as language constants after tlua stopped
// recognizing them. Every word the grammar colours as a value literal
// (constant.language) or a type keyword (support.type) must be a scanner keyword.
func TestGrammarLiteralWordsAreKeywords(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile(filepath.Join(repo.RootPath(), "_extension", "syntaxes", "tlua.tmLanguage.json"))
	assert.NilError(t, err)
	var grammar any
	assert.NilError(t, json.Unmarshal(data, &grammar))

	words := map[string]string{}
	var walk func(node any)
	walk = func(node any) {
		switch node := node.(type) {
		case map[string]any:
			name, _ := node["name"].(string)
			match, _ := node["match"].(string)
			if match != "" && (strings.HasPrefix(name, "constant.language") || strings.HasPrefix(name, "support.type")) {
				for _, group := range grammarWordPattern.FindAllStringSubmatch(match, -1) {
					for word := range strings.SplitSeq(group[1], "|") {
						words[word] = name
					}
				}
			}
			for _, child := range node {
				walk(child)
			}
		case []any:
			for _, child := range node {
				walk(child)
			}
		}
	}
	walk(grammar)

	// A change to how the grammar spells its word guard would otherwise make this
	// pass by finding nothing.
	assert.Assert(t, len(words) >= 15, "found only %d literal words; has the grammar's word guard changed?", len(words))
	for _, word := range slices.Sorted(maps.Keys(words)) {
		_, ok := textToKeyword[word]
		assert.Assert(t, ok, "the grammar colours %q as %s, but the scanner does not recognize it", word, words[word])
	}
}
