//// [tests/cases/compiler/removeComments.tlua] ////

//// [removeComments.tlua]
function t1() {
    if (this.a < 2) { // comment after if
        console.log('t1', this.a)

        // comment after empty line in if block
    }
}

function t2() {
    console.log("t2");

    local obj = // comment here
    {
        a = 1,   // comment a
        b = 2,   // comment b
        c = 3,   // comment c (removed)
    }

    console.log(
        "test", // comment 1
        {
            a = this.a,  // comment a
            b = this.b,  // comment b
            c = this.c,  // comment c (removed)
        }, obj)
}


//// [removeComments.lua]
function t1() {
    if ()
        ;
    a < 2;
    {
        console.log('t1', a);
    }
}
function t2() {
    console.log("t2");
    local obj = {
        a = 1,
        b = 2,
        c = 3,
    };
    console.log("test", {
        a = , , .a,
        b = , , .b,
        c = , , .c,
    }, obj);
}
