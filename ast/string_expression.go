package ast

// StringExpression ...
type StringExpression struct {
	StringNode           StringValue
	StringExpressionNode Expression
}

// NewStringExpression ...
func NewStringExpression() *StringExpression {
	return &StringExpression{}
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
