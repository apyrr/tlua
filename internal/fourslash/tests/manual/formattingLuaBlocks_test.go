package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

// Lua keyword-blocks (`function`/`do`/`then`/`repeat` ... `end`/`until`) have no
// braces, so the block's opening delimiter -- not a `{` -- is the indentation
// anchor. This locks that the body of every block form indents one level and the
// closing keyword returns to the block's own level.
func TestFormattingLuaBlocks(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `local function f()
local x = 1
return x
end

if a then
x()
elseif b then
y()
else
z()
end

while c do
w()
end

for i = 1, 10 do
print(i)
end

repeat
r()
until done

do
local y = 2
end

local function outer()
do
local z = 3
end
return z
end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.VerifyCurrentFileContent(t, `local function f()
    local x = 1
    return x
end

if a then
    x()
elseif b then
    y()
else
    z()
end

while c do
    w()
end

for i = 1, 10 do
    print(i)
end

repeat
    r()
until done

do
    local y = 2
end

local function outer()
    do
        local z = 3
    end
    return z
end`)
}
