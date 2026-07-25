//// [tests/cases/compiler/tluaTableAugmentationAlias.tlua] ////

//// [tluaTableAugmentationAlias.tlua]
declare function condition(): boolean;
declare function make(): {};

// Augmentation is order-independent: a member declared anywhere in the program
// belongs to every constructor its path can reach. A replacement that only may
// execute therefore does not erase the member from either table.
BranchRoot = { initial = true };
if condition() then
  BranchRoot = { replaced = true };
end
BranchRoot.value = 1;
local branchValue: number = BranchRoot.value;

LoopRoot = { initial = true };
while condition() do
  LoopRoot = { replaced = true };
  break;
end
LoopRoot.value = "loop";
local loopValue: string = LoopRoot.value;

GotoRoot = { initial = true };
if condition() then
  goto afterReplacement;
end
GotoRoot = { replaced = true };
::afterReplacement::
GotoRoot.value = true;
local gotoValue: boolean = GotoRoot.value;

// A stable local alias names the same constructor, so a write through either
// spelling declares the member on one table.
local aliasRoot = {};
local alias = aliasRoot;
alias.member = 1;
local aliasMember: number = aliasRoot.member;

// Rebinding a local ends its aliasing: the write below must not leak `leaked`
// onto sealedRoot's constructor.
local sealedRoot = {};
local rebound = sealedRoot;
rebound = 0;
rebound.leaked = true; -- error
local sealedRootUse: {} = sealedRoot;

// A nested function sees the enclosing constructor and can augment it.
NestedRoot = {};
local function addNested()
  NestedRoot.nested = "inner";
end
local nestedUse: string = NestedRoot.nested;

// An annotation seals the table.
local annotated: { x: number } = { x = 1 };
annotated.y = 2; -- error

// A non-constructor initializer seals it too, in any file order.
local fromCall = make();
fromCall.z = 3; -- error

// A local rebound to another constructor holds both, so a member declared
// through it belongs to every table it can name. Only a non-constructor store
// seals, which is what separates this from `rebound` above.
local reassigned = {};
reassigned = {};
reassigned.value = 1;
local reassignedUse: number = reassigned.value;

// The same holds through a target list, where the store shares its statement
// with an unrelated target.
local listReassigned = {};
listReassigned, listSide = {}, 0;
listReassigned.value = "list";
local listReassignedUse: string = listReassigned.value;


//// [tluaTableAugmentationAlias.lua]
-- Augmentation is order-independent: a member declared anywhere in the program
-- belongs to every constructor its path can reach. A replacement that only may
-- execute therefore does not erase the member from either table.
BranchRoot = { initial = true };
if condition() then
  BranchRoot = { replaced = true };
end
BranchRoot.value = 1;
local branchValue = BranchRoot.value;
LoopRoot = { initial = true };
while condition() do
  LoopRoot = { replaced = true };
  break;
end
LoopRoot.value = "loop";
local loopValue = LoopRoot.value;
GotoRoot = { initial = true };
if condition() then
  goto afterReplacement;
end
GotoRoot = { replaced = true };
::afterReplacement::
GotoRoot.value = true;
local gotoValue = GotoRoot.value;
-- A stable local alias names the same constructor, so a write through either
-- spelling declares the member on one table.
local aliasRoot = {};
local alias = aliasRoot;
alias.member = 1;
local aliasMember = aliasRoot.member;
-- Rebinding a local ends its aliasing: the write below must not leak `leaked`
-- onto sealedRoot's constructor.
local sealedRoot = {};
local rebound = sealedRoot;
rebound = 0;
rebound.leaked = true; -- error
local sealedRootUse = sealedRoot;
-- A nested function sees the enclosing constructor and can augment it.
NestedRoot = {};
local function addNested()
  NestedRoot.nested = "inner";
end
local nestedUse = NestedRoot.nested;
-- An annotation seals the table.
local annotated = { x = 1 };
annotated.y = 2; -- error
-- A non-constructor initializer seals it too, in any file order.
local fromCall = make();
fromCall.z = 3; -- error
-- A local rebound to another constructor holds both, so a member declared
-- through it belongs to every table it can name. Only a non-constructor store
-- seals, which is what separates this from `rebound` above.
local reassigned = {};
reassigned = {};
reassigned.value = 1;
local reassignedUse = reassigned.value;
-- The same holds through a target list, where the store shares its statement
-- with an unrelated target.
local listReassigned = {};
listReassigned, listSide = {}, 0;
listReassigned.value = "list";
local listReassignedUse = listReassigned.value;
