package ast

// NullStatement represents an empty statement
type NullStatement struct {
}

// NewNullStatement ...
func NewNullStatement() *NullStatement {
	return &NullStatement{}
}

// AsString return the node as a string
func (s NullStatement) AsString(indent string) string {
	return indent + "NullStatement"
}
