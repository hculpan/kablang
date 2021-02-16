package ast

// BlockStack is a LIFO stack
// for statement blocks, used by
// both the parser and executor
type BlockStack struct {
	blocks []*Block
}

// NewBlockStack ...
func NewBlockStack() *BlockStack {
	return &BlockStack{blocks: make([]*Block, 30)}
}

// Pop returns the top block and removes
// it from the stack
func (b *BlockStack) Pop() *Block {
	var result *Block = nil
	if len(b.blocks) > 0 {
		result = b.blocks[len(b.blocks)-1]
		b.blocks = b.blocks[:len(b.blocks)-1]
	}
	return result
}

// Peek returns the top block without
// removing it from the stack
func (b *BlockStack) Peek() *Block {
	var result *Block = nil
	if len(b.blocks) > 0 {
		result = b.blocks[len(b.blocks)-1]
	}
	return result
}

// Push adds a block to the stack
func (b *BlockStack) Push(block *Block) {
	b.blocks = append(b.blocks, block)
}
