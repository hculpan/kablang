package lexer

import (
	"fmt"
	"sort"
)

// Token represents a token within a string
type Token struct {
	TypeID TokenType
	Value  string
	Name   string

	Line int
	Col  int
}

// NewToken ...
func NewToken(typeID TokenType, v string, name string, line int, col int) *Token {
	return &Token{Value: v, TypeID: typeID, Name: name, Line: line, Col: col}
}

var initalized bool = false

// Lex is the main entry point into the lexer.
// Returns a list of all the tokens within the
// given string.  This is meant to be a more generic
// lexer, thus the keywords are passed in.
func Lex(s string, currLine int) ([]Token, error) {
	initTokenDefinitions()

	result := []Token{}

	currLoc := 0
	for currLoc < len(s) {
		if s[currLoc] == ' ' || s[currLoc] == '\t' {
			currLoc++
			continue
		}

		found := false
		for _, t := range tokenDefs {
			if t.exp == nil {
				continue
			}

			if i := t.exp.FindStringIndex(s[currLoc:]); i != nil {
				result = append(result, *NewToken(t.TypeID, s[currLoc:currLoc+i[1]], t.Name, currLine, currLoc+1))
				currLoc += i[1]
				found = true
			}
		}

		if !found {
			return result, fmt.Errorf("No token match for '%s'", s[currLoc:])
		}
	}

	return result, nil
}

// initTokenDefinitions must be called before the
// lexer is used.  It sorts
func initTokenDefinitions() {
	if !initalized {
		sort.Sort(byPriority(tokenDefs))
		initalized = true
	}
}

// Equals returns whether the two tokens are equal
func (t Token) Equals(t2 Token) bool {
	if t.TypeID == Newline {
		return t.TypeID == t2.TypeID
	}

	return t.TypeID == t2.TypeID && t.Value == t2.Value
}
