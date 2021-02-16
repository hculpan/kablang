package ast

// Block represents a collection
// of statements executed in sequence
type Block struct {
	StatementsNode *Statements
	Symbols        *SymbolTable
}

// NewBlock ...
func NewBlock(parent *Block) *Block {
	if parent == nil {
		return &Block{Symbols: NewSymbolTable(nil)}
	}
	return &Block{Symbols: NewSymbolTable(parent.Symbols)}
}

// AddSymbol adds a symbol to the internal map
func (b *Block) AddSymbol(s Symbol) {
	if s != nil {
		b.Symbols.Add(s.GetName(), s)
	}
}

// RemoveSymbol removes a symbol from the internal map
func (b *Block) RemoveSymbol(name string) {
	b.Symbols.Delete(name)
}

// AsString return the node as a string
func (b *Block) AsString(indent string) string {
	result := indent + "Block"
	if b.StatementsNode != nil {
		result += "\n" + b.StatementsNode.AsString(indent+"  ")
	}
	return result
}
