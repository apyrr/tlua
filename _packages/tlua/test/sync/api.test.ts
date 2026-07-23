//
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !!! THIS FILE IS AUTO-GENERATED - DO NOT EDIT !!!
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
//
// Source: test/async/api.test.ts
// Regenerate: npm run generate (from _packages/tlua)
//
import {
    type __String,
    cast,
    escapeLeadingUnderscores,
    type Expression,
    getJSDocTags,
    getSynthesizedDeepClone,
    InternalSymbolName,
    isCallExpression,
    isFunctionDeclaration,
    isIdentifier,
    isJSDocParameterTag,
    isReturnStatement,
    isStringLiteral,
    isTemplateHead,
    isTemplateMiddle,
    isTemplateTail,
    isTypeAliasDeclaration,
    isTypeNode,
    isVariableDeclarationList,
    isVariableStatement,
    type Node,
    type NodeArray,
    NodeFlags,
    SyntaxKind,
    unescapeLeadingUnderscores,
} from "@tlua/cli/unstable/ast";
import {
    createArrayTypeNode,
    createFunctionTypeNode,
    createIdentifier,
    createKeywordTypeNode,
    createParameterDeclaration,
    createToken,
    createTypeAliasDeclaration,
    createTypeReferenceNode,
    createUnionTypeNode,
} from "@tlua/cli/unstable/ast/factory";
import { visitEachChild } from "@tlua/cli/unstable/ast/visitor";
import { createVirtualFileSystem } from "@tlua/cli/unstable/fs";
import type { FileSystem } from "@tlua/cli/unstable/fs";
import {
    API,
    type ConditionalType,
    DiagnosticCategory,
    type FreshableType,
    type IndexedAccessType,
    type IndexType,
    type InterfaceType,
    type IntrinsicType,
    isErrorType,
    type LiteralType,
    ModifierFlags,
    ModuleKind,
    ObjectFlags,
    SignatureKind,
    type StringMappingType,
    SymbolFlags,
    type TemplateLiteralType,
    TypeFlags,
    type TypeParameter,
    TypePredicateKind,
    type TypeReference,
    type UnionOrIntersectionType,
} from "@tlua/cli/unstable/sync";
import assert from "node:assert";
import { globSync } from "node:fs";
import { resolve } from "node:path";
import {
    describe,
    test,
} from "node:test";
import { fileURLToPath } from "node:url";
import { runBenchmarks } from "./api.bench.ts";

const defaultFiles = {
    "/tluaconfig.json": "{}",
    "/src/index.tlua": `local foo = require('src.foo').foo;`,
    "/src/foo.tlua": `local foo = 42;\nreturn { foo = foo };`,
};

describe("API", () => {
    test("parseConfigFile", () => {
        const api = spawnAPI();
        try {
            const config = api.parseConfigFile("/tluaconfig.json");
            assert.deepEqual(config.fileNames, ["/src/index.tlua", "/src/foo.tlua"]);
            assert.deepEqual(config.options, { configFilePath: "/tluaconfig.json" });
        }
        finally {
            api.close();
        }
    });
});

describe("Snapshot", () => {
    test("updateSnapshot returns snapshot with projects", () => {
        const api = spawnAPI();
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            assert.ok(snapshot);
            assert.ok(snapshot.id);
            assert.ok(snapshot.getProjects().length > 0);
            assert.ok(snapshot.getProject("/tluaconfig.json"));
        }
        finally {
            api.close();
        }
    });

    test("getSymbolAtPosition", () => {
        const api = spawnAPI();
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const symbol = project.checker.getSymbolAtPosition("/src/index.tlua", 7);
            assert.ok(symbol);
            assert.equal(symbol.name, "foo");
            assert.ok(symbol.flags & SymbolFlags.BlockScopedVariable);
        }
        finally {
            api.close();
        }
    });

    test("getSymbolAtLocation", () => {
        const api = spawnAPI();
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/index.tlua");
            assert.ok(sourceFile);
            const node = cast(
                cast(sourceFile.statements[0], isVariableStatement).declarationList.declarations[0].name,
                isIdentifier,
            );
            assert.ok(node);
            const symbol = project.checker.getSymbolAtLocation(node);
            assert.ok(symbol);
            assert.equal(symbol.name, "foo");
            assert.ok(symbol.flags & SymbolFlags.BlockScopedVariable);
        }
        finally {
            api.close();
        }
    });

    test("getTypeOfSymbol", () => {
        const api = spawnAPI();
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const symbol = project.checker.getSymbolAtPosition("/src/index.tlua", 7);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            // The module return table widens the const literal to number.
            assert.ok(type.flags & TypeFlags.Number);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getApparentType", () => {
    test("returns the apparent type of a literal type", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `local x = "hello" as const;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = `local x = "hello" as const;`.indexOf("x =");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            assert.equal(type.isLiteralType(), true);
            assert.equal(type.isStringLiteralType(), true);
            assert.equal(type.isIntrinsicType(), false);
            const apparent = project.checker.getApparentType(type);
            assert.ok(apparent);
            assert.ok(apparent.id > 0);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getMemberInModuleExports", () => {
    test("returns a named export when present", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/index.tlua": `local direct = 1;\nreturn { direct = direct };`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/index.tlua");
            assert.ok(sourceFile);
            const moduleSymbol = project.checker.getSymbolAtLocation(sourceFile);
            assert.ok(moduleSymbol);
            // getExportsOfModule resolves the export= target, and a return
            // table has no named exports of its own.
            const found = project.checker.getMemberInModuleExports(moduleSymbol, "direct");
            assert.equal(found, undefined);
            // The unresolved exports table still carries the export= symbol.
            const exports = moduleSymbol.getExports();
            const names = [...exports.values()].map(e => e.name);
            assert.ok(names.includes("export="), "module carries its return-table export");
        }
        finally {
            api.close();
        }
    });
});

describe("SourceFile", () => {
    test("getSourceFileNames returns all program files, not just root files", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({
                compilerOptions: {
                    noLib: true,
                },
            }),
            "/src/index.tlua": `local foo = require("src.foo");\nlocal mylib = require("my-lib");\nlocal result = foo.foo;`,
            "/src/foo.tlua": `local foo = 42;\nreturn { foo = foo };`,
            "/node_modules/my-lib/package.json": JSON.stringify({ name: "my-lib" }),
            "/node_modules/my-lib/index.tlua": `local bar: number = 1;\nreturn { bar = bar };`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const fileNames = project.program.getSourceFileNames();
            assert.deepEqual(fileNames, [
                "/src/foo.tlua",
                "/node_modules/my-lib/index.tlua",
                "/src/index.tlua",
            ]);
        }
        finally {
            api.close();
        }
    });

    test("source file metadata identifies external library and default library files", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({
                compilerOptions: {},
            }),
            "/src/index.tlua": `local mylib = require("my-lib");\nlocal result = 1;`,
            "/node_modules/my-lib/package.json": JSON.stringify({ name: "my-lib" }),
            "/node_modules/my-lib/index.tlua": `local bar: number = 1;\nreturn { bar = bar };`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const program = project.program;

            const index = program.getSourceFile("/src/index.tlua");
            assert.ok(index);
            assert.equal(program.isSourceFileFromExternalLibrary(index), false);
            assert.equal(program.isSourceFileDefaultLibrary(index), false);

            const lib = program.getSourceFile("/node_modules/my-lib/index.tlua");
            assert.ok(lib);
            assert.equal(program.isSourceFileFromExternalLibrary(lib), true);
            assert.equal(program.isSourceFileDefaultLibrary(lib), false);

            const fileNames = program.getSourceFileNames();
            const defaultLibName = fileNames.find(name => name.endsWith("lib.d.ts") || name.includes("/lib."));
            assert.ok(defaultLibName, "expected a default library file in the program");
            const defaultLib = program.getSourceFile(defaultLibName);
            assert.ok(defaultLib);
            assert.equal(program.isSourceFileDefaultLibrary(defaultLib), true);
        }
        finally {
            api.close();
        }
    });

    test("source file metadata reports implied node format", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({
                compilerOptions: {
                    module: "nodenext",
                },
            }),
            "/src/index.tlua": `local x = 1;`,
            "/esm/package.json": JSON.stringify({ type: "module" }),
            "/esm/index.tlua": `local m = 1;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const program = snapshot.getProject("/tluaconfig.json")!.program;

            // A plain .ts file with no nearby `"type": "module"` is CommonJS.
            const index = program.getSourceFile("/src/index.tlua");
            assert.ok(index);
            assert.equal((program.getSourceFileMetadata(index.fileName))?.impliedNodeFormat, ModuleKind.CommonJS);

            // A plain .ts file under a `"type": "module"` package is ESM.
            const esmIndex = program.getSourceFile("/esm/index.tlua");
            assert.ok(esmIndex);
            assert.equal((program.getSourceFileMetadata(esmIndex.fileName))?.impliedNodeFormat, ModuleKind.ESNext);
            assert.equal((program.getSourceFileMetadataByPath(esmIndex.path))?.impliedNodeFormat, ModuleKind.ESNext);
        }
        finally {
            api.close();
        }
    });

    test("file properties", () => {
        const api = spawnAPI();
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/index.tlua");

            assert.ok(sourceFile);
            assert.equal(sourceFile.text, defaultFiles["/src/index.tlua"]);
            assert.equal(sourceFile.fileName, "/src/index.tlua");
        }
        finally {
            api.close();
        }
    });

    test("extended data", () => {
        const api = spawnAPI();
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/index.tlua");

            assert.ok(sourceFile);
            let nodeCount = 1;
            sourceFile.forEachChild(function visit(node) {
                if (isTemplateHead(node)) {
                    assert.equal(node.text, "head ");
                    assert.equal(node.rawText, "head ");
                    assert.equal(node.templateFlags, 0);
                }
                else if (isTemplateMiddle(node)) {
                    assert.equal(node.text, "middle");
                    assert.equal(node.rawText, "middle");
                    assert.equal(node.templateFlags, 0);
                }
                else if (isTemplateTail(node)) {
                    assert.equal(node.text, " tail");
                    assert.equal(node.rawText, " tail");
                    assert.equal(node.templateFlags, 0);
                }
                nodeCount++;
                node.forEachChild(visit);
            });
            // sourceFile + VariableStatement, declarationList, declaration,
            // name, property access, require call, require, specifier, .foo,
            // EndOfFileToken
            assert.equal(nodeCount, 11);
        }
        finally {
            api.close();
        }
    });

    test("forEachChild with visitList does not visit array children twice", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ files: ["/input.tlua"] }),
            "/input.tlua": `local arrow = function() end`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/input.tlua");

            assert.ok(sourceFile);

            const visited: { kind: SyntaxKind; pos: number; end: number; }[] = [];
            (function walk(node: Node): void {
                visited.push({ kind: node.kind, pos: node.pos, end: node.end });
                node.forEachChild(walk, (nodes: NodeArray<Node>) => {
                    for (let i = 0; i < nodes.length; i++) {
                        walk(nodes[i]);
                    }
                    return undefined;
                });
            })(sourceFile);

            // Each node should be visited exactly once, even when a visitList callback
            // is supplied. Previously array children were visited twice.
            const seen = new Set<string>();
            for (const { kind, pos, end } of visited) {
                const key = `${kind}.${pos}.${end}`;
                assert.ok(!seen.has(key), `Node ${key} was visited more than once`);
                seen.add(key);
            }
        }
        finally {
            api.close();
        }
    });
});

test("unicode escapes", () => {
    const api = spawnAPI({
        "/tluaconfig.json": "{}",
        "/src/1.tlua": `"😃"`,
        "/src/2.tlua": `"\\ud83d\\ude03"`,
        "/src/3.tlua": `"\\ud800a\\udc00"`,
    });
    try {
        const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
        const project = snapshot.getProject("/tluaconfig.json")!;
        const expectedTexts = new Map([
            ["/src/1.tlua", "😃"],
            ["/src/2.tlua", "😃"],
            ["/src/3.tlua", "\ud800a\udc00"],
        ]);

        for (const file of expectedTexts.keys()) {
            const sourceFile = project.program.getSourceFile(file);
            assert.ok(sourceFile);

            sourceFile.forEachChild(function visit(node) {
                if (isStringLiteral(node)) {
                    assert.equal(node.text, expectedTexts.get(file));
                }
                node.forEachChild(visit);
            });
        }
    }
    finally {
        api.close();
    }
});

test("template unicode escapes", () => {
    const api = spawnAPI({
        "/tluaconfig.json": "{}",
        "/src/index.tlua": "`\\ud800${0}\\udc00`",
    });
    try {
        const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
        const project = snapshot.getProject("/tluaconfig.json")!;
        const sourceFile = project.program.getSourceFile("/src/index.tlua");
        assert.ok(sourceFile);

        let sawHead = false;
        let sawTail = false;
        sourceFile.forEachChild(function visit(node) {
            if (isTemplateHead(node)) {
                assert.equal(node.text, "\ud800");
                sawHead = true;
            }
            else if (isTemplateTail(node)) {
                assert.equal(node.text, "\udc00");
                sawTail = true;
            }
            node.forEachChild(visit);
        });
        assert.ok(sawHead);
        assert.ok(sawTail);
    }
    finally {
        api.close();
    }
});

test("Object equality", () => {
    const api = spawnAPI();
    try {
        const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
        const project = snapshot.getProject("/tluaconfig.json")!;
        // Same symbol returned from same snapshot's checker
        assert.strictEqual(
            project.checker.getSymbolAtPosition("/src/index.tlua", 6),
            project.checker.getSymbolAtPosition("/src/index.tlua", 8),
        );
    }
    finally {
        api.close();
    }
});

test("Snapshot dispose", () => {
    const api = spawnAPI();
    try {
        const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
        const project = snapshot.getProject("/tluaconfig.json")!;
        const symbol = project.checker.getSymbolAtPosition("/src/index.tlua", 7);
        assert.ok(symbol);

        // Snapshot dispose should release server-side resources
        assert.ok(snapshot.isDisposed() === false);
        snapshot.dispose();
        assert.ok(snapshot.isDisposed() === true);

        // After dispose, snapshot methods should throw
        assert.throws(() => {
            snapshot.getProject("/tluaconfig.json");
        }, {
            name: "Error",
            message: "Snapshot is disposed",
        });
    }
    finally {
        api.close();
    }
});

describe("Multiple snapshots", () => {
    test("two snapshots work independently", () => {
        const api = spawnAPI();
        try {
            const snap1 = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const snap2 = api.updateSnapshot({ openProject: "/tluaconfig.json" });

            // Both can fetch source files
            const sf1 = snap1.getProject("/tluaconfig.json")!.program.getSourceFile("/src/index.tlua");
            const sf2 = snap2.getProject("/tluaconfig.json")!.program.getSourceFile("/src/index.tlua");
            assert.ok(sf1);
            assert.ok(sf2);

            // Disposing one doesn't break the other
            snap1.dispose();
            assert.ok(snap1.isDisposed());
            assert.ok(!snap2.isDisposed());

            // snap2 still works after snap1 is disposed
            const symbol = snap2.getProject("/tluaconfig.json")!.checker.getSymbolAtPosition("/src/index.tlua", 7);
            assert.ok(symbol);
            assert.equal(symbol.name, "foo");
        }
        finally {
            api.close();
        }
    });

    test("each snapshot has its own server-side lifecycle", () => {
        const { api, fs } = spawnAPIWithFS();
        try {
            const snap1 = api.updateSnapshot({ openProject: "/tluaconfig.json" });

            // Verify initial state
            const sf1 = snap1.getProject("/tluaconfig.json")!.program.getSourceFile("/src/foo.tlua");
            assert.ok(sf1);
            assert.equal(sf1.text, `local foo = 42;\nreturn { foo = foo };`);

            // Mutate the file and create a new snapshot with the change
            fs.writeFile!("/src/foo.tlua", `local foo = "changed";\nreturn { foo = foo };`);
            const snap2 = api.updateSnapshot({
                fileChanges: { changed: ["/src/foo.tlua"] },
            });

            // snap2 should reflect the updated content
            const sf2 = snap2.getProject("/tluaconfig.json")!.program.getSourceFile("/src/foo.tlua");
            assert.ok(sf2);
            assert.equal(sf2.text, `local foo = "changed";\nreturn { foo = foo };`);

            // snap1's source file should still have the original content
            assert.equal(sf1.text, `local foo = 42;\nreturn { foo = foo };`);

            snap1.dispose();

            // snap2 still works independently after snap1 is disposed
            const symbol = snap2.getProject("/tluaconfig.json")!.checker.getSymbolAtPosition("/src/index.tlua", 7);
            assert.ok(symbol);

            snap2.dispose();

            // Both are disposed, new snapshot works fine with latest content
            const snap3 = api.updateSnapshot();
            const sf3 = snap3.getProject("/tluaconfig.json")!.program.getSourceFile("/src/foo.tlua");
            assert.ok(sf3);
            assert.equal(sf3.text, `local foo = "changed";\nreturn { foo = foo };`);
        }
        finally {
            api.close();
        }
    });

    test("adding a new file is reflected in the next snapshot", () => {
        const { api, fs } = spawnAPIWithFS();
        try {
            const snap1 = api.updateSnapshot({ openProject: "/tluaconfig.json" });

            // Add a brand new file
            fs.writeFile!("/src/bar.tlua", `local bar = true;\nreturn { bar = bar };`);
            const snap2 = api.updateSnapshot({
                fileChanges: { created: ["/src/bar.tlua"] },
            });

            const sf = snap2.getProject("/tluaconfig.json")!.program.getSourceFile("/src/bar.tlua");
            assert.ok(sf);
            assert.equal(sf.text, `local bar = true;\nreturn { bar = bar };`);

            // Original snapshot shouldn't have the new file
            const sfOld = snap1.getProject("/tluaconfig.json")!.program.getSourceFile("/src/bar.tlua");
            assert.equal(sfOld, undefined);
        }
        finally {
            api.close();
        }
    });

    test("multiple sequential edits produce correct snapshots", () => {
        const { api, fs } = spawnAPIWithFS();
        try {
            api.updateSnapshot({ openProject: "/tluaconfig.json" });

            const versions = [
                `local foo = 1;\nreturn { foo = foo };`,
                `local foo = 2;\nreturn { foo = foo };`,
                `local foo = 3;\nreturn { foo = foo };`,
            ];

            for (const version of versions) {
                fs.writeFile!("/src/foo.tlua", version);
                const snap = api.updateSnapshot({
                    fileChanges: { changed: ["/src/foo.tlua"] },
                });
                const sf = snap.getProject("/tluaconfig.json")!.program.getSourceFile("/src/foo.tlua");
                assert.ok(sf);
                assert.equal(sf.text, version);
            }
        }
        finally {
            api.close();
        }
    });
});

describe("Source file caching", () => {
    test("same file from same snapshot returns cached object", () => {
        const api = spawnAPI();
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sf1 = project.program.getSourceFile("/src/index.tlua");
            const sf2 = project.program.getSourceFile("/src/index.tlua");
            assert.ok(sf1);
            assert.strictEqual(sf1, sf2, "Same source file should be returned from cache");
        }
        finally {
            api.close();
        }
    });

    test("same file from two snapshots (same content) returns cached object", () => {
        const api = spawnAPI();
        try {
            const snap1 = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const snap2 = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            // Fetch from snap1 first (populates cache), then snap2 (cache hit via hash)
            const sf1 = snap1.getProject("/tluaconfig.json")!.program.getSourceFile("/src/index.tlua");
            const sf2 = snap2.getProject("/tluaconfig.json")!.program.getSourceFile("/src/index.tlua");
            assert.ok(sf1);
            assert.ok(sf2);
            // Same content hash → cache hit → same object
            assert.strictEqual(sf1, sf2, "Same file with same content should share cached object");
        }
        finally {
            api.close();
        }
    });

    test("modified file returns a different source file object", () => {
        const { api, fs } = spawnAPIWithFS();
        try {
            const snap1 = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const sf1 = snap1.getProject("/tluaconfig.json")!.program.getSourceFile("/src/foo.tlua");
            assert.ok(sf1);
            assert.equal(sf1.text, `local foo = 42;\nreturn { foo = foo };`);

            // Mutate the file in the VFS
            fs.writeFile!("/src/foo.tlua", `local foo = 100;\nreturn { foo = foo };`);

            // Notify the server about the change
            const snap2 = api.updateSnapshot({
                fileChanges: { changed: ["/src/foo.tlua"] },
            });
            const sf2 = snap2.getProject("/tluaconfig.json")!.program.getSourceFile("/src/foo.tlua");
            assert.ok(sf2);
            assert.equal(sf2.text, `local foo = 100;\nreturn { foo = foo };`);

            // Different content → different object
            assert.notStrictEqual(sf1, sf2, "Modified file should return a new source file object");
        }
        finally {
            api.close();
        }
    });

    test("unmodified file retains cached object across file change notification", () => {
        const { api, fs } = spawnAPIWithFS();
        try {
            const snap1 = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const sf1 = snap1.getProject("/tluaconfig.json")!.program.getSourceFile("/src/index.tlua");
            assert.ok(sf1);

            // Mutate a different file
            fs.writeFile!("/src/foo.tlua", `local foo = 999;\nreturn { foo = foo };`);

            // Notify the server about the change to foo.ts only
            const snap2 = api.updateSnapshot({
                fileChanges: { changed: ["/src/foo.tlua"] },
            });
            const sf2 = snap2.getProject("/tluaconfig.json")!.program.getSourceFile("/src/index.tlua");
            assert.ok(sf2);

            // index.ts wasn't changed — should still get cached object
            assert.strictEqual(sf1, sf2, "Unchanged file should return cached object across snapshots");
        }
        finally {
            api.close();
        }
    });

    test("cache entries survive when one of two snapshots is disposed", () => {
        const api = spawnAPI();
        try {
            const snap1 = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            // Fetch from snap1 to populate cache
            const sf1 = snap1.getProject("/tluaconfig.json")!.program.getSourceFile("/src/index.tlua");
            assert.ok(sf1);

            // snap2 retains snap1's cache refs for unchanged files via snapshot changes
            const snap2 = api.updateSnapshot({ openProject: "/tluaconfig.json" });

            // Dispose snap1 — snap2 still holds a ref, so the entry survives
            snap1.dispose();

            // Fetching from snap2 should still return the cached object
            const sf2 = snap2.getProject("/tluaconfig.json")!.program.getSourceFile("/src/index.tlua");
            assert.ok(sf2);
            assert.strictEqual(sf1, sf2, "Cache entry should survive when retained by the next snapshot");
        }
        finally {
            api.close();
        }
    });

    test("invalidateAll causes all files to be re-fetched", () => {
        const { api, fs } = spawnAPIWithFS();
        try {
            const snap1 = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const sf1 = snap1.getProject("/tluaconfig.json")!.program.getSourceFile("/src/foo.tlua");
            assert.ok(sf1);
            assert.equal(sf1.text, `local foo = 42;\nreturn { foo = foo };`);

            // Mutate the file
            fs.writeFile!("/src/foo.tlua", `local foo = "hello";\nreturn { foo = foo };`);

            // Use invalidateAll to force re-fetch
            const snap2 = api.updateSnapshot({
                fileChanges: { invalidateAll: true },
            });
            const sf2 = snap2.getProject("/tluaconfig.json")!.program.getSourceFile("/src/foo.tlua");
            assert.ok(sf2);
            assert.equal(sf2.text, `local foo = "hello";\nreturn { foo = foo };`);
            assert.notStrictEqual(sf1, sf2, "invalidateAll should produce new source file objects");
        }
        finally {
            api.close();
        }
    });

    test("node handles from a cached source file should be valid in a new snapshot", () => {
        const { api, fs } = spawnAPIWithFS({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `function foo(x: number) end\nfoo(42);`,
            "/src/other.tlua": `local x = 1;\nreturn { x = x };`,
        });
        try {
            // Snapshot 1: get a node and verify getContextualType works
            const snap1 = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const proj1 = snap1.getProject("/tluaconfig.json")!;

            const sf1 = proj1.program.getSourceFile("/src/main.tlua");
            assert.ok(sf1);

            let numLiteral: Expression | undefined;
            sf1.forEachChild(function visit(node) {
                if (isCallExpression(node)) numLiteral = node.arguments[0];
                node.forEachChild(visit);
            });
            assert.ok(numLiteral, "should find the 42 argument");

            const type1 = proj1.checker.getContextualType(numLiteral);
            assert.ok(type1);
            assert.ok(type1.flags & TypeFlags.Number);

            // Snapshot 2: change a different file
            fs.writeFile!("/src/other.tlua", `local x = 2;\nreturn { x = x };`);
            const snap2 = api.updateSnapshot({
                fileChanges: { changed: ["/src/other.tlua"] },
            });
            const proj2 = snap2.getProject("/tluaconfig.json")!;

            // main.ts is unchanged — client returns the cached SourceFile (same object)
            const sf2 = proj2.program.getSourceFile("/src/main.tlua");
            assert.ok(sf2);
            assert.strictEqual(sf1, sf2, "unchanged file should be served from client cache");

            let numLiteral2: Expression | undefined;
            sf2.forEachChild(function visit(node) {
                if (isCallExpression(node)) numLiteral2 = node.arguments[0];
                node.forEachChild(visit);
            });
            assert.ok(numLiteral2, "should find the 42 argument");
            assert.strictEqual(numLiteral, numLiteral2, "unchanged file should be served from client cache");

            // A type from new snapshot should be resolved
            const type2 = proj2.checker.getContextualType(numLiteral);
            assert.ok(type2);
            assert.ok(type2.flags & TypeFlags.Number);
        }
        finally {
            api.close();
        }
    });
});

describe("Snapshot disposal", () => {
    test("dispose is idempotent", () => {
        const api = spawnAPI();
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            snapshot.dispose();
            assert.ok(snapshot.isDisposed());
            // Second dispose should not throw
            snapshot.dispose();
            assert.ok(snapshot.isDisposed());
        }
        finally {
            api.close();
        }
    });

    test("api.close disposes all active snapshots", () => {
        const api = spawnAPI();
        const snap1 = api.updateSnapshot({ openProject: "/tluaconfig.json" });
        const snap2 = api.updateSnapshot({ openProject: "/tluaconfig.json" });
        assert.ok(!snap1.isDisposed());
        assert.ok(!snap2.isDisposed());
        api.close();
        assert.ok(snap1.isDisposed());
        assert.ok(snap2.isDisposed());
    });
});

describe("Source file cache keying across projects", () => {
    // Three projects share the same file (/src/shared.ts).
    // The file sits inside a package.json scope with "type": "module".
    //
    // Parsing no longer depends on compiler options (every non-declaration
    // file is a module), so all three projects must share one cached
    // source file object regardless of their module settings.
    const multiProjectFiles: Record<string, string> = {
        "/package.json": JSON.stringify({ type: "module" }),
        "/src/shared.tlua": `local x = 1;\nreturn { x = x };`,
        // Project A – bundler, auto detection (default)
        "/projectA/tluaconfig.json": JSON.stringify({
            compilerOptions: {
                module: "esnext",
                strict: true,
            },
            files: ["../src/shared.tlua"],
        }),
        // Project B – bundler
        "/projectB/tluaconfig.json": JSON.stringify({
            compilerOptions: {
                module: "esnext",
                strict: true,
            },
            files: ["../src/shared.tlua"],
        }),
        // Project C – nodenext (type:module → module)
        "/projectC/tluaconfig.json": JSON.stringify({
            compilerOptions: {
                module: "nodenext",
                strict: true,
            },
            files: ["../src/shared.tlua"],
        }),
    };

    test("projects with different options share the same cached source file", () => {
        const api = spawnAPI(multiProjectFiles);
        try {
            // Open all three projects
            api.updateSnapshot({ openProject: "/projectA/tluaconfig.json" });
            api.updateSnapshot({ openProject: "/projectB/tluaconfig.json" });
            const snapshot = api.updateSnapshot({ openProject: "/projectC/tluaconfig.json" });

            const projectA = snapshot.getProject("/projectA/tluaconfig.json")!;
            const projectB = snapshot.getProject("/projectB/tluaconfig.json")!;
            const projectC = snapshot.getProject("/projectC/tluaconfig.json")!;
            assert.ok(projectA, "projectA should exist");
            assert.ok(projectB, "projectB should exist");
            assert.ok(projectC, "projectC should exist");

            // Fetch the shared file from each project
            const sfA = projectA.program.getSourceFile("/src/shared.tlua");
            const sfB = projectB.program.getSourceFile("/src/shared.tlua");
            const sfC = projectC.program.getSourceFile("/src/shared.tlua");
            assert.ok(sfA, "sfA should exist");
            assert.ok(sfB, "sfB should exist");
            assert.ok(sfC, "sfC should exist");

            // Parse options no longer vary by project, so all three share one cached object.
            assert.strictEqual(sfA, sfB, "projectA and projectB should share the same cached source file");
            assert.strictEqual(sfB, sfC, "projectB and projectC should share the same cached source file");
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - symbol identity across projects", () => {
    const sharedSymbolFiles = {
        "/projectA/tluaconfig.json": JSON.stringify({ files: ["../src/shared.tlua"] }),
        "/projectB/tluaconfig.json": JSON.stringify({ files: ["../src/shared.tlua"] }),
        "/src/shared.tlua": `local sharedVar = 42;\nreturn { sharedVar = sharedVar };`,
    };

    test("getSymbolAtPosition returns same Symbol instance across projects", () => {
        const api = spawnAPI(sharedSymbolFiles);
        try {
            api.updateSnapshot({ openProject: "/projectA/tluaconfig.json" });
            const snapshot = api.updateSnapshot({ openProject: "/projectB/tluaconfig.json" });

            const projectA = snapshot.getProject("/projectA/tluaconfig.json")!;
            const projectB = snapshot.getProject("/projectB/tluaconfig.json")!;
            assert.ok(projectA, "projectA should exist");
            assert.ok(projectB, "projectB should exist");

            const src = sharedSymbolFiles["/src/shared.tlua"];
            const varPos = src.indexOf("sharedVar");

            const symbolA = projectA.checker.getSymbolAtPosition("/src/shared.tlua", varPos);
            const symbolB = projectB.checker.getSymbolAtPosition("/src/shared.tlua", varPos);

            assert.ok(symbolA, "symbolA should exist");
            assert.ok(symbolB, "symbolB should exist");
            assert.equal(symbolA.name, "sharedVar");
            assert.equal(symbolB.name, "sharedVar");

            assert.strictEqual(symbolA, symbolB, "Same source symbol queried from two projects should be the same object");
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - types and signatures", () => {
    const checkerFiles = {
        "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
        "/src/main.tlua": `
local x = 42 as const;
function add(a: number, b: number, ...rest: number[]): number return a + b end
interface Widget {
    value: string;
}
declare MyClass: {
    new (): Widget;
};
local holder = {
    getValue = function(): string return "" end,
};
`,
    };

    test("getTypeAtPosition", () => {
        const api = spawnAPI(checkerFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = checkerFiles["/src/main.tlua"];
            const xPos = src.indexOf("x = 42");
            const type = project.checker.getTypeAtPosition("/src/main.tlua", xPos);
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.NumberLiteral);
        }
        finally {
            api.close();
        }
    });

    test("getTypeAtPosition batched", () => {
        const api = spawnAPI(checkerFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = checkerFiles["/src/main.tlua"];
            const xPos = src.indexOf("x = 42");
            const addPos = src.indexOf("add(");
            const types = project.checker.getTypeAtPosition("/src/main.tlua", [xPos, addPos]);
            assert.equal(types.length, 2);
            assert.ok(types[0]);
            assert.ok(types[1]);
        }
        finally {
            api.close();
        }
    });

    test("getTypeAtLocation", () => {
        const api = spawnAPI(checkerFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);
            const firstVarDecl = sourceFile.statements[2]; // "interface Widget"
            assert.ok(firstVarDecl);
            const type = project.checker.getTypeAtLocation(firstVarDecl);
            assert.ok(type);
        }
        finally {
            api.close();
        }
    });

    test("getTypeAtLocation returns property type for parenthesized and chained access (issue #3938)", () => {
        const files = {
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `
interface A {
    a: number;
    b: { c: number };
}
local obj: A = { a = 1, b = { c = 2 } };
local a1 = obj.a;
local a2 = (obj).a;
local c = obj.b.c;
`,
        };

        const api = spawnAPI(files);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);

            const targetNodes = new Map<string, Node>();
            sourceFile.forEachChild(function visit(node) {
                if (node.kind === SyntaxKind.PropertyAccessExpression) {
                    const text = sourceFile.text.slice(node.pos, node.end).trim();
                    if (text === "obj.a" || text === "(obj).a" || text === "obj.b.c") {
                        targetNodes.set(text, node);
                    }
                }
                node.forEachChild(visit);
            });

            assert.equal(targetNodes.size, 3, "Should find all target property access expressions");
            for (const expr of ["obj.a", "(obj).a", "obj.b.c"] as const) {
                const node = targetNodes.get(expr);
                assert.ok(node, `Should find expression '${expr}'`);
                const type = project.checker.getTypeAtLocation(node);
                assert.ok(type, `Should get a type for '${expr}'`);
                assert.ok(type.flags & TypeFlags.Number, `Expected '${expr}' to have number type flags, got ${type.flags}`);
            }
        }
        finally {
            api.close();
        }
    });

    test("getTypeAtLocation returns call result type for a call expression (issue #4041)", () => {
        const files = {
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true, target: "esnext" } }),
            "/src/main.tlua": `
type Result = { readonly value: string };

function buildCapabilities(): Result
    return { value = "ok" };
end

function run(): Result
    return buildCapabilities();
end
`,
        };

        const api = spawnAPI(files);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);

            let callNode: import("@tlua/cli/unstable/ast").CallExpression | undefined;
            sourceFile.forEachChild(function visit(node) {
                if (isCallExpression(node)) {
                    const text = sourceFile.text.slice(node.pos, node.end).trim();
                    if (text === "buildCapabilities()") {
                        callNode = node;
                    }
                }
                node.forEachChild(visit);
            });

            assert.ok(callNode, "Should find the call expression");
            const callType = project.checker.getTypeAtLocation(callNode);
            const calleeType = project.checker.getTypeAtLocation(callNode.expression);
            assert.ok(callType, "Should get type for call expression");
            assert.ok(calleeType, "Should get type for callee expression");

            const callSignatures = project.checker.getSignaturesOfType(calleeType, SignatureKind.Call);
            assert.ok(callSignatures.length > 0, "Callee should be callable");
            const returnType = project.checker.getReturnTypeOfSignature(callSignatures[0]);
            assert.ok(returnType, "Should get return type for the call");
            const callExprSignatures = project.checker.getSignaturesOfType(callType, SignatureKind.Call);

            assert.ok(callType.flags & TypeFlags.Object, `Expected call expression type to be object-like, got ${callType.flags}`);
            assert.ok(returnType.flags & TypeFlags.Object, `Expected return type to be object-like, got ${returnType.flags}`);
            assert.equal(callType.flags, returnType.flags, "Call expression type should have same flags as method return type");
            assert.equal(callExprSignatures.length, 0, "Call expression result type should not itself be callable");
        }
        finally {
            api.close();
        }
    });

    test("getSignaturesOfType - call signatures", () => {
        const api = spawnAPI(checkerFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = checkerFiles["/src/main.tlua"];
            const addPos = src.indexOf("add(");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", addPos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const callSigs = project.checker.getSignaturesOfType(type, SignatureKind.Call);
            assert.ok(callSigs.length > 0);
            const sig = callSigs[0];
            assert.ok(sig.id);
            assert.ok(sig.parameters.length >= 2);
            assert.ok(sig.hasRestParameter);
            assert.ok(!sig.isConstruct);
            assert.ok(!sig.isAbstract);
        }
        finally {
            api.close();
        }
    });

    test("getSignaturesOfType - construct signatures", () => {
        const api = spawnAPI(checkerFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = checkerFiles["/src/main.tlua"];
            const classPos = src.indexOf("MyClass");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", classPos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const constructSigs = project.checker.getSignaturesOfType(type, SignatureKind.Construct);
            assert.ok(constructSigs.length > 0);
            const sig = constructSigs[0];
            assert.ok(sig.isConstruct);
        }
        finally {
            api.close();
        }
    });

    test("Signature declaration can be resolved", () => {
        const api = spawnAPI(checkerFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = checkerFiles["/src/main.tlua"];
            const addPos = src.indexOf("add(");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", addPos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const callSigs = project.checker.getSignaturesOfType(type, SignatureKind.Call);
            assert.ok(callSigs.length > 0);
            const sig = callSigs[0];
            assert.ok(sig.declaration);
            const node = sig.declaration.resolve(project);
            assert.ok(node);
            // The handle remembers its canonical project, so resolve() works without an argument.
            const nodeFromCanonical = sig.declaration.resolve();
            assert.ok(nodeFromCanonical);
            assert.strictEqual(nodeFromCanonical.kind, node.kind);

            const methodPos = src.indexOf("getValue");
            const methodSymbol = project.checker.getSymbolAtPosition("/src/main.tlua", methodPos);
            assert.ok(methodSymbol);
            assert.ok(methodSymbol.valueDeclaration);
            const methodNode = methodSymbol.valueDeclaration.resolve(project);
            assert.ok(methodNode);
            assert.strictEqual(methodNode.parent.kind, SyntaxKind.ObjectLiteralExpression);
            assert.strictEqual(methodNode.parent.parent.kind, SyntaxKind.VariableDeclaration);
            // A symbol's declaration handles default to the symbol's canonical project.
            const methodNodeFromCanonical = methodSymbol.valueDeclaration.resolve();
            assert.ok(methodNodeFromCanonical);
            assert.strictEqual(methodNodeFromCanonical.kind, methodNode.kind);
        }
        finally {
            api.close();
        }
    });

    test("getSignaturesOfType - signature type parameters", () => {
        const mainFile = `
            interface Operator<T, R> {
            }
            interface Observable<T> {
                lift<R>(operator: Operator<T, R>): Observable<R>;
            }
            `;
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": mainFile,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const liftPos = mainFile.indexOf("lift");
            const type = project.checker.getTypeAtPosition("/src/main.tlua", liftPos);
            assert.ok(type);
            const callSigs = project.checker.getSignaturesOfType(type, SignatureKind.Call);
            assert.ok(callSigs.length === 1, "should have exactly one call signature, found: " + callSigs.length);
            const sig = callSigs[0];
            assert.ok(sig.typeParameters?.length === 1, "should have exactly one type parameter, found: " + sig.typeParameters?.length);
            const typeParams = sig.getTypeParameters();
            const typeParam = typeParams[0];
            assert.ok(typeParam, "should have type parameter");
            const name = (typeParam.getSymbol())?.name;
            assert.ok(name === "R", "should be named R, instead: " + name);
            assert.ok(typeParam.flags & TypeFlags.TypeParameter, "should be a type parameter, instead flags: " + typeParam.flags);
        }
        finally {
            api.close();
        }
    });

    test("Signature.getParameters() returns parameter symbols with correct names", () => {
        const api = spawnAPI(checkerFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = checkerFiles["/src/main.tlua"];
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("add("));
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const sigs = project.checker.getSignaturesOfType(type, SignatureKind.Call);
            assert.ok(sigs.length > 0);
            const params = sigs[0].getParameters();
            assert.equal(params.length, 3);
            assert.equal(params[0].name, "a");
            assert.equal(params[1].name, "b");
            assert.equal(params[2].name, "rest");
            assert.ok(params[0].flags & SymbolFlags.FunctionScopedVariable, `expected FunctionScopedVariable on 'a', got ${params[0].flags}`);
        }
        finally {
            api.close();
        }
    });

    test("Signature.getThisParameter() returns undefined when no explicit this parameter", () => {
        const api = spawnAPI(checkerFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = checkerFiles["/src/main.tlua"];
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("add("));
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const sigs = project.checker.getSignaturesOfType(type, SignatureKind.Call);
            assert.ok(sigs.length > 0);
            const thisParam = sigs[0].getThisParameter();
            assert.strictEqual(thisParam, undefined, "add() has no explicit this parameter");
        }
        finally {
            api.close();
        }
    });

    test("Signature.getTarget() returns undefined for a non-instantiated signature", () => {
        const api = spawnAPI(checkerFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = checkerFiles["/src/main.tlua"];
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("add("));
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const sigs = project.checker.getSignaturesOfType(type, SignatureKind.Call);
            assert.ok(sigs.length > 0);
            const target = sigs[0].getTarget();
            assert.strictEqual(target, undefined, "add() is not an instantiated signature");
        }
        finally {
            api.close();
        }
    });

    test("Signature.getTarget() returns the generic source signature for an instantiated call", () => {
        const src = `
            function identity<T>(x: T): T return x; end
            identity<string>("hello");
        `;
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);
            let callNode: Node | undefined;
            sourceFile.forEachChild(function visit(node) {
                if (isCallExpression(node)) callNode = node;
                node.forEachChild(visit);
            });
            assert.ok(callNode, "should find a call expression");
            const sig = project.checker.getResolvedSignature(callNode);
            assert.ok(sig, "should resolve a signature for the call");
            assert.ok(sig.target !== undefined, "instantiated call should have a target ID");
            const target = sig.getTarget();
            assert.ok(target, "getTarget() should return the generic signature");
            assert.ok(target.typeParameters && target.typeParameters.length > 0, "target should have type parameters");
        }
        finally {
            api.close();
        }
    });
});

describe("Symbol - parent, members, exports", () => {
    const symbolFiles = {
        "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
        "/src/mod.tlua": `
interface Animal {
    name: string;
    speak(): void;
}
local value = 1;
return { value = value };
`,
    };

    test("getMembers returns interface members", () => {
        const api = spawnAPI(symbolFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = symbolFiles["/src/mod.tlua"];
            const animalPos = src.indexOf("Animal");
            const symbol = project.checker.getSymbolAtPosition("/src/mod.tlua", animalPos);
            assert.ok(symbol);
            const members = symbol.getMembers();
            assert.ok(members.size > 0);
            const memberNames = [...members.values()].map(m => m.name);
            assert.ok(memberNames.includes("name"), "should have 'name' member");
            assert.ok(memberNames.includes("speak"), "should have 'speak' member");
        }
        finally {
            api.close();
        }
    });

    test("getExports returns module exports via sourceFile symbol", () => {
        const api = spawnAPI(symbolFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/mod.tlua");
            assert.ok(sourceFile);
            const moduleSymbol = project.checker.getSymbolAtLocation(sourceFile);
            assert.ok(moduleSymbol);
            const exports = moduleSymbol.getExports();
            assert.ok(exports.size > 0);
            const exportNames = [...exports.values()].map(e => e.name);
            // A Lua module's only export is the export= symbol carrying the
            // return table; types are global, not exported.
            assert.ok(exportNames.includes("export="), "should have the return-table export");
        }
        finally {
            api.close();
        }
    });

    test("getParent returns containing symbol", () => {
        const api = spawnAPI(symbolFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = symbolFiles["/src/mod.tlua"];
            const namePos = src.indexOf("name:");
            const nameSymbol = project.checker.getSymbolAtPosition("/src/mod.tlua", namePos);
            assert.ok(nameSymbol);
            assert.equal(nameSymbol.name, "name");
            const parent = nameSymbol.getParent();
            assert.ok(parent);
            assert.equal(parent.name, "Animal");
        }
        finally {
            api.close();
        }
    });
});

describe("Type - getSymbol", () => {
    test("getSymbol returns the symbol of a type", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/types.d.tlua": `
interface Foo {
    x: number;
}
declare instance: Foo;
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `\ninterface Foo {\n    x: number;\n}\ndeclare instance: Foo;\n`;
            const instancePos = src.indexOf("instance");
            const symbol = project.checker.getSymbolAtPosition("/src/types.d.tlua", instancePos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const typeSymbol = type.getSymbol();
            assert.ok(typeSymbol);
            assert.equal(typeSymbol.name, "Foo");
        }
        finally {
            api.close();
        }
    });
});

describe("Type - sub-property fetchers", () => {
    const typeFiles = {
        "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true, target: "esnext" } }),
        "/src/types.d.tlua": `
declare arr: Array<number>;
declare union: string | number;
declare intersection: { a: number } & { b: string };
type KeyOf<T> = keyof T;
type Lookup<T, K extends keyof T> = T[K];
type Cond<T> = T extends string ? "yes" : "no";
declare tpl: \`hello \${string}\`;
type Upper = Uppercase<"hello">;
declare tuple: readonly [number, string?, ...boolean[]];
`,
    };

    function getTypeAtName(api: API, name: string) {
        const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
        const project = snapshot.getProject("/tluaconfig.json")!;
        const src = typeFiles["/src/types.d.tlua"];
        const pos = src.indexOf(name);
        assert.ok(pos >= 0, `Could not find "${name}" in source`);
        const symbol = project.checker.getSymbolAtPosition("/src/types.d.tlua", pos);
        assert.ok(symbol, `No symbol found at "${name}"`);
        const type = project.checker.getTypeOfSymbol(symbol);
        assert.ok(type, `No type found for symbol "${name}"`);
        return { type, project, snapshot, api };
    }

    test("TypeReference.getTarget() returns the generic target", () => {
        const { type, api } = getTypeAtName(spawnAPI(typeFiles), "arr:");
        try {
            assert.ok(type.flags & TypeFlags.Object);
            const ref = type as TypeReference;
            assert.ok(ref.objectFlags & ObjectFlags.Reference);
            assert.equal(type.isObjectType(), true);
            assert.equal(type.isTypeReference(), true);
            assert.equal(type.isLiteralType(), false);
            const target = ref.getTarget();
            assert.ok(target);
            assert.ok(target.flags & TypeFlags.Object);
        }
        finally {
            api.close();
        }
    });

    test("UnionOrIntersectionType.getTypes() returns union members", () => {
        const { type, api } = getTypeAtName(spawnAPI(typeFiles), "union:");
        try {
            assert.ok(type.flags & TypeFlags.Union);
            const union = type as UnionOrIntersectionType;
            const types = union.getTypes();
            assert.ok(types.length >= 2);
            assert.equal(type.isUnionType(), true);
            assert.equal(type.isIntersectionType(), false);
        }
        finally {
            api.close();
        }
    });

    test("UnionOrIntersectionType.getTypes() returns intersection members", () => {
        const { type, api } = getTypeAtName(spawnAPI(typeFiles), "intersection:");
        try {
            assert.ok(type.flags & TypeFlags.Intersection);
            const inter = type as UnionOrIntersectionType;
            const types = inter.getTypes();
            assert.ok(types.length >= 2);
        }
        finally {
            api.close();
        }
    });

    test("UnionOrIntersectionType.getTypes() on a wrongly-cast type returns undefined without hitting the server", () => {
        const src = `local s: string = ""; local u: string | number = "";`;
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;

            // `string` is neither a union/intersection nor a template literal type,
            // so it has no constituent types. The client guards on the type's flags
            // and returns undefined without ever sending a request the server cannot satisfy.
            const sSymbol = project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("s:"));
            assert.ok(sSymbol);
            const sType = project.checker.getTypeOfSymbol(sSymbol);
            assert.ok(sType);
            assert.equal((sType as unknown as UnionOrIntersectionType).getTypes(), undefined);

            // A real union still returns its constituents.
            const uSymbol = project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("u:"));
            assert.ok(uSymbol);
            const uType = project.checker.getTypeOfSymbol(uSymbol);
            assert.ok(uType);
            assert.equal(((uType as UnionOrIntersectionType).getTypes()).length, 2);
        }
        finally {
            api.close();
        }
    });

    test("IndexType.getTarget() returns the target type", () => {
        const api = spawnAPI(typeFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const symbol = project.checker.resolveName("KeyOf", SymbolFlags.TypeAlias, { document: "/src/types.d.tlua", position: 0 });
            assert.ok(symbol);
            const type = project.checker.getDeclaredTypeOfSymbol(symbol);
            assert.ok(type);
            // KeyOf<T> = keyof T — this is an IndexType
            assert.ok(type.flags & TypeFlags.Index, `Expected IndexType, got flags ${type.flags}`);
            const indexType = type as IndexType;
            const target = indexType.getTarget();
            assert.ok(target);
        }
        finally {
            api.close();
        }
    });

    test("IndexedAccessType.getObjectType() and getIndexType()", () => {
        const api = spawnAPI(typeFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const symbol = project.checker.resolveName("Lookup", SymbolFlags.TypeAlias, { document: "/src/types.d.tlua", position: 0 });
            assert.ok(symbol);
            const type = project.checker.getDeclaredTypeOfSymbol(symbol);
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.IndexedAccess, `Expected IndexedAccessType, got flags ${type.flags}`);
            const ia = type as IndexedAccessType;
            assert.equal(type.isIndexedAccessType(), true);
            const objectType = ia.getObjectType();
            assert.ok(objectType);
            const indexType = ia.getIndexType();
            assert.ok(indexType);
        }
        finally {
            api.close();
        }
    });

    test("ConditionalType.getCheckType() and getExtendsType()", () => {
        const api = spawnAPI(typeFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const symbol = project.checker.resolveName("Cond", SymbolFlags.TypeAlias, { document: "/src/types.d.tlua", position: 0 });
            assert.ok(symbol);
            const type = project.checker.getDeclaredTypeOfSymbol(symbol);
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.Conditional, `Expected ConditionalType, got flags ${type.flags}`);
            const cond = type as ConditionalType;
            assert.equal(type.isConditionalType(), true);
            const checkType = cond.getCheckType();
            assert.ok(checkType);
            const extendsType = cond.getExtendsType();
            assert.ok(extendsType);
        }
        finally {
            api.close();
        }
    });

    test("ConditionalType.getTrueType() and getFalseType()", () => {
        const api = spawnAPI(typeFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const symbol = project.checker.resolveName("Cond", SymbolFlags.TypeAlias, { document: "/src/types.d.tlua", position: 0 });
            assert.ok(symbol);
            const type = project.checker.getDeclaredTypeOfSymbol(symbol);
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.Conditional, `Expected ConditionalType, got flags ${type.flags}`);

            const trueType = (type as ConditionalType).getTrueType();
            assert.ok(trueType, "should return the true-branch type");
            assert.ok(trueType.flags & TypeFlags.StringLiteral, `Expected StringLiteral for true branch, got flags ${trueType.flags}`);
            assert.equal((trueType as LiteralType).value, "yes");

            const falseType = (type as ConditionalType).getFalseType();
            assert.ok(falseType, "should return the false-branch type");
            assert.ok(falseType.flags & TypeFlags.StringLiteral, `Expected StringLiteral for false branch, got flags ${falseType.flags}`);
            assert.equal((falseType as LiteralType).value, "no");
        }
        finally {
            api.close();
        }
    });

    test("TemplateLiteralType.texts and getTypes()", () => {
        const { type, api } = getTypeAtName(spawnAPI(typeFiles), "tpl:");
        try {
            assert.ok(type.flags & TypeFlags.TemplateLiteral, `Expected TemplateLiteralType, got flags ${type.flags}`);
            const tpl = type as TemplateLiteralType;
            assert.ok(tpl.texts);
            assert.ok(tpl.texts.length >= 2);
            assert.equal(tpl.texts[0], "hello ");
            const types = tpl.getTypes();
            assert.ok(types.length >= 1);
        }
        finally {
            api.close();
        }
    });

    test("StringMappingType.getTarget() returns the mapped type", () => {
        const api = spawnAPI(typeFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = typeFiles["/src/types.d.tlua"];
            const pos = src.indexOf("Upper");
            const symbol = project.checker.getSymbolAtPosition("/src/types.d.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            // Uppercase<"hello"> may resolve to a StringMappingType or a string literal
            if (type.flags & TypeFlags.StringMapping) {
                const sm = type as StringMappingType;
                const target = sm.getTarget();
                assert.ok(target);
            }
            // If it resolved to "HELLO" literal, that's fine too — it means eager evaluation
        }
        finally {
            api.close();
        }
    });

    test("TupleType properties", () => {
        const { type, api } = getTypeAtName(spawnAPI(typeFiles), "tuple:");
        try {
            assert.ok(type.flags & TypeFlags.Object);
            const ref = type as TypeReference;
            assert.ok(ref.objectFlags & ObjectFlags.Reference);
            const target = ref.getTarget();
            assert.ok(target);
            assert.ok(target.flags & TypeFlags.Object);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - intrinsic type getters", () => {
    const intrinsicFiles = {
        "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
        "/src/main.tlua": `local x = 1;\nreturn { x = x };`,
    };

    test("getAnyType returns a type with Any flag", () => {
        const api = spawnAPI(intrinsicFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const type = project.checker.getAnyType();
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.Any);
        }
        finally {
            api.close();
        }
    });

    test("getStringType returns a type with String flag", () => {
        const api = spawnAPI(intrinsicFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const type = project.checker.getStringType();
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.String);
        }
        finally {
            api.close();
        }
    });

    test("getNumberType returns a type with Number flag", () => {
        const api = spawnAPI(intrinsicFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const type = project.checker.getNumberType();
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.Number);
        }
        finally {
            api.close();
        }
    });

    test("getBooleanType returns a type with Boolean flag", () => {
        const api = spawnAPI(intrinsicFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const type = project.checker.getBooleanType();
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.Boolean);
        }
        finally {
            api.close();
        }
    });

    test("getVoidType returns a type with Void flag", () => {
        const api = spawnAPI(intrinsicFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const type = project.checker.getVoidType();
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.Void);
        }
        finally {
            api.close();
        }
    });

    test("getUndefinedType returns a type with Nil flag", () => {
        const api = spawnAPI(intrinsicFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const type = project.checker.getUndefinedType();
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.Nil);
        }
        finally {
            api.close();
        }
    });

    test("getNullType returns a type with Nil flag", () => {
        const api = spawnAPI(intrinsicFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const type = project.checker.getNullType();
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.Nil);
        }
        finally {
            api.close();
        }
    });

    test("getNeverType returns a type with Never flag", () => {
        const api = spawnAPI(intrinsicFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const type = project.checker.getNeverType();
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.Never);
        }
        finally {
            api.close();
        }
    });

    test("getUnknownType returns a type with Unknown flag", () => {
        const api = spawnAPI(intrinsicFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const type = project.checker.getUnknownType();
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.Unknown);
        }
        finally {
            api.close();
        }
    });

    test("getESSymbolType returns a type with ESSymbol flag", () => {
        const api = spawnAPI(intrinsicFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const type = project.checker.getESSymbolType();
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.ESSymbol);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - multi-project type ID uniqueness", () => {
    test("intrinsic types from 3 projects in the same snapshot have non-colliding IDs", () => {
        const api = spawnAPI({
            "/proj1/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/proj1/src/index.tlua": `local x = 1;\nreturn { x = x };`,
            "/proj2/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/proj2/src/index.tlua": `local y = "hello";\nreturn { y = y };`,
            "/proj3/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/proj3/src/index.tlua": `local z = true;\nreturn { z = z };`,
        });
        try {
            // Open all 3 projects — each updateSnapshot accumulates open projects
            api.updateSnapshot({ openProject: "/proj1/tluaconfig.json" });
            api.updateSnapshot({ openProject: "/proj2/tluaconfig.json" });
            const snapshot = api.updateSnapshot({ openProject: "/proj3/tluaconfig.json" });

            const proj1 = snapshot.getProject("/proj1/tluaconfig.json")!;
            const proj2 = snapshot.getProject("/proj2/tluaconfig.json")!;
            const proj3 = snapshot.getProject("/proj3/tluaconfig.json")!;
            assert.ok(proj1, "proj1 should be in final snapshot");
            assert.ok(proj2, "proj2 should be in final snapshot");
            assert.ok(proj3, "proj3 should be in final snapshot");

            // Fetch several intrinsic types from each checker.
            // If type IDs collide across checkers, registerType panics → API error.
            const num1 = proj1.checker.getNumberType();
            const str1 = proj1.checker.getStringType();
            const bool1 = proj1.checker.getBooleanType();
            const any1 = proj1.checker.getAnyType();
            const num2 = proj2.checker.getNumberType();
            const str2 = proj2.checker.getStringType();
            const bool2 = proj2.checker.getBooleanType();
            const any2 = proj2.checker.getAnyType();
            const num3 = proj3.checker.getNumberType();
            const str3 = proj3.checker.getStringType();
            const bool3 = proj3.checker.getBooleanType();
            const any3 = proj3.checker.getAnyType();

            assert.ok(num1.flags & TypeFlags.Number, "proj1 number type");
            assert.ok(str1.flags & TypeFlags.String, "proj1 string type");
            assert.ok(bool1.flags & TypeFlags.Boolean, "proj1 boolean type");
            assert.ok(any1.flags & TypeFlags.Any, "proj1 any type");

            assert.ok(num2.flags & TypeFlags.Number, "proj2 number type");
            assert.ok(str2.flags & TypeFlags.String, "proj2 string type");
            assert.ok(bool2.flags & TypeFlags.Boolean, "proj2 boolean type");
            assert.ok(any2.flags & TypeFlags.Any, "proj2 any type");

            assert.ok(num3.flags & TypeFlags.Number, "proj3 number type");
            assert.ok(str3.flags & TypeFlags.String, "proj3 string type");
            assert.ok(bool3.flags & TypeFlags.Boolean, "proj3 boolean type");
            assert.ok(any3.flags & TypeFlags.Any, "proj3 any type");
        }
        finally {
            api.close();
        }
    });

    test("symbol and signature handles from 3 projects in the same snapshot have non-colliding IDs", () => {
        const api = spawnAPI({
            "/proj1/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/proj1/src/index.tlua": `function add(a: number, b: number): number return a + b; end`,
            "/proj2/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/proj2/src/index.tlua": `function greet(name: string): string return "hello " + name; end`,
            "/proj3/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/proj3/src/index.tlua": `function toggle(b: boolean): boolean return !b; end`,
        });
        try {
            api.updateSnapshot({ openProject: "/proj1/tluaconfig.json" });
            api.updateSnapshot({ openProject: "/proj2/tluaconfig.json" });
            const snapshot = api.updateSnapshot({ openProject: "/proj3/tluaconfig.json" });

            const proj1 = snapshot.getProject("/proj1/tluaconfig.json")!;
            const proj2 = snapshot.getProject("/proj2/tluaconfig.json")!;
            const proj3 = snapshot.getProject("/proj3/tluaconfig.json")!;

            // Get a symbol from each project (exercises symbol registry)
            const src1 = `function add(a: number, b: number): number return a + b; end`;
            const src2 = `function greet(name: string): string return "hello " + name; end`;
            const src3 = `function toggle(b: boolean): boolean return !b; end`;

            const sym1 = proj1.checker.getSymbolAtPosition("/proj1/src/index.tlua", src1.indexOf("add"));
            const sym2 = proj2.checker.getSymbolAtPosition("/proj2/src/index.tlua", src2.indexOf("greet"));
            const sym3 = proj3.checker.getSymbolAtPosition("/proj3/src/index.tlua", src3.indexOf("toggle"));
            assert.ok(sym1, "proj1 symbol");
            assert.ok(sym2, "proj2 symbol");
            assert.ok(sym3, "proj3 symbol");
            assert.equal(sym1.name, "add", "proj1 symbol name");
            assert.equal(sym2.name, "greet", "proj2 symbol name");
            assert.equal(sym3.name, "toggle", "proj3 symbol name");

            // Get type of each symbol, then signatures (exercises type + signature registries)
            const type1 = proj1.checker.getTypeOfSymbol(sym1);
            const type2 = proj2.checker.getTypeOfSymbol(sym2);
            const type3 = proj3.checker.getTypeOfSymbol(sym3);
            assert.ok(type1, "proj1 function type");
            assert.ok(type2, "proj2 function type");
            assert.ok(type3, "proj3 function type");

            const sigs1 = proj1.checker.getSignaturesOfType(type1, SignatureKind.Call);
            const sigs2 = proj2.checker.getSignaturesOfType(type2, SignatureKind.Call);
            const sigs3 = proj3.checker.getSignaturesOfType(type3, SignatureKind.Call);
            assert.equal(sigs1.length, 1, "proj1 has 1 call signature");
            assert.equal(sigs2.length, 1, "proj2 has 1 call signature");
            assert.equal(sigs3.length, 1, "proj3 has 1 call signature");
            assert.equal(sigs1[0].parameters.length, 2, "proj1 add() has 2 params");
            assert.equal(sigs2[0].parameters.length, 1, "proj2 greet() has 1 param");
            assert.equal(sigs3[0].parameters.length, 1, "proj3 toggle() has 1 param");
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getBaseTypeOfLiteralType", () => {
    test("number literal widens to number", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `local x = 42 as const;\nreturn { x = x };`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `local x = 42 as const;\nreturn { x = x };`;
            const pos = src.indexOf("x =");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const literalType = project.checker.getTypeOfSymbol(symbol);
            assert.ok(literalType);
            assert.ok(literalType.flags & TypeFlags.NumberLiteral);
            const baseType = project.checker.getBaseTypeOfLiteralType(literalType);
            assert.ok(baseType);
            assert.ok(baseType.flags & TypeFlags.Number);
        }
        finally {
            api.close();
        }
    });

    test("string literal widens to string", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `local s = "hello" as const;\nreturn { s = s };`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `local s = "hello" as const;\nreturn { s = s };`;
            const pos = src.indexOf("s ");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const literalType = project.checker.getTypeOfSymbol(symbol);
            assert.ok(literalType);
            assert.ok(literalType.flags & TypeFlags.StringLiteral);
            const baseType = project.checker.getBaseTypeOfLiteralType(literalType);
            assert.ok(baseType);
            assert.ok(baseType.flags & TypeFlags.String);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getContextualType", () => {
    test("contextual type from function parameter", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `
function foo(x: number) end
foo(42);
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;

            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);

            // Find the argument "42" in foo(42)
            // statement[1] = foo(42); which is an ExpressionStatement -> CallExpression
            const callStmt = sourceFile.statements[1];
            assert.ok(callStmt);
            let numLiteral: import("@tlua/cli/unstable/ast").Expression | undefined;
            callStmt.forEachChild(function visit(node) {
                if (isCallExpression(node)) {
                    // First argument
                    numLiteral = node.arguments[0];
                }
                node.forEachChild(visit);
            });
            assert.ok(numLiteral);
            const contextualType = project.checker.getContextualType(numLiteral);
            assert.ok(contextualType);
            assert.ok(contextualType.flags & TypeFlags.Number);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getTypeOfSymbolAtLocation", () => {
    test("narrowed type via type() check", () => {
        const api = spawnAPI({
            // `type` is a lualib global, so the LuaJIT declarations must be in the program.
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true, lib: ["esnext", "luajit"] } }),
            "/src/main.tlua": `
function check(x: string | number)
    if type(x) == "string" then
        return x;
    end
    return x;
end
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `\nfunction check(x: string | number)\n    if type(x) == "string" then\n        return x;\n    end\n    return x;\nend\n`;

            // Get the symbol for parameter "x"
            const paramPos = src.indexOf("x:");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", paramPos);
            assert.ok(symbol);
            assert.equal(symbol.name, "x");

            // Get the type of "x" at the wider function scope — should be string | number
            const wideType = project.checker.getTypeOfSymbol(symbol);
            assert.ok(wideType);
            assert.ok(wideType.flags & TypeFlags.Union, `Expected union type, got flags ${wideType.flags}`);

            // Get the narrowed return x inside the if block
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);

            // statement[0] is the function declaration
            const funcDecl = sourceFile.statements[0];
            assert.ok(funcDecl);
            // Walk to find the first "return x" — inside the if, x should be narrowed to string
            let firstReturnX: import("@tlua/cli/unstable/ast").Node | undefined;
            funcDecl.forEachChild(function visit(node) {
                if (isReturnStatement(node) && !firstReturnX) {
                    // The expression of the return statement is the identifier "x"
                    if (node.expression) {
                        firstReturnX = node.expression;
                    }
                }
                node.forEachChild(visit);
            });
            assert.ok(firstReturnX);
            const narrowedType = project.checker.getTypeOfSymbolAtLocation(symbol, firstReturnX);
            assert.ok(narrowedType);
            // Inside the if (typeof x === "string") branch, x should be narrowed to string
            assert.ok(narrowedType.flags & TypeFlags.String, `Expected string type, got flags ${narrowedType.flags}`);
        }
        finally {
            api.close();
        }
    });
});

describe("readFile callback semantics", () => {
    test("readFile: string returns content, null blocks fallback, undefined falls through to real FS", () => {
        const virtualFiles: Record<string, string> = {
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/index.tlua": `local x: number = 1;`,
        };
        const vfs = createVirtualFileSystem(virtualFiles);
        const blockedPath = "/src/blocked.tlua";

        const fs: FileSystem = {
            ...vfs,
            readFile: (fileName: string) => {
                if (fileName === blockedPath) {
                    // null = file not found, don't fall back to real FS
                    return null;
                }
                // Try the VFS first; if it has the file, return its content (string).
                // Otherwise return undefined to fall through to the real FS.
                return vfs.readFile!(fileName);
            },
        };

        const api = new API({
            cwd: fileURLToPath(new URL("../../../../", import.meta.url).toString()),
            fs,
        });

        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;

            // 1. String content: virtual file is found
            const sf = project.program.getSourceFile("/src/index.tlua");
            assert.ok(sf, "Virtual file should be found");
            assert.equal(sf.text, virtualFiles["/src/index.tlua"]);

            // 2. undefined fallback: lib files from the real FS should be present.
            //    If readFile returned null for unknowns, lib files would be missing
            //    and `number` would not resolve — this was the original async bug.
            //    Verify by checking that `number` resolves to a proper type (not error).
            const pos = virtualFiles["/src/index.tlua"].indexOf("x:");
            const type = project.checker.getTypeAtPosition("/src/index.tlua", pos);
            assert.ok(type, "Type should resolve");
            assert.ok(type.flags & TypeFlags.Number, `Expected number type, got flags ${type.flags}`);

            // 3. null blocks fallback: blocked file should not be found
            const blockedSf = project.program.getSourceFile(blockedPath);
            assert.equal(blockedSf, undefined, "Blocked file should not be found (null prevents fallback)");
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - isArrayType / isTupleType", () => {
    test("number[] is array, not tuple", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `local xs: number[] = [];`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `local xs: number[] = [];`;
            const pos = src.indexOf("xs");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            assert.equal(project.checker.isArrayType(type), true);
            assert.equal(project.checker.isTupleType(type), false);
        }
        finally {
            api.close();
        }
    });

    test("readonly number[] is array", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `local xs: readonly number[] = [];`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `local xs: readonly number[] = [];`;
            const pos = src.indexOf("xs");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            assert.equal(project.checker.isArrayType(type), true);
            assert.equal(project.checker.isTupleType(type), false);
        }
        finally {
            api.close();
        }
    });

    test("Array<number> is distinct from the T[] table type", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `local xs: Array<number> = [];`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `local xs: Array<number> = [];`;
            const pos = src.indexOf("xs");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            // Array<T> is inert and distinct from the T[] table type: not isArrayType.
            assert.equal(project.checker.isArrayType(type), false);
            assert.equal(project.checker.isTupleType(type), false);
        }
        finally {
            api.close();
        }
    });

    test("[number, string] is tuple, not array", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            // Tuple type syntax parses only in declaration files after the removal.
            "/src/shapes.d.tlua": `declare tup: [number, string];`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `declare tup: [number, string];`;
            const pos = src.indexOf("tup");
            const symbol = project.checker.getSymbolAtPosition("/src/shapes.d.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            assert.equal(project.checker.isArrayType(type), false);
            assert.equal(project.checker.isTupleType(type), true);
        }
        finally {
            api.close();
        }
    });

    test("readonly [number, string] is tuple, not array", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            // Tuple type syntax parses only in declaration files after the removal.
            "/src/shapes.d.tlua": `declare tup: readonly [number, string];`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `declare tup: readonly [number, string];`;
            const pos = src.indexOf("tup");
            const symbol = project.checker.getSymbolAtPosition("/src/shapes.d.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            assert.equal(project.checker.isArrayType(type), false);
            assert.equal(project.checker.isTupleType(type), true);
        }
        finally {
            api.close();
        }
    });

    test("string is neither array nor tuple", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `local str: string = "";`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `local str: string = "";`;
            const pos = src.indexOf("str");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            assert.equal(project.checker.isArrayType(type), false);
            assert.equal(project.checker.isTupleType(type), false);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getReturnTypeOfSignature", () => {
    test("returns the return type of a function signature", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `function add(a: number, b: number): number return a + b; end`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `function add(a: number, b: number): number return a + b; end`;
            const pos = src.indexOf("add(");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const sigs = project.checker.getSignaturesOfType(type, SignatureKind.Call);
            assert.ok(sigs.length > 0);
            const returnType = project.checker.getReturnTypeOfSignature(sigs[0]);
            assert.ok(returnType);
            assert.ok(returnType.flags & TypeFlags.Number, `Expected number, got flags ${returnType.flags}`);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getRestTypeOfSignature", () => {
    test("returns the rest type of a signature with rest parameter", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `function sum(...nums: number[]): number return nums.reduce(function(a, b) return a + b end, 0); end`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `function sum(...nums: number[]): number return nums.reduce(function(a, b) return a + b end, 0); end`;
            const pos = src.indexOf("sum(");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const sigs = project.checker.getSignaturesOfType(type, SignatureKind.Call);
            assert.ok(sigs.length > 0);
            const restType = project.checker.getRestTypeOfSignature(sigs[0]);
            assert.ok(restType);
            assert.ok(restType.flags & TypeFlags.Number, `Expected number type, got flags ${restType.flags}`);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getTypePredicateOfSignature", () => {
    test("returns type predicate for 'x is T' guard", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `function isString(x: unknown): x is string return typeof x === "string"; end`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `function isString(x: unknown): x is string return typeof x === "string"; end`;
            const pos = src.indexOf("isString(");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const sigs = project.checker.getSignaturesOfType(type, SignatureKind.Call);
            assert.ok(sigs.length > 0);
            const predicate = project.checker.getTypePredicateOfSignature(sigs[0]);
            assert.ok(predicate);
            assert.equal(predicate.kind, TypePredicateKind.Identifier);
            assert.equal(predicate.parameterName, "x");
            assert.equal(predicate.parameterIndex, 0);
            assert.ok(predicate.type);
            assert.ok(predicate.type.flags & TypeFlags.String);
        }
        finally {
            api.close();
        }
    });

    test("returns type predicate for 'asserts x is T'", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `function assertIsString(x: unknown): asserts x is string if (typeof x ~= "string") then throw new Error(); end end`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `function assertIsString(x: unknown): asserts x is string if (typeof x ~= "string") then throw new Error(); end end`;
            const pos = src.indexOf("assertIsString(");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const sigs = project.checker.getSignaturesOfType(type, SignatureKind.Call);
            assert.ok(sigs.length > 0);
            const predicate = project.checker.getTypePredicateOfSignature(sigs[0]);
            assert.ok(predicate);
            assert.equal(predicate.kind, TypePredicateKind.AssertsIdentifier);
            assert.equal(predicate.parameterName, "x");
            assert.equal(predicate.parameterIndex, 0);
            assert.ok(predicate.type);
            assert.ok(predicate.type.flags & TypeFlags.String);
        }
        finally {
            api.close();
        }
    });

    test("returns undefined for signature without type predicate", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `function add(a: number, b: number): number return a + b; end`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `function add(a: number, b: number): number return a + b; end`;
            const pos = src.indexOf("add(");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const sigs = project.checker.getSignaturesOfType(type, SignatureKind.Call);
            assert.ok(sigs.length > 0);
            const predicate = project.checker.getTypePredicateOfSignature(sigs[0]);
            assert.equal(predicate, undefined);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getBaseTypes", () => {
    // `class` is removed in tlua; `getBaseTypes` on class types was redundant with
    // the interface case below, which exercises the same InterfaceType path.
    test("returns base types of an interface", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `
interface Animal {
    name: string;
}
interface Dog extends Animal {
    bark(): void;
}
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `\ninterface Animal {\n    name: string;\n}\ninterface Dog extends Animal {\n    bark(): void;\n}\n`;
            const pos = src.indexOf("Dog");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getDeclaredTypeOfSymbol(symbol);
            assert.ok(type);
            const baseTypes = project.checker.getBaseTypes(type as InterfaceType);
            assert.ok(baseTypes.length > 0, "Should have at least one base type");
            const baseSymbol = baseTypes[0].getSymbol();
            assert.ok(baseSymbol);
            assert.equal(baseSymbol.name, "Animal");
        }
        finally {
            api.close();
        }
    });

    test("does not panic for a type alias to a generic interface instantiation", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `
interface Box<T> {
    value: T;
}
type BoxOfString = Box<string>;
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);
            const typeAlias = sourceFile.statements.find(isTypeAliasDeclaration);
            assert.ok(typeAlias);
            const type = project.checker.getTypeAtLocation(typeAlias);
            assert.ok(type);
            // A generic interface instantiation produces a type reference, not an
            // interface type, so it has no base types and yields [].
            const baseTypes = project.checker.getBaseTypes(type as InterfaceType);
            assert.deepEqual(baseTypes, []);
        }
        finally {
            api.close();
        }
    });
});

describe("Type - getBaseTypes", () => {
    test("returns base types for a class/interface type and undefined for a non-class/interface", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `
interface Base {
    x: number;
}
interface Derived extends Base {
    y: string;
}
local n: number = 0;
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `\ninterface Base {\n    x: number;\n}\ninterface Derived extends Base {\n    y: string;\n}\nlocal n: number = 0;\n`;

            const derivedSymbol = project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("Derived"));
            assert.ok(derivedSymbol);
            const derivedType = project.checker.getDeclaredTypeOfSymbol(derivedSymbol);
            assert.ok(derivedType.isClassOrInterface(), "Derived should be a class/interface type");
            const baseTypes = derivedType.getBaseTypes();
            assert.ok(baseTypes && baseTypes.length > 0, "interface type should have base types");
            const baseSymbol = baseTypes![0].getSymbol();
            assert.equal(baseSymbol?.name, "Base");

            // A primitive type is not a class/interface, so getBaseTypes() is undefined.
            const numberSymbol = project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("n:"));
            assert.ok(numberSymbol);
            const numberType = project.checker.getTypeOfSymbol(numberSymbol);
            assert.ok(numberType);
            assert.equal(numberType.isClassOrInterface(), false);
            assert.equal(numberType.getBaseTypes(), undefined);
        }
        finally {
            api.close();
        }
    });
});

describe("Type - isErrorType", () => {
    test("identifies the error type from an unresolvable annotation", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `
declare local good: string;
declare local bad: ThisTypeDoesNotExist;
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `\ndeclare local good: string;\ndeclare local bad: ThisTypeDoesNotExist;\n`;

            const badSymbol = project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("bad"));
            assert.ok(badSymbol);
            const badType = project.checker.getTypeOfSymbol(badSymbol);
            assert.ok(badType);
            assert.ok(badType.isErrorType(), "unresolved annotation should yield the error type");
            assert.ok(isErrorType(badType));

            const goodSymbol = project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("good"));
            assert.ok(goodSymbol);
            const goodType = project.checker.getTypeOfSymbol(goodSymbol);
            assert.ok(goodType);
            assert.equal(goodType.isErrorType(), false, "string type is not the error type");
            assert.equal(isErrorType(goodType), false);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - well-known symbols", () => {
    test("isUnknownSymbol identifies the aliased unknown symbol", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `
local value = 1;
type Alias = typeof value;
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `\nlocal value = 1;\ntype Alias = typeof value;\n`;

            // A real symbol is not the unknown/undefined symbol.
            const valueSymbol = project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("value"));
            assert.ok(valueSymbol);
            assert.equal(project.checker.isUnknownSymbol(valueSymbol), false);
            assert.equal(project.checker.isUndefinedSymbol(valueSymbol), false);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - well-known signatures", () => {
    test("isUnknownSignature identifies an unresolvable call", () => {
        const src = `
local ok = function(x: number) return x end;
ok(1);
local notCallable = 1;
notCallable();
`;
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);
            const calls: Node[] = [];
            sourceFile.forEachChild(function visit(node) {
                if (isCallExpression(node)) calls.push(node);
                node.forEachChild(visit);
            });
            assert.equal(calls.length, 2, "should find two call expressions");

            // The valid call resolves to a real signature, not the unknown signature.
            const okSig = project.checker.getResolvedSignature(calls[0]);
            assert.equal(project.checker.isUnknownSignature(okSig), false);

            // The call to a non-callable value yields the unknown signature.
            const badSig = project.checker.getResolvedSignature(calls[1]);
            assert.equal(project.checker.isUnknownSignature(badSig), true);
        }
        finally {
            api.close();
        }
    });
});

describe("Symbol - escaped names and tables", () => {
    test("getExports/getMembers return a cached Map keyed by escaped name", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `
interface Animal {
    name: string;
    speak(): void;
}
local value = 1;
return { value = value };
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);
            const moduleSymbol = project.checker.getSymbolAtLocation(sourceFile);
            assert.ok(moduleSymbol);

            const exports = moduleSymbol.getExports();
            assert.ok(exports.get(escapeLeadingUnderscores("export=")));
            // Calling again returns the same cached Map instance.
            assert.strictEqual(moduleSymbol.getExports(), exports);

            const src = `
interface Animal {
    name: string;
    speak(): void;
}
local value = 1;
return { value = value };
`;
            const animal = (project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("Animal")))!;
            assert.ok(animal);
            const members = animal.getMembers();
            assert.ok(members.get(escapeLeadingUnderscores("name")));
            assert.ok(members.get(escapeLeadingUnderscores("speak")));
            assert.strictEqual(animal.getMembers(), members);
        }
        finally {
            api.close();
        }
    });

    test("anonymous type symbol has a __-escaped name, never the \\xFE sentinel", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `
local obj: { a: number } = { a = 1 };
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `\nlocal obj: { a: number } = { a = 1 };\n`;
            const objSymbol = project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("obj"));
            assert.ok(objSymbol);
            const objType = project.checker.getTypeOfSymbol(objSymbol);
            assert.ok(objType);
            const typeSymbol = objType.getSymbol();
            assert.ok(typeSymbol);
            // The anonymous type literal symbol is internally named; over the wire
            // it must be the "__type" escaped form, not the raw "\xFE" sentinel.
            assert.equal(typeSymbol.escapedName, InternalSymbolName.Type);
            assert.ok(!typeSymbol.escapedName.includes("\xFE"));
        }
        finally {
            api.close();
        }
    });
});

describe("ast - escapeLeadingUnderscores", () => {
    test("round-trips display and escaped names", () => {
        assert.equal(escapeLeadingUnderscores("foo"), "foo");
        assert.equal(escapeLeadingUnderscores("_foo"), "_foo");
        assert.equal(escapeLeadingUnderscores("__foo"), "___foo");
        assert.equal(unescapeLeadingUnderscores("foo" as __String), "foo");
        assert.equal(unescapeLeadingUnderscores("__type" as __String), "__type");
        assert.equal(unescapeLeadingUnderscores("___foo" as __String), "__foo");
    });
});

describe("ast - getJSDocTags", () => {
    test("returns a node's own tags, and inherited @param / @template tags", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `
/**
 * Adds two numbers.
 * @param a the first number
 * @param b the second number
 * @returns the sum
 */
function add(a: number, b: number): number
    return a + b
end

/**
 * @template T the element type
 * @param x a value
 */
function identity<T>(x: T): T
    return x
end

/** @deprecated use add */
local total = add(1, 2);
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);
            const functions = [...sourceFile.statements].filter(isFunctionDeclaration);
            const add = functions[0];
            const identity = functions[1];
            assert.ok(add);
            assert.ok(identity);

            // A function reports its own JSDoc tags.
            assert.deepEqual(getJSDocTags(add).map(t => t.tagName.text), ["param", "param", "returns"]);

            // A parameter inherits the matching @param tag from its containing
            // signature.
            const paramA = add.parameters[0];
            const paramATags = getJSDocTags(paramA);
            assert.deepEqual(paramATags.map(t => t.tagName.text), ["param"]);
            const paramATag = paramATags[0];
            assert.ok(isJSDocParameterTag(paramATag));
            assert.ok(isIdentifier(paramATag.name));
            assert.equal(paramATag.name.text, "a");

            const paramB = add.parameters[1];
            const paramBTags = getJSDocTags(paramB);
            assert.deepEqual(paramBTags.map(t => t.tagName.text), ["param"]);
            const paramBTag = paramBTags[0];
            assert.ok(isJSDocParameterTag(paramBTag));
            assert.ok(isIdentifier(paramBTag.name));
            assert.equal(paramBTag.name.text, "b");

            // A type parameter inherits the matching @template tag.
            const typeParam = identity.typeParameters![0];
            assert.deepEqual(getJSDocTags(typeParam).map(t => t.tagName.text), ["template"]);

            // The value parameter of `identity` inherits its @param tag.
            const identityParam = identity.parameters[0];
            assert.deepEqual(getJSDocTags(identityParam).map(t => t.tagName.text), ["param"]);

            // A variable declaration walks up its comment-location chain
            // (declaration -> declaration list -> statement) to the JSDoc on the
            // containing variable statement.
            const variable = sourceFile.statements.find(isVariableStatement);
            assert.ok(variable);
            const declaration = variable.declarationList.declarations[0];
            assert.deepEqual(getJSDocTags(declaration).map(t => t.tagName.text), ["deprecated"]);
        }
        finally {
            api.close();
        }
    });

    test("a function-expression parameter inherits @param tags from the variable statement", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: {} }),
            "/src/main.tlua": `
/**
 * @param {string} name a name
 * @returns {number} the length
 */
local measure = function (name)
    return name.length;
end;
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);
            const variable = sourceFile.statements.find(isVariableStatement);
            assert.ok(variable);
            const declaration = variable.declarationList.declarations[0];
            const funcExpr = declaration.initializer;
            assert.ok(funcExpr);

            // The variable statement carries the @param and @returns tags.
            assert.deepEqual(getJSDocTags(declaration).map(t => t.tagName.text), ["param", "returns"]);

            // The function expression's parameter walks declaration -> list ->
            // statement and inherits the matching @param tag.
            const param = (funcExpr as unknown as { parameters: NodeArray<Node>; }).parameters[0];
            const paramTags = getJSDocTags(param);
            assert.deepEqual(paramTags.map(t => t.tagName.text), ["param"]);
            const paramTag = paramTags[0];
            assert.ok(isJSDocParameterTag(paramTag));
            assert.ok(isIdentifier(paramTag.name));
            assert.equal(paramTag.name.text, "name");
        }
        finally {
            api.close();
        }
    });

    test("a @type cast tag is owned by its parenthesized expression, not the declaration", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: {} }),
            "/src/main.tlua": `
/** @type {string} */
local value = "hello";

local cast = /** @type {number} */ (someValue);
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);
            const statements = [...sourceFile.statements].filter(isVariableStatement);

            // A @type tag directly on a declaration is reported for that declaration.
            const valueDecl = statements[0].declarationList.declarations[0];
            assert.deepEqual(getJSDocTags(valueDecl).map(t => t.tagName.text), ["type"]);

            // A @type cast tag attached to a parenthesized expression is owned by
            // that expression and is not reported for the enclosing declaration.
            const castDecl = statements[1].declarationList.declarations[0];
            assert.deepEqual(getJSDocTags(castDecl).map(t => t.tagName.text), []);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getPropertiesOfType", () => {
    test("returns properties of an object type", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `
interface Person {
    name: string;
    age: number;
    greet(): void;
}
declare p: Person;
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `\ninterface Person {\n    name: string;\n    age: number;\n    greet(): void;\n}\ndeclare p: Person;\n`;
            const pos = src.indexOf("p: Person");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const props = project.checker.getPropertiesOfType(type);
            assert.ok(props.length >= 3, `Expected at least 3 properties, got ${props.length}`);
            const names = props.map(p => p.name);
            assert.ok(names.includes("name"), "should have 'name' property");
            assert.ok(names.includes("age"), "should have 'age' property");
            assert.ok(names.includes("greet"), "should have 'greet' property");
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getIndexInfosOfType", () => {
    test("returns index signatures of an indexed type", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `
interface StringMap {
    [key: string]: number;
}
declare m: StringMap;
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `\ninterface StringMap {\n    [key: string]: number;\n}\ndeclare m: StringMap;\n`;
            const pos = src.indexOf("m: StringMap");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const indexInfos = project.checker.getIndexInfosOfType(type);
            assert.ok(indexInfos.length > 0, "Should have at least one index info");
            const info = indexInfos[0];
            assert.ok(info.keyType);
            assert.ok(info.keyType.flags & TypeFlags.String, `Expected string key type, got flags ${info.keyType.flags}`);
            assert.ok(info.valueType);
            assert.ok(info.valueType.flags & TypeFlags.Number, `Expected number value type, got flags ${info.valueType.flags}`);
            assert.equal(info.isReadonly, false);
        }
        finally {
            api.close();
        }
    });

    test("readonly index signature reports isReadonly true", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `
interface ReadonlyMap {
    readonly [key: string]: number;
}
declare m: ReadonlyMap;
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `\ninterface ReadonlyMap {\n    readonly [key: string]: number;\n}\ndeclare m: ReadonlyMap;\n`;
            const pos = src.indexOf("m: ReadonlyMap");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const indexInfos = project.checker.getIndexInfosOfType(type);
            assert.ok(indexInfos.length > 0);
            assert.equal(indexInfos[0].isReadonly, true);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getConstraintOfTypeParameter", () => {
    test("returns constraint of a type parameter", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `function identity<T extends string>(x: T): T return x; end`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `function identity<T extends string>(x: T): T return x; end`;
            const pos = src.indexOf("identity<");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const sigs = project.checker.getSignaturesOfType(type, SignatureKind.Call);
            assert.ok(sigs.length > 0);
            assert.ok(sigs[0].typeParameters && sigs[0].typeParameters.length > 0, "Should have type parameters");
            const typeParams = sigs[0].getTypeParameters();
            const constraint = project.checker.getConstraintOfTypeParameter(typeParams[0]);
            assert.ok(constraint);
            assert.ok(constraint.flags & TypeFlags.String, `Expected string constraint, got flags ${constraint.flags}`);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getTypeArguments", () => {
    test("returns type arguments of a generic instantiation", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `declare arr: Array<number>;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `declare arr: Array<number>;`;
            const pos = src.indexOf("arr:");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const typeArgs = project.checker.getTypeArguments(type as TypeReference);
            assert.ok(typeArgs.length > 0, "Should have type arguments");
            assert.ok(typeArgs[0].flags & TypeFlags.Number, `Expected number type argument, got flags ${typeArgs[0].flags}`);
        }
        finally {
            api.close();
        }
    });

    test("a wrongly-typed call throws on the client without taking down the server", () => {
        const src = `local s: string = ""; local arr: Array<number> = [1];`;
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;

            // `string` is not a type reference. When getTypeArguments is reached
            // with one, the server panics, but the per-request panic recovery
            // converts that into an error response rather than crashing the process.
            const sSymbol = project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("s:"));
            assert.ok(sSymbol);
            const sType = project.checker.getTypeOfSymbol(sSymbol);
            assert.ok(sType);
            assert.throws(() => project.checker.getTypeArguments(sType as unknown as TypeReference));

            // The server survived: a subsequent valid request still succeeds.
            const arrSymbol = project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("arr:"));
            assert.ok(arrSymbol);
            const arrType = project.checker.getTypeOfSymbol(arrSymbol);
            assert.ok(arrType);
            const typeArgs = project.checker.getTypeArguments(arrType as TypeReference);
            assert.ok(typeArgs.length > 0, "Server should still serve valid requests");
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getBaseConstraintOfType", () => {
    test("returns the base constraint of a type parameter", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `function identity<T extends string>(x: T): T return x; end`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `function identity<T extends string>(x: T): T return x; end`;
            const pos = src.indexOf("identity<");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const sigs = project.checker.getSignaturesOfType(type, SignatureKind.Call);
            assert.ok(sigs.length > 0);
            const typeParams = sigs[0].getTypeParameters();
            const constraint = project.checker.getBaseConstraintOfType(typeParams[0]);
            assert.ok(constraint, "Should resolve a base constraint");
            assert.ok(constraint.flags & TypeFlags.String, `Expected string constraint, got flags ${constraint.flags}`);
        }
        finally {
            api.close();
        }
    });

    test("returns undefined for a non-instantiable type", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `local x: number = 1;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = `local x: number = 1;`.indexOf("x:");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const constraint = project.checker.getBaseConstraintOfType(type);
            assert.equal(constraint, undefined);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getPropertyOfType", () => {
    test("returns a named property symbol of a type", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `
interface Person {
    name: string;
    age: number;
}
declare p: Person;
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const src = `\ninterface Person {\n    name: string;\n    age: number;\n}\ndeclare p: Person;\n`;
            const pos = src.indexOf("p: Person");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const nameProp = project.checker.getPropertyOfType(type, "name");
            assert.ok(nameProp, "Should find 'name' property");
            assert.equal(nameProp.name, "name");
            const missing = project.checker.getPropertyOfType(type, "doesNotExist");
            assert.equal(missing, undefined);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getSignatureFromDeclaration", () => {
    test("returns the signature of a function declaration", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `function add(a: number, b: number): number return a + b; end`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);
            let funcDecl: Node | undefined;
            sourceFile.forEachChild(function visit(node) {
                if (isFunctionDeclaration(node)) funcDecl = node;
                node.forEachChild(visit);
            });
            assert.ok(funcDecl, "Should find the function declaration");
            const sig = project.checker.getSignatureFromDeclaration(funcDecl);
            assert.ok(sig, "Should resolve a signature");
            assert.equal(sig.parameters.length, 2);
            const returnType = project.checker.getReturnTypeOfSignature(sig);
            assert.ok(returnType);
            assert.ok(returnType.flags & TypeFlags.Number);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getAliasedSymbol", () => {
    test("resolves the top-level return alias to its target symbol", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/foo.tlua": `local foo = 42;\nreturn { foo = foo };`,
            "/src/main.tlua": `local m = require("src.foo");\nreturn m;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);
            const moduleSymbol = project.checker.getSymbolAtLocation(sourceFile);
            assert.ok(moduleSymbol);
            // `return m;` binds the export= symbol as an alias to the local m.
            const aliasSymbol = [...(moduleSymbol.getExports()).values()].find(e => e.name === "export=");
            assert.ok(aliasSymbol);
            assert.ok(aliasSymbol.flags & SymbolFlags.Alias, "Top-level return of a name should be an alias");
            const aliased = project.checker.getAliasedSymbol(aliasSymbol);
            assert.ok(aliased, "Should resolve the aliased symbol");
            assert.equal(aliased.name, "m");
            assert.ok(!(aliased.flags & SymbolFlags.Alias), "Target should not be an alias");
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getExportsOfModule", () => {
    test("returns the export= symbol of a Lua module", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/inner.tlua": `local innerValue = 1;\nreturn { innerValue = innerValue };`,
            "/src/index.tlua": `
local direct = 1;
local inner = require("src.inner");
return { direct = direct, innerValue = inner.innerValue };
`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/index.tlua");
            assert.ok(sourceFile);
            const moduleSymbol = project.checker.getSymbolAtLocation(sourceFile);
            assert.ok(moduleSymbol, "Source file should have a module symbol");
            // getExportsOfModule resolves through export=; a return table has
            // no named exports, so the resolved list is empty.
            const exports = project.checker.getExportsOfModule(moduleSymbol);
            assert.deepEqual(exports, []);
            // The export= symbol itself is visible on the module symbol.
            const raw = moduleSymbol.getExports();
            assert.ok([...raw.values()].some(e => e.name === "export="));
        }
        finally {
            api.close();
        }
    });
});

describe("Symbol - getDocumentationComment and getJsDocTags", () => {
    const docFiles = {
        "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
        "/src/main.tlua": `
/**
 * Adds two numbers together.
 * @param a the first number
 * @returns the sum
 */
function add(a: number, b: number): number return a + b; end
`,
    };

    test("getDocumentationComment returns the leading comment text", () => {
        const api = spawnAPI(docFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = docFiles["/src/main.tlua"].indexOf("add(a");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const doc = symbol.getDocumentationComment(project.checker);
            assert.ok(doc.includes("Adds two numbers together"), `Expected documentation, got: ${doc}`);
            assert.ok(!doc.includes("@param"), "Documentation comment should not include tags");
        }
        finally {
            api.close();
        }
    });

    test("getJsDocTags returns structured tag name/text pairs", () => {
        const api = spawnAPI(docFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = docFiles["/src/main.tlua"].indexOf("add(a");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const tags = symbol.getJsDocTags(project.checker);
            const param = tags.find(t => t.name === "param");
            assert.ok(param, `Expected a @param tag, got: ${JSON.stringify(tags)}`);
            assert.equal(param.text, "a the first number");
            const returns = tags.find(t => t.name === "returns");
            assert.ok(returns, `Expected a @returns tag, got: ${JSON.stringify(tags)}`);
            assert.equal(returns.text, "the sum");
        }
        finally {
            api.close();
        }
    });
});

describe("TypeParameter - isSelfType", () => {
    test("isSelfType identifies an interface's polymorphic self type", () => {
        const src = `\ninterface Fluent {\n    clone(): self;\n}\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = src.indexOf("self");
            const type = project.checker.getTypeAtPosition("/src/main.tlua", pos);
            assert.ok(type, "Expected a type at the 'self' position");
            assert.ok(type.flags & TypeFlags.TypeParameter, "Expected TypeParameter");
            const typeParam = type as TypeParameter;
            assert.equal(typeParam.isSelfType, true, "Expected the hidden interface receiver type");
        }
        finally {
            api.close();
        }
    });

    test("isSelfType is absent for a regular generic type parameter", () => {
        const src = `\nfunction identity<T>(x: T): T return x; end\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            // Point to 'T' in the type parameter declaration '<T>' — getTypeAtPosition
            // on a type annotation reference doesn't resolve to TypeParameter, but
            // the declaration position does.
            const pos = src.indexOf("<T>") + 1;
            const type = project.checker.getTypeAtPosition("/src/main.tlua", pos);
            assert.ok(type, "Expected a type at the 'T' position");
            assert.ok(type.flags & TypeFlags.TypeParameter, "Expected TypeParameter");
            const typeParam = type as TypeParameter;
            assert.ok(!typeParam.isSelfType, "Expected isSelfType to be absent/false for a regular type parameter");
        }
        finally {
            api.close();
        }
    });
});

describe("Type - getAliasTypeArguments", () => {
    test("returns the type arguments of a single-param generic type alias", () => {
        const src = `\ntype Box<T> = { value: T };\nlocal x: Box<string> = { value = "hi" };\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = src.indexOf("x:");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const aliasArgs = type.getAliasTypeArguments();
            assert.equal(aliasArgs.length, 1, "Expected 1 alias type argument");
            assert.ok(aliasArgs[0].flags & TypeFlags.String, `Expected string, got flags ${aliasArgs[0].flags}`);
        }
        finally {
            api.close();
        }
    });

    test("returns multiple type arguments for a multi-param generic type alias", () => {
        const src = `\ntype Pair<A, B> = { first: A; second: B };\nlocal p: Pair<string, number> = { first = "hello", second = 42 };\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = src.indexOf("p:");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const aliasArgs = type.getAliasTypeArguments();
            assert.equal(aliasArgs.length, 2, "Expected 2 alias type arguments");
            assert.ok(aliasArgs[0].flags & TypeFlags.String, `Expected first arg to be string, got flags ${aliasArgs[0].flags}`);
            assert.ok(aliasArgs[1].flags & TypeFlags.Number, `Expected second arg to be number, got flags ${aliasArgs[1].flags}`);
        }
        finally {
            api.close();
        }
    });

    test("returns empty array for a non-alias generic type", () => {
        const src = `\nlocal arr: Array<string> = ["hello"];\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = src.indexOf("arr:");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const aliasArgs = type.getAliasTypeArguments();
            assert.equal(aliasArgs.length, 0, "Expected no alias type arguments for a direct generic reference");
        }
        finally {
            api.close();
        }
    });
});

describe("Type - getAliasSymbol", () => {
    test("returns the symbol for a non-generic type alias", () => {
        // Object-type aliases preserve aliasSymbol; primitive aliases (type Foo = string) do not.
        const src = `\ntype Point = { x: number; y: number };\nlocal p: Point = { x = 1, y = 2 };\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = src.indexOf("p:");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const aliasSymbol = type.getAliasSymbol();
            assert.ok(aliasSymbol, "Expected alias symbol to exist");
            assert.equal(aliasSymbol.name, "Point");
        }
        finally {
            api.close();
        }
    });

    test("returns the symbol for a generic type alias", () => {
        const src = `\ntype Container<T> = { item: T };\nlocal c: Container<number> = { item = 42 };\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = src.indexOf("c:");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const aliasSymbol = type.getAliasSymbol();
            assert.ok(aliasSymbol, "Expected alias symbol for generic alias");
            assert.equal(aliasSymbol.name, "Container");
        }
        finally {
            api.close();
        }
    });

    test("returns undefined for a non-alias type", () => {
        const src = `\nlocal str: string = "test";\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = src.indexOf("str:");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const aliasSymbol = type.getAliasSymbol();
            assert.equal(aliasSymbol, undefined, "Expected no alias symbol for primitive type");
        }
        finally {
            api.close();
        }
    });
});

describe("IntrinsicType - intrinsicName", () => {
    test("intrinsicName matches the primitive type name", () => {
        const src = `\nlocal x: string = "hello";\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const stringType = project.checker.getStringType();
            assert.equal((stringType as IntrinsicType).intrinsicName, "string");
            const anyType = project.checker.getAnyType();
            assert.equal((anyType as IntrinsicType).intrinsicName, "any");
            const neverType = project.checker.getNeverType();
            assert.equal((neverType as IntrinsicType).intrinsicName, "never");
            const pos = src.indexOf("x:");
            const sym = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(sym);
            const litType = project.checker.getTypeOfSymbol(sym);
            assert.ok(litType);
            assert.ok(litType.flags & TypeFlags.Intrinsic);
            assert.equal((litType as IntrinsicType).intrinsicName, "string");
        }
        finally {
            api.close();
        }
    });
});

describe("FreshableType - getFreshType and getRegularType", () => {
    test("LiteralType.value is empty string for the empty-string literal type", () => {
        const src = `\nlocal empty: "" = "";\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", src.indexOf("empty:"));
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.StringLiteral, "Expected StringLiteral");
            const literal = type as LiteralType;
            assert.equal(literal.value, "", "value should be empty string, not undefined");
        }
        finally {
            api.close();
        }
    });

    test("LiteralType.value is accessible via the FreshableType hierarchy", () => {
        const src = `\nlocal greeting: "hello" = "hello";\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = src.indexOf("greeting:");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.StringLiteral, "Expected StringLiteral");
            const literal = type as LiteralType;
            assert.equal(literal.value, "hello");
        }
        finally {
            api.close();
        }
    });

    test("getFreshType() returns a fresh twin with matching value", () => {
        const src = `\nlocal greeting: "hello" = "hello";\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = src.indexOf("greeting:");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.Freshable, "Type should be a freshable type");
            const fresh = (type as FreshableType).getFreshType();
            assert.ok(fresh, "Expected getFreshType() to return non-undefined for a literal type");
            assert.ok(fresh.flags & TypeFlags.StringLiteral, "Fresh type should be a StringLiteral");
            assert.equal((fresh as LiteralType).value, "hello", "Fresh type should carry the same value");
            assert.notEqual(fresh.id, type.id, "Fresh type should not be the original type");
            const freshFresh = fresh.getFreshType();
            assert.ok(freshFresh, "Expected getFreshType() to return non-undefined for a fresh type");
            assert.equal(freshFresh.id, fresh.id, "Fresh type of a fresh type should be the fresh type");
        }
        finally {
            api.close();
        }
    });

    test("getRegularType() on a fresh literal returns the regular twin", () => {
        // The initial type response from getTypeOfSymbol does not always include the
        // regularType handle, so getRegularType() is tested via the fresh twin which
        // always includes its regularType in its own response.
        const src = `\nlocal greeting: "hello" = "hello";\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = src.indexOf("greeting:");
            const symbol = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(symbol);
            const type = project.checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            assert.ok(type.flags & TypeFlags.Freshable, "Type should be a freshable type");
            const fresh = (type as FreshableType).getFreshType();
            assert.ok(fresh, "Need fresh type for this test");
            assert.ok(fresh.flags & TypeFlags.Freshable, "Fresh type should be a freshable type");
            const regular = fresh.getRegularType();
            assert.ok(regular, "Expected getRegularType() on the fresh twin to return non-undefined");
            assert.ok(regular.flags & TypeFlags.StringLiteral, "Regular type should be a StringLiteral");
            assert.equal((regular as LiteralType).value, "hello", "Regular type should carry the same value");
            assert.equal(regular.id, type.id, "Regular type should be the original type");
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - isContextSensitive", () => {
    test("arrow function with no type annotation is context sensitive", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `local fn = function(x) return x end;\nreturn { fn = fn };`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);
            // Find the function expression node
            let functionExpr: import("@tlua/cli/unstable/ast").Node | undefined;
            sourceFile.forEachChild(function visit(node) {
                if (node.kind === SyntaxKind.FunctionExpression) {
                    functionExpr = node;
                }
                node.forEachChild(visit);
            });
            assert.ok(functionExpr, "Should find a function expression");
            const result = project.checker.isContextSensitive(functionExpr);
            assert.equal(result, true);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - isTypeAssignableTo", () => {
    test("returns true when source is assignable to target", () => {
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": `local nothing = 1;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const stringType = project.checker.getStringType();
            const anyType = project.checker.getAnyType();
            const neverType = project.checker.getNeverType();
            assert.ok(project.checker.isTypeAssignableTo(stringType, stringType), "string assignable to string");
            assert.ok(project.checker.isTypeAssignableTo(stringType, anyType), "string assignable to any");
            assert.ok(project.checker.isTypeAssignableTo(neverType, stringType), "never assignable to string (bottom type)");
        }
        finally {
            api.close();
        }
    });

    test("returns false when source is not assignable to target", () => {
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": `local nothing = 1;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const stringType = project.checker.getStringType();
            const numberType = project.checker.getNumberType();
            assert.ok(!project.checker.isTypeAssignableTo(numberType, stringType), "number not assignable to string");
            assert.ok(!project.checker.isTypeAssignableTo(stringType, numberType), "string not assignable to number");
        }
        finally {
            api.close();
        }
    });

    test("a string literal type is assignable to string but not number", () => {
        const src = `\nlocal x: "hello" = "hello";\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = src.indexOf("x:");
            const sym = project.checker.getSymbolAtPosition("/src/main.tlua", pos);
            assert.ok(sym);
            const litType = project.checker.getTypeOfSymbol(sym);
            assert.ok(litType);
            assert.ok(litType.flags & TypeFlags.StringLiteral);
            const stringType = project.checker.getStringType();
            const numberType = project.checker.getNumberType();
            assert.ok(project.checker.isTypeAssignableTo(litType, stringType));
            assert.ok(!project.checker.isTypeAssignableTo(litType, numberType));
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getCompletionsAtPosition", () => {
    test("returns member completions after a dot", () => {
        const src = `\nlocal obj = { name = "hello", age = 42 };\nobj.\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            // Position right after "obj." — member completion trigger
            const pos = src.indexOf("obj.") + "obj.".length;
            const completions = project.checker.getCompletionsAtPosition("/src/main.tlua", pos, { triggerCharacter: "." });
            assert.ok(completions, "Expected completions to be returned");
            assert.ok(completions.entries.length > 0, "Expected at least one completion entry");
            assert.ok(completions.entries.some(e => e.name === "name"), "Expected 'name' property in completions");
            assert.ok(completions.entries.some(e => e.name === "age"), "Expected 'age' property in completions");
            assert.ok(completions.entries.every(e => e.symbol === undefined), "Expected no symbol information");
        }
        finally {
            api.close();
        }
    });

    test("completion entries include sortText", () => {
        const src = `\nlocal obj = { value = 1 };\nobj.\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = src.indexOf("obj.") + "obj.".length;
            const completions = project.checker.getCompletionsAtPosition("/src/main.tlua", pos, { triggerCharacter: "." });
            assert.ok(completions);
            assert.ok(completions.entries.length > 0);
            assert.ok(completions.entries.some(e => e.sortText !== undefined), "Expected sortText on all entries");
        }
        finally {
            api.close();
        }
    });

    test("returns undefined for a non-existent file", () => {
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": `local nothing = 1;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const completions = project.checker.getCompletionsAtPosition("/src/does-not-exist.tlua", 0);
            assert.equal(completions, undefined, "Expected undefined for non-existent file");
        }
        finally {
            api.close();
        }
    });

    test("includeSymbol: true populates symbol on property completions", () => {
        const src = `\nlocal obj = { name = "hello", age = 42 };\nobj.\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/main.tlua": src,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const pos = src.indexOf("obj.") + "obj.".length;
            const completions = project.checker.getCompletionsAtPosition("/src/main.tlua", pos, { triggerCharacter: ".", includeSymbol: true });
            assert.ok(completions, "Expected completions");
            const nameEntry = completions.entries.find(e => e.name === "name");
            assert.ok(nameEntry, "Expected 'name' entry");
            assert.ok(nameEntry.symbol, "Expected symbol to be set on 'name' entry when includeSymbol: true");
            assert.equal(nameEntry.symbol.name, "name", "Symbol name should match completion name");
        }
        finally {
            api.close();
        }
    });
});

describe("Emitter - printNode", () => {
    const emitterFiles = {
        "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
        "/src/main.tlua": `
local x = 42;
function greet(name: string): string return name; end
type Pair = [string, number];
local obj = { m = 1, s = "hi", b = true };
`,
    };

    test("printNode with factory-created keyword type", () => {
        const api = spawnAPI(emitterFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const node = createKeywordTypeNode(SyntaxKind.StringKeyword);
            const text = project.emitter.printNode(node);
            assert.strictEqual(text, "string");
        }
        finally {
            api.close();
        }
    });

    test("printNode with factory-created union type", () => {
        const api = spawnAPI(emitterFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const node = createUnionTypeNode([
                createKeywordTypeNode(SyntaxKind.StringKeyword),
                createKeywordTypeNode(SyntaxKind.NumberKeyword),
            ]);
            const text = project.emitter.printNode(node);
            assert.strictEqual(text, "string | number");
        }
        finally {
            api.close();
        }
    });

    test("printNode with factory-created function type", () => {
        const api = spawnAPI(emitterFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const param = createParameterDeclaration(
                undefined,
                undefined,
                createIdentifier("x"),
                undefined,
                createKeywordTypeNode(SyntaxKind.StringKeyword),
                undefined,
            );
            const node = createFunctionTypeNode(
                undefined,
                undefined,
                [param],
                createKeywordTypeNode(SyntaxKind.NumberKeyword),
            );
            const text = project.emitter.printNode(node);
            assert.strictEqual(text, "(x: string) => number");
        }
        finally {
            api.close();
        }
    });

    test("printNode with factory-created type reference", () => {
        const api = spawnAPI(emitterFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const node = createTypeReferenceNode(createIdentifier("Array"), [
                createKeywordTypeNode(SyntaxKind.StringKeyword),
            ]);
            const text = project.emitter.printNode(node);
            assert.strictEqual(text, "Array<string>");
        }
        finally {
            api.close();
        }
    });

    test("printNode with factory-created array type", () => {
        const api = spawnAPI(emitterFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const node = createArrayTypeNode(createKeywordTypeNode(SyntaxKind.NumberKeyword));
            const text = project.emitter.printNode(node);
            assert.strictEqual(text, "number[]");
        }
        finally {
            api.close();
        }
    });

    test("typeToTypeNode + printNode round-trip", () => {
        const api = spawnAPI(emitterFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const { checker, emitter } = snapshot.getProject("/tluaconfig.json")!;
            const src = emitterFiles["/src/main.tlua"];

            const greetPos = src.indexOf("greet(");
            const symbol = checker.getSymbolAtPosition("/src/main.tlua", greetPos);
            assert.ok(symbol);
            const type = checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const typeNode = checker.typeToTypeNode(type);
            assert.ok(typeNode);
            const text = emitter.printNode(typeNode);
            assert.ok(text);
            assert.strictEqual(text, "(name: string) => string");
        }
        finally {
            api.close();
        }
    });

    test("visitEachChild on typeToTypeNode result with keyword types", () => {
        const api = spawnAPI(emitterFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const { checker } = snapshot.getProject("/tluaconfig.json")!;
            const src = emitterFiles["/src/main.tlua"];
            const objPos = src.indexOf("obj");
            const symbol = checker.getSymbolAtPosition("/src/main.tlua", objPos);
            assert.ok(symbol, "should find symbol for obj");
            const type = checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const typeNode = checker.typeToTypeNode(type);
            assert.ok(typeNode, "typeToTypeNode should return a type node");

            // Recursively visit to reach PropertySignature.type where isTypeNode is checked.
            const visited = (function visit(node: Node): Node {
                return visitEachChild(node, visit);
            })(typeNode);
            assert.ok(visited, "visitEachChild should not throw");

            const kinds = [
                SyntaxKind.NumberKeyword,
                SyntaxKind.StringKeyword,
                SyntaxKind.BooleanKeyword,
                SyntaxKind.AnyKeyword,
                SyntaxKind.VoidKeyword,
                SyntaxKind.NilKeyword,
                SyntaxKind.NeverKeyword,
                SyntaxKind.UnknownKeyword,
                SyntaxKind.ObjectKeyword,
                SyntaxKind.SymbolKeyword,
                SyntaxKind.SelfKeyword,
                SyntaxKind.IntrinsicKeyword,
                SyntaxKind.ExpressionWithTypeArguments,
                SyntaxKind.JSDocAllType,
                SyntaxKind.JSDocNullableType,
                SyntaxKind.JSDocNonNullableType,
                SyntaxKind.JSDocOptionalType,
                SyntaxKind.JSDocVariadicType,
                SyntaxKind.JSDocTypeExpression,
                SyntaxKind.JSDocTypeLiteral,
                SyntaxKind.JSDocSignature,
            ];
            for (const kind of kinds) {
                assert.ok(isTypeNode({ kind } as any), `isTypeNode should accept ${SyntaxKind[kind]}`);
            }
        }
        finally {
            api.close();
        }
    });

    test("typeToString", () => {
        const api = spawnAPI(emitterFiles);
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const { checker } = snapshot.getProject("/tluaconfig.json")!;
            const src = emitterFiles["/src/main.tlua"];

            const greetPos = src.indexOf("greet(");
            const symbol = checker.getSymbolAtPosition("/src/main.tlua", greetPos);
            assert.ok(symbol);
            const type = checker.getTypeOfSymbol(symbol);
            assert.ok(type);
            const text = checker.typeToString(type);
            assert.strictEqual(text, "(name: string) => string");
        }
        finally {
            api.close();
        }
    });

    test("printNode with terminateUnterminatedLiterals option", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/main.tlua": `local foo = /asdfasf;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/main.tlua");
            assert.ok(sourceFile);

            // Find the regex literal node
            let regexNode: import("@tlua/cli/unstable/ast").Node | undefined;
            sourceFile.forEachChild(function visit(node) {
                if (node.kind === SyntaxKind.RegularExpressionLiteral) {
                    regexNode = node;
                    return;
                }
                node.forEachChild(visit);
            });
            assert.ok(regexNode, "Should find a regex literal");

            // Without the option, regex is printed as-is
            const textWithout = project.emitter.printNode(regexNode);
            assert.strictEqual(textWithout, "/asdfasf");

            // With the option, the closing slash is added
            const textWith = project.emitter.printNode(regexNode, { terminateUnterminatedLiterals: true });
            assert.strictEqual(textWith, "/asdfasf/");
        }
        finally {
            api.close();
        }
    });
});

describe("modifierFlags", () => {
    test("async function has Async flag", () => {
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/index.tlua": `async function foo() end`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/index.tlua");
            assert.ok(sourceFile);

            let fnNode: import("@tlua/cli/unstable/ast").FunctionDeclaration | undefined;
            sourceFile.forEachChild(function visit(node) {
                if (isFunctionDeclaration(node)) {
                    fnNode = node;
                }
                node.forEachChild(visit);
            });
            assert.ok(fnNode, "Should find a function declaration");
            assert.ok(fnNode.modifierFlags & ModifierFlags.Async, "Should have Async flag");
            assert.strictEqual(fnNode.modifierFlags, ModifierFlags.Async);
        }
        finally {
            api.close();
        }
    });

    test("node without modifiers has ModifierFlags.None", () => {
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/index.tlua": `function bar() end`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/index.tlua");
            assert.ok(sourceFile);

            let fnNode: import("@tlua/cli/unstable/ast").FunctionDeclaration | undefined;
            sourceFile.forEachChild(function visit(node) {
                if (isFunctionDeclaration(node)) {
                    fnNode = node;
                }
                node.forEachChild(visit);
            });
            assert.ok(fnNode, "Should find a function declaration");
            assert.strictEqual(fnNode.modifierFlags, ModifierFlags.None);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getResolvedSymbol", () => {
    test("resolves variable reference to its declaration symbol", () => {
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/index.tlua": `local x = 1;\nlocal y = x;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/index.tlua");
            assert.ok(sourceFile);

            // Find the 'x' identifier in `local y = x`
            let refNode: import("@tlua/cli/unstable/ast").Identifier | undefined;
            sourceFile.forEachChild(function visit(node) {
                if (isIdentifier(node) && node.text === "x") {
                    // We want the reference, not the declaration - take the last one
                    refNode = node;
                }
                node.forEachChild(visit);
            });
            assert.ok(refNode, "Should find identifier 'x'");

            const symbol = project.checker.getResolvedSymbol(refNode);
            assert.ok(symbol, "Should resolve symbol for 'x'");
            assert.equal(symbol.name, "x");
        }
        finally {
            api.close();
        }
    });
});

describe("VariableDeclarationList - BlockScoped flags", () => {
    test("local declaration has LuaLocal flag", () => {
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/index.tlua": `local x = 1;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/index.tlua");
            assert.ok(sourceFile);

            let declList: import("@tlua/cli/unstable/ast").Node | undefined;
            sourceFile.forEachChild(function visit(node) {
                if (isVariableDeclarationList(node)) {
                    declList = node;
                }
                node.forEachChild(visit);
            });
            assert.ok(declList, "Should find VariableDeclarationList");
            assert.ok(declList.flags & NodeFlags.LuaLocal, "Should have LuaLocal flag");
        }
        finally {
            api.close();
        }
    });

    test("local declaration is not Let or Const flagged", () => {
        // tlua deleted the let/const surface syntax; a `local` list is marked
        // with LuaLocal (block-scoped via the Lua-local binder path), not with
        // the Let/Const scope bits that the old let/const declarations carried.
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/index.tlua": `local x = 1;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/index.tlua");
            assert.ok(sourceFile);

            let declList: import("@tlua/cli/unstable/ast").Node | undefined;
            sourceFile.forEachChild(function visit(node) {
                if (isVariableDeclarationList(node)) {
                    declList = node;
                }
                node.forEachChild(visit);
            });
            assert.ok(declList, "Should find VariableDeclarationList");
            assert.ok(declList.flags & NodeFlags.LuaLocal, "Should have LuaLocal flag");
            assert.equal(declList.flags & NodeFlags.BlockScoped, 0, "Should not have Let/Const flags");
        }
        finally {
            api.close();
        }
    });
});

test("TypeOperator operator kind", () => {
    const api = spawnAPI({
        "/tluaconfig.json": "{}",
        "/src/index.tlua": `function test(arg: readonly number[])\nend\n`,
    });
    try {
        const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
        const project = snapshot.getProject("/tluaconfig.json")!;
        const sourceFile = project.program.getSourceFile("/src/index.tlua");
        assert(sourceFile);
        const param = (sourceFile.statements[0] as import("@tlua/cli/unstable/ast").FunctionDeclaration).parameters[0];
        assert(param);
        const type = param.type as import("@tlua/cli/unstable/ast").TypeOperatorNode;
        assert(type);
        assert.equal(type.kind, SyntaxKind.TypeOperator);
        assert.equal(type.operator, SyntaxKind.ReadonlyKeyword);
        const printed = project.emitter.printNode(sourceFile);
        assert.equal(sourceFile.text, printed);
    }
    finally {
        api.close();
    }
});

test("VarargExpression roundtrip", () => {
    const api = spawnAPI({
        "/tluaconfig.json": "{}",
        "/src/index.tlua": `local thing = { ... };\n`,
    });
    try {
        const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
        const project = snapshot.getProject("/tluaconfig.json")!;
        const sourceFile = project.program.getSourceFile("/src/index.tlua");
        assert(sourceFile);
        const stmt = sourceFile.statements[0] as import("@tlua/cli/unstable/ast").VariableStatement;
        const object = stmt.declarationList.declarations[0].initializer as import("@tlua/cli/unstable/ast").ObjectLiteralExpression;
        const entry = object.properties[0] as import("@tlua/cli/unstable/ast").TableEntry;
        assert(entry);
        assert.equal(entry.kind, SyntaxKind.TableEntry);
        const expr = entry.expression;
        assert(expr);
        assert.equal(expr.kind, SyntaxKind.VarargExpression);
        const printed = project.emitter.printNode(sourceFile);
        assert.equal(sourceFile.text, printed);
    }
    finally {
        api.close();
    }
});

test("VariableDeclarationList local flag clone", () => {
    const api = spawnAPI({
        "/tluaconfig.json": "{}",
        "/src/index.tlua": `local thing = 123;\n`,
    });
    try {
        const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
        const project = snapshot.getProject("/tluaconfig.json")!;
        const sourceFile = project.program.getSourceFile("/src/index.tlua");
        assert(sourceFile);
        {
            const stmt = sourceFile.statements[0] as import("@tlua/cli/unstable/ast").VariableStatement;
            const list = stmt.declarationList;
            assert(list.flags & NodeFlags.LuaLocal);
        }
        const cloned = getSynthesizedDeepClone(sourceFile);
        {
            const stmt = cloned.statements[0] as import("@tlua/cli/unstable/ast").VariableStatement;
            const list = stmt.declarationList;
            assert(list.flags & NodeFlags.LuaLocal);
        }
        const printed = project.emitter.printNode(cloned);
        assert.equal(sourceFile.text, printed);
    }
    finally {
        api.close();
    }
});

test("JSDoc before ExpressionStatement allowed", () => {
    const api = spawnAPI({
        "/tluaconfig.json": "{}",
        "/src/index.tlua": `
/**
 * A doc.
 */
doThing();
        `,
    });
    try {
        const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
        const project = snapshot.getProject("/tluaconfig.json")!;
        const sourceFile = project.program.getSourceFile("/src/index.tlua");
        assert(sourceFile);
        const printed = project.emitter.printNode(sourceFile);
        // tlua prints comments in Lua style only, so the JSDoc block normalizes to a
        // `--[[ ]]` long comment rather than round-tripping back to `/** */`.
        assert.equal(printed.trim(), "--[[*\n * A doc.\n ]]\ndoThing();");
    }
    finally {
        api.close();
    }
});

test("Factory ModifierList auto-conversion", () => {
    const api = spawnAPI();
    try {
        const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
        const project = snapshot.getProject("/tluaconfig.json")!;
        const node = createTypeAliasDeclaration(
            [createToken(SyntaxKind.ExportKeyword)],
            createIdentifier("Test"),
            undefined,
            createKeywordTypeNode(SyntaxKind.AnyKeyword),
        );

        assert.equal(project.emitter.printNode(node), "export type Test = any;");

        const cloned = getSynthesizedDeepClone(node);
        assert.equal(project.emitter.printNode(cloned), "export type Test = any;");
    }
    finally {
        api.close();
    }
});

test("Parse-clone-emit roundtrip", () => {
    const tsSource = fileURLToPath(new URL("../../../../_submodules/TypeScript/src", import.meta.url).toString());
    const api = new API({
        cwd: tsSource,
    });
    const target = {
        cloneCrashed: 0,
        printCrashed: 0,
        clonePrintCrashed: 0,
    };
    const errors = { ...target };
    try {
        for (const tsconfig of globSync("**/tsconfig.json", { cwd: tsSource })) {
            const snapshot = api.updateSnapshot({ openProject: resolve(tsSource, tsconfig) });
            const project = snapshot.getProject(tsconfig);
            assert(project);
            for (const file of project.rootFiles) {
                const source = project.program.getSourceFile(file);
                assert(source);
                // These TypeScript-only forms are not representable by tlua's AST.
                if (/\benum\s+[$A-Za-z_]|\bundefined\b/.test(source.text)) continue;
                let clone: typeof source;

                try {
                    project.emitter.printNode(source);
                }
                catch {
                    errors.printCrashed++;
                    continue;
                }

                try {
                    clone = getSynthesizedDeepClone(source);
                }
                catch {
                    errors.cloneCrashed++;
                    continue;
                }

                try {
                    project.emitter.printNode(clone);
                }
                catch {
                    errors.clonePrintCrashed++;
                    continue;
                }
            }
        }
    }
    finally {
        api.close();
    }
    assert.deepEqual(errors, target);
});

describe("Program - diagnostics", () => {
    test("getSyntacticDiagnostics", () => {
        const source = `local x: = 1;`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/index.tlua": source,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const diags = project.program.getSyntacticDiagnostics("/src/index.tlua");
            assert.deepEqual(diags, [{
                fileName: "/src/index.tlua",
                ...rangeOf(source, "="),
                code: 1110,
                category: DiagnosticCategory.Error,
                text: "Type expected.",
            }]);
        }
        finally {
            api.close();
        }
    });

    test("getSemanticDiagnostics with messageChain and relatedInformation", () => {
        const source = `interface Props { callback: (x: string) => void }\nlocal p: Props = { callback = function(x: number) end };`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/index.tlua": source,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const diags = project.program.getSemanticDiagnostics("/src/index.tlua");
            const declRange = rangeOf(source, "callback", 0);
            const assignRange = rangeOf(source, "callback", 1);
            assert.deepEqual(diags, [{
                fileName: "/src/index.tlua",
                ...assignRange,
                code: 2322,
                category: DiagnosticCategory.Error,
                text: "Type '(x: number) => void' is not assignable to type '(x: string) => void'.",
                messageChain: [{
                    fileName: "/src/index.tlua",
                    ...assignRange,
                    code: 2328,
                    category: DiagnosticCategory.Error,
                    text: "Types of parameters 'x' and 'x' are incompatible.",
                    messageChain: [{
                        fileName: "/src/index.tlua",
                        ...assignRange,
                        code: 2322,
                        category: DiagnosticCategory.Error,
                        text: "Type 'string' is not assignable to type 'number'.",
                    }],
                }],
                relatedInformation: [{
                    fileName: "/src/index.tlua",
                    ...declRange,
                    code: 6500,
                    category: DiagnosticCategory.Message,
                    text: "The expected type comes from property 'callback' which is declared here on type 'Props'",
                }],
            }]);
        }
        finally {
            api.close();
        }
    });

    test("getSuggestionDiagnostics", () => {
        const source = `function f() local x = 1; return x; end\nlocal _unused = 1;\nreturn { f = f };\n`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/index.tlua": source,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const diags = project.program.getSuggestionDiagnostics("/src/index.tlua");
            assert.deepEqual(diags, [{
                fileName: "/src/index.tlua",
                ...rangeOf(source, "_unused"),
                code: 6133,
                category: DiagnosticCategory.Suggestion,
                text: "'_unused' is declared but its value is never read.",
                reportsUnnecessary: true,
            }]);
        }
        finally {
            api.close();
        }
    });

    test("getConfigFileParsingDiagnostics", () => {
        const config = `{ "compilerOptions": { "target": "invalid" } }`;
        const api = spawnAPI({
            "/tluaconfig.json": config,
            "/src/index.tlua": `local x = 1;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const diags = project.program.getConfigFileParsingDiagnostics();
            assert.deepEqual(diags, [{
                fileName: "/tluaconfig.json",
                ...rangeOf(config, `"invalid"`),
                code: 6046,
                category: DiagnosticCategory.Error,
                text: "Argument for '--target' option must be: 'es6', 'es2015', 'es2016', 'es2017', 'es2018', 'es2019', 'es2020', 'es2021', 'es2022', 'es2023', 'es2024', 'es2025', 'esnext'.",
            }]);
        }
        finally {
            api.close();
        }
    });

    test("getDeclarationDiagnostics", () => {
        const api = spawnAPI({
            "/tluaconfig.json": `{ "compilerOptions": { "declaration": true } }`,
            "/src/index.tlua": `local x: number = 1;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const diags = project.program.getDeclarationDiagnostics("/src/index.tlua");
            // Declaration emit is not supported for a Lua module yet.
            assert.equal(diags.length, 1);
            assert.equal(diags[0].code, 100054);
        }
        finally {
            api.close();
        }
    });

    test("getBindDiagnostics", () => {
        const source = `declare x: number;\ndeclare x: number;`;
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/index.tlua": source,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const diags = project.program.getBindDiagnostics("/src/index.tlua");
            assert.deepEqual(diags, [
                {
                    fileName: "/src/index.tlua",
                    ...rangeOf(source, "x", 0),
                    code: 2451,
                    category: DiagnosticCategory.Error,
                    text: "Cannot redeclare block-scoped variable 'x'.",
                },
                {
                    fileName: "/src/index.tlua",
                    ...rangeOf(source, "x", 1),
                    code: 2451,
                    category: DiagnosticCategory.Error,
                    text: "Cannot redeclare block-scoped variable 'x'.",
                },
            ]);
        }
        finally {
            api.close();
        }
    });

    test("getProgramDiagnostics", () => {
        const config = `{ "compilerOptions": { "exactOptionalPropertyTypes": true, "strictNullChecks": false } }`;
        const api = spawnAPI({
            "/tluaconfig.json": config,
            "/src/index.tlua": `local x = 1;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const diags = project.program.getProgramDiagnostics();
            assert.deepEqual(diags, [
                {
                    fileName: "/tluaconfig.json",
                    ...rangeOf(config, `"exactOptionalPropertyTypes"`),
                    code: 5052,
                    category: DiagnosticCategory.Error,
                    text: "Option 'exactOptionalPropertyTypes' cannot be specified without specifying option 'strictNullChecks'.",
                },
            ]);
        }
        finally {
            api.close();
        }
    });

    test("getGlobalDiagnostics", () => {
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/index.tlua": `local x = 1;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const diags = project.program.getGlobalDiagnostics();
            assert.deepEqual(diags, []);
        }
        finally {
            api.close();
        }
    });

    test("getGlobalDiagnostics returns file-less diagnostics from the checker", () => {
        const api = spawnAPI({
            "/tluaconfig.json": `{ "compilerOptions": { "noLib": true } }`,
            "/src/index.tlua": `local x = [1, 2, 3];\nreturn { x = x };`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const diags = project.program.getGlobalDiagnostics();
            // With noLib, the checker reports "Cannot find global type" diagnostics that
            // are not associated with any source file.
            assert.ok(diags.length > 0, "expected global diagnostics to be reported");
            for (const diag of diags) {
                assert.equal(diag.fileName, undefined);
                assert.equal(diag.code, 2318);
                assert.equal(diag.category, DiagnosticCategory.Error);
            }
            assert.ok(
                diags.some(d => d.text === "Cannot find global type 'Array'."),
                "expected a global diagnostic for the 'Array' type",
            );
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getReferencedSymbolsForNode", () => {
    test("getReferencedSymbolsForNode", () => {
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/index.tlua": `function greet(name: string) return name; end\ngreet("world");`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/index.tlua");
            assert.ok(sourceFile);
            const funcDecl = cast(sourceFile.statements[0], isFunctionDeclaration);
            const funcName = funcDecl.name!;
            const refs = project.checker.getReferencedSymbolsForNode(funcName, funcName.pos);
            assert.ok(refs.length > 0);
            // Each entry should have a definition and references
            const entry = refs[0];
            assert.ok(entry.definition);
            assert.ok(entry.references.length > 0);
        }
        finally {
            api.close();
        }
    });
});

describe("Checker - getSignatureUsage", () => {
    test("getSignatureUsage", () => {
        const api = spawnAPI({
            "/tluaconfig.json": "{}",
            "/src/index.tlua": `function greet(name: string) return name; end\ngreet("world");`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/index.tlua");
            assert.ok(sourceFile);
            const funcDecl = cast(sourceFile.statements[0], isFunctionDeclaration);
            const usages = project.checker.getSignatureUsage(funcDecl);
            assert.ok(usages.length > 0);
            // The call site should have a call expression
            const usage = usages.find(u => u.call !== undefined);
            assert.ok(usage, "Expected at least one usage with a call expression");
        }
        finally {
            api.close();
        }
    });
});

describe("getDefaultProjectForFile", () => {
    test("finds inferred project for d.ts in node_modules after openFiles", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/index.tlua": `local x = 1;`,
            "/node_modules/my-lib/package.json": JSON.stringify({ name: "my-lib", types: "./index.d.ts" }),
            "/node_modules/my-lib/index.d.tlua": `declare foo: string;`,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;

            // The d.ts is not imported, so it is not in the project's program
            const dtsSf = project.program.getSourceFile("/node_modules/my-lib/index.d.tlua");
            assert.equal(dtsSf, undefined, "d.ts not in import graph should not be found via project.program.getSourceFile");

            // Before opening the file, getDefaultProjectForFile returns undefined (no error)
            const noProject = snapshot.getDefaultProjectForFile("/node_modules/my-lib/index.d.tlua");
            assert.equal(noProject, undefined, "getDefaultProjectForFile returns undefined for unloaded file");

            // Load the file into the inferred project via updateSnapshot openFiles
            const snapshot2 = api.updateSnapshot({ openFiles: ["/node_modules/my-lib/index.d.tlua"] });
            const defaultProject = snapshot2.getDefaultProjectForFile("/node_modules/my-lib/index.d.tlua");
            assert.ok(defaultProject, "getDefaultProjectForFile should find inferred project after openFiles");

            const fooPos = `declare foo: string;`.indexOf("foo");
            const fooType = defaultProject.checker.getTypeAtPosition("/node_modules/my-lib/index.d.tlua", fooPos);
            assert.ok(fooType);
            assert.ok(fooType.flags & TypeFlags.String);
        }
        finally {
            api.close();
        }
    });

    test("keeps previously opened files open across subsequent openFiles calls", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/index.tlua": `local x = 1;`,
            "/node_modules/my-lib/package.json": JSON.stringify({ name: "my-lib", types: "./index.d.ts" }),
            "/node_modules/my-lib/index.d.tlua": `declare foo: string;`,
            "/node_modules/other-lib/package.json": JSON.stringify({ name: "other-lib", types: "./index.d.ts" }),
            "/node_modules/other-lib/index.d.tlua": `declare bar: number;`,
        });
        try {
            api.updateSnapshot({ openProject: "/tluaconfig.json" });
            api.updateSnapshot({ openFiles: ["/node_modules/my-lib/index.d.tlua"] });

            // Opening a second file in a later snapshot must not close the first one.
            const snapshot = api.updateSnapshot({ openFiles: ["/node_modules/other-lib/index.d.tlua"] });

            const firstProject = snapshot.getDefaultProjectForFile("/node_modules/my-lib/index.d.tlua");
            assert.ok(firstProject, "previously opened file should remain in the inferred project");
            const secondProject = snapshot.getDefaultProjectForFile("/node_modules/other-lib/index.d.tlua");
            assert.ok(secondProject, "newly opened file should be in the inferred project");
        }
        finally {
            api.close();
        }
    });

    test("opening a file resolves to a configured project via ancestor search", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/index.tlua": `local x = 1;`,
        });
        try {
            // Open the file without first opening the project. Like LSP's didOpen, this
            // should search ancestor directories for a tsconfig that contains the file.
            const snapshot = api.updateSnapshot({ openFiles: ["/src/index.tlua"] });
            const defaultProject = snapshot.getDefaultProjectForFile("/src/index.tlua");
            assert.ok(defaultProject, "should find a project for the opened file");
            assert.equal(
                defaultProject.configFileName,
                "/tluaconfig.json",
                "opened file should resolve to the containing configured project, not the inferred project",
            );
        }
        finally {
            api.close();
        }
    });

    test("closeProjects releases a project opened via openProjects", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/index.tlua": `local x = 1;`,
        });
        try {
            const opened = api.updateSnapshot({ openProjects: ["/tluaconfig.json"] });
            assert.ok(opened.getProject("/tluaconfig.json"), "project should be open after openProjects");

            const closed = api.updateSnapshot({ closeProjects: ["/tluaconfig.json"] });
            assert.equal(
                closed.getProject("/tluaconfig.json"),
                undefined,
                "project should be unloaded after closeProjects",
            );
        }
        finally {
            api.close();
        }
    });

    test("closeFiles releases a file opened via openFiles", () => {
        const api = spawnAPI({
            "/tluaconfig.json": JSON.stringify({ compilerOptions: { strict: true } }),
            "/src/index.tlua": `local x = 1;`,
            "/node_modules/my-lib/package.json": JSON.stringify({ name: "my-lib", types: "./index.d.ts" }),
            "/node_modules/my-lib/index.d.tlua": `declare foo: string;`,
        });
        try {
            api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const opened = api.updateSnapshot({ openFiles: ["/node_modules/my-lib/index.d.tlua"] });
            assert.ok(
                opened.getDefaultProjectForFile("/node_modules/my-lib/index.d.tlua"),
                "file should resolve to a project after openFiles",
            );

            const closed = api.updateSnapshot({ closeFiles: ["/node_modules/my-lib/index.d.tlua"] });
            assert.equal(
                closed.getDefaultProjectForFile("/node_modules/my-lib/index.d.tlua"),
                undefined,
                "file should no longer resolve to a project after closeFiles",
            );
        }
        finally {
            api.close();
        }
    });
});

test("Benchmarks", () => {
    runBenchmarks({ singleIteration: true });
});

describe("Timing", () => {
    test("collects combined client, server, and transport timing info", () => {
        const api = new API({
            cwd: fileURLToPath(new URL("../../../../", import.meta.url).toString()),
            fs: createVirtualFileSystem({ ...defaultFiles }),
            collectTiming: true,
        });
        try {
            // Baseline: enabled, but nothing measured yet.
            let info = api.getTimingInfo();
            assert.equal(info.enabled, true);
            assert.equal(info.totals.requestCount, 0);
            assert.equal(info.recentRequests.length, 0);

            // Exercise a JSON request and a binary source-file request.
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/index.tlua");
            assert.ok(sourceFile);

            // Client-side timing: round-trip latency and byte counts.
            info = api.getTimingInfo();
            assert.ok(info.totals.requestCount >= 2, "expected at least two measured requests");
            assert.ok(info.totals.bytesSent > 0);
            assert.ok(info.totals.bytesReceived > 0);
            assert.ok(info.totals.roundTripMs >= 0);

            // Server-side timing is folded into the same snapshot.
            assert.ok(info.totals.serverTimeMs >= 0);
            assert.ok(info.totals.transportOverheadMs >= 0);
            assert.equal(
                info.totals.transportOverheadMs,
                Math.max(0, info.totals.roundTripMs - info.totals.serverTimeMs),
            );

            // The ring buffer retains at most the 5 most recent requests.
            assert.ok(info.recentRequests.length > 0);
            assert.ok(info.recentRequests.length <= 5);
            assert.ok(info.recentRequests.length <= info.totals.requestCount);

            for (const r of info.recentRequests) {
                assert.ok(r.roundTripMs >= 0);
                assert.ok(r.bytesSent >= 0);
                assert.ok(r.bytesReceived >= 0);
                assert.ok(typeof r.method === "string" && r.method.length > 0);
                assert.equal(typeof r.timestamp, "number");
                // Transport overhead is present exactly when server time is.
                assert.equal(r.transportOverheadMs === undefined, r.serverTimeMs === undefined);
                if (r.serverTimeMs !== undefined) {
                    assert.ok(r.serverTimeMs >= 0);
                    assert.equal(r.transportOverheadMs, Math.max(0, r.roundTripMs - r.serverTimeMs));
                }
            }

            // Server processing time is folded in for the recent requests.
            assert.ok(info.recentRequests.every(r => r.serverTimeMs !== undefined), "server time should be reported");

            // Reset clears totals and history on both client and server.
            api.resetTimingInfo();
            info = api.getTimingInfo();
            assert.equal(info.enabled, true);
            assert.equal(info.totals.requestCount, 0);
            assert.equal(info.totals.serverTimeMs, 0);
            assert.equal(info.totals.transportOverheadMs, 0);
            assert.equal(info.recentRequests.length, 0);
        }
        finally {
            api.close();
        }
    });

    test("tracks source-file node materialization", () => {
        const api = new API({
            cwd: fileURLToPath(new URL("../../../../", import.meta.url).toString()),
            fs: createVirtualFileSystem({ ...defaultFiles }),
            collectTiming: true,
        });
        try {
            const snapshot = api.updateSnapshot({ openProject: "/tluaconfig.json" });
            const project = snapshot.getProject("/tluaconfig.json")!;
            const sourceFile = project.program.getSourceFile("/src/index.tlua");
            assert.ok(sourceFile);

            // The source file node itself is pre-cached, so before walking the
            // tree no descendant nodes have been materialized.
            let info = api.getTimingInfo();
            const before = info.totals.nodesMaterialized;
            assert.equal(info.totals.sourceFilesFetched, 1, "one source file was fetched");
            assert.ok(info.totals.nodesFetched > 0, "the fetched file contributes materializable nodes");
            assert.equal(info.totals.nodesMaterialized, 0, "nothing walked yet, so no nodes materialized");

            // Walk the whole tree to force lazy materialization of every node.
            let visited = 0;
            sourceFile.forEachChild(function visit(node) {
                visited++;
                node.forEachChild(visit);
            });
            assert.ok(visited > 0, "expected to visit at least one node");

            info = api.getTimingInfo();
            assert.ok(
                info.totals.nodesMaterialized > before,
                "walking the tree should materialize nodes",
            );
            assert.ok(
                info.totals.nodesMaterialized >= visited,
                "every visited node should have been materialized",
            );

            // A full walk materializes (nearly) every fetched node, so the share
            // of fetched nodes materialized should be substantial.
            assert.equal(info.totals.sourceFilesFetched, 1);
            assert.ok(
                info.totals.nodesMaterialized > 0
                    && info.totals.nodesMaterialized <= info.totals.nodesFetched,
                "materialized nodes should be in (0, nodesFetched]",
            );

            // Reset clears materialization totals along with the rest.
            api.resetTimingInfo();
            info = api.getTimingInfo();
            assert.equal(info.totals.nodesMaterialized, 0);
            assert.equal(info.totals.sourceFilesFetched, 0);
            assert.equal(info.totals.nodesFetched, 0);
        }
        finally {
            api.close();
        }
    });

    test("is disabled by default", () => {
        const api = spawnAPI();
        try {
            api.parseConfigFile("/tluaconfig.json");
            const info = api.getTimingInfo();
            assert.equal(info.enabled, false);
            assert.equal(info.totals.requestCount, 0);
            assert.equal(info.totals.nodesMaterialized, 0);
            assert.equal(info.totals.sourceFilesFetched, 0);
            assert.equal(info.totals.nodesFetched, 0);
            assert.equal(info.recentRequests.length, 0);
        }
        finally {
            api.close();
        }
    });
});

function spawnAPI(files: Record<string, string> = { ...defaultFiles }) {
    return new API({
        cwd: fileURLToPath(new URL("../../../../", import.meta.url).toString()),
        fs: createVirtualFileSystem(files),
    });
}

function spawnAPIWithFS(files: Record<string, string> = { ...defaultFiles }): { api: API; fs: FileSystem; } {
    const fs = createVirtualFileSystem(files);
    const api = new API({
        cwd: fileURLToPath(new URL("../../../../", import.meta.url).toString()),
        fs,
    });
    return { api, fs };
}

/** Returns `{ pos, end }` for the nth (0-based, default 0) occurrence of `searchString` in `source`. */
function rangeOf(source: string, searchString: string, occurrence: number = 0): { pos: number; end: number; } {
    let index = -1;
    for (let i = 0; i <= occurrence; i++) {
        index = source.indexOf(searchString, index + 1);
        if (index === -1) {
            throw new Error(`Occurrence ${occurrence} of "${searchString}" not found in source`);
        }
    }
    return { pos: index, end: index + searchString.length };
}
