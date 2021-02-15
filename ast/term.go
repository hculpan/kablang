package ast

import "fmt"

// Term ...
type Term struct {
	FactorNode *Factor
	Operator   int
	TermNode   *Term
}

// NewTerm ...
func NewTerm() *Term {
	return &Term{Operator: NoOperator}
}

// AsString returns a string representation of the node
func (t *Term) AsString(indent string) string {
	result := indent + "Term"

	if t.FactorNode != nil {
		result += fmt.Sprintf("\n%s", t.FactorNode.AsString("  "+indent))
	}

	if t.Operator == MultOperator {
		result += "\n" + indent + "    *"
		result += fmt.Sprintf("\n%s", t.TermNode.AsString("    "+indent))
	} else if t.Operator == DivOperator {
		result += "\n" + indent + "    /"
		result += fmt.Sprintf("\n%s", t.TermNode.AsString("    "+indent))
	}

	return result
}
