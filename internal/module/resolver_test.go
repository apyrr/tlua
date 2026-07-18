package module_test

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/module"
	"github.com/apyrr/tlua/internal/vfs"
	"github.com/apyrr/tlua/internal/vfs/vfstest"
)

type resolutionHostStub struct {
	fs  vfs.FS
	cwd string
}

func (h *resolutionHostStub) FS() vfs.FS                  { return h.fs }
func (h *resolutionHostStub) GetCurrentDirectory() string { return h.cwd }

// A bare package name is a valid Lua module name and resolves through
// node_modules (here via package.json `main`). A path-shaped spelling like
// `pkg/` is not a Lua module name and must not resolve at all.
func TestResolveModuleNamePackageMain(t *testing.T) {
	t.Parallel()

	fs := vfstest.FromMap(map[string]string{
		"/repo/node_modules/pkg/package.json": `{"name":"pkg","main":"main.lua"}`,
		"/repo/node_modules/pkg/main.lua":     "exports.x = 1;",
		"/repo/src/file.tlua":                 "",
	}, true)
	host := &resolutionHostStub{fs: fs, cwd: "/repo"}
	opts := &core.CompilerOptions{
		Module: core.ModuleKindESNext,
		Target: core.ScriptTargetESNext,
	}
	resolver := module.NewResolver(host, opts, "", "")

	r, _ := resolver.ResolveModuleName("pkg", "/repo/src/file.tlua", core.ModuleKindESNext, nil)
	if !r.IsResolved() {
		t.Error(`"pkg" failed to resolve`)
	}
	r, _ = resolver.ResolveModuleName("pkg/", "/repo/src/file.tlua", core.ModuleKindESNext, nil)
	if r.IsResolved() {
		t.Error(`path-shaped "pkg/" resolved; it is not a Lua module name`)
	}
}

// blockingFS wraps a vfs.FS and forces FileExists calls for `targetPath` to
// block on `gate` until released. Each caller sends on `arrived` when it
// reaches the gate. This is used to deterministically reproduce the
// `package.json` info-cache insert race described in
// https://github.com/microsoft/typescript-go/issues/3526.
type blockingFS struct {
	vfs.FS
	targetPath string
	gate       chan struct{}
	arrived    chan struct{} // each blocked goroutine sends one value
}

// waitForSignal waits for a synchronization point in these race regression
// tests and converts deadlocks into deterministic test failures.
func waitForSignal(t *testing.T, ch <-chan struct{}, description string) {
	t.Helper()
	select {
	case <-ch:
		return
	case <-t.Context().Done():
		t.Fatalf("timed out waiting for %s", description)
	}
}

func (f *blockingFS) FileExists(path string) bool {
	if path == f.targetPath {
		f.arrived <- struct{}{}
		<-f.gate
	}
	return f.FS.FileExists(path)
}

// flipFileExistsFS wraps a vfs.FS and returns false for the first
// FileExists call to `targetPath`, then true for the second. Both calls
// signal arrival via channel then block until released via their respective
// gate channels. ReadFile for the target path also signals arrival then
// blocks, so the "file doesn't exist" Set completes before the "file exists"
// Set (reproducing the LoadOrStore race).
type flipFileExistsFS struct {
	vfs.FS
	targetPath    string
	callCount     atomic.Int32
	firstArrived  chan struct{} // closed when the first FileExists caller arrives
	secondArrived chan struct{} // closed when the second FileExists caller arrives
	firstGate     chan struct{}
	secondGate    chan struct{}
	readArrived   chan struct{} // closed when ReadFile caller arrives
	readGate      chan struct{}
}

func (f *flipFileExistsFS) FileExists(path string) bool {
	if path == f.targetPath {
		n := f.callCount.Add(1)
		if n == 1 {
			close(f.firstArrived)
			<-f.firstGate
			return false // first caller: simulate "file not yet visible"
		}
		if n == 2 {
			close(f.secondArrived)
			<-f.secondGate
			return f.FS.FileExists(path) // second caller: file is visible
		}
	}
	return f.FS.FileExists(path)
}

func (f *flipFileExistsFS) ReadFile(path string) (string, bool) {
	if path == f.targetPath {
		close(f.readArrived)
		<-f.readGate
	}
	return f.FS.ReadFile(path)
}

// Regression test for the package.json info-cache insert race described in
// https://github.com/microsoft/typescript-go/issues/3526 (re-expressed for
// Lua module names: the trailing-slash spelling that originally triggered
// the mismatch is now rejected before resolution, but the concurrent
// LoadOrStore on the info cache is still live for bare package names).
//
// Two goroutines resolve `pkg` from different containing directories. A
// blocking FS holds both at the `FileExists` check for `package.json` —
// *after* each has confirmed a `package.json` info-cache miss but *before*
// either has called `Set`. When released, both proceed to `LoadOrStore` and
// one of them loses; the loser must still resolve through the winner's
// cached entry.
func TestResolvePackageJsonInfoCacheRace(t *testing.T) {
	t.Parallel()

	const pkgJSONPath = "/repo/node_modules/pkg/package.json"
	files := map[string]string{
		// `main` points at a file that is not discoverable through any
		// fallback path: there is no `index.*` and no `init.*`. The only way
		// to resolve `pkg` is via the package.json `main` field inside
		// `loadNodeModuleFromDirectoryWorker`.
		pkgJSONPath:                            `{"name":"pkg","main":"./lib/entry.lua"}`,
		"/repo/node_modules/pkg/lib/entry.lua": "exports.x = 1;",
		// Distinct containing files so each `ResolveModuleName` call has a
		// unique module-resolution-cache key.
		"/repo/src/a/file.tlua": "",
		"/repo/src/b/file.tlua": "",
	}
	fs := &blockingFS{
		FS:         vfstest.FromMap(files, true),
		targetPath: pkgJSONPath,
		gate:       make(chan struct{}),
		arrived:    make(chan struct{}, 2),
	}
	host := &resolutionHostStub{fs: fs, cwd: "/repo"}
	opts := &core.CompilerOptions{
		Module: core.ModuleKindESNext,
		Target: core.ScriptTargetESNext,
	}
	resolver := module.NewResolver(host, opts, "", "")

	type resolutionResult struct {
		name     string
		resolved bool
	}
	results := make(chan resolutionResult, 2)
	var wg sync.WaitGroup
	for _, containingFile := range []string{"/repo/src/a/file.tlua", "/repo/src/b/file.tlua"} {
		wg.Go(func() {
			r, _ := resolver.ResolveModuleName("pkg", containingFile, core.ModuleKindESNext, nil)
			results <- resolutionResult{containingFile, r.IsResolved()}
		})
	}

	// Wait for both goroutines to reach the FileExists gate, guaranteeing
	// both have observed a package.json info-cache miss.
	waitForSignal(t, fs.arrived, "first FileExists gate arrival")
	waitForSignal(t, fs.arrived, "second FileExists gate arrival")
	close(fs.gate)

	wg.Wait()
	close(results)
	for r := range results {
		if !r.resolved {
			t.Errorf("%q failed to resolve", r.name)
		}
	}
}

// Regression test for https://github.com/microsoft/typescript-go/issues/1290
// (re-expressed for Lua module names: `pkg/sub` became the dotted `pkg.sub`,
// which maps into the package tree and consults the package.json of the
// directory it lands on).
//
// Two goroutines resolve `pkg.sub` concurrently. Both miss the package.json
// info-cache for the `pkg/sub` directory. A `flipFileExistsFS` forces the
// first goroutine's `FileExists` to return false (simulating the file not yet
// being visible), so it stores a nil-Contents cache entry. The second
// goroutine's `FileExists` returns true, but its `Set` call (`LoadOrStore`)
// returns the first goroutine's nil-Contents entry. Without the `Exists()`
// guard on the `typesVersions` lookup, `packageInfo.Contents.GetVersionPaths`
// dereferences nil and panics. With the guard the nil-Contents entry is safely
// skipped.
func TestResolveSubpathNilContentsRace(t *testing.T) {
	t.Parallel()

	const subPkgJSON = "/repo/node_modules/pkg/sub/package.json"
	files := map[string]string{
		subPkgJSON:                             `{"name":"pkg-sub","version":"1.0.0"}`,
		"/repo/node_modules/pkg/sub/index.lua": "exports.sub = 1;",
		"/repo/src/a/file.tlua":                "",
		"/repo/src/b/file.tlua":                "",
	}
	fs := &flipFileExistsFS{
		FS:            vfstest.FromMap(files, true),
		targetPath:    subPkgJSON,
		firstArrived:  make(chan struct{}),
		secondArrived: make(chan struct{}),
		firstGate:     make(chan struct{}),
		secondGate:    make(chan struct{}),
		readArrived:   make(chan struct{}),
		readGate:      make(chan struct{}),
	}
	host := &resolutionHostStub{fs: fs, cwd: "/repo"}
	opts := &core.CompilerOptions{
		Module: core.ModuleKindESNext,
		Target: core.ScriptTargetESNext,
	}
	resolver := module.NewResolver(host, opts, "", "")

	var panicked atomic.Bool
	type resolutionResult struct {
		containingFile string
		resolved       bool
	}
	results := make(chan resolutionResult, 2)
	var wg sync.WaitGroup
	// Two goroutines both resolve "pkg.sub". Each calls getPackageJsonInfo
	// for the pkg/sub directory, reaching FileExists for subPkgJSON.
	for _, containingFile := range []string{"/repo/src/a/file.tlua", "/repo/src/b/file.tlua"} {
		wg.Go(func() {
			resolved := false
			defer func() {
				if r := recover(); r != nil {
					panicked.Store(true)
				}
				results <- resolutionResult{containingFile: containingFile, resolved: resolved}
			}()
			r, _ := resolver.ResolveModuleName("pkg.sub", containingFile, core.ModuleKindESNext, nil)
			resolved = r.IsResolved()
		})
	}

	// Phase 1: Wait for both goroutines to reach FileExists for the root
	// package.json, guaranteeing both have observed a cache miss.
	waitForSignal(t, fs.firstArrived, "first pkg/sub package.json FileExists arrival")
	waitForSignal(t, fs.secondArrived, "second pkg/sub package.json FileExists arrival")

	// Phase 2: Release the first FileExists caller (returns false).
	// It enters the "file not found" branch and stores a nil-Contents entry
	// via Set — this is nearly instant (no ReadFile).
	close(fs.firstGate)

	// Phase 3: Release the second FileExists caller (returns true).
	// It proceeds to ReadFile, which we gate separately to ensure the first
	// goroutine's nil-Contents Set has completed.
	close(fs.secondGate)

	// Phase 4: Wait for the second goroutine to reach ReadFile, then release.
	// By this point the first goroutine has stored its nil-Contents entry.
	// The second goroutine's Set (LoadOrStore) will return that stale entry.
	waitForSignal(t, fs.readArrived, "pkg/sub package.json ReadFile arrival")
	close(fs.readGate)

	wg.Wait()
	close(results)
	if panicked.Load() {
		t.Fatal("resolver panicked due to nil Contents dereference in loadModuleFromSpecificNodeModulesDirectory")
	}
	for r := range results {
		if !r.resolved {
			t.Fatalf("%q failed to resolve pkg.sub", r.containingFile)
		}
	}
}

func TestParseNodeModuleFromPath(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		path     string
		isFolder bool
		want     string
	}{
		{"file in package", "/a/node_modules/b/lib/index.d.tlua", false, "/a/node_modules/b"},
		{"file in scoped package", "/a/node_modules/@scope/b/lib/index.d.tlua", false, "/a/node_modules/@scope/b"},
		{"folder subpath", "/a/node_modules/b/lib/File", true, "/a/node_modules/b"},
		{"folder subpath scoped", "/a/node_modules/@scope/b/lib/File", true, "/a/node_modules/@scope/b"},
		{"package root folder", "/a/node_modules/b", true, "/a/node_modules/b"},
		{"scoped package root folder", "/a/node_modules/@scope/b", true, "/a/node_modules/@scope/b"},
		// A bare scope directory has no package name; must not panic (https://github.com/microsoft/typescript-go/issues/4373).
		{"scope-only folder", "/a/node_modules/@scope", true, "/a/node_modules/@scope"},
		{"types scope-only folder", "/a/node_modules/@types", true, "/a/node_modules/@types"},
		{"not in node_modules", "/a/src/index.tlua", false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := module.ParseNodeModuleFromPath(tt.path, tt.isFolder); got != tt.want {
				t.Errorf("ParseNodeModuleFromPath(%q, %v) = %q, want %q", tt.path, tt.isFolder, got, tt.want)
			}
		})
	}
}

// Regression test for https://github.com/microsoft/typescript-go/issues/4478
// (re-expressed through type-reference resolution: require() no longer
// computes package IDs, but type-reference directives still resolve through
// the node_modules machinery that reads peerDependencies).
//
// While resolving a package with peerDependencies, two goroutines look up the
// peer package's package.json concurrently. A `flipFileExistsFS` forces the
// first lookup to cache a nil-Contents entry and the second lookup to receive
// that stale entry from `Set`. The resolver must not dereference the peer
// package.json contents unless the entry actually exists.
func TestResolvePeerDependencyNilContentsRace(t *testing.T) {
	t.Parallel()

	const peerPkgJSON = "/repo/node_modules/peer/package.json"
	files := map[string]string{
		"/repo/node_modules/pkg/package.json": `{"name":"pkg","version":"1.0.0","types":"index.d.tlua","peerDependencies":{"peer":"*"}}`,
		"/repo/node_modules/pkg/index.d.tlua": "export declare local x: number;",
		peerPkgJSON:                           `{"name":"peer","version":"2.0.0"}`,
		"/repo/src/a/file.tlua":               "",
		"/repo/src/b/file.tlua":               "",
	}
	fs := &flipFileExistsFS{
		FS:            vfstest.FromMap(files, true),
		targetPath:    peerPkgJSON,
		firstArrived:  make(chan struct{}),
		secondArrived: make(chan struct{}),
		firstGate:     make(chan struct{}),
		secondGate:    make(chan struct{}),
		readArrived:   make(chan struct{}),
		readGate:      make(chan struct{}),
	}
	host := &resolutionHostStub{fs: fs, cwd: "/repo"}
	opts := &core.CompilerOptions{
		Module: core.ModuleKindESNext,
		Target: core.ScriptTargetESNext,
	}
	resolver := module.NewResolver(host, opts, "", "")

	var panicked atomic.Bool
	type resolutionResult struct {
		containingFile string
		resolved       bool
	}
	results := make(chan resolutionResult, 2)
	var wg sync.WaitGroup
	for _, containingFile := range []string{"/repo/src/a/file.tlua", "/repo/src/b/file.tlua"} {
		wg.Go(func() {
			resolved := false
			defer func() {
				if r := recover(); r != nil {
					panicked.Store(true)
				}
				results <- resolutionResult{containingFile: containingFile, resolved: resolved}
			}()
			r, _ := resolver.ResolveTypeReferenceDirective("pkg", containingFile, core.ModuleKindNone, nil)
			resolved = r.IsResolved()
		})
	}

	waitForSignal(t, fs.firstArrived, "first peer package.json FileExists arrival")
	waitForSignal(t, fs.secondArrived, "second peer package.json FileExists arrival")
	close(fs.firstGate)
	var firstResult resolutionResult
	select {
	case result := <-results:
		firstResult = result
	case <-t.Context().Done():
		t.Fatal("timed out waiting for first peer package.json lookup to finish")
	}
	close(fs.secondGate)
	waitForSignal(t, fs.readArrived, "peer package.json ReadFile arrival")
	close(fs.readGate)

	wg.Wait()
	close(results)
	if panicked.Load() {
		t.Fatal("resolver panicked due to nil Contents dereference in readPackageJsonPeerDependencies")
	}
	if !firstResult.resolved {
		t.Fatalf("%q failed to resolve pkg", firstResult.containingFile)
	}
	for r := range results {
		if !r.resolved {
			t.Fatalf("%q failed to resolve pkg", r.containingFile)
		}
	}
}
