import * as vscode from "vscode";

import type { ConfigurationMiddleware } from "vscode-languageclient/node";
import type { MessageSignature } from "vscode-languageserver-protocol";

/**
 * Configuration middleware for the tlua language server.
 *
 * The default vscode-languageclient handler uses `getConfiguration().get()`, which
 * returns "fully resolved" values including VS Code defaults. This is problematic
 * because the server has its own defaults, and receiving VS Code's prevents the
 * server from distinguishing user-set values from defaults.
 *
 * This module instead uses `inspect()` to retrieve both explicitly-set values (user/
 * workspace/workspace-folder settings) and VS Code default values from the single
 * `tlua` configuration section, keeping explicit > default precedence.
 *
 * The server still requests its preferences under the `js/ts`, `typescript`, and
 * `javascript` section names (hardcoded server-side); we answer each of those with
 * the user's `tlua.*` configuration so one namespace drives everything.
 *
 * Both the `workspace/configuration` (pull) and `workspace/didChangeConfiguration`
 * (push) middlewares return/send the same merged object for every requested section.
 */

// The section user settings are read from.
const readSection = "tlua";

// The section names the server requests configuration for; each is answered with
// the merged `tlua` configuration.
const serverConfigSections = ["js/ts", "typescript", "javascript"];

const configAliases = [
    ["validate.enabled", "validate.enable"],
    ["format.enabled", "format.enable"],
    ["autoClosingTags.enabled", "autoClosingTags"],
    ["suggest.jsdoc.enabled", "suggest.completeJSDocs"],
] as const;

/**
 * Build a single merged configuration object from the `tlua` config section.
 *
 * For each key, the explicitly-set value (if any) wins over the declared default,
 * so user-set values always take precedence and declared-but-unset settings still
 * get their default.
 */
function getMergedConfiguration(resource: vscode.Uri | undefined): Record<string, any> {
    const { explicit, defaults } = getInspectedConfiguration(readSection, resource);

    const mergedDefaults = defaults ?? Object.create(null);
    const mergedExplicit = explicit ?? Object.create(null);
    applyConfigAliases(mergedExplicit);

    return deepMerge(mergedDefaults, mergedExplicit);
}

/**
 * Given a configuration section name (e.g., "typescript"), use vscode's
 * inspect API to collect both explicitly-set values and default values,
 * returning them as separate nested objects.
 */
function getInspectedConfiguration(
    section: string,
    resource: vscode.Uri | undefined,
): { explicit: Record<string, any> | null; defaults: Record<string, any> | null; } {
    const config = vscode.workspace.getConfiguration(section, resource);
    // Use Object.create(null) so these objects have no prototype to pollute.
    const explicit: Record<string, any> = Object.create(null);
    const defaults: Record<string, any> = Object.create(null);
    let hasExplicit = false;
    let hasDefaults = false;

    const allKeys = collectConfigurationKeys(config);

    for (const key of allKeys) {
        const inspection = config.inspect(key);
        if (!inspection) {
            continue;
        }

        // Pick the most specific explicitly-set value.
        // Language-specific overrides (e.g. [typescript]) take precedence
        // over non-language values at the same scope.
        const explicitValue = inspection.workspaceFolderLanguageValue
            ?? inspection.workspaceFolderValue
            ?? inspection.workspaceLanguageValue
            ?? inspection.workspaceValue
            ?? inspection.globalLanguageValue
            ?? inspection.globalValue;

        if (explicitValue !== undefined) {
            setNestedValue(explicit, key, toJSONObject(explicitValue));
            hasExplicit = true;
        }
        else if (inspection.defaultValue !== undefined) {
            setNestedValue(defaults, key, toJSONObject(inspection.defaultValue));
            hasDefaults = true;
        }
    }

    return {
        explicit: hasExplicit ? explicit : null,
        defaults: hasDefaults ? defaults : null,
    };
}

/**
 * Collect all leaf key paths from a workspace configuration section.
 */
function collectConfigurationKeys(config: vscode.WorkspaceConfiguration): string[] {
    const keys: string[] = [];
    const configMethods = new Set(["get", "has", "inspect", "update"]);

    function walk(obj: any, prefix: string) {
        if (obj === null || obj === undefined || typeof obj !== "object" || Array.isArray(obj)) {
            return;
        }
        for (const key of Object.keys(obj)) {
            if (configMethods.has(key)) {
                continue;
            }
            const fullKey = prefix ? `${prefix}.${key}` : key;
            const value = obj[key];
            if (value !== null && typeof value === "object" && !Array.isArray(value)) {
                walk(value, fullKey);
            }
            else {
                keys.push(fullKey);
            }
        }
    }

    walk(config, "");
    return keys;
}

const prototypeKeys = new Set(["__proto__", "constructor", "prototype"]);

function setNestedValue(obj: Record<string, any>, dottedKey: string, value: any): void {
    const parts = dottedKey.split(".");
    let current = obj;
    for (let i = 0; i < parts.length - 1; i++) {
        const part = parts[i];
        if (prototypeKeys.has(part)) {
            return;
        }
        if (!(part in current) || typeof current[part] !== "object" || current[part] === null) {
            current[part] = Object.create(null);
        }
        current = current[part];
    }
    const lastPart = parts[parts.length - 1];
    if (!prototypeKeys.has(lastPart)) {
        current[lastPart] = value;
    }
}

function getNestedValue(obj: Record<string, any>, dottedKey: string): any {
    const parts = dottedKey.split(".");
    let current: any = obj;
    for (const part of parts) {
        if (prototypeKeys.has(part) || current === null || typeof current !== "object" || !(part in current)) {
            return undefined;
        }
        current = current[part];
    }
    return current;
}

function applyConfigAliases(explicit: Record<string, any>): void {
    for (const [preferred, fallback] of configAliases) {
        if (getNestedValue(explicit, preferred) === undefined) {
            const fallbackValue = getNestedValue(explicit, fallback);
            if (fallbackValue !== undefined) {
                setNestedValue(explicit, preferred, fallbackValue);
            }
        }
    }
}

/**
 * Deep merge b into a. Values in b take precedence over values in a.
 * Returns a new object; does not mutate inputs.
 */
function deepMerge(a: Record<string, any>, b: Record<string, any>): Record<string, any> {
    // Use Object.create(null) so the result has no prototype to pollute.
    const result: Record<string, any> = Object.create(null);
    Object.assign(result, a);
    for (const key of Object.keys(b)) {
        if (prototypeKeys.has(key)) {
            continue;
        }
        if (
            key in result
            && result[key] !== null && typeof result[key] === "object" && !Array.isArray(result[key])
            && b[key] !== null && typeof b[key] === "object" && !Array.isArray(b[key])
        ) {
            result[key] = deepMerge(result[key], b[key]);
        }
        else {
            result[key] = b[key];
        }
    }
    return result;
}

function toJSONObject(obj: any): any {
    if (obj === null || obj === undefined) {
        return obj;
    }
    if (Array.isArray(obj)) {
        return obj.map(toJSONObject);
    }
    if (typeof obj === "object") {
        const res: Record<string, any> = Object.create(null);
        for (const key in obj) {
            if (Object.prototype.hasOwnProperty.call(obj, key)) {
                res[key] = toJSONObject(obj[key]);
            }
        }
        return res;
    }
    return obj;
}

const serverConfigSectionsSet: ReadonlySet<string> = new Set<string>(serverConfigSections);

/**
 * Intercepts workspace/configuration requests. For items requesting one of the
 * server's config sections, returns the merged `tlua` configuration. For any other
 * section (e.g. `editor`), delegates to the default handler via `next`.
 */
export const configurationMiddleware: ConfigurationMiddleware = {
    async configuration(params, token, next) {
        const hasOther = params.items.some(
            item => item.section === undefined || !serverConfigSectionsSet.has(item.section),
        );

        // If all items are handled sections, no need to call next.
        let defaultResults: any[] | undefined;
        if (hasOther) {
            const res = await next(params, token);
            if (Array.isArray(res)) {
                defaultResults = res;
            }
        }

        // Cache merged config per resource URI to avoid redundant recalculation.
        const mergedCache = new Map<string, Record<string, any>>();
        function getMergedCached(resource: vscode.Uri | undefined): Record<string, any> {
            const key = resource?.toString() ?? "";
            let cached = mergedCache.get(key);
            if (cached === undefined) {
                cached = getMergedConfiguration(resource);
                mergedCache.set(key, cached);
            }
            return cached;
        }

        const result: any[] = params.items.map((item, i) => {
            if (item.section !== undefined && serverConfigSectionsSet.has(item.section)) {
                const resource = item.scopeUri ? vscode.Uri.parse(item.scopeUri) : undefined;
                return getMergedCached(resource);
            }
            return defaultResults?.[i] ?? null;
        });

        return result;
    },
};

/**
 * Intercepts outgoing workspace/didChangeConfiguration notifications.
 * Replaces the default settings (which include VS Code defaults) with
 * the merged configuration, keyed by section name.
 *
 * This is typed as returning `Promise<void>` rather than `void` because the
 * `didChangeConfiguration` notification is misannotated in vscode-languageclient
 * as returning void, so we must go through `sendNotification` instead.
 */
export function sendNotificationMiddleware(
    type: string | MessageSignature,
    next: (type: string | MessageSignature, params?: any) => Promise<void>,
    params: any,
): Promise<void> {
    const method = typeof type === "string" ? type : type.method;
    if (method === "workspace/didChangeConfiguration") {
        const merged = getMergedConfiguration(undefined);
        const settings: Record<string, any> = {};
        for (const section of serverConfigSections) {
            settings[section] = merged;
        }
        return next(type, { settings });
    }
    return next(type, params);
}
