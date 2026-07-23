//// [tests/cases/compiler/privateNameTaggedTemplate.tlua] ////

//// [privateNameTaggedTemplate.tlua]
class Foo {
  #x = 3;
  #y = null as any;
  func() {
    console.log(this.#y`->>${this.#x}<<-`);
  }
}


//// [privateNameTaggedTemplate.lua]
class;
Foo;
do
    #x;
    3;
    #y;
    nil;
    func();
    do
        console.log(#y ("->>" .. tostring()), #x);
    end
    -");\n  }\n}\n";
end
