package ast

import "fmt"

// StringExpression ...
type StringExpression struct {
	StringNode           *String
	StringExpressionNode Expression
}

// NewStringExpression ...
func NewStringExpression(s *String, exp Expression) (*StringExpression, error) {
	result := StringExpression{StringNode: s, StringExpressionNode: nil}

	switch exp.(type) {
	case StringExpression:
		result.StringExpressionNode = exp
	case nil:
		result.StringExpressionNode = nil
	default:
		return &result, fmt.Errorf("Expecting string expression, found type %T", exp)
	}

	return &result, nil
}

// AsString return the node as a string
func (s StringExpression) AsString(indent string) string {
	result := indent + "StringExpression"

	if s.StringNode != nil {
		result += "\n" + s.StringNode.AsString("  "+indent)
	}

	if s.StringExpressionNode != nil {
		result += "\n  " + indent + "+\n" + s.StringExpressionNode.AsString("  "+indent)
	}

	return result
}
