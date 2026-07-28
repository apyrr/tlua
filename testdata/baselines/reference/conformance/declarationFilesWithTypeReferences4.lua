//// [tests/cases/conformance/ported/declarationFilesWithTypeReferences4.tlua] ////

//// [index.d.tlua]
interface Error {
    stack2: string
}

//// [app.tlua]
/// <reference types="node"/>
function foo(): Error
    return nil as any
end


//// [app.lua]
-- / <reference types="node"/>
function foo()
  return nil;
end
