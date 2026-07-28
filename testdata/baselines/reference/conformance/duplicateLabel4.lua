//// [tests/cases/conformance/ported/duplicateLabel4.tlua] ////

//// [duplicateLabel4.tlua]
-- ported from tests/cases/compiler/duplicateLabel4.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target)
-- note: a JavaScript label scopes to the single statement it prefixes, so
--   upstream's two sequential labeled loops reuse the name without error. A Lua
--   label scopes to its enclosing block, so the port gives each label its own
--   `do ... end` block; that preserves the upstream no-error intent and keeps
--   the emitted Lua valid.


do
    ::target::
    while true do
    end
end

do
    ::target::
    while true do
    end
end


//// [duplicateLabel4.lua]
-- ported from tests/cases/compiler/duplicateLabel4.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target)
-- note: a JavaScript label scopes to the single statement it prefixes, so
--   upstream's two sequential labeled loops reuse the name without error. A Lua
--   label scopes to its enclosing block, so the port gives each label its own
--   `do ... end` block; that preserves the upstream no-error intent and keeps
--   the emitted Lua valid.
do
  ::target::
  while true do
  end
end
do
  ::target::
  while true do
  end
end
