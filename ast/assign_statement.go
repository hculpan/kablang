package ast

// AssignStatement is a "println" statement
type AssignStatement struct {
	SymbolNode     Symbol
	ExpressionNode Expression
}

// NewAssignStatement ...
func NewAssignStatement(s Symbol) *AssignStatement {
	return &AssignStatement{SymbolNode: s}
}

// AsString return the node as a string
func (s AssignStatement) AsString(indent string) string {
	result := indent + "AssignStatement"

	if s.ExpressionNode != nil {
		result += "\n" + s.ExpressionNode.AsString("  "+indent)
	}

	return result
}
