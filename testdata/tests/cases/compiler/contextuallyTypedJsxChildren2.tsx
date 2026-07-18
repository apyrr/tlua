// @target: es2015
// @strict: true
// @jsx: react
// @noEmit: true

// https://github.com/microsoft/typescript-go/issues/2802
// Children function parameters must be contextually typed through NoInfer.

interface JsxElement {}
interface JsxElementChildrenAttribute { children: {}; }
interface JsxIntrinsicElements { div: { children?: any }; }
declare React: { createElement: any, Fragment: any };

declare TestComponentWithChildren: <T, TParam>(props: {
  state: T;
  selector?: (state: NoInfer<T>) => TParam;
  children?: (state: NoInfer<TParam>) => JsxElement | null;
}) => JsxElement;

declare TestComponentWithoutChildren: <T, TParam>(props: {
  state: T;
  selector?: (state: NoInfer<T>) => TParam;
  notChildren?: (state: NoInfer<TParam>) => JsxElement | null;
}) => JsxElement;

local App = () => {
  return (
    <>
      <TestComponentWithChildren state={{ foo = 123 }} selector={(state) => state.foo}>
        {(selected) => <div>{Math.max(selected, 0)}</div>}
      </TestComponentWithChildren>

      <TestComponentWithoutChildren
        state={{ foo = 123 }}
        selector={(state) => state.foo}
        notChildren={(selected) => <div>{Math.max(selected, 0)}</div>}
      />
    </>
  );
};
