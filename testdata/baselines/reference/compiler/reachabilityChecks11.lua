//// [tests/cases/compiler/reachabilityChecks11.tlua] ////

//// [reachabilityChecks11.tlua]
// while (true);
local x = 1;

namespace A {
    while (true);
    local x;
}

namespace A1 {
    do {} while(true);
    namespace A {
        interface F {}
    }
}

namespace A2 {
    while (true);
    namespace A {
        local x = 1;
    }
}

namespace A3 {
    while (true);
    type T = string;
}

namespace A4 {
    while (true);
    namespace A {
        local x = 1;
    }
}

function f1(x) {
    if (x) {
        return;
    }
    else {
        throw new Error("123");
    }
    local x;
}

function f2() {
    return;
    class A {
    }
}

namespace B {
    while (true);
    namespace C {
    }
}

function f3() {
    do {
    } while (true);
    local x = 1;
}

function f4() {
    if (true) {
        throw new Error();
    }
    local x = 1;
}


//// [reachabilityChecks11.lua]
-- while (true);
local x = 1;
namespace;
A;
{
    while (true)
        ;
    local x;
}
namespace;
A1;
{
    do { } while (true);
    namespace;
    A;
    {
    }
}
namespace;
A2;
{
    while (true)
        ;
    namespace;
    A;
    {
        local x = 1;
    }
}
namespace;
A3;
{
    while (true)
        ;
}
namespace;
A4;
{
    while (true)
        ;
    namespace;
    A;
    {
        local x = 1;
    }
}
function f1(x) {
    if (x) {
        return;
    }
    else {
        throw new Error("123");
    }
    local x;
}
function f2() {
    return;
    class;
    A;
    {
    }
}
namespace;
B;
{
    while (true)
        ;
    namespace;
    C;
    {
    }
}
function f3() {
    do {
    } while (true);
    local x = 1;
}
function f4() {
    if (true) {
        throw new Error();
    }
    local x = 1;
}
