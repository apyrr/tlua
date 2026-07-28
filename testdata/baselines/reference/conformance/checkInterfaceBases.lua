//// [tests/cases/conformance/ported/checkInterfaceBases.tlua] ////

//// [jquery.d.tlua]
interface JQueryEventObjectTest {
    data: any
    which: number
    metaKey: any
}

//// [app.tlua]
///<reference path='jquery.d.tlua' />
interface SecondEvent {
    data: any
}
interface Third extends JQueryEventObjectTest, SecondEvent {}


//// [app.lua]
-- /<reference path='jquery.d.tlua' />
