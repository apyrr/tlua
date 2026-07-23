//// [tests/cases/compiler/tluaGenericPackDeclarationEmit.tlua] ////

//// [tluaGenericPackDeclarationEmit.tlua]
// A generic pack signature round-trips through declaration emit as `<...A>` and
// `...: A` / `...R`, the same spelling the source uses.

declare function apply<...A, ...R>(f: (...: A) => (...R), ...: A): (...R);

function wrap<...A, ...R>(f: (...: A) => (...R)): (...: A) => (...R)
  return f;
end


//// [tluaGenericPackDeclarationEmit.lua]
-- A generic pack signature round-trips through declaration emit as `<...A>` and
-- `...: A` / `...R`, the same spelling the source uses.
function wrap(f)
  return f;
end
