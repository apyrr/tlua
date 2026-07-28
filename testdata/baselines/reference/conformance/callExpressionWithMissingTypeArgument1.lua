//// [tests/cases/conformance/ported/callExpressionWithMissingTypeArgument1.tlua] ////

//// [callExpressionWithMissingTypeArgument1.tlua]
-- ported from tests/cases/compiler/callExpressionWithMissingTypeArgument1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

Foo<a,,b>()


//// [callExpressionWithMissingTypeArgument1.lua]
-- ported from tests/cases/compiler/callExpressionWithMissingTypeArgument1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
Foo();
