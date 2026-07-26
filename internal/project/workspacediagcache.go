package project

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"maps"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/collections"
	"github.com/apyrr/tlua/internal/compiler"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/json"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/tspath"
	"github.com/zeebo/xxh3"
)

// WorkspaceDiagCache backs the `workspace/diagnostic` pull handler.
//
// The client re-pulls the whole workspace every couple of seconds, so the
// common case -- nothing changed since the last pull -- must be near free. The
// cache therefore does no work of its own in the background: every pull calls
// Reconcile, which diffs the project's current program against what was cached
// the last time the project was reconciled and marks the affected files
// invalid. The handler then recomputes diagnostics only for invalid files and
// answers `unchanged` (with the previously handed-out result ID) for the rest.
//
// Reconcile's fast path is a generation check: Project.ProgramLastUpdate is the
// snapshot ID that built the program, so an untouched project carries the same
// generation across pulls and the diff is skipped entirely. When the generation
// moves, the diff is driven by content hashes (ast.SourceFile.Hash) rather than
// declaration signatures: d.tlua signatures are not available for Lua modules
// (TS100054), so there is no signature stopper and invalidation walks the full
// transitive reverse-import closure of every changed file. A change to a file
// that contributes to the global scope, a change to the file set, or a change
// of compiler options invalidates the whole project.
//
// All state lives behind mu. Nothing under mu type-checks or touches the file
// system; the expensive parts of a pull (running the checker) happen in the
// handler, between Reconcile and StoreComputed.
type WorkspaceDiagCache struct {
	mu       sync.Mutex
	projects map[tspath.Path]*projectDiagState
	nextID   atomic.Uint64
	// nonce distinguishes result IDs minted by this process from result IDs a
	// client kept across a server restart.
	nonce string
	// prefsSet records whether styleChecksAsWarnings has been observed yet, so
	// the first pull of a session does not have to guess a baseline.
	prefsSet bool
	// styleChecksAsWarnings is the preference the cached items were converted
	// with (see ReconcilePreferences).
	styleChecksAsWarnings bool
}

// projectDiagState is the cached diagnostic state of one project, keyed by the
// project's configFilePath (stable across snapshots, synthetic but stable for
// the inferred project).
type projectDiagState struct {
	// generation is the Project.ProgramLastUpdate this state was last
	// reconciled against.
	generation uint64
	// options is compared by pointer identity: a new options object means the
	// program was reconfigured, and every file is invalidated.
	options *core.CompilerOptions
	files   map[tspath.Path]*fileDiagEntry
	// excluded holds the content hash of every program file the pull does not
	// report on. Those files are not tracked as entries, but they are part of
	// the program: a project reference's source that the user edits in its own
	// project, or a changed file under node_modules, moves the generation
	// without appearing in the diff, and the cache would answer `unchanged`
	// with diagnostics computed against the old content. Any difference here is
	// a full invalidation.
	excluded map[tspath.Path]xxh3.Uint128
	// revDeps maps an importee's path to the paths of the files importing it.
	// Built lazily, only when a selective invalidation needs it.
	revDeps    map[tspath.Path][]tspath.Path
	revDepsGen uint64
}

// fileDiagEntry is the cached diagnostic state of one checked file.
type fileDiagEntry struct {
	// fileHash is the ast.SourceFile.Hash the items were computed against.
	fileHash           xxh3.Uint128
	contributesGlobals bool
	// resultID is the opaque ID handed to the client for these items. Empty
	// until diagnostics have been computed once.
	resultID string
	// items is nil when the file has no diagnostics.
	items []*lsproto.Diagnostic
	// valid is false when the items must be recomputed before answering.
	valid bool
	// reported records that a full report was emitted for this file at least
	// once, so a file that becomes clean can be cleared on the client.
	reported bool
}

// FileDiagStatus is a copy-out snapshot of one file's cache entry.
type FileDiagStatus struct {
	Valid    bool
	ResultID string
	// Items is shared with the cache. Callers must not mutate it or its
	// elements.
	Items      []*lsproto.Diagnostic
	Reported   bool
	Generation uint64
	FileHash   xxh3.Uint128
}

// NewWorkspaceDiagCache creates an empty cache with a fresh per-process nonce.
func NewWorkspaceDiagCache() *WorkspaceDiagCache {
	var buf [8]byte
	// crypto/rand.Read never fails on any supported platform.
	_, _ = rand.Read(buf[:])
	return &WorkspaceDiagCache{
		projects: make(map[tspath.Path]*projectDiagState),
		nonce:    hex.EncodeToString(buf[:]),
	}
}

// IsWorkspaceDiagnosticFile reports whether a file's diagnostics are reported by
// a `workspace/diagnostic` pull. It is the single definition of the pull's file
// set: the cache tracks exactly these files, so the handler must ask the same
// question or its Status lookups will miss.
//
// Files the program does not type-check (`noCheck`, or a user declaration file
// under skipLibCheck) are deliberately still in the set: the compiler's own
// per-file pass gates itself on SkipTypeChecking
// (Program.getBindAndCheckDiagnosticsWithChecker, internal/compiler/program.go),
// so an unchecked file contributes exactly its syntactic diagnostics instead of
// disappearing from the pull. That matches what a `textDocument/diagnostic`
// pull answers for the same file; excluding them here silently dropped syntax
// errors of noCheck projects and of the user's own declaration files.
//
// What is excluded is what is not the user's to fix, or not this project's to
// report.
func IsWorkspaceDiagnosticFile(program *compiler.Program, file *ast.SourceFile) bool {
	// Bundled default libraries are excluded unconditionally: skipLibCheck is
	// not on by default for tluaconfig.json, and without this a cold pull would
	// type-check the whole lualib once per program generation just to report
	// nothing (lib diagnostics are not the user's to fix).
	if program.IsSourceFileDefaultLibrary(file.Path()) {
		return false
	}
	// A project reference's sources are reported by the project that owns them,
	// not by every project that consumes their output.
	if program.IsSourceFromProjectReference(file.Path()) {
		return false
	}
	// The literal /node_modules/ test mirrors Program.collectPackageNames:
	// locally installed typings are added as root files and so do not pass
	// IsSourceFileFromExternalLibrary.
	return !program.IsSourceFileFromExternalLibrary(file) &&
		!strings.Contains(file.FileName(), "/node_modules/")
}

// Reconcile diffs the project's current program against the cached state and
// marks the entries whose diagnostics can no longer be reused invalid. It is
// the only place cache invalidation happens; call it once per project per pull,
// before reading any Status.
func (c *WorkspaceDiagCache) Reconcile(project *Project) {
	program := project.GetProgram()
	if program == nil {
		return
	}
	gen := project.ProgramLastUpdate
	opts := program.CommandLine().CompilerOptions()

	// Cheap generation check first: an untouched project reconciles to nothing.
	c.mu.Lock()
	if state, ok := c.projects[project.configFilePath]; ok && state.generation == gen && state.options == opts {
		c.mu.Unlock()
		return
	}
	c.mu.Unlock()

	// Collecting the reported files touches no cache state, so do it outside the
	// lock.
	cur := make(map[tspath.Path]*ast.SourceFile, len(program.GetSourceFiles()))
	excludedCur := make(map[tspath.Path]xxh3.Uint128)
	for _, file := range program.GetSourceFiles() {
		if !IsWorkspaceDiagnosticFile(program, file) {
			excludedCur[file.Path()] = file.Hash
			continue
		}
		cur[file.Path()] = file
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	state, ok := c.projects[project.configFilePath]
	if !ok {
		state = &projectDiagState{files: make(map[tspath.Path]*fileDiagEntry)}
		c.projects[project.configFilePath] = state
	} else if state.generation == gen && state.options == opts {
		// Another pull reconciled this project while we were collecting files.
		return
	}

	// Options change, any file added, removed or renamed, or a change to a file
	// the pull does not report on: everything goes. The excluded set is almost
	// always the bundled libraries, whose hashes never move, so the common case
	// is one map comparison per program generation.
	full := state.options != opts || len(state.files) != len(cur) ||
		!maps.Equal(state.excluded, excludedCur)
	if !full {
		for path := range state.files {
			if _, ok := cur[path]; !ok {
				full = true
				break
			}
		}
	}

	var changed []tspath.Path
	if !full {
		for path, file := range cur {
			if state.files[path].fileHash != file.Hash {
				changed = append(changed, path)
			}
		}
		// Global-scope rule, on the union of the old and new shape of each
		// changed file: a file that contributed globals before, or contributes
		// them now, can affect the meaning of every other file.
		for _, path := range changed {
			if state.files[path].contributesGlobals || fileContributesToGlobalScope(cur[path]) {
				full = true
				break
			}
		}
	}

	switch {
	case full:
		fullInvalidate(state, cur)
	case len(changed) > 0:
		if state.revDeps == nil || state.revDepsGen != gen {
			state.revDeps = buildReverseImportGraph(program, cur)
			state.revDepsGen = gen
		}
		// Transitive reverse-dependency closure, including the changed files
		// themselves. There is no signature stopper: Lua modules have no usable
		// declaration signature, so any importer of a changed file may see a
		// different type.
		seen := collections.Set[tspath.Path]{}
		queue := make([]tspath.Path, 0, len(changed))
		for _, path := range changed {
			if seen.AddIfAbsent(path) {
				queue = append(queue, path)
			}
		}
		for len(queue) > 0 {
			path := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			if entry := state.files[path]; entry != nil {
				entry.valid = false
			}
			for _, importer := range state.revDeps[path] {
				if seen.AddIfAbsent(importer) {
					queue = append(queue, importer)
				}
			}
		}
		for _, path := range changed {
			entry, file := state.files[path], cur[path]
			entry.fileHash = file.Hash
			entry.contributesGlobals = fileContributesToGlobalScope(file)
		}
	}

	state.generation = gen
	state.options = opts
	state.excluded = excludedCur
}

// ReconcilePreferences invalidates every cached entry when a preference that is
// baked into the converted diagnostics changes. ReportStyleChecksAsWarnings is
// the only such preference: client capabilities are fixed for the session, and
// EnableValidation is handled by the pull itself.
//
// The entries keep their result ID, items and reported flag: recompute-dedup
// compares the newly converted items against them, so a new ID is minted
// exactly where a severity actually moved.
func (c *WorkspaceDiagCache) ReconcilePreferences(reportStyleChecksAsWarnings bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.prefsSet && c.styleChecksAsWarnings == reportStyleChecksAsWarnings {
		return
	}
	c.prefsSet = true
	c.styleChecksAsWarnings = reportStyleChecksAsWarnings
	for _, state := range c.projects {
		for _, entry := range state.files {
			entry.valid = false
		}
	}
}

// fullInvalidate rebuilds state.files against the program's current file set,
// invalidating every entry. Entries for surviving files are reused so their
// result ID and reported flag are preserved: recompute-dedup still recognizes
// unchanged diagnostics, and a file that becomes clean can still be cleared on
// the client. Entries for files that left the program are dropped, so a file
// that comes back is always sent a full report.
func fullInvalidate(state *projectDiagState, cur map[tspath.Path]*ast.SourceFile) {
	files := make(map[tspath.Path]*fileDiagEntry, len(cur))
	for path, file := range cur {
		entry := state.files[path]
		if entry == nil {
			entry = &fileDiagEntry{}
		}
		entry.valid = false
		entry.fileHash = file.Hash
		entry.contributesGlobals = fileContributesToGlobalScope(file)
		files[path] = entry
	}
	state.files = files
	state.revDeps = nil
	state.revDepsGen = 0
}

// Status returns a copy of the cached state of one file. The second result is
// false when the project or the file is not cached, which means the file is not
// type-checked in that project (or Reconcile has never run for it).
func (c *WorkspaceDiagCache) Status(projectPath tspath.Path, filePath tspath.Path) (FileDiagStatus, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	state, ok := c.projects[projectPath]
	if !ok {
		return FileDiagStatus{}, false
	}
	entry, ok := state.files[filePath]
	if !ok {
		return FileDiagStatus{}, false
	}
	return FileDiagStatus{
		Valid:      entry.valid,
		ResultID:   entry.resultID,
		Items:      entry.items,
		Reported:   entry.reported,
		Generation: state.generation,
		FileHash:   entry.fileHash,
	}, true
}

// StoreComputed records freshly computed diagnostics for a file. `gen` and
// `hash` are the generation and file hash the diagnostics were computed
// against, as read from Status before checking: if either no longer matches the
// cache, the program moved under the computation and the result is discarded
// (ok is false).
//
// `changed` reports whether the items differ from what the client was last
// given. When they do not, the previous result ID is kept, so the handler can
// answer `unchanged` instead of resending an identical report.
func (c *WorkspaceDiagCache) StoreComputed(
	projectPath tspath.Path,
	filePath tspath.Path,
	gen uint64,
	hash xxh3.Uint128,
	items []*lsproto.Diagnostic,
) (resultID string, changed bool, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	state, ok := c.projects[projectPath]
	if !ok || state.generation != gen {
		return "", false, false
	}
	entry, ok := state.files[filePath]
	if !ok || entry.fileHash != hash {
		return "", false, false
	}
	if entry.resultID != "" && diagnosticsEqual(entry.items, items) {
		entry.valid = true
		return entry.resultID, false, true
	}
	entry.resultID = c.nonce + "-" + strconv.FormatUint(c.nextID.Add(1), 36)
	entry.items = items
	entry.valid = true
	return entry.resultID, true, true
}

// MarkReported records that a full report carrying `resultID` was sent to the
// client. It is a no-op if the entry has moved on to a newer result.
func (c *WorkspaceDiagCache) MarkReported(projectPath tspath.Path, filePath tspath.Path, resultID string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if state, ok := c.projects[projectPath]; ok {
		if entry, ok := state.files[filePath]; ok && entry.resultID == resultID {
			entry.reported = true
		}
	}
}

// Prune drops the cached state of projects that no longer exist in the given
// snapshot.
func (c *WorkspaceDiagCache) Prune(snapshot *Snapshot) {
	if snapshot == nil {
		return
	}
	live := collections.Set[tspath.Path]{}
	for _, project := range snapshot.ProjectCollection.Projects() {
		live.Add(project.configFilePath)
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	for path := range c.projects {
		if !live.Has(path) {
			delete(c.projects, path)
		}
	}
}

// fileContributesToGlobalScope reports whether a file adds anything to the
// global symbol table, in which case editing it can change the meaning of every
// other file in the program and diagnostics cannot be invalidated selectively.
//
// This mirrors two checker passes and must stay in sync with both:
//
//   - Checker.initializeChecker (internal/checker/checker.go), which hoists a
//     module's bare top-level `type`/`interface` declarations into the global
//     type table. There is no type-export syntax, so an ordinary module file
//     contributes globals -- which is also why this must NOT be replaced by the
//     incremental builder's fileAffectsGlobalScope, a heuristic that does not
//     know the rule.
//   - Checker.initializeLuaAugmentations (internal/checker/lua_augmentation.go),
//     which turns assignment forms into global *values*: a bare store to a name
//     no `local` declares creates an implicit global, and a write through `_G`
//     or through anything rooted at a global augments that global.
//
// Only reads are performed here: file.Locals, file.GlobalExports,
// file.ModuleAugmentations, file.LuaWriteCandidates and symbol
// flags/declarations are all frozen once binding has finished, and every file
// in a program is bound.
func fileContributesToGlobalScope(file *ast.SourceFile) bool {
	// UMD global exports (guarded like initializeChecker guards them).
	if file.Symbol != nil && len(file.GlobalExports) > 0 {
		return true
	}
	// ModuleAugmentations holds the name node of each augmentation; the
	// declaration itself is its parent. Only `declare global` can be parsed
	// (string-named ambient modules are gone), but check anyway.
	for _, moduleName := range file.ModuleAugmentations {
		if ast.IsGlobalScopeAugmentation(moduleName.Parent) {
			return true
		}
	}
	if !ast.IsExternalOrCommonJSModule(file) {
		// A script has no module scope: all of its top-level locals are global.
		if len(file.Locals) > 0 {
			return true
		}
	} else {
		// A module hoists its purely type-side top-level declarations, unless
		// they are marked `local` (mirrors checker.symbolHasLuaLocalDeclaration).
		for _, symbol := range file.Locals {
			if symbol.Flags&ast.SymbolFlagsValue == 0 && symbol.Flags&ast.SymbolFlagsType != 0 &&
				!core.Some(symbol.Declarations, ast.IsLuaLocal) {
				return true
			}
		}
	}
	// Lua assignment forms create global values without any declaration the
	// checks above can see. This holds for modules and scripts alike: a module
	// scope binds `local`s, not stores to undeclared names.
	for _, candidate := range file.LuaWriteCandidates {
		if luaCandidateWritesGlobal(candidate, 0 /*depth*/) {
			return true
		}
	}
	return false
}

// luaAliasChaseDepth bounds how far luaCandidateWritesGlobal follows a chain of
// local aliases. Chains this long do not occur in practice; the limit is only
// there so a cycle cannot spin, and it errs towards reporting a global.
const luaAliasChaseDepth = 4

// luaCandidateWritesGlobal reports whether one recorded Lua write creates or
// augments a global value, which is what makes its file global-contributing.
// It mirrors the resolution initializeLuaAugmentations performs, without a
// checker: only the write's root identifier and the lexical bindings visible at
// the write matter.
func luaCandidateWritesGlobal(candidate ast.LuaWriteCandidate, depth int) bool {
	rootName, reference, isBare, ok := luaWriteCandidateRoot(candidate)
	if !ok {
		return false
	}
	return luaWriteRootIsGlobal(rootName, reference, isBare, depth)
}

func luaWriteRootIsGlobal(rootName string, reference *ast.Node, isBare bool, depth int) bool {
	binding := luaRootBinding(rootName, reference)
	if binding == nil {
		// Nothing declares the root: a bare store creates an implicit global,
		// and a path write augments one. This also covers `_G.x = ...` — the
		// environment table has no lexical binding, so it resolves to nil here
		// exactly as the checker's resolveName resolves it to the environment
		// symbol. A file-local `local _G = {}` shadows the environment for the
		// checker and for this walk alike, so its writes stay local.
		return true
	}
	if isBare {
		// A bare store to a resolved local declares no member and contributes
		// only to that local's own storage.
		return false
	}
	if depth >= luaAliasChaseDepth {
		return true
	}
	// `local S = Shared; S.count = 1` augments a global through an alias, so a
	// path write rooted at a local has to follow that local's initializer.
	// Anything else -- a table constructor, a function, a literal, a call
	// (`require` included) -- roots the write in a value this file owns.
	//
	// A require alias is safe to treat as owned: writing a new member through
	// it (`local m = require("x"); m.f = ...`) does not augment module x's
	// type in tlua -- the checker reports TS2339 on the write -- so there is
	// no cross-file effect to invalidate for.
	declaration := binding.ValueDeclaration
	if declaration == nil || !ast.IsVariableDeclaration(declaration) || !ast.IsLuaLocal(declaration) {
		// Parameters and other non-variable declarations are not stable
		// augmentation roots.
		return false
	}
	initializer := declaration.Initializer()
	if initializer == nil {
		return false
	}
	root := luaStaticAccessRoot(initializer)
	if root == nil {
		return false
	}
	return luaWriteRootIsGlobal(root.Text(), root, false /*isBare*/, depth+1)
}

// luaWriteCandidateRoot returns the root identifier of a write's target, the
// node whose position decides which `local`s are visible, and whether the write
// is bare (a plain identifier target, declaring no member). It mirrors
// checker.luaAugmentationPath, which derives the same root from Target, falling
// back to the declaration for `function t.m()` writes where Target is nil.
//
// A dynamically keyed target (`t[k] = 1`) keeps its root: the checker forms no
// augmentation path for it at all, so reporting the root as global is only ever
// coarser than the checker, never wrong.
func luaWriteCandidateRoot(candidate ast.LuaWriteCandidate) (rootName string, reference *ast.Node, isBare bool, ok bool) {
	node := candidate.Source
	expression := candidate.Target
	isMember := false
	if ast.IsFunctionDeclaration(node) {
		// `function t.m()` writes a member of its receiver; the member name
		// comes from the declaration rather than from an lvalue.
		expression = node.AsFunctionDeclaration().Target
		if expression == nil || ast.GetNameOfDeclaration(node) == nil {
			return "", nil, false, false
		}
		isMember = true
	} else if !ast.IsBinaryExpression(node) || expression == nil {
		return "", nil, false, false
	}
	root := luaStaticAccessRoot(expression)
	if root == nil {
		return "", nil, false, false
	}
	// The checker resolves the root from the target when there is one, and from
	// the declaration otherwise, so lexical visibility is judged at that node.
	reference = candidate.Target
	if reference == nil {
		reference = candidate.Source
	}
	return root.Text(), reference, !isMember && ast.IsIdentifier(expression), true
}

// luaStaticAccessRoot walks an access chain leftward to its root identifier,
// like checker.luaEntityNamePath. It returns nil when the root is a dynamic
// expression (a call, say), which forms no static path.
func luaStaticAccessRoot(expression *ast.Node) *ast.Node {
	// Parentheses are as transparent here as they are to the checker's alias
	// resolution: `local S = (Shared)` still aliases the global.
	expression = ast.SkipOuterExpressions(expression, ast.OEKParentheses|ast.OEKAssertions)
	for ast.IsAccessExpression(expression) {
		expression = ast.SkipOuterExpressions(expression.Expression(), ast.OEKParentheses|ast.OEKAssertions)
	}
	if !ast.IsIdentifier(expression) {
		return nil
	}
	return expression
}

// luaRootBinding resolves a write's root name to the value binding visible at
// `reference`, mirroring the scope walk of checker.resolveName without a
// checker: each enclosing locals container is consulted for an ordinary value
// symbol and, position-sensitively, for a Lua `local`. A nil result means the
// name is global.
func luaRootBinding(rootName string, reference *ast.Node) *ast.Symbol {
	for node := reference; node != nil; node = node.Parent {
		if !ast.IsLocalsContainer(node) {
			continue
		}
		// Reads only: Node.Locals and Node.LuaLocals return the binder's tables
		// without allocating, and indexing a nil map is legal.
		if symbol := node.Locals()[rootName]; symbol != nil && symbol.Flags&ast.SymbolFlagsValue != 0 {
			return symbol
		}
		if symbol := ast.LookupLuaLocal(node, rootName, reference, ast.SymbolFlagsValue); symbol != nil {
			return symbol
		}
	}
	return nil
}

// buildReverseImportGraph maps each file's path to the paths of the checked
// files that reference it, through `require`/import specifiers, triple-slash
// references and type reference directives. Only checked files appear as
// importers: a skipped file's diagnostics are never reported, so nothing needs
// to be invalidated on its behalf.
func buildReverseImportGraph(program *compiler.Program, files map[tspath.Path]*ast.SourceFile) map[tspath.Path][]tspath.Path {
	revDeps := make(map[tspath.Path][]tspath.Path, len(files))
	typeRefs := program.GetResolvedTypeReferenceDirectives()
	targets := collections.Set[tspath.Path]{}
	for path, file := range files {
		targets.Clear()
		containingDirectory := tspath.GetDirectoryPath(file.FileName())
		for _, moduleSpecifier := range file.Imports() {
			if resolved := program.GetResolvedModuleFromModuleSpecifier(file, moduleSpecifier); resolved.IsResolved() {
				targets.Add(referencedFilePath(program, resolved.ResolvedFileName, containingDirectory))
			}
		}
		for _, referencedFile := range file.ReferencedFiles {
			targets.Add(referencedFilePath(program, referencedFile.FileName, containingDirectory))
		}
		for _, typeRef := range typeRefs[file.Path()] {
			if typeRef.ResolvedFileName != "" {
				targets.Add(referencedFilePath(program, typeRef.ResolvedFileName, containingDirectory))
			}
		}
		for target := range targets.Keys() {
			if target != path {
				revDeps[target] = append(revDeps[target], path)
			}
		}
	}
	return revDeps
}

// referencedFilePath converts a referenced file name to the path the program
// knows it by, following project-reference redirects the way the incremental
// builder's referenced-files collection does.
func referencedFilePath(program *compiler.Program, fileName string, containingDirectory string) tspath.Path {
	if redirect := program.GetParseFileRedirect(fileName); redirect != "" {
		return tspath.ToPath(redirect, program.GetCurrentDirectory(), program.UseCaseSensitiveFileNames())
	}
	return tspath.ToPath(fileName, containingDirectory, program.UseCaseSensitiveFileNames())
}

// diagnosticsEqual reports whether two diagnostic lists would serialize
// identically. Comparing the wire form sidesteps the union types inside
// lsproto.Diagnostic; it only ever runs on files whose diagnostics were just
// recomputed, so the cost is noise next to the check itself.
func diagnosticsEqual(a []*lsproto.Diagnostic, b []*lsproto.Diagnostic) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	encodedA, err := json.Marshal(a)
	if err != nil {
		return false
	}
	encodedB, err := json.Marshal(b)
	if err != nil {
		return false
	}
	return bytes.Equal(encodedA, encodedB)
}
