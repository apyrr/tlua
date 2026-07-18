//// [tests/cases/compiler/jsxNestedIndentation.tsx] ////

//// [jsxNestedIndentation.tsx]
declare React: any;
declare function Child(props: { children?: any }): any;
function Test() {
    return <Child>
        <Child>
            <Child></Child>
        </Child>
    </Child>
}


//// [jsxNestedIndentation.lua]
function Test() {
    return React.createElement(Child, nil,
        React.createElement(Child, nil,
            React.createElement(Child, nil)));
}
