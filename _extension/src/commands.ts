import * as vscode from "vscode";

import type {
    DocumentUri,
    Location,
    Position,
} from "vscode-languageclient";

import type * as tr from "./telemetryReporting";

export function registerEnablementCommands(
    context: vscode.ExtensionContext,
    telemetryReporter: tr.TelemetryReporter,
    startSession: () => Promise<void>,
    stopSession: () => Promise<void>,
): void {
    context.subscriptions.push(vscode.commands.registerCommand("tlua.enable", async () => {
        telemetryReporter.sendTelemetryEvent("command.enable");
        await startSession();
    }));

    context.subscriptions.push(vscode.commands.registerCommand("tlua.disable", async () => {
        telemetryReporter.sendTelemetryEvent("command.disable");
        await stopSession();
    }));
}

export const codeLensShowLocationsCommandName = "tlua.codeLens.showLocations";
export function registerCodeLensShowLocationsCommand(): vscode.Disposable {
    return vscode.commands.registerCommand(codeLensShowLocationsCommandName, showCodeLensLocations);

    function showCodeLensLocations(...args: unknown[]): void {
        if (args.length !== 3) {
            throw new Error(vscode.l10n.t("Unexpected number of arguments."));
        }

        const lspUri = args[0] as DocumentUri;
        const lspPosition = args[1] as Position;
        const lspLocations = args[2] as Location[];

        const editorUri = vscode.Uri.parse(lspUri);
        const editorPosition = new vscode.Position(lspPosition.line, lspPosition.character);
        const editorLocations = lspLocations.map(loc =>
            new vscode.Location(
                vscode.Uri.parse(loc.uri),
                new vscode.Range(
                    new vscode.Position(loc.range.start.line, loc.range.start.character),
                    new vscode.Position(loc.range.end.line, loc.range.end.character),
                ),
            )
        );

        vscode.commands.executeCommand("editor.action.showReferences", editorUri, editorPosition, editorLocations);
    }
}
