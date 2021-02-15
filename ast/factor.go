package ast

import "fmt"

// Factor ...
type Factor struct {
	NumberNode NumberValue
	ParenNode  *NumExpression
}

// NewFactor ...
func NewFactor() *Factor {
	return &Factor{}
}

// AsString returns a string representation of th node
func (f *Factor) AsString(indent string) string {
	result := indent + "Factor"

	if f.NumberNode != nil {
		result += fmt.Sprintf("\n%s", f.NumberNode.AsString("  "+indent))
	} else if f.ParenNode != nil {
		result += fmt.Sprintf("\n%s", f.ParenNode.AsString("  "+indent))
	}

	return result
}
