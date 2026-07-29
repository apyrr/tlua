//// [tests/cases/compiler/tluaSuspendDeclarationEmit.tlua] ////

//// [net.d.tlua]
declare suspend function fetchOk(url: string): (boolean, any);
declare pollBudget: suspend (n: number) => number;

//// [main.tlua]
// Error: direct call from a sync function.
function syncCaller(): void
  fetchOk("nope");
end

// OK: calls from a suspend context, through both ambient forms.
suspend function suspendCaller(): number
  local ok, res = fetchOk("https://example.com");
  if ok then
    return pollBudget(1);
  end
  return 0;
end

// Error: the declared suspend value keeps the contract too.
function budgetSyncCaller(): void
  pollBudget(2);
end


//// [main.lua]
-- Error: direct call from a sync function.
function syncCaller()
  fetchOk("nope");
end
-- OK: calls from a suspend context, through both ambient forms.
function suspendCaller()
  local ok, res = fetchOk("https://example.com");
  if ok then
    return pollBudget(1);
  end
  return 0;
end
-- Error: the declared suspend value keeps the contract too.
function budgetSyncCaller()
  pollBudget(2);
end
