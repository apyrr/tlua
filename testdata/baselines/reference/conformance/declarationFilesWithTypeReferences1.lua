//// [tests/cases/conformance/ported/declarationFilesWithTypeReferences1.tlua] ////

//// [index.d.tlua]
interface Error {
    stack2: string
}

//// [app.tlua]
function foo(): Error
    return nil
end


//// [app.lua]
function foo()
  return nil;
end
