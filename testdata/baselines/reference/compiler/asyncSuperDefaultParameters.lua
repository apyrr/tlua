//// [tests/cases/compiler/asyncSuperDefaultParameters.tlua] ////

//// [asyncSuperDefaultParameters.tlua]
class B {
    m() {
        return 1;
    }
}

class C extends B {
    f() {
        local g = async (b = super.m()) => b;
        return g();
    }

    async h(b = super.m()) {
        return b;
    }
}


//// [asyncSuperDefaultParameters.lua]
class;
B;
{
    m();
    {
        return 1;
    }
}
class;
C;
B;
{
    f();
    {
        local g = async function(b = super.m()) return b end;
        return g();
    }
    async;
    h(b = super.m());
    {
        return b;
    }
}
