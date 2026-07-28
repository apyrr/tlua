//// [tests/cases/conformance/ported/recurringTypeParamForContainerOfBase01.tlua] ////

//// [recurringTypeParamForContainerOfBase01.tlua]
-- ported from tests/cases/conformance/types/typeParameters/recurringTypeParamForContainerOfBase01.ts
-- dropped: @target: es2015 and @declaration directives (tlua uses the latest target and does not compare declaration emit)

interface BoxOfFoo<T extends Foo<T>> {
    item: T
}

interface Foo<T extends Foo<T>> {
    self: T
}

interface Bar<T extends Bar<T>> extends Foo<T> {
    other: BoxOfFoo<T>
}


//// [recurringTypeParamForContainerOfBase01.lua]
-- ported from tests/cases/conformance/types/typeParameters/recurringTypeParamForContainerOfBase01.ts
-- dropped: @target: es2015 and @declaration directives (tlua uses the latest target and does not compare declaration emit)
