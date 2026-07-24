//// [tests/cases/compiler/tluaTableAugmentationDefaultedGuardCrossFile.tlua] ////

//// [roots.tlua]
Forward = Forward or {};
MixedKey = MixedKey or {};
Parenthesized = Parenthesized or {};

//// [members.tlua]
Forward.zod = Forward.zod or {};
Forward.zod.value = true;

MixedKey.zod = MixedKey["zod"] or {};
MixedKey.zod.value = 1;

// Parenthesized self-reads are not the exact defaulted-guard idiom.
Parenthesized.zod = (Parenthesized.zod) or {};

//// [reverseMember.tlua]
// Root resolution is program-wide, not source-file-order dependent.
Reverse.zod = Reverse.zod or {};
Reverse.zod.value = "reverse";

//// [reverseRoot.tlua]
Reverse = Reverse or {};

//// [reads.tlua]
local forward: boolean = Forward.zod.value;
local mixedKey: number = MixedKey.zod.value;
local reverse: string = Reverse.zod.value;


//// [roots.lua]
Forward = Forward or {};
MixedKey = MixedKey or {};
Parenthesized = Parenthesized or {};
//// [members.lua]
Forward.zod = Forward.zod or {};
Forward.zod.value = true;
MixedKey.zod = MixedKey["zod"] or {};
MixedKey.zod.value = 1;
-- Parenthesized self-reads are not the exact defaulted-guard idiom.
Parenthesized.zod = (Parenthesized.zod) or {};
//// [reverseMember.lua]
-- Root resolution is program-wide, not source-file-order dependent.
Reverse.zod = Reverse.zod or {};
Reverse.zod.value = "reverse";
//// [reverseRoot.lua]
Reverse = Reverse or {};
//// [reads.lua]
local forward = Forward.zod.value;
local mixedKey = MixedKey.zod.value;
local reverse = Reverse.zod.value;
