//// [tests/cases/conformance/ported/parseErrorDoubleCommaInCall.tlua] ////

//// [parseErrorDoubleCommaInCall.tlua]
-- ported from tests/cases/compiler/parseErrorDoubleCommaInCall.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: Boolean JS runtime wrapper; replaced with the Lua print builtin

print({
    x = 0,,
})


//// [parseErrorDoubleCommaInCall.lua]
-- ported from tests/cases/compiler/parseErrorDoubleCommaInCall.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: Boolean JS runtime wrapper; replaced with the Lua print builtin
print({
  x = 0,
});
