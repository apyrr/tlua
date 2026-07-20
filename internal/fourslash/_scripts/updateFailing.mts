#!/usr/bin/env -S node --experimental-strip-types --no-warnings

// Usage: node --experimental-strip-types --no-warnings updateFailing.mts
//
// CAUTION: this rewrites failingTests.txt WHOLESALE from one suite run, and
// SKIPPED tests count as not-failing. Running it under any skip condition
// (missing embedded libs, @lib-gated tests, a partial -run filter) silently
// un-quarantines tests that never executed, which then fail the suite when
// they become runnable. Prefer regenerating from a full clean run:
//   TLUA_FOURSLASH_IGNORE_FAILING=1 go test ./internal/fourslash/...
// and writing exactly the FAILING test names; single-name additions and
// removals are safest done by hand.

import * as cp from "child_process";
import * as fs from "fs";
import path from "path";
import * as readline from "readline";
import which from "which";

const failingTestsPath = path.join(import.meta.dirname, "failingTests.txt");
const crashingTestsPath = path.join(import.meta.dirname, "crashingTests.txt");
const generatedTestsPath = path.join(import.meta.dirname, "..", "tests", "gen");

interface TestEvent {
    Time?: string;
    Action: string;
    Package?: string;
    Test?: string;
    Output?: string;
    Elapsed?: number;
}

async function main() {
    const go = which.sync("go");
    const existingFailingTests = fs.existsSync(failingTestsPath)
        ? fs.readFileSync(failingTestsPath, "utf-8").split(/\r?\n/).filter(Boolean)
        : [];

    let testProcess: cp.ChildProcess;
    try {
        // Run tests with TLUA_FOURSLASH_IGNORE_FAILING=1 to run all tests including those in failingTests.txt
        testProcess = cp.spawn(go, ["test", "-json", "./internal/fourslash/tests/gen"], {
            stdio: ["ignore", "pipe", "pipe"],
            env: { ...process.env, TLUA_FOURSLASH_IGNORE_FAILING: "1" },
        });
    }
    catch (error) {
        throw new Error("Failed to spawn test process: " + error);
    }

    if (!testProcess.stdout || !testProcess.stderr) {
        throw new Error("Test process stdout or stderr is null");
    }

    const failingTests: string[] = [];
    const crashingTests: string[] = [];
    const testOutputs = new Map<string, string[]>();
    const allOutputs: string[] = [];
    let hadPanic = false;

    const rl = readline.createInterface({
        input: testProcess.stdout,
        crlfDelay: Infinity,
    });

    rl.on("line", line => {
        try {
            const event: TestEvent = JSON.parse(line);

            // Collect output for each test
            if (event.Action === "output" && event.Output) {
                allOutputs.push(event.Output);
                if (event.Test) {
                    if (!testOutputs.has(event.Test)) {
                        testOutputs.set(event.Test, []);
                    }
                    testOutputs.get(event.Test)!.push(event.Output);
                }

                // Check for panics
                if (/^panic/m.test(event.Output)) {
                    hadPanic = true;
                }
            }

            // Process failed tests
            if (event.Action === "fail" && event.Test) {
                const outputs = testOutputs.get(event.Test) || [];

                // Check if this is a crashing test (contains InternalError)
                const hasCrash = outputs.some(line => line.includes("InternalError"));
                if (hasCrash) {
                    crashingTests.push(event.Test);
                }

                // Unsupported upstream syntax can fail only through a baseline mismatch.
                // Record every failure so recovery output is skipped rather than accepted.
                failingTests.push(event.Test);
            }
        }
        catch (e) {
            // Not JSON, possibly stderr or other output - ignore
        }
    });

    testProcess.stderr.on("data", data => {
        // Check stderr for panics too
        const output = data.toString();
        allOutputs.push(output);
        if (/^panic/m.test(output)) {
            hadPanic = true;
        }
    });

    await new Promise<void>((resolve, reject) => {
        testProcess.on("close", code => {
            if (hadPanic) {
                reject(new Error("Unrecovered panic detected in tests\n" + allOutputs.join("")));
                return;
            }

            const generatedTestNames = new Set<string>();
            for (const entry of fs.readdirSync(generatedTestsPath, { withFileTypes: true })) {
                if (!entry.isFile() || !entry.name.endsWith("_test.go")) continue;
                const content = fs.readFileSync(path.join(generatedTestsPath, entry.name), "utf-8");
                for (const match of content.matchAll(/^func (Test\w+)\(/gm)) {
                    generatedTestNames.add(match[1]);
                }
            }
            // updateFailing runs only the generated package. Preserve approvals for
            // hand-maintained tests, which cannot be rediscovered by this run.
            const preservedTests = existingFailingTests.filter(test => !generatedTestNames.has(test));
            const allFailingTests = [...new Set([...preservedTests, ...failingTests])];
            fs.writeFileSync(failingTestsPath, allFailingTests.sort((a, b) => a.localeCompare(b, "en-US")).join("\n") + "\n", "utf-8");
            fs.writeFileSync(crashingTestsPath, crashingTests.sort((a, b) => a.localeCompare(b, "en-US")).join("\n") + "\n", "utf-8");
            resolve();
        });

        testProcess.on("error", error => {
            reject(error);
        });
    });
}

main().catch(error => {
    console.error("Error:", error);
    process.exit(1);
});
