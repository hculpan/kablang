package ast

// SymbolTable contains a table of symbols
type SymbolTable struct {
	symbols map[string]Symbol
	parent  *SymbolTable
}

// NewSymbolTable ...
func NewSymbolTable(parent *SymbolTable) *SymbolTable {
	return &SymbolTable{parent: parent, symbols: make(map[string]Symbol, 50)}
}

// GetSymbols provides direct access to the internal map.
// This is here so that I can output the symbols from the
// main package
func (s *SymbolTable) GetSymbols() map[string]Symbol {
	return s.symbols
}

// Add adds a symbol to the table
func (s *SymbolTable) Add(name string, symbol Symbol) {
	s.symbols[name] = symbol
}

// ExistsLocal checks if a symbol exists in current symbol table
func (s *SymbolTable) ExistsLocal(name string) bool {
	_, exists := s.symbols[name]
	return exists
}

// Exists checks if symbol exists in any symbol table
func (s *SymbolTable) Exists(name string) bool {
	if !s.ExistsLocal(name) && s.parent != nil {
		return s.parent.Exists(name)
	}

	return true
}

// GetLocal retrieves the symbol from the local symbol
// table only
func (s *SymbolTable) GetLocal(name string) (Symbol, bool) {
	if s.ExistsLocal(name) {
		return s.symbols[name], true
	}

	return nil, false
}

// Get returns the first occurence of this symbol,
// starting with local and proceeding up the parent chain
func (s *SymbolTable) Get(name string) (Symbol, bool) {
	symbol, exists := s.GetLocal(name)

	if !exists && s.parent != nil {
		symbol, exists = s.parent.Get(name)
	}

	return symbol, exists
}

// Delete removes a symbol from the local symbol
// table only
func (s *SymbolTable) Delete(name string) bool {
	if s.ExistsLocal(name) {
		delete(s.symbols, name)
		return true
	}

	return false
}
