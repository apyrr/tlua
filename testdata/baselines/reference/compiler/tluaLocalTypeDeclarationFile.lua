//// [tests/cases/compiler/tluaLocalTypeDeclarationFile.tlua] ////

//// [lib.d.tlua]
// A declaration file is a non-module script with no module scope: its top-level
// locals live only in the global table. `local` therefore has no boundary to be
// private to and degrades to global here -- both `Priv` and `Pub` are global.
// (Module privacy for `local type` requires a module, i.e. a regular .tlua file.)
local type Priv = number;
type Pub = string;
declare g: Priv;

//// [main.tlua]
// Both are visible: `local` did not make `Priv` module-private in a script.
local a: Pub = "x";
local b: Priv = 1;


//// [main.lua]
-- Both are visible: `local` did not make `Priv` module-private in a script.
local a = "x";
local b = 1;
