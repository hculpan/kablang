package ast

// PrintlnStatement is a "println" statement
type PrintlnStatement struct {
	StringExpressionNode *StringExpression
	NumExpressionNode    *NumExpression

	ExpressionTypeID int
}

// NewStringPrintlnStatement ...
func NewStringPrintlnStatement(s *StringExpression) *PrintlnStatement {
	return &PrintlnStatement{StringExpressionNode: s, ExpressionTypeID: StringExpressionType}
}

// NewNumPrintlnStatement ...
func NewNumPrintlnStatement(n *NumExpression) *PrintlnStatement {
	return &PrintlnStatement{NumExpressionNode: n, ExpressionTypeID: NumExpressionType}
}

// NewEmptyPrintlnStatement ...
func NewEmptyPrintlnStatement() *PrintlnStatement {
	return &PrintlnStatement{ExpressionTypeID: EmptyExpressionType}
}

// AsString return the node as a string
func (s *PrintlnStatement) AsString(indent string) string {
	result := indent + "PrintlnStatement"

	switch s.ExpressionTypeID {
	case StringExpressionType:
		result += "\n" + s.StringExpressionNode.AsString("  "+indent)
	case NumExpressionType:
		result += "\n" + s.NumExpressionNode.AsString("  "+indent)
	}

	return result
}
