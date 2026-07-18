//// [tests/cases/compiler/jsDocCodeblockParsing.tlua] ////

//// [jsDocCodeblockParsing.tlua]
/**
 * text
 * @example Foo
 * ```
 * @Embed[asfasdfasf]
 * ```
 * becomes
 * ```html
 * <div></div>
 * ```
 */
local x = 1;

/**
 * Some text
 * ```
 * @tag inside code
 * ```
 * @param y - a number
 */
function foo(y: number) {}


//// [jsDocCodeblockParsing.lua]
--[[*
 * text
 * @example Foo
 * ```
 * @Embed[asfasdfasf]
 * ```
 * becomes
 * ```html
 * <div></div>
 * ```
 ]]
local x = 1;
--[[*
 * Some text
 * ```
 * @tag inside code
 * ```
 * @param y - a number
 ]]
function foo(y) { }
