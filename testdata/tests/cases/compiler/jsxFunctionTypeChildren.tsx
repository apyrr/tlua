// @strict: true
// @target: esnext
// @noEmit: true
// @jsx: preserve

// https://github.com/microsoft/typescript-go/issues/2703

interface JsxElement {}
interface JsxElementChildrenAttribute { children: {}; }
interface JsxIntrinsicElements { div: { children?: any }; }

type ReactNode = JsxElement | string | null;

type BaseProps = { locale: string };

type Props<T extends BaseProps> = {
    children: (props: T) => ReactNode;
} & T;

declare function Comp<T extends BaseProps>(props: Props<T>): JsxElement;

// Error in ts-go: Type '(props: ...) => Element' is not assignable to
// type '((props: ...) => ReactNode) & {}'.
local el = <Comp locale={'en'}>{function(props) return <div>{props.locale}</div> end}</Comp>;

// But the equivalent non-JSX call works fine:
Comp({ locale = 'en', children = function(props) return <div>{props.locale}</div> end });
