//// [tests/cases/compiler/classExpressionWithComputedPropertyInLoop.tlua] ////

//// [classExpressionWithComputedPropertyInLoop.tlua]
// Class expression in a loop with a computed property name on an instance field.
// The class temp variable should be block-scoped (let) to ensure each iteration
// gets its own binding, matching the behavior when BlockScopedBindingInLoop is set.

local array: any[] = [];
local key = "myKey";
for i = 0, 2 do
    array.push(class C {
        [key] = i;
        #field = i;
    });
end


//// [classExpressionWithComputedPropertyInLoop.lua]
-- Class expression in a loop with a computed property name on an instance field.
-- The class temp variable should be block-scoped (let) to ensure each iteration
-- gets its own binding, matching the behavior when BlockScopedBindingInLoop is set.
local array = [];
local key = "myKey";
for i = 0, 2 do
    array.push(class, C, {
        [key] = i,
        #field, i,
    });
end
