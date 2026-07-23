//// [tests/cases/compiler/asyncSuperDefaultParameters.tlua] ////

//// [asyncSuperDefaultParameters.tlua]
class B {
    m() {
        return 1;
    }
}

class C extends B {
    f() {
        local g = async function(b = super.m()) return b end;
        return g();
    }

    async h(b = super.m()) {
        return b;
    }
}


//// [asyncSuperDefaultParameters.lua]
class;
B;
do
  m();
  do
    return 1;
  end
end
class;
C;
B;
do
  f();
  do
    local g = async function(b = super.m())
      return b;
    end;
    return g();
  end
  async;
  h(b = super.m());
  do
    return b;
  end
end
