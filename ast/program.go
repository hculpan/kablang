package ast

// Program is the top-level entry
// into a program
type Program struct {
	BlockNode *Block
}

// NewProgram ...
func NewProgram(blockNode *Block) *Program {
	return &Program{BlockNode: blockNode}
}

// AsString return the node as a string
func (p *Program) AsString(indent string) string {
	result := indent + "Program"
	if p.BlockNode != nil {
		result += "\n" + p.BlockNode.AsString(indent+"  ")
	}
	return result
}
