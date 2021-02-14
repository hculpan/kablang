package ast

// Statement is the interface for AST
// statements
type Statement interface {
	AsString(indent string) string
}
