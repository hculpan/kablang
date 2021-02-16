package ast

// VarStatement is a "println" statement
type VarStatement struct {
	SymbolNode     Symbol
	ExpressionNode Expression
}

// NewVarStatement ...
func NewVarStatement(name string, typeID int) *VarStatement {
	return &VarStatement{SymbolNode: NewSymbol(name, typeID)}
}

// AsString return the node as a string
func (s VarStatement) AsString(indent string) string {
	result := indent + "VarStatement : " + s.SymbolNode.AsString("")

	if s.ExpressionNode != nil {
		result += "\n" + "  " + indent + "=\n" + s.ExpressionNode.AsString("  "+indent)
	}

	return result
}
