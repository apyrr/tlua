import * as vscode from "vscode";
import {
    ExeInfo,
    tluaLanguageModes,
} from "./util";

export function setupStatusBar(exe: ExeInfo): vscode.Disposable {
    const statusItem = vscode.languages.createLanguageStatusItem("tlua.status", tluaLanguageModes);
    statusItem.name = vscode.l10n.t("tlua");
    statusItem.text = exe.isLocal ? vscode.l10n.t("$(beaker) {0}", exe.version) : exe.version;
    statusItem.detail = vscode.l10n.t("tlua Language Server");
    statusItem.command = {
        title: vscode.l10n.t("Show Menu"),
        command: "tlua.showMenu",
    };
    return statusItem;
}
