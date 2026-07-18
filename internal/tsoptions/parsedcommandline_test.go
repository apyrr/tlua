package tsoptions_test

import (
	"slices"
	"testing"

	"github.com/apyrr/tlua/internal/tsoptions"
	"github.com/apyrr/tlua/internal/tsoptions/tsoptionstest"
	"github.com/apyrr/tlua/internal/vfs/vfstest"
	"gotest.tools/v3/assert"
)

func TestParsedCommandLine(t *testing.T) {
	t.Parallel()
	t.Run("PossiblyMatchesFileName", func(t *testing.T) {
		t.Parallel()

		noFiles := map[string]string{}
		noFilesFS := vfstest.FromMap(noFiles, true)

		files := map[string]string{
			"/dev/a.tlua":        "",
			"/dev/a.d.tlua":      "",
			"/dev/a.lua":         "",
			"/dev/b.tlua":        "",
			"/dev/b.lua":         "",
			"/dev/c.d.tlua":      "",
			"/dev/z/a.tlua":      "",
			"/dev/z/abz.tlua":    "",
			"/dev/z/aba.tlua":    "",
			"/dev/z/b.tlua":      "",
			"/dev/z/bbz.tlua":    "",
			"/dev/z/bba.tlua":    "",
			"/dev/x/a.tlua":      "",
			"/dev/x/aa.tlua":     "",
			"/dev/x/b.tlua":      "",
			"/dev/x/y/a.tlua":    "",
			"/dev/x/y/b.tlua":    "",
			"/dev/js/a.lua":      "",
			"/dev/js/b.lua":      "",
			"/dev/js/d.min.lua":  "",
			"/dev/js/ab.min.lua": "",
			"/ext/ext.tlua":      "",
			"/ext/b/a..b.tlua":   "",
		}

		assertMatches := func(t *testing.T, parsedCommandLine *tsoptions.ParsedCommandLine, files map[string]string, matches []string) {
			t.Helper()
			for fileName := range files {
				actual := parsedCommandLine.PossiblyMatchesFileName(fileName)
				expected := slices.Contains(matches, fileName)
				assert.Equal(t, actual, expected, "fileName: %s", fileName)
			}
			for _, fileName := range matches {
				if _, ok := files[fileName]; !ok {
					actual := parsedCommandLine.PossiblyMatchesFileName(fileName)
					assert.Equal(t, actual, true, "fileName: %s", fileName)
				}
			}
		}

		t.Run("with literal file list", func(t *testing.T) {
			t.Parallel()
			t.Run("without exclude", func(t *testing.T) {
				t.Parallel()
				parsedCommandLine := tsoptionstest.GetParsedCommandLine(
					t,
					`{
						"files": [
							"a.tlua",
							"b.tlua"
						]
					}`,
					files,
					"/dev",
					/*useCaseSensitiveFileNames*/ true,
				)

				assertMatches(t, parsedCommandLine, files, []string{
					"/dev/a.tlua",
					"/dev/b.tlua",
				})
			})

			t.Run("are not removed due to excludes", func(t *testing.T) {
				t.Parallel()
				parsedCommandLine := tsoptionstest.GetParsedCommandLine(
					t,
					`{
						"files": [
							"a.tlua",
							"b.tlua"
						],
						"exclude": [
							"b.tlua"
						]
					}`,
					files,
					"/dev",
					/*useCaseSensitiveFileNames*/ true,
				)

				assertMatches(t, parsedCommandLine, files, []string{
					"/dev/a.tlua",
					"/dev/b.tlua",
				})

				emptyParsedCommandLine := parsedCommandLine.ReloadFileNamesOfParsedCommandLine(noFilesFS)
				assertMatches(t, emptyParsedCommandLine, noFiles, []string{
					"/dev/a.tlua",
					"/dev/b.tlua",
				})
			})

			t.Run("duplicates", func(t *testing.T) {
				t.Parallel()
				parsedCommandLine := tsoptionstest.GetParsedCommandLine(
					t,
					`{
						"files": [
							"a.tlua",
							"a.tlua",
							"b.tlua",
						]
					}`,
					files,
					"/dev",
					/*useCaseSensitiveFileNames*/ true,
				)

				assert.DeepEqual(t, parsedCommandLine.LiteralFileNames(), []string{
					"/dev/a.tlua",
					"/dev/b.tlua",
				})
			})
		})

		t.Run("with literal include list", func(t *testing.T) {
			t.Parallel()
			t.Run("without exclude", func(t *testing.T) {
				t.Parallel()
				parsedCommandLine := tsoptionstest.GetParsedCommandLine(
					t,
					`{
						"include": [
							"a.tlua",
							"b.tlua"
						]
					}`,
					files,
					"/dev",
					/*useCaseSensitiveFileNames*/ true,
				)

				assertMatches(t, parsedCommandLine, files, []string{
					"/dev/a.tlua",
					"/dev/b.tlua",
				})

				emptyParsedCommandLine := parsedCommandLine.ReloadFileNamesOfParsedCommandLine(noFilesFS)
				assertMatches(t, emptyParsedCommandLine, noFiles, []string{
					"/dev/a.tlua",
					"/dev/b.tlua",
				})
			})
		})
	})
}
