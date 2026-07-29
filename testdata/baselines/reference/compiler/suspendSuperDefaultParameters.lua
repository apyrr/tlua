//// [tests/cases/compiler/suspendSuperDefaultParameters.tlua] ////

//// [suspendSuperDefaultParameters.tlua]
class B {
    m() {
        return 1;
    }
}

class C extends B {
    f() {
        local g = suspend function(b = super.m()) return b end;
        return g();
    }

    suspend h(b = super.m()) {
        return b;
    }
}


//// [suspendSuperDefaultParameters.lua]
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
    local g = function(b)
      if b == nil then
        b = super.m();
      end
      return b;
    end;
    return g();
  end
  suspend;
  h(b = super.m());
  do
    return b;
  end
end
