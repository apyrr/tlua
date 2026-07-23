//// [tests/cases/compiler/tluaNoClass.tlua] ////

//// [tluaNoClass.tlua]
// Classes are removed in tlua: `class` is an ordinary identifier (Lua has no
// `class` keyword) and class declarations/expressions no longer parse. Each form
// below produces natural parse errors (no tlua-specific diagnostic).

class C {}

class D extends C {}

abstract class E {}

class F {}

local g = class {};

class WithMembers {
  x = 1;
  m() {}
  constructor() {}
  static s() {}
  #p = 2;
  get a() {
    return 1;
  }
}

// `class` is now an ordinary identifier.
function usesClassName(): number
  local class = 3;
  return class;
end


//// [tluaNoClass.lua]
-- Classes are removed in tlua: `class` is an ordinary identifier (Lua has no
-- `class` keyword) and class declarations/expressions no longer parse. Each form
-- below produces natural parse errors (no tlua-specific diagnostic).
class;
C;
do
end
class;
D;
C;
do
end
abstract;
class;
E;
do
end
class;
F;
do
end
local g = class;
do
end
;
class;
WithMembers;
do
    x = 1;
    m();
    do
    end
    constructor();
    do
    end
    s();
    do
    end
    #p;
    2;
    get;
    a();
    do
        return 1;
    end
end
-- `class` is now an ordinary identifier.
function usesClassName()
    local class = 3;
    return class;
end
