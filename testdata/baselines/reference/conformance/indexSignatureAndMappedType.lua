//// [tests/cases/conformance/ported/indexSignatureAndMappedType.tlua] ////

//// [indexSignatureAndMappedType.tlua]
-- ported from tests/cases/compiler/indexSignatureAndMappedType.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: declaration output is unsupported for this Lua module; checker coverage is retained

-- A mapped type { [P in K]: X }, where K is a generic type, is related to
-- { [key: string]: Y } if X is related to Y.

function f1<T, K extends string>(x: { [key: string]: T }, y: Record<K, T>)
    x = y
    y = x -- Error
end

function f2<T>(x: { [key: string]: T }, y: Record<string, T>)
    x = y
    y = x
end

function f3<T, U, K extends string>(x: { [key: string]: T }, y: Record<K, U>)
    x = y -- Error
    y = x -- Error
end

-- Repro from #14548

type Dictionary = {
    [key: string]: string;
}

interface IBaseEntity {
    name: string;
    properties: Dictionary;
}

interface IEntity<T extends string> extends IBaseEntity {
    properties: Record<T, string>;
}


//// [indexSignatureAndMappedType.lua]
-- ported from tests/cases/compiler/indexSignatureAndMappedType.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: declaration output is unsupported for this Lua module; checker coverage is retained
-- A mapped type { [P in K]: X }, where K is a generic type, is related to
-- { [key: string]: Y } if X is related to Y.
function f1(x, y)
  x = y;
  y = x; -- Error
end
function f2(x, y)
  x = y;
  y = x;
end
function f3(x, y)
  x = y; -- Error
  y = x; -- Error
end
