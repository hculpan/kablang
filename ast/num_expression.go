package ast

import "fmt"

// NumExpression ...
type NumExpression struct {
	numNode           SignedNumber
	numExpressionNode Expression
}

// NewNumExpression ...
func NewNumExpression(s SignedNumber, exp Expression) (NumExpression, error) {
	result := NumExpression{numNode: s, numExpressionNode: nil}

	switch exp.(type) {
	case NumExpression:
		result.numExpressionNode = exp
	default:
		return result, fmt.Errorf("Expecting number expression, found type %T", exp)
	}

	return result, nil
}

// AsString return the node as a string
func (s NumExpression) AsString(indent string) string {
	result := indent + "NumExpression"
	return result
}
