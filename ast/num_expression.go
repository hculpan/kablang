package ast

import "fmt"

// List the posible operations
// allowable in a numeric expression
const (
	NoOperator = iota
	PlusOperator
	MinusOperator
	MultOperator
	DivOperator
	ModuloOperator
)

// NumExpression ...
type NumExpression struct {
	TermNode          *Term
	Operator          int
	NumExpressionNode Expression
}

// NewNumExpression ...
func NewNumExpression(t *Term, exp Expression) (*NumExpression, error) {
	result := &NumExpression{TermNode: t, NumExpressionNode: nil, Operator: NoOperator}

	switch exp.(type) {
	case NumExpression:
		result.NumExpressionNode = exp
	default:
		return result, fmt.Errorf("Expecting number expression, found type %T", exp)
	}

	return result, nil
}

// AsString return the node as a string
func (s NumExpression) AsString(indent string) string {
	result := indent + "NumExpression"

	if s.TermNode != nil {
		result += fmt.Sprintf("\n%s", s.TermNode.AsString("  "+indent))
	}

	if s.Operator == PlusOperator {
		result += "\n" + indent + "    +"
		result += fmt.Sprintf("\n%s", s.NumExpressionNode.AsString("    "+indent))
	} else if s.Operator == MinusOperator {
		result += "\n" + indent + "    -"
		result += fmt.Sprintf("\n%s", s.NumExpressionNode.AsString("    "+indent))
	}

	return result
}
