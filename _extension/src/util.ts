import * as path from "path";
import * as vscode from "vscode";

export const aiConnectionString = "0c6ae279ed8443289764825290e4f9e2-1a736e7c-1324-4338-be46-fc2a58ae4d14-7255";

export const languageClientName = "tlua Language Server";

// The language IDs this extension serves. tlua owns its own language (contributed
// in package.json), so there is no coexistence with VS Code's built-in TypeScript.
export const tluaLanguageModes = [
    "tlua",
];

/**
 * URI schemes for which language features should be disabled.
 */
export const disabledSchemes = new Set([
    "git",
    "vsls",
    "github",
    "azurerepos",
    "chat-editing-text-model",
]);

export function isSupportedLanguageMode(doc: vscode.TextDocument): boolean {
    return tluaLanguageModes.includes(doc.languageId);
}

const configRegex = /^tluaconfig\.(.+\.)?json$/i;
export function isTsConfigFileName(fileName: string): boolean {
    return configRegex.test(path.basename(fileName));
}

export interface ExeInfo {
    path: string;
    version: string;
    name: string;
    isLocal?: boolean;
}

const packagedExeBaseNames = ["tlua"];

async function getBuiltinExePath(context: vscode.ExtensionContext): Promise<ExeInfo> {
    if (context.extensionMode === vscode.ExtensionMode.Development) {
        const exeName = `tlua${process.platform === "win32" ? ".exe" : ""}`;
        const exe = context.asAbsolutePath(path.join("../", "built", "local", exeName));
        try {
            await vscode.workspace.fs.stat(vscode.Uri.file(exe));
            return { path: exe, version: "(local)", name: "tlua", isLocal: true };
        }
        catch {}
    }
    return getPackagedExePath(context.extension.extensionUri, getBundledVersion(context.extension.packageJSON));
}

async function getPackagedExePath(extensionUri: vscode.Uri, version: unknown): Promise<ExeInfo> {
    const exe = await tryGetPackagedExePath(extensionUri, version);
    if (exe) {
        return exe;
    }
    throw new Error(vscode.l10n.t("Could not find a tlua executable in the extension package."));
}

function getBundledVersion(packageJSON: unknown): string {
    if (packageJSON && typeof packageJSON === "object" && "version" in packageJSON) {
        const version = packageJSON.version;
        if (typeof version === "string") {
            return version;
        }
    }
    return "unknown";
}

async function tryGetPackagedExePath(extensionUri: vscode.Uri, version: unknown): Promise<ExeInfo | undefined> {
    for (const baseName of packagedExeBaseNames) {
        const exeName = `${baseName}${process.platform === "win32" ? ".exe" : ""}`;
        const exePath = vscode.Uri.joinPath(extensionUri, "lib", exeName);
        try {
            await vscode.workspace.fs.stat(exePath);
            return {
                path: withLongPathPrefix(exePath.fsPath),
                version: typeof version === "string" ? version : "unknown",
                name: baseName,
            };
        }
        catch {}
    }
    return undefined;
}

/**
 * Returns the base directory for resolving relative paths in workspace config.
 * - Multi-root workspace: the directory containing the `.code-workspace` file.
 * - Single-root workspace: the lone workspace folder.
 * - No workspace: undefined.
 */
function workspaceConfigBase(): vscode.Uri | undefined {
    const wsFile = vscode.workspace.workspaceFile;
    if (wsFile && wsFile.scheme === "file") {
        return vscode.Uri.file(path.dirname(wsFile.fsPath));
    }
    if (vscode.workspace.workspaceFolders && vscode.workspace.workspaceFolders.length > 0) {
        return vscode.workspace.workspaceFolders[0].uri;
    }
    return undefined;
}

function workspaceResolve(relativePath: string): vscode.Uri {
    if (path.isAbsolute(relativePath)) {
        return vscode.Uri.file(relativePath);
    }
    const base = workspaceConfigBase();
    if (base) {
        return vscode.Uri.joinPath(base, relativePath);
    }
    return vscode.Uri.file(relativePath);
}

/**
 * Memento used to control whether the user has opted into using a sdk location defined
 * in workspace settings. This is *not* a trust boundary - workspace trust is required
 * before the extension will prompt to set this memento to true. It provides users a way
 * to opt out of using the workspace-provided sdk without changing committed workspace
 * settings. Since the stored value is only a boolean, it does not protect against
 * executing a different sdk than the one the user originally opted into if the workspace
 * settings or node_modules content changes - that's why workspace trust is always
 * required, and why the prompts that set this value should not be interpreted as
 * indicating trust for a specific sdk installation.
 */
const useWorkspaceSdkStorageKey = "tlua.useWorkspaceSdk";

export async function getExe(context: vscode.ExtensionContext): Promise<ExeInfo> {
    for (const candidate of getTrustedSdkCandidates(context, await getSdkCandidates())) {
        const exe = await resolveSdkPathToExe(candidate.value);
        if (exe) {
            return exe;
        }
    }

    return getBuiltinExePath(context);
}

function getTrustedSdkCandidates(context: vscode.ExtensionContext, sdkCandidates: ExplicitConfigValue<string>[]): ExplicitConfigValue<string>[] {
    // If sdk is set at the workspace level, require both workspace trust and
    // explicit user opt-in. Workspace trust can be revoked after the memento is
    // set, so we must always check both.
    if (sdkCandidates.some(candidate => candidate.target !== vscode.ConfigurationTarget.Global)) {
        if (!vscode.workspace.isTrusted || !context.workspaceState.get<boolean>(useWorkspaceSdkStorageKey, false)) {
            return sdkCandidates.filter(candidate => candidate.target === vscode.ConfigurationTarget.Global);
        }
    }
    return sdkCandidates;
}

interface ExplicitConfigValue<T> {
    value: T;
    target: vscode.ConfigurationTarget;
    order: number;
}

export function readTluaConfig<T>(key: string, defaultValue: T): T {
    return vscode.workspace.getConfiguration("tlua").get<T>(key, defaultValue);
}

function getExplicitConfigValues<T>(section: string, key: string): ExplicitConfigValue<T>[] {
    const inspection = vscode.workspace.getConfiguration(section).inspect<T>(key);
    if (!inspection) {
        return [];
    }

    const candidates: Array<{ value: T | undefined; target: vscode.ConfigurationTarget; order: number; }> = [
        { value: inspection.workspaceFolderLanguageValue, target: vscode.ConfigurationTarget.WorkspaceFolder, order: 0 },
        { value: inspection.workspaceFolderValue, target: vscode.ConfigurationTarget.WorkspaceFolder, order: 1 },
        { value: inspection.workspaceLanguageValue, target: vscode.ConfigurationTarget.Workspace, order: 2 },
        { value: inspection.workspaceValue, target: vscode.ConfigurationTarget.Workspace, order: 3 },
        { value: inspection.globalLanguageValue, target: vscode.ConfigurationTarget.Global, order: 4 },
        { value: inspection.globalValue, target: vscode.ConfigurationTarget.Global, order: 5 },
    ];

    const result: ExplicitConfigValue<T>[] = [];
    for (const candidate of candidates) {
        if (candidate.value !== undefined) {
            result.push({
                value: candidate.value,
                target: candidate.target,
                order: candidate.order,
            });
        }
    }
    return result.sort(compareExplicitConfigValues);
}

function compareExplicitConfigValues<T>(a: ExplicitConfigValue<T>, b: ExplicitConfigValue<T>): number {
    return b.target - a.target || a.order - b.order;
}

async function getSdkCandidates(): Promise<ExplicitConfigValue<string>[]> {
    // A single source: `tlua.sdk.path`. Points at a directory containing a tlua
    // binary (or an installed package that ships one).
    return getExplicitConfigValues<string>("tlua", "sdk.path")
        .filter(candidate => typeof candidate.value === "string" && !!candidate.value);
}

async function resolveSdkPathToExe(sdkPath: string): Promise<ExeInfo | undefined> {
    const resolved = workspaceResolve(sdkPath);
    for (const packagePath of [resolved, vscode.Uri.joinPath(resolved, "..")]) {
        try {
            const packageJsonPath = vscode.Uri.joinPath(packagePath, "package.json");
            const packageJson = JSON.parse(await vscode.workspace.fs.readFile(packageJsonPath).then(buffer => buffer.toString()));
            const name: unknown = packageJson.name;
            const bin: unknown = packageJson.bin;
            if (typeof name !== "string" || !bin || typeof bin !== "object") continue;

            const baseName = name.startsWith("@") ? name.split("/")[1] : name;
            if (!baseName) continue;
            // Validate the package declares its npx-command bin, but the actual Go
            // binary file is always named `tlua` (kept in sync with getExePath.js).
            const expectedBinName = "tlua";
            if (!Object.prototype.hasOwnProperty.call(bin, expectedBinName)) continue;

            const exeName = `tlua${process.platform === "win32" ? ".exe" : ""}`;
            const platformPackage = `${baseName}-${process.platform}-${process.arch}`;
            const nodeModules = name.startsWith("@")
                ? vscode.Uri.joinPath(packagePath, "..", "..")
                : vscode.Uri.joinPath(packagePath, "..");
            const exePath = vscode.Uri.joinPath(nodeModules, "@tlua", platformPackage, "lib", exeName);
            await vscode.workspace.fs.stat(exePath);
            return { path: withLongPathPrefix(exePath.fsPath), version: typeof packageJson.version === "string" ? packageJson.version : "unknown", name: expectedBinName };
        }
        catch {}
    }
    for (const baseName of packagedExeBaseNames) {
        try {
            const exePath = vscode.Uri.joinPath(resolved, `${baseName}${process.platform === "win32" ? ".exe" : ""}`);
            await vscode.workspace.fs.stat(exePath);
            return { path: withLongPathPrefix(exePath.fsPath), version: "(local)", name: baseName, isLocal: true };
        }
        catch {}
    }
}

function withLongPathPrefix(exePath: string): string {
    if (process.platform === "win32" && exePath.length >= 248 && !exePath.startsWith("\\\\?\\")) {
        return "\\\\?\\" + exePath;
    }
    return exePath;
}

export function getLanguageForUri(uri: vscode.Uri): string | undefined {
    // `.d.tlua` also reports extension `.tlua` (extname keys off the last dot).
    return path.posix.extname(uri.path) === ".tlua" ? "tlua" : undefined;
}
