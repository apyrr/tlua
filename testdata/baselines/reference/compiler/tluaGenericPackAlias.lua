//// [tests/cases/compiler/tluaGenericPackAlias.tlua] ////

//// [tluaGenericPackAlias.tlua]
type Handler<...A> = (...: A) => void;
type PairHandler = Handler<(number, string)>;
type EmptyHandler = Handler<()>;
type OneHandler = Handler<number>;
type StringHandlers = Handler<(...string)>;
type ForwardedHandler<...A> = Handler<(...A)>;

local pair: Handler<(number, string)> = function(a, b): void
    local n: number = a;
    local s: string = b;
end;
pair(1, "x");

local empty: Handler<()> = function(): void end;
empty();

local one: Handler<number> = function(value): void
    local n: number = value;
end;
one(1);

local strings: Handler<(...string)> = function(...: string): void end;
strings("a", "b");

local forward = function<...A>(handler: Handler<(...A)>, ...: A): void
    handler(...);
end;
forward(pair, 1, "x");

type Scalar<T> = T;
type WrongAliasArgument = Scalar<(number, string)>;

declare function scalar<T>(value: T): void;
scalar<(number, string)>(1);

interface StillRejected<...A> {}

type BarePackArgument<...A> = Handler<A>;


//// [tluaGenericPackAlias.lua]
local pair = function(a, b)
    local n = a;
    local s = b;
end;
pair(1, "x");
local empty = function()
end;
empty();
local one = function(value)
    local n = value;
end;
one(1);
local strings = function(...)
end;
strings("a", "b");
local forward = function(handler, ...)
    handler(...);
end;
forward(pair, 1, "x");
scalar(1);
