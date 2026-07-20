//// [tests/cases/compiler/sliceTupleTypeOutOfBounds.tlua] ////

//// [sliceTupleTypeOutOfBounds.tlua]
type Middle<T> = T extends [unknown, ... infer X, unknown] ? X: never;
type Example = Middle<[1]>;


//// [sliceTupleTypeOutOfBounds.lua]
