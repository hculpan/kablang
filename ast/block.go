package ast

// Block represents a collection
// of statements executed in sequence
type Block struct {
	StatementsNode *Statements
}

// NewBlock ...
func NewBlock(stmts *Statements) *Block {
	return &Block{StatementsNode: stmts}
}

// AsString return the node as a string
func (b *Block) AsString(indent string) string {
	result := indent + "Block"
	if b.StatementsNode != nil {
		result += "\n" + b.StatementsNode.AsString(indent+"  ")
	}
	return result
}
