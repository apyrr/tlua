import * as vscode from "vscode";

import { registerEnablementCommands } from "./commands";
import { aiConnectionString } from "./util";

import { TelemetryReporter as VSCodeTelemetryReporter } from "@vscode/extension-telemetry";
import { SessionManager } from "./session";

import { ExperimentationService } from "./experimentationService";
import { createTelemetryReporter } from "./telemetryReporting";

import assert from "node:assert";

export interface ExtensionAPI {
    onLanguageServerInitialized: vscode.Event<void>;
    initializeAPIConnection(pipe?: string): Promise<string>;
}

export async function activate(context: vscode.ExtensionContext): Promise<ExtensionAPI> {
    await vscode.commands.executeCommand("setContext", "tlua.serverRunning", false);

    const telemetryReporter = createTelemetryReporter(new VSCodeTelemetryReporter(aiConnectionString));
    context.subscriptions.push(telemetryReporter);

    const version = context.extension.packageJSON.version;
    assert(typeof version === "string");
    // Constructing the experimentation service actually sets shared properties
    // so that events include context on treatments/flights.
    // If we actually need to read treatment variables we would hold onto this instance,
    // but for now we just construct it to ensure shared properties are set for telemetry.
    void new ExperimentationService(telemetryReporter, context.extension.id, version, context.globalState);

    const output = vscode.window.createOutputChannel("tlua", { log: true });
    context.subscriptions.push(output);

    const languageServerInitializedEventEmitter = new vscode.EventEmitter<void>();
    context.subscriptions.push(languageServerInitializedEventEmitter);

    const sessionManager = new SessionManager(context, output, languageServerInitializedEventEmitter, telemetryReporter);
    context.subscriptions.push(sessionManager);
    registerEnablementCommands(
        context,
        telemetryReporter,
        () => sessionManager.start(context),
        () => sessionManager.stop(),
    );

    function onLanguageServerInitialized(listener: () => void): vscode.Disposable {
        if (sessionManager.currentSession?.client.isInitialized) {
            listener();
        }
        return languageServerInitializedEventEmitter.event(listener);
    }

    const api: ExtensionAPI = {
        onLanguageServerInitialized: onLanguageServerInitialized,
        async initializeAPIConnection(pipe?: string): Promise<string> {
            return sessionManager.initializeAPIConnection(pipe);
        },
    };

    // tlua owns its own language, so there is no built-in extension to defer to:
    // start the language server as soon as a tlua file activates the extension.
    await sessionManager.start(context);

    return api;
}
