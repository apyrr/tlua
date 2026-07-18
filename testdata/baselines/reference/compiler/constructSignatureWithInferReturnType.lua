//// [tests/cases/compiler/constructSignatureWithInferReturnType.tlua] ////

//// [constructSignatureWithInferReturnType.tlua]
type ExtractReturn<T> = T extends { new(): infer R } ? R : never;


//// [constructSignatureWithInferReturnType.lua]
