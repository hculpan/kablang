package parser

import (
	"github.com/hculpan/kablang/lexer"
)

// LexerHandler is the interface from the
// parser to the lexer
type LexerHandler struct {
	tokens         []lexer.Token
	Errors         []error
	currTokenIndex int
}

// NewLexerHandler ...
func NewLexerHandler(lines []string, keywords []lexer.TokenDef) *LexerHandler {
	result := &LexerHandler{tokens: []lexer.Token{}, Errors: []error{}, currTokenIndex: 0}

	for i, l := range lines {
		tokens, err := lexer.Lex(l, keywords, i+1)
		if err != nil {
			result.Errors = append(result.Errors, err)
		}
		tokens = append(tokens, *lexer.NewToken(lexer.NEWLINE, "\n", "Newline", i+1, len(l)))
		result.tokens = append(result.tokens, tokens...)
	}
	result.tokens = append(result.tokens, *lexer.NewToken(lexer.END_TOKEN_LIST, "END_TOKENS", "End of tokens", len(lines), 0))

	return result
}

// Pop returns the next token and increments the
// index
func (l *LexerHandler) Pop() lexer.Token {
	result := l.tokens[l.currTokenIndex]
	if result.TypeID != lexer.END_TOKEN_LIST {
		l.currTokenIndex++
	}
	return result
}

// Peek looks at the next token, but does not
// increment the index
func (l *LexerHandler) Peek() lexer.Token {
	return l.tokens[l.currTokenIndex]
}

// Push restores the last token popped
func (l *LexerHandler) Push() lexer.Token {
	l.currTokenIndex--
	if l.currTokenIndex < 0 {
		l.currTokenIndex = 0
	}
	return l.tokens[l.currTokenIndex]
}

// Swallow will consume the next token if it
// matches the expected type
func (l *LexerHandler) Swallow(typeID int) bool {
	t := l.Pop()
	if t.TypeID != typeID {
		l.Push()
		return false
	}
	return true
}
