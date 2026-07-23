//// [tests/cases/compiler/tluaNoLabeledStatement.tlua] ////

//// [tluaNoLabeledStatement.tlua]
// tlua has no labeled statement and no labeled jump: `::name::` and `goto` took
// their place. Each of these is now a parse error.

// `name:` no longer starts a labeled statement.
outer: while true do
end

// break takes no label.
function labeledBreak(n: number): void
    while n > 0 do
        break outer
    end
end

// continue takes no label.
function labeledContinue(n: number): void
    while n > 0 do
        continue outer
    end
end

// `goto` is a reserved word, so it cannot be a binding name.
local goto = 1;


//// [tluaNoLabeledStatement.lua]
-- tlua has no labeled statement and no labeled jump: `::name::` and `goto` took
-- their place. Each of these is now a parse error.
-- `name:` no longer starts a labeled statement.
outer;
while true do
end
-- break takes no label.
function labeledBreak(n)
  while n > 0 do
    break;
    outer;
  end
end
-- continue takes no label.
function labeledContinue(n)
  while n > 0 do
    continue;
    outer;
  end
end
-- `goto` is a reserved word, so it cannot be a binding name.
local ;
goto ;
1;
