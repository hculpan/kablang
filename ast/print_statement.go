package ast

// PrintStatement is a "println" statement
type PrintStatement struct {
	StringExpressionNode *StringExpression
	NumExpressionNode    *NumExpression

	ExpressionTypeID int

	WithEndline bool
}

// NewStringPrintStatement ...
func NewStringPrintStatement(s *StringExpression, endline bool) *PrintStatement {
	return &PrintStatement{StringExpressionNode: s, ExpressionTypeID: StringExpressionType, WithEndline: endline}
}

// NewNumPrintStatement ...
func NewNumPrintStatement(n *NumExpression, endline bool) *PrintStatement {
	return &PrintStatement{NumExpressionNode: n, ExpressionTypeID: NumExpressionType, WithEndline: endline}
}

// NewEmptyPrintStatement ...
func NewEmptyPrintStatement(endline bool) *PrintStatement {
	return &PrintStatement{NumExpressionNode: nil, ExpressionTypeID: EmptyExpressionType, WithEndline: endline}
}

// AsString return the node as a string
func (s PrintStatement) AsString(indent string) string {
	var result string
	if s.WithEndline {
		result = indent + "PrintlnStatement"
	} else {
		result = indent + "PrintStatement"
	}

	switch s.ExpressionTypeID {
	case StringExpressionType:
		if s.StringExpressionNode != nil {
			result += "\n" + s.StringExpressionNode.AsString("  "+indent)
		}
	case NumExpressionType:
		if s.NumExpressionNode != nil {
			result += "\n" + s.NumExpressionNode.AsString("  "+indent)
		}
	}

	return result
}
