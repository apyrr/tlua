//// [tests/cases/conformance/ported/computedPropertyNames35_ES6.tlua] ////

//// [computedPropertyNames35_ES6.tlua]
-- ported from tests/cases/conformance/es6/computedProperties/computedPropertyNames35_ES6.ts
-- dropped: @target: es6 directive (tlua defaults to latest target)

function foo<T>() return '' end
interface I<T> {
    bar(): string
    [foo<T>()](): void
}


//// [computedPropertyNames35_ES6.lua]
-- ported from tests/cases/conformance/es6/computedProperties/computedPropertyNames35_ES6.ts
-- dropped: @target: es6 directive (tlua defaults to latest target)
function foo()
  return '';
end
