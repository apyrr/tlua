package project_test

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/project"
	"github.com/apyrr/tlua/internal/testutil/projecttestutil"
	"github.com/apyrr/tlua/internal/tspath"
	"gotest.tools/v3/assert"
)

const wsDiagConfigPath = tspath.Path("/project/tluaconfig.json")

// wsDiagFiles builds a project over the given sources. skipLibCheck keeps the
// `require` shim from being type-checked; it is still part of the workspace
// pull's file set (an unchecked file contributes its syntactic diagnostics), so
// the cache tracks it alongside the sources named here.
func wsDiagFiles(sources map[string]string) map[string]any {
	files := map[string]any{
		"/project/tluaconfig.json": `{
			"compilerOptions": {
				"noLib": true,
				"skipLibCheck": true
			}
		}`,
		"/project/require.d.tlua": `declare function require(module: string): any;`,
	}
	for name, content := range sources {
		files["/project/"+name] = content
	}
	return files
}

type wsDiagFixture struct {
	t       *testing.T
	session *project.Session
	utils   *projecttestutil.SessionUtils
	cache   *project.WorkspaceDiagCache
	openURI lsproto.DocumentUri
	version int32
	project *project.Project
}

// newWSDiagFixture sets up a session over the given sources, opens `openFile`
// and builds the program, mirroring what a first `workspace/diagnostic` pull
// would find.
func newWSDiagFixture(t *testing.T, sources map[string]string, openFile string) *wsDiagFixture {
	t.Helper()
	files := wsDiagFiles(sources)
	session, utils := projecttestutil.Setup(files)
	fixture := &wsDiagFixture{
		t:       t,
		session: session,
		utils:   utils,
		cache:   session.WorkspaceDiagnosticsCache(),
		openURI: lsproto.DocumentUri("file:///project/" + openFile),
		version: 1,
	}
	session.DidOpenFile(context.Background(), fixture.openURI, fixture.version, sources[openFile], lsproto.LanguageKindTypeScript)
	fixture.refresh()

	// Guard the fixtures themselves: a source that did not parse, or that never
	// made it into the program, would make every assertion below vacuous.
	program := fixture.project.GetProgram()
	for name := range sources {
		file := program.GetSourceFile("/project/" + name)
		assert.Assert(t, file != nil, "%s should be part of the program", name)
		assert.Equal(t, len(program.GetSyntacticDiagnostics(projecttestutil.WithRequestID(context.Background()), file)), 0,
			"%s should parse cleanly", name)
	}
	return fixture
}

// refresh re-fetches the current project from the session's snapshot: every
// change produces a new snapshot with a new *Project, and the cache is keyed by
// the (stable) config file path.
func (f *wsDiagFixture) refresh() {
	f.t.Helper()
	_, err := f.session.GetLanguageService(context.Background(), f.openURI)
	assert.NilError(f.t, err)
	f.session.WaitForBackgroundTasks()
	snapshot := f.session.Snapshot()
	proj := snapshot.ProjectCollection.ConfiguredProject(wsDiagConfigPath)
	assert.Assert(f.t, proj != nil)
	assert.Assert(f.t, proj.GetProgram() != nil)
	f.project = proj
}

// reconcile does what one pull does before reading any status: take the current
// project and diff it against the cache.
func (f *wsDiagFixture) reconcile() {
	f.t.Helper()
	f.refresh()
	f.cache.Reconcile(f.project)
}

func (f *wsDiagFixture) path(name string) tspath.Path {
	return f.utils.ToPath("/project/" + name)
}

func (f *wsDiagFixture) status(name string) project.FileDiagStatus {
	f.t.Helper()
	status, ok := f.cache.Status(f.project.Id(), f.path(name))
	assert.Assert(f.t, ok, "expected a cache entry for %s", name)
	return status
}

func (f *wsDiagFixture) cached(name string) bool {
	_, ok := f.cache.Status(f.project.Id(), f.path(name))
	return ok
}

// store records diagnostics for a file the way the handler does: read the
// status, "compute", store against the generation and hash that were read.
func (f *wsDiagFixture) store(name string, items []*lsproto.Diagnostic) (string, bool) {
	f.t.Helper()
	status := f.status(name)
	resultID, changed, ok := f.cache.StoreComputed(f.project.Id(), f.path(name), status.Generation, status.FileHash, items)
	assert.Assert(f.t, ok, "expected StoreComputed to be accepted for %s", name)
	return resultID, changed
}

// storeAll gives every named file one synthetic diagnostic and returns the
// result IDs it was handed.
func (f *wsDiagFixture) storeAll(names ...string) map[string]string {
	f.t.Helper()
	ids := make(map[string]string, len(names))
	for _, name := range names {
		resultID, _ := f.store(name, wsDiagItems(name))
		ids[name] = resultID
	}
	return ids
}

func (f *wsDiagFixture) assertValid(names ...string) {
	f.t.Helper()
	for _, name := range names {
		assert.Assert(f.t, f.status(name).Valid, "expected %s to be valid", name)
	}
}

func (f *wsDiagFixture) assertInvalid(names ...string) {
	f.t.Helper()
	for _, name := range names {
		assert.Assert(f.t, !f.status(name).Valid, "expected %s to be invalid", name)
	}
}

// writeFile edits a file on disk and reports the change the way a file watcher
// would.
func (f *wsDiagFixture) writeFile(name string, content string) {
	f.t.Helper()
	assert.NilError(f.t, f.utils.FS().WriteFile("/project/"+name, content))
	f.session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{{
		Type: lsproto.FileChangeTypeChanged,
		Uri:  lsproto.DocumentUri("file:///project/" + name),
	}})
}

func (f *wsDiagFixture) watchedEvent(name string, kind lsproto.FileChangeType) {
	f.t.Helper()
	f.session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{{
		Type: kind,
		Uri:  lsproto.DocumentUri("file:///project/" + name),
	}})
}

// changeOverlay edits the open document without touching disk.
func (f *wsDiagFixture) changeOverlay(content string) {
	f.t.Helper()
	f.version++
	f.session.DidChangeFile(context.Background(), f.openURI, f.version, []lsproto.TextDocumentContentChangePartialOrWholeDocument{{
		WholeDocument: &lsproto.TextDocumentContentChangeWholeDocument{Text: content},
	}})
}

func wsDiagItems(messages ...string) []*lsproto.Diagnostic {
	items := make([]*lsproto.Diagnostic, 0, len(messages))
	for _, message := range messages {
		items = append(items, &lsproto.Diagnostic{
			Range: lsproto.Range{
				Start: lsproto.Position{Line: 0, Character: 0},
				End:   lsproto.Position{Line: 0, Character: 1},
			},
			Message: lsproto.StringOrMarkupContent{String: new(message)},
		})
	}
	return items
}

// Value-only modules: none of these contribute to the global scope, so a change
// to one of them can be invalidated selectively.
const (
	wsDiagModuleA = `local b = require("b")
local function useB()
	return b.f()
end
return { useB = useB }
`
	wsDiagModuleB = `local function f()
	return 1
end
return { f = f }
`
	wsDiagModuleC = `local function unrelated()
	return 1
end
return { unrelated = unrelated }
`
)

func TestWorkspaceDiagCache(t *testing.T) {
	t.Parallel()

	t.Run("fresh reconcile", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
		}, "a.tlua")

		// Nothing is cached until the first Reconcile.
		assert.Assert(t, !f.cached("a.tlua"))

		f.reconcile()
		for _, name := range []string{"a.tlua", "b.tlua", "c.tlua"} {
			status := f.status(name)
			assert.Assert(t, !status.Valid, "%s should start out invalid", name)
			assert.Equal(t, status.ResultID, "")
			assert.Assert(t, !status.Reported)
		}
		// The `require` declaration is not type-checked under skipLibCheck, but
		// its syntactic diagnostics are still the pull's to report, so it is
		// tracked like any other file.
		assert.Assert(t, f.cached("require.d.tlua"))
		assert.Assert(t, !f.status("require.d.tlua").Valid)

		ids := f.storeAll("a.tlua", "b.tlua", "c.tlua")
		for name, id := range ids {
			assert.Assert(t, id != "", "expected a result ID for %s", name)
			assert.Assert(t, f.status(name).Valid)
		}
		assert.Assert(t, ids["a.tlua"] != ids["b.tlua"])

		// A second pull with nothing changed is a no-op.
		f.reconcile()
		for name, id := range ids {
			status := f.status(name)
			assert.Assert(t, status.Valid, "%s should still be valid", name)
			assert.Equal(t, status.ResultID, id)
		}
	})

	t.Run("selective invalidation of importers", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
		}, "a.tlua")
		f.reconcile()
		ids := f.storeAll("a.tlua", "b.tlua", "c.tlua")

		f.writeFile("b.tlua", `local function f()
	return 2
end
return { f = f }
`)
		f.reconcile()

		// b changed and a imports it; c is untouched and keeps its result.
		f.assertInvalid("b.tlua", "a.tlua")
		cStatus := f.status("c.tlua")
		assert.Assert(t, cStatus.Valid, "unrelated file should stay valid")
		assert.Equal(t, cStatus.ResultID, ids["c.tlua"])
	})

	t.Run("transitive reverse-import closure", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": `local b = require("b")
return { useB = b.f }
`,
			"b.tlua": `local c = require("c")
local function f()
	return c.unrelated()
end
return { f = f }
`,
			"c.tlua": wsDiagModuleC,
		}, "a.tlua")
		f.reconcile()
		f.storeAll("a.tlua", "b.tlua", "c.tlua")

		f.writeFile("c.tlua", `local function unrelated()
	return 2
end
return { unrelated = unrelated }
`)
		f.reconcile()

		f.assertInvalid("a.tlua", "b.tlua", "c.tlua")
	})

	t.Run("global type edit invalidates everything", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
			// A bare top-level type in a module is hoisted into the global
			// type table, so editing this file can change any other file.
			"d.tlua": `type Glob = number
local x: Glob = 1
return { x = x }
`,
		}, "a.tlua")
		f.reconcile()
		f.storeAll("a.tlua", "b.tlua", "c.tlua", "d.tlua")

		f.writeFile("d.tlua", `type Glob = string
local x: Glob = "1"
return { x = x }
`)
		f.reconcile()

		f.assertInvalid("a.tlua", "b.tlua", "c.tlua", "d.tlua")
	})

	t.Run("local type edit invalidates selectively", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
			// `local` keeps the type out of the global table.
			"d.tlua": `local type Priv = number
local x: Priv = 1
return { x = x }
`,
		}, "a.tlua")
		f.reconcile()
		ids := f.storeAll("a.tlua", "b.tlua", "c.tlua", "d.tlua")

		f.writeFile("d.tlua", `local type Priv = string
local x: Priv = "1"
return { x = x }
`)
		f.reconcile()

		f.assertInvalid("d.tlua")
		for _, name := range []string{"a.tlua", "b.tlua", "c.tlua"} {
			status := f.status(name)
			assert.Assert(t, status.Valid, "%s should stay valid", name)
			assert.Equal(t, status.ResultID, ids[name])
		}
	})

	// The Lua assignment forms below create global *values* the checker only
	// resolves in initializeLuaAugmentations: no declaration in the file's
	// locals names them, so the type-hoisting rules above cannot see them.
	t.Run("implicit global write invalidates everything", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
			// No `local Shared`: the bare store creates a global.
			"d.tlua": `Shared = 1
return { ok = 1 }
`,
		}, "a.tlua")
		f.reconcile()
		f.storeAll("a.tlua", "b.tlua", "c.tlua", "d.tlua")

		f.writeFile("d.tlua", `Shared = 2
return { ok = 1 }
`)
		f.reconcile()

		f.assertInvalid("a.tlua", "b.tlua", "c.tlua", "d.tlua")
	})

	t.Run("rebinding a local invalidates selectively", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
			// A bare store to a name a `local` declares stays local.
			"d.tlua": `local x = 0
x = 2
return { x = x }
`,
		}, "a.tlua")
		f.reconcile()
		ids := f.storeAll("a.tlua", "b.tlua", "c.tlua", "d.tlua")

		f.writeFile("d.tlua", `local x = 0
x = 3
return { x = x }
`)
		f.reconcile()

		f.assertInvalid("d.tlua")
		for _, name := range []string{"a.tlua", "b.tlua", "c.tlua"} {
			assert.Assert(t, f.status(name).Valid, "%s should stay valid", name)
			assert.Equal(t, f.status(name).ResultID, ids[name])
		}
	})

	t.Run("environment write invalidates everything", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
			"d.tlua": `_G.Flag = true
return { ok = 1 }
`,
		}, "a.tlua")
		f.reconcile()
		f.storeAll("a.tlua", "b.tlua", "c.tlua", "d.tlua")

		f.writeFile("d.tlua", `_G.Flag = false
return { ok = 1 }
`)
		f.reconcile()

		f.assertInvalid("a.tlua", "b.tlua", "c.tlua", "d.tlua")
	})

	t.Run("method write through a local invalidates selectively", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
			// The ordinary module idiom: every write is rooted at a local
			// table constructor, so nothing global comes of it.
			"d.tlua": `local M = {}
function M.helper()
	return 1
end
return M
`,
		}, "a.tlua")
		f.reconcile()
		ids := f.storeAll("a.tlua", "b.tlua", "c.tlua", "d.tlua")

		f.writeFile("d.tlua", `local M = {}
function M.helper()
	return 2
end
return M
`)
		f.reconcile()

		f.assertInvalid("d.tlua")
		for _, name := range []string{"a.tlua", "b.tlua", "c.tlua"} {
			assert.Assert(t, f.status(name).Valid, "%s should stay valid", name)
			assert.Equal(t, f.status(name).ResultID, ids[name])
		}
	})

	t.Run("write through a parameter invalidates selectively", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
			// A parameter is not a stable augmentation root.
			"d.tlua": `local function fill(t: { x: number })
	t.x = 1
end
return { fill = fill }
`,
		}, "a.tlua")
		f.reconcile()
		ids := f.storeAll("a.tlua", "b.tlua", "c.tlua", "d.tlua")

		f.writeFile("d.tlua", `local function fill(t: { x: number })
	t.x = 2
end
return { fill = fill }
`)
		f.reconcile()

		f.assertInvalid("d.tlua")
		for _, name := range []string{"a.tlua", "b.tlua", "c.tlua"} {
			assert.Assert(t, f.status(name).Valid, "%s should stay valid", name)
			assert.Equal(t, f.status(name).ResultID, ids[name])
		}
	})

	t.Run("augmenting a global through an alias invalidates everything", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
			"d.tlua": `Shared = {}
return { ok = 1 }
`,
			// The local alias resolves to nothing lexical, so the write lands
			// on the global `Shared` declares.
			"e.tlua": `local S = Shared
S.count = 1
return { count = S.count }
`,
		}, "a.tlua")
		f.reconcile()
		f.storeAll("a.tlua", "b.tlua", "c.tlua", "d.tlua", "e.tlua")

		f.writeFile("e.tlua", `local S = Shared
S.count = 2
return { count = S.count }
`)
		f.reconcile()

		f.assertInvalid("a.tlua", "b.tlua", "c.tlua", "d.tlua", "e.tlua")
	})

	t.Run("a parenthesized alias is as global as a bare one", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
			"d.tlua": `Shared = {}
return { ok = 1 }
`,
			// Parentheses are transparent to the checker's alias resolution, so
			// the write still augments the global.
			"e.tlua": `local S = (Shared)
S.count = 1
return { count = S.count }
`,
		}, "a.tlua")
		f.reconcile()
		f.storeAll("a.tlua", "b.tlua", "c.tlua", "d.tlua", "e.tlua")

		f.writeFile("e.tlua", `local S = (Shared)
S.count = 2
return { count = S.count }
`)
		f.reconcile()

		f.assertInvalid("a.tlua", "b.tlua", "c.tlua", "d.tlua", "e.tlua")
	})

	t.Run("a preference change invalidates every entry", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
		}, "a.tlua")
		f.reconcile()
		// The first observation of the preference is the session's baseline.
		f.cache.ReconcilePreferences(false)
		ids := f.storeAll("a.tlua", "b.tlua", "c.tlua")

		// Same value: nothing to redo.
		f.cache.ReconcilePreferences(false)
		f.assertValid("a.tlua", "b.tlua", "c.tlua")

		// The severities baked into the cached items are no longer the ones the
		// client would be given.
		f.cache.ReconcilePreferences(true)
		f.assertInvalid("a.tlua", "b.tlua", "c.tlua")
		// The result IDs survive, so an unchanged recompute can still dedup.
		for name, id := range ids {
			assert.Equal(t, f.status(name).ResultID, id)
		}

		f.cache.ReconcilePreferences(true)
		f.assertInvalid("a.tlua", "b.tlua", "c.tlua")
	})

	t.Run("an unchecked file is still tracked", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
		}, "a.tlua")
		f.reconcile()
		f.storeAll("a.tlua", "b.tlua", "require.d.tlua")

		// The shim is a declaration script, so its content is global: editing
		// it invalidates everything.
		f.writeFile("require.d.tlua", `declare function require(module: string): any;
declare function print(value: any): nil;
`)
		f.reconcile()

		f.assertInvalid("a.tlua", "b.tlua", "require.d.tlua")
	})

	t.Run("removing a global contribution still invalidates everything", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
			"d.tlua": `type Glob = number
local x: Glob = 1
return { x = x }
`,
		}, "a.tlua")
		f.reconcile()
		f.storeAll("a.tlua", "b.tlua", "c.tlua", "d.tlua")

		// The new shape contributes nothing global, but the old one did.
		f.writeFile("d.tlua", `local x = 1
return { x = x }
`)
		f.reconcile()

		f.assertInvalid("a.tlua", "b.tlua", "c.tlua", "d.tlua")

		// And now that neither shape contributes globals, the next edit is
		// selective again.
		f.storeAll("a.tlua", "b.tlua", "c.tlua", "d.tlua")
		f.writeFile("d.tlua", `local x = 2
return { x = x }
`)
		f.reconcile()
		f.assertInvalid("d.tlua")
		f.assertValid("a.tlua", "b.tlua", "c.tlua")
	})

	t.Run("file added and removed invalidates everything", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
		}, "a.tlua")
		f.reconcile()
		ids := f.storeAll("a.tlua", "b.tlua", "c.tlua")

		assert.NilError(t, f.utils.FS().WriteFile("/project/e.tlua", `local e = 1
return { e = e }
`))
		f.watchedEvent("e.tlua", lsproto.FileChangeTypeCreated)
		f.reconcile()

		assert.Assert(t, f.cached("e.tlua"), "the new file should be tracked")
		f.assertInvalid("a.tlua", "b.tlua", "c.tlua", "e.tlua")
		// A surviving file keeps its result ID across a full invalidation, so
		// recompute-dedup can still answer `unchanged`.
		assert.Equal(t, f.status("a.tlua").ResultID, ids["a.tlua"])
		assert.Equal(t, f.status("e.tlua").ResultID, "")

		f.storeAll("a.tlua", "b.tlua", "c.tlua", "e.tlua")
		assert.NilError(t, f.utils.FS().Remove("/project/e.tlua"))
		f.watchedEvent("e.tlua", lsproto.FileChangeTypeDeleted)
		f.reconcile()

		assert.Assert(t, !f.cached("e.tlua"), "the removed file's entry should be dropped")
		f.assertInvalid("a.tlua", "b.tlua", "c.tlua")
	})

	t.Run("a change to an excluded file invalidates everything", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": `local pkg = require("node_modules.pkg.index")
return { v = pkg.f() }
`,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
			// An installed package: part of the program, never reported on, and
			// so not an entry the hash diff would look at. It stands in for the
			// case that motivates tracking excluded hashes at all -- a project
			// reference's source, edited in the project that owns it.
			"node_modules/pkg/index.tlua": wsDiagModuleB,
		}, "a.tlua")
		f.reconcile()
		assert.Assert(t, !f.cached("node_modules/pkg/index.tlua"),
			"an installed package is not the user's to fix")
		f.storeAll("a.tlua", "b.tlua", "c.tlua")

		f.writeFile("node_modules/pkg/index.tlua", `local function f()
	return "1"
end
return { f = f }
`)
		f.reconcile()

		f.assertInvalid("a.tlua", "b.tlua", "c.tlua")
	})

	t.Run("compiler option change invalidates everything", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
		}, "a.tlua")
		f.reconcile()
		f.storeAll("a.tlua", "b.tlua", "c.tlua")

		assert.NilError(t, f.utils.FS().WriteFile("/project/tluaconfig.json", `{
			"compilerOptions": {
				"noLib": true,
				"skipLibCheck": true,
				"noImplicitAny": false
			}
		}`))
		f.watchedEvent("tluaconfig.json", lsproto.FileChangeTypeChanged)
		f.reconcile()

		f.assertInvalid("a.tlua", "b.tlua", "c.tlua")
	})

	t.Run("StoreComputed dedups identical diagnostics", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
		}, "a.tlua")
		f.reconcile()
		ids := f.storeAll("a.tlua", "b.tlua", "c.tlua")

		// A whitespace-only edit: the hash moves, the diagnostics do not.
		f.writeFile("c.tlua", wsDiagModuleC+"\n")
		f.reconcile()
		f.assertInvalid("c.tlua")

		resultID, changed := f.store("c.tlua", wsDiagItems("c.tlua"))
		assert.Equal(t, resultID, ids["c.tlua"], "an identical recompute must keep the result ID")
		assert.Assert(t, !changed)
		assert.Assert(t, f.status("c.tlua").Valid)

		// Different diagnostics do mint a new ID.
		f.writeFile("c.tlua", wsDiagModuleC+"\n\n")
		f.reconcile()
		resultID, changed = f.store("c.tlua", wsDiagItems("c.tlua", "extra"))
		assert.Assert(t, changed)
		assert.Assert(t, resultID != ids["c.tlua"])
	})

	t.Run("StoreComputed rejects a stale computation", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
		}, "a.tlua")
		f.reconcile()
		f.storeAll("a.tlua", "b.tlua", "c.tlua")

		// A pull reads the status, then the program moves before the check
		// finishes.
		stale := f.status("c.tlua")
		f.writeFile("c.tlua", wsDiagModuleC+"\n")
		f.reconcile()
		before := f.status("c.tlua")
		assert.Assert(t, !before.Valid)

		resultID, changed, ok := f.cache.StoreComputed(f.project.Id(), f.path("c.tlua"), stale.Generation, stale.FileHash, wsDiagItems("stale"))
		assert.Assert(t, !ok, "a stale generation must be refused")
		assert.Equal(t, resultID, "")
		assert.Assert(t, !changed)

		after := f.status("c.tlua")
		assert.Assert(t, !after.Valid, "the refused store must not validate the entry")
		assert.Equal(t, after.ResultID, before.ResultID)

		// The current generation with a stale hash is refused too.
		_, _, ok = f.cache.StoreComputed(f.project.Id(), f.path("c.tlua"), after.Generation, stale.FileHash, wsDiagItems("stale"))
		assert.Assert(t, !ok, "a stale file hash must be refused")
		assert.Assert(t, !f.status("c.tlua").Valid)
	})

	t.Run("overlay edit without a save invalidates", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
		}, "a.tlua")
		f.reconcile()
		ids := f.storeAll("a.tlua", "b.tlua", "c.tlua")

		f.changeOverlay(`local b = require("b")
local x: number = "oops"
return { useB = b.f, x = x }
`)
		f.reconcile()

		f.assertInvalid("a.tlua")
		for _, name := range []string{"b.tlua", "c.tlua"} {
			status := f.status(name)
			assert.Assert(t, status.Valid, "%s should stay valid", name)
			assert.Equal(t, status.ResultID, ids[name])
		}
	})

	t.Run("MarkReported", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
		}, "a.tlua")
		f.reconcile()
		ids := f.storeAll("a.tlua", "b.tlua")

		assert.Assert(t, !f.status("a.tlua").Reported)
		f.cache.MarkReported(f.project.Id(), f.path("a.tlua"), "some-other-id")
		assert.Assert(t, !f.status("a.tlua").Reported, "a stale result ID must not mark the entry reported")
		f.cache.MarkReported(f.project.Id(), f.path("a.tlua"), ids["a.tlua"])
		assert.Assert(t, f.status("a.tlua").Reported)

		// The reported flag survives invalidation, so a file that goes clean
		// can still be cleared on the client. a.tlua is the open document, so
		// the edit has to go through the overlay to be seen.
		f.changeOverlay(wsDiagModuleA + "\n")
		f.reconcile()
		status := f.status("a.tlua")
		assert.Assert(t, !status.Valid)
		assert.Assert(t, status.Reported)
	})

	t.Run("a local _G shadows the environment", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
			"c.tlua": wsDiagModuleC,
			// The write goes through a file-local table that happens to be
			// named _G; the checker resolves it lexically, and so does the
			// detector — this file contributes nothing global.
			"d.tlua": `local _G = {}
_G.localOnly = true
return { ok = _G.localOnly }
`,
		}, "a.tlua")
		f.reconcile()
		f.storeAll("a.tlua", "b.tlua", "c.tlua", "d.tlua")

		f.writeFile("d.tlua", `local _G = {}
_G.localOnly = false
return { ok = _G.localOnly }
`)
		f.reconcile()

		f.assertInvalid("d.tlua")
		f.assertValid("a.tlua", "b.tlua", "c.tlua")
	})

	t.Run("Prune drops projects missing from the snapshot", func(t *testing.T) {
		t.Parallel()
		f := newWSDiagFixture(t, map[string]string{
			"a.tlua": wsDiagModuleA,
			"b.tlua": wsDiagModuleB,
		}, "a.tlua")
		f.reconcile()
		f.storeAll("a.tlua", "b.tlua")

		// A nil snapshot and a snapshot that still has the project both leave
		// the state alone.
		f.cache.Prune(nil)
		f.cache.Prune(f.session.Snapshot())
		assert.Assert(t, f.cached("a.tlua"))

		// A snapshot without the project drops it.
		empty, _ := projecttestutil.Setup(wsDiagFiles(map[string]string{"a.tlua": wsDiagModuleA}))
		emptySnapshot := empty.Snapshot()
		assert.Equal(t, len(emptySnapshot.ProjectCollection.Projects()), 0)
		f.cache.Prune(emptySnapshot)
		assert.Assert(t, !f.cached("a.tlua"))
		assert.Assert(t, !f.cached("b.tlua"))
	})
}

func TestWorkspaceDiagCacheConcurrentAccess(t *testing.T) {
	t.Parallel()

	sources := map[string]string{
		"a.tlua": wsDiagModuleA,
		"b.tlua": wsDiagModuleB,
		"c.tlua": wsDiagModuleC,
	}
	f := newWSDiagFixture(t, sources, "a.tlua")
	f.reconcile()
	f.storeAll("a.tlua", "b.tlua", "c.tlua")

	names := []string{"a.tlua", "b.tlua", "c.tlua"}
	proj := f.project
	projectID := proj.Id()

	var wg sync.WaitGroup
	for i := range 8 {
		wg.Go(func() {
			for j := range 10 {
				name := names[(i+j)%len(names)]
				path := f.path(name)
				switch j % 4 {
				case 0:
					f.cache.Reconcile(proj)
				case 1:
					f.cache.Status(projectID, path)
				case 2:
					if status, ok := f.cache.Status(projectID, path); ok {
						f.cache.StoreComputed(projectID, path, status.Generation, status.FileHash, wsDiagItems(fmt.Sprintf("%s-%d", name, i)))
					}
				case 3:
					if status, ok := f.cache.Status(projectID, path); ok {
						f.cache.MarkReported(projectID, path, status.ResultID)
					}
				}
			}
		})
	}
	wg.Wait()

	for _, name := range names {
		status, ok := f.cache.Status(projectID, f.path(name))
		assert.Assert(t, ok)
		assert.Assert(t, status.ResultID != "")
	}
}
