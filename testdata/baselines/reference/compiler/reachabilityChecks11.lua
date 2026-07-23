//// [tests/cases/compiler/reachabilityChecks11.tlua] ////

//// [reachabilityChecks11.tlua]
// while (true);
local x = 1;

namespace A {
    while true do end
    local x;
}

namespace A1 {
    while true do end
    namespace A {
        interface F {}
    }
}

namespace A2 {
    while true do end
    namespace A {
        local x = 1;
    }
}

namespace A3 {
    while true do end
    type T = string;
}

namespace A4 {
    while true do end
    namespace A {
        local x = 1;
    }
}

function f1(x)
    if x then
        return;
    else
        throw new Error("123");
    end
    local x;
end

function f2()
    return;
    class A {
    }
end

namespace B {
    while true do end
    namespace C {
    }
}

function f3()
    while true do
    end
    local x = 1;
end

function f4()
    if true then
        throw new Error();
    end
    local x = 1;
end


//// [reachabilityChecks11.lua]
-- while (true);
local x = 1;
namespace;
A;
do
  while true do
  end
  local x;
end
namespace;
A1;
do
  while true do
  end
  namespace;
  A;
  do
  end
end
namespace;
A2;
do
  while true do
  end
  namespace;
  A;
  do
    local x = 1;
  end
end
namespace;
A3;
do
  while true do
  end
end
namespace;
A4;
do
  while true do
  end
  namespace;
  A;
  do
    local x = 1;
  end
end
function f1(x)
  if x then
    return;
  else
    throw new Error("123");
  end
  local x;
end
function f2()
  return;
  class;
  A;
  do
  end
end
namespace;
B;
do
  while true do
  end
  namespace;
  C;
  do
  end
end
function f3()
  while true do
  end
  local x = 1;
end
function f4()
  if true then
    throw new Error();
  end
  local x = 1;
end
