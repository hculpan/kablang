package ast

// Block represents a collection
// of statements executed in sequence
type Block struct {
	StatementsNode *Statements
	Symbols        map[string]Symbol
}

// NewBlock ...
func NewBlock() *Block {
	return &Block{Symbols: map[string]Symbol{}}
}

// AddSymbol adds a symbol to the internal map
func (b *Block) AddSymbol(s Symbol) {
	if s != nil {
		b.Symbols[s.GetName()] = s
	}
}

// RemoveSymbol removes a symbol from the internal map
func (b *Block) RemoveSymbol(name string) {
	delete(b.Symbols, name)
}

// AsString return the node as a string
func (b *Block) AsString(indent string) string {
	result := indent + "Block"
	if b.StatementsNode != nil {
		result += "\n" + b.StatementsNode.AsString(indent+"  ")
	}
	return result
}
