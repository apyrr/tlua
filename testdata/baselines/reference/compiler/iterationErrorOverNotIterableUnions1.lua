//// [tests/cases/compiler/iterationErrorOverNotIterableUnions1.tlua] ////

//// [iterationErrorOverNotIterableUnions1.tlua]
type A = { a: string };
type B = { b: string };

declare data: A[] | B;

for _, item in ipairs(data) do
    item.b;
end

for _, ignoredItem in ipairs(data) do
    ignoredItem.b;
end

local [el] = data;
el.b;

local el2 = data[0];
el2.b;


//// [iterationErrorOverNotIterableUnions1.lua]
for _, item in ipairs(data) do
    item.b;
end
for _, ignoredItem in ipairs(data) do
    ignoredItem.b;
end
local ;
el;
data;
el.b;
local el2 = data[0];
el2.b;
