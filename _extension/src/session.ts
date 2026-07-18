import * as vscode from "vscode";
import { ActiveJsTsEditorTracker } from "./activeJsTsEditorTracker";
import { Client } from "./client";
import { registerCodeLensShowLocationsCommand } from "./commands";
import { ProjectStatus } from "./projectStatus";
import { setupStatusBar } from "./statusBar";
import { TelemetryReporter } from "./telemetryReporting";
import {
    getExe,
    readTluaConfig,
} from "./util";

/**
 * SessionManager's lifetime is equal to that of the extension. It is responsible
 * for starting, restarting, replacing, and disposing the Session.
 */
export class SessionManager implements vscode.Disposable {
    currentSession?: Session;
    private disposables: vscode.Disposable[] = [];
    private outputChannel: vscode.LogOutputChannel;
    private initializedEventEmitter: vscode.EventEmitter<void>;
    private telemetryReporter: TelemetryReporter;

    constructor(
        context: vscode.ExtensionContext,
        outputChannel: vscode.LogOutputChannel,
        initializedEventEmitter: vscode.EventEmitter<void>,
        telemetryReporter: TelemetryReporter,
    ) {
        this.outputChannel = outputChannel;
        this.telemetryReporter = telemetryReporter;
        this.initializedEventEmitter = initializedEventEmitter;
    }

    start(context: vscode.ExtensionContext): Promise<void> {
        return this.restart(context);
    }

    async restart(context: vscode.ExtensionContext): Promise<void> {
        if (this.currentSession) {
            this.outputChannel.appendLine("Restarting tlua language server...");
            await this.currentSession.stop();
        }
        this.currentSession = new Session(context, this.outputChannel, this.initializedEventEmitter, this.telemetryReporter, () => this.stop(), () => this.restart(context));
        return this.currentSession.start(context);
    }

    async stop(): Promise<void> {
        if (this.currentSession) {
            await this.currentSession.stop();
            this.currentSession = undefined;
        }
    }

    async initializeAPIConnection(pipe?: string): Promise<string> {
        if (!this.currentSession) {
            throw new Error(vscode.l10n.t("Language server is not running."));
        }
        const result = await this.currentSession.client.initializeAPISession(pipe);
        return result.pipe;
    }

    async dispose(): Promise<void> {
        await this.currentSession?.dispose();
        await Promise.all(this.disposables.map(d => d.dispose()));
    }
}

/**
 * Session's lifetime is equal to that of its LanguageClient. The LanguageClient
 * can be restarted within the same Session only if the underlying exe path/version
 * has not changed. Otherwise, a new Session must be created. Since Session only
 * exists while the LSP server is running (or actively starting/restarting/stopping),
 * it also owns the commands and UI elements that should only be active while the
 * server is running.
 */
class Session implements vscode.Disposable {
    client: Client;
    private disposables: vscode.Disposable[] = [];
    private context: vscode.ExtensionContext;
    private outputChannel: vscode.LogOutputChannel;
    private telemetryReporter: TelemetryReporter;
    private initializedEventEmitter: vscode.EventEmitter<void>;
    private stopSession: () => Promise<void>;
    private restartSession: () => Promise<void>;

    constructor(
        context: vscode.ExtensionContext,
        outputChannel: vscode.LogOutputChannel,
        initializedEventEmitter: vscode.EventEmitter<void>,
        telemetryReporter: TelemetryReporter,
        stopSession: () => Promise<void>,
        restartSession: () => Promise<void>,
    ) {
        this.client = new Client(outputChannel, initializedEventEmitter, telemetryReporter);
        this.context = context;
        this.outputChannel = outputChannel;
        this.telemetryReporter = telemetryReporter;
        this.initializedEventEmitter = initializedEventEmitter;
        this.stopSession = stopSession;
        this.restartSession = restartSession;
        this.registerCommands();
    }

    async start(context: vscode.ExtensionContext): Promise<void> {
        await vscode.commands.executeCommand("setContext", "tlua.isManagedFile", false);
        const exe = await getExe(context);
        await this.client.start(exe);
        this.disposables.push(setupStatusBar(exe));

        // Set up active editor tracker and UI features
        const activeEditorTracker = new ActiveJsTsEditorTracker();
        this.disposables.push(activeEditorTracker);

        const projectStatus = new ProjectStatus(this.client, activeEditorTracker, this.initializedEventEmitter.event);
        this.disposables.push(projectStatus);

        // If already initialized, fire immediately so projectStatus picks it up
        if (this.client.isInitialized) {
            this.initializedEventEmitter.fire();
        }

        await vscode.commands.executeCommand("setContext", "tlua.serverRunning", true);
    }

    tryRestartClient(context: vscode.ExtensionContext): Promise<boolean> {
        return this.client.tryRestart(context);
    }

    registerCommands(): void {
        this.disposables.push(registerCodeLensShowLocationsCommand());

        this.disposables.push(vscode.commands.registerCommand("tlua.restart", async () => {
            this.telemetryReporter.sendTelemetryEvent("command.restartLanguageServer");
            if (await this.tryRestartClient(this.context)) {
                return;
            }

            await this.restartSession();
        }));

        this.disposables.push(vscode.commands.registerCommand("tlua.output.focus", () => {
            this.outputChannel.show();
        }));

        this.disposables.push(vscode.commands.registerCommand("tlua.showMenu", () => {
            showCommands(this.client);
        }));

        this.disposables.push(vscode.commands.registerCommand("tlua.sortImports", () => {
            return vscode.commands.executeCommand("editor.action.sourceAction", {
                kind: "source.sortImports",
                apply: "first",
            });
        }));

        this.disposables.push(vscode.commands.registerCommand("tlua.removeUnusedImports", () => {
            return vscode.commands.executeCommand("editor.action.sourceAction", {
                kind: "source.removeUnusedImports",
                apply: "first",
            });
        }));

        this.disposables.push(vscode.commands.registerCommand("tlua.reportIssue", () => {
            this.telemetryReporter.sendTelemetryEvent("command.reportIssue");
            vscode.commands.executeCommand("workbench.action.openIssueReporter", {
                extensionId: this.context.extension.id,
            });
        }));

        this.disposables.push(vscode.commands.registerCommand("tlua.dev.runGC", async () => {
            try {
                await this.client.runGC();
                vscode.window.showInformationMessage(vscode.l10n.t(`Garbage collection triggered`));
            }
            catch (error) {
                vscode.window.showErrorMessage(vscode.l10n.t(`Failed to run GC: {0}`, String(error)));
            }
        }));

        this.disposables.push(vscode.commands.registerCommand("tlua.dev.saveHeapProfile", async () => {
            const dir = await promptForProfileDirectory();
            if (!dir) return;
            try {
                const file = await this.client.saveHeapProfile(dir);
                vscode.window.showInformationMessage(vscode.l10n.t(`Heap profile saved to: {0}`, file));
            }
            catch (error) {
                vscode.window.showErrorMessage(vscode.l10n.t(`Failed to save heap profile: {0}`, String(error)));
            }
        }));

        this.disposables.push(vscode.commands.registerCommand("tlua.dev.saveAllocProfile", async () => {
            const dir = await promptForProfileDirectory();
            if (!dir) return;
            try {
                const file = await this.client.saveAllocProfile(dir);
                vscode.window.showInformationMessage(vscode.l10n.t(`Allocation profile saved to: {0}`, file));
            }
            catch (error) {
                vscode.window.showErrorMessage(vscode.l10n.t(`Failed to save allocation profile: {0}`, String(error)));
            }
        }));

        this.disposables.push(vscode.commands.registerCommand("tlua.dev.startCPUProfile", async () => {
            const dir = await promptForProfileDirectory();
            if (!dir) return;
            try {
                await this.client.startCPUProfile(dir);
                vscode.commands.executeCommand("setContext", "tlua.cpuProfileRunning", true);
                vscode.window.showInformationMessage(vscode.l10n.t(`CPU profiling started. Profile will be saved to: {0}`, dir));
            }
            catch (error) {
                vscode.window.showErrorMessage(vscode.l10n.t(`Failed to start CPU profile: {0}`, String(error)));
                vscode.commands.executeCommand("setContext", "tlua.cpuProfileRunning", false);
            }
        }));

        this.disposables.push(vscode.commands.registerCommand("tlua.dev.stopCPUProfile", async () => {
            try {
                const file = await this.client.stopCPUProfile();
                vscode.commands.executeCommand("setContext", "tlua.cpuProfileRunning", false);
                vscode.window.showInformationMessage(vscode.l10n.t(`CPU profile saved to: {0}`, file));
            }
            catch (error) {
                vscode.window.showErrorMessage(vscode.l10n.t(`Failed to stop CPU profile: {0}`, String(error)));
            }
        }));

        this.disposables.push(vscode.commands.registerCommand("tlua.initializeAPIConnection", async () => {
            const result = await this.client.initializeAPISession();
            return result.pipe;
        }));

        this.disposables.push(vscode.commands.registerCommand("tlua.initializeAPIConnection.ui", async () => {
            try {
                const result = await this.client.initializeAPISession();
                const copyString = vscode.l10n.t("Copy");
                const copy = await vscode.window.showInformationMessage(vscode.l10n.t(`API session initialized. Listening on: {0}`, result.pipe), copyString);
                if (copy === copyString) {
                    await vscode.env.clipboard.writeText(result.pipe);
                }
            }
            catch (error) {
                const message = error instanceof Error ? error.message : String(error);
                vscode.window.showErrorMessage(vscode.l10n.t(`Failed to initialize API session: {0}`, message));
            }
        }));
    }

    private async disposeSessionState(): Promise<void> {
        await vscode.commands.executeCommand("setContext", "tlua.serverRunning", false);
        await vscode.commands.executeCommand("setContext", "tlua.cpuProfileRunning", false);
        const disposables = this.disposables.splice(0);
        await Promise.all(disposables.map(d => d.dispose()));
    }

    async stop(): Promise<void> {
        await this.disposeSessionState();
        await this.client.stop();
    }

    async dispose(): Promise<void> {
        await this.disposeSessionState();
        await this.client.dispose();
    }
}

async function showCommands(client: Client): Promise<void> {
    interface CommandItem {
        label: string;
        description?: string;
        kind?: vscode.QuickPickItemKind;
        command?: string;
        action?: () => Promise<void>;
    }
    const commands: CommandItem[] = [
        {
            label: vscode.l10n.t("$(refresh) Restart Server"),
            description: vscode.l10n.t("Restart the tlua language server"),
            command: "tlua.restart",
        },
        {
            label: vscode.l10n.t("$(output) Show Output"),
            description: vscode.l10n.t("Show the tlua output log"),
            command: "tlua.output.focus",
        },
        {
            label: vscode.l10n.t("$(report) Report Issue"),
            description: vscode.l10n.t("Report an issue with tlua"),
            command: "tlua.reportIssue",
        },
        {
            label: vscode.l10n.t("$(stop-circle) Disable tlua"),
            description: vscode.l10n.t("Stop the tlua language server"),
            command: "tlua.disable",
        },
    ];

    const showDebugInfo = readTluaConfig("showDebugInfo", false);
    if (showDebugInfo) {
        const exe = client.getCurrentExe();
        const pid = client.serverPid;
        commands.push({ label: "", kind: vscode.QuickPickItemKind.Separator });
        if (exe) {
            commands.push({
                label: vscode.l10n.t(`Executable`),
                description: exe.path,
                action: async () => {
                    await vscode.env.clipboard.writeText(exe.path);
                    vscode.window.showInformationMessage(vscode.l10n.t("Executable path copied to clipboard."));
                },
            });
        }
        if (pid) {
            commands.push({
                label: vscode.l10n.t(`PID`),
                description: `${pid}`,
                action: async () => {
                    await vscode.env.clipboard.writeText(`${pid}`);
                    vscode.window.showInformationMessage(vscode.l10n.t("Server PID copied to clipboard."));
                },
            });
        }
    }

    const selected = await vscode.window.showQuickPick(commands, {
        placeHolder: vscode.l10n.t("{0} Commands", "tlua"),
    });

    if (selected) {
        if (selected.action) {
            await selected.action();
        }
        else if (selected.command) {
            await vscode.commands.executeCommand(selected.command);
        }
    }
}

async function promptForProfileDirectory(): Promise<string | undefined> {
    const defaultDir = vscode.workspace.workspaceFolders?.[0]?.uri.fsPath ?? "";
    const dir = await vscode.window.showInputBox({
        prompt: vscode.l10n.t("Enter directory path for profile output"),
        value: defaultDir,
        validateInput: value => {
            if (!value.trim()) {
                return vscode.l10n.t("Directory path is required");
            }
            return undefined;
        },
    });
    return dir?.trim();
}
