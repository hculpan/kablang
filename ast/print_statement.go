package ast

// PrintStatement is a "println" statement
type PrintStatement struct {
	StringExpressionNode *StringExpression
	NumExpressionNode    *NumExpression

	ExpressionTypeID int
}

// NewStringPrintStatement ...
func NewStringPrintStatement(s *StringExpression) *PrintStatement {
	return &PrintStatement{StringExpressionNode: s, ExpressionTypeID: StringExpressionType}
}

// NewNumPrintStatement ...
func NewNumPrintStatement(n *NumExpression) PrintStatement {
	return PrintStatement{NumExpressionNode: n, ExpressionTypeID: NumExpressionType}
}

// AsString return the node as a string
func (s PrintStatement) AsString(indent string) string {
	result := indent + "PrintStatement"

	switch s.ExpressionTypeID {
	case StringExpressionType:
		result += "\n" + s.StringExpressionNode.AsString("  "+indent)
	case NumExpressionType:
		result += "\n" + s.NumExpressionNode.AsString("  "+indent)
	}

	return result
}
