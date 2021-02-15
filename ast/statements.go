package ast

// Statements represents a series of statements
type Statements struct {
	StatementListNode []Statement
}

// NewStatements ...
func NewStatements(stmts []Statement) *Statements {
	return &Statements{StatementListNode: stmts}
}

// AsString return the node as a string
func (s *Statements) AsString(indent string) string {
	result := indent + "Statements"
	if s.StatementListNode != nil {
		for _, v := range s.StatementListNode {
			if v != nil {
				result += "\n" + v.AsString(indent+"  ")
			}
		}
	}
	return result
}
