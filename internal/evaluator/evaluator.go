package evaluator

import (
	"strings"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/jsnum"
)

type Evaluator func(expr *ast.Node, location *ast.Node) any

func NewEvaluator(evaluateEntity Evaluator, outerExpressionsToSkip ast.OuterExpressionKinds) Evaluator {
	var evaluate Evaluator
	evaluate = func(expr *ast.Node, location *ast.Node) any {
		// Type assertions intentionally break literal evaluation.
		expr = ast.SkipOuterExpressions(expr, outerExpressionsToSkip|ast.OEKParentheses)
		switch expr.Kind {
		case ast.KindPrefixUnaryExpression:
			if value, ok := evaluate(expr.AsPrefixUnaryExpression().Operand, location).(jsnum.Number); ok {
				switch expr.AsPrefixUnaryExpression().Operator {
				case ast.KindPlusToken:
					return value
				case ast.KindMinusToken:
					return -value
				}
			}
		case ast.KindBinaryExpression:
			left := evaluate(expr.AsBinaryExpression().Left, location)
			right := evaluate(expr.AsBinaryExpression().Right, location)
			operator := expr.AsBinaryExpression().OperatorToken.Kind
			leftNum, leftIsNum := left.(jsnum.Number)
			rightNum, rightIsNum := right.(jsnum.Number)
			if leftIsNum && rightIsNum {
				switch operator {
				case ast.KindAsteriskToken:
					return leftNum * rightNum
				case ast.KindSlashToken:
					return leftNum / rightNum
				case ast.KindPlusToken:
					return leftNum + rightNum
				case ast.KindMinusToken:
					return leftNum - rightNum
				case ast.KindPercentToken:
					return leftNum.Remainder(rightNum)
				case ast.KindAsteriskAsteriskToken:
					return leftNum.Exponentiate(rightNum)
				}
			}
			leftStr, leftIsStr := left.(string)
			rightStr, rightIsStr := right.(string)
			if (leftIsStr || leftIsNum) && (rightIsStr || rightIsNum) && operator == ast.KindPlusToken {
				if leftIsNum {
					leftStr = leftNum.String()
				}
				if rightIsNum {
					rightStr = rightNum.String()
				}
				return leftStr + rightStr
			}
		case ast.KindStringLiteral, ast.KindNoSubstitutionTemplateLiteral:
			return expr.Text()
		case ast.KindTemplateExpression:
			return evaluateTemplateExpression(expr, location, evaluate)
		case ast.KindNumericLiteral:
			return jsnum.FromString(expr.Text())
		case ast.KindIdentifier:
			return evaluateEntity(expr, location)
		}
		return nil
	}
	return evaluate
}

func evaluateTemplateExpression(expr *ast.Node, location *ast.Node, evaluate Evaluator) any {
	var sb strings.Builder
	sb.WriteString(expr.AsTemplateExpression().Head.Text())
	for _, span := range expr.AsTemplateExpression().TemplateSpans.Nodes {
		spanResult := evaluate(span.Expression(), location)
		if spanResult == nil {
			return nil
		}
		sb.WriteString(AnyToString(spanResult))
		sb.WriteString(span.AsTemplateSpan().Literal.Text())
	}
	return sb.String()
}

func AnyToString(v any) string {
	switch v := v.(type) {
	case string:
		return v
	case jsnum.Number:
		return v.String()
	case bool:
		return core.IfElse(v, "true", "false")
	}
	panic("Unhandled case in AnyToString")
}
