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
{
    #x;
    3;
    #y;
    nil;
    func();
    {
        console.log(#y ("->>" .. tostring()), #x);
    }
    -");\n  }\n}\n";
}
