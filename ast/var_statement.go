package ast

// VarStatement is a "println" statement
type VarStatement struct {
	SymbolNode Symbol
}

// NewVarStatement ...
func NewVarStatement(name string, typeID int) *VarStatement {
	return &VarStatement{SymbolNode: NewSymbol(name, typeID)}
}

// AsString return the node as a string
func (s VarStatement) AsString(indent string) string {
	return indent + "VarStatement : " + s.SymbolNode.AsString("")
}
