//// [tests/cases/conformance/ported/inferentialTypingWithObjectLiteralProperties.tlua] ////

//// [inferentialTypingWithObjectLiteralProperties.tlua]
-- ported from tests/cases/compiler/inferentialTypingWithObjectLiteralProperties.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function f<T>(x: T, y: T): T
    return x
end

f({ x = { nil } }, { x = { 1 } }).x[1] = "" -- ok
f({ x = { 1 } }, { x = { nil } }).x[1] = "" -- intentional type error


//// [inferentialTypingWithObjectLiteralProperties.lua]
-- ported from tests/cases/compiler/inferentialTypingWithObjectLiteralProperties.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function f(x, y)
  return x;
end
f({ x = { nil } }, { x = { 1 } }).x[1] = ""; -- ok
f({ x = { 1 } }, { x = { nil } }).x[1] = ""; -- intentional type error
