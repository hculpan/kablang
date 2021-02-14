package ast

// Type of expressions
const (
	StringExpressionType = iota
	NumExpressionType
	EmptyExpressionType
)

// Expression interface represents a generic
// expression
type Expression interface {
	AsString(indent string) string
}
