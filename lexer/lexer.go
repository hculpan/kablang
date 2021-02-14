package lexer

import (
	"fmt"
	"unicode"
)

// Token represents a token within a string
type Token struct {
	TypeID int
	Value  string
	Name   string
}

// NewToken ...
func NewToken(typeID int, v string, name string) *Token {
	return &Token{Value: v, TypeID: typeID, Name: name}
}

// Lex is the main entry point into the lexer.
// Returns a list of all the tokens within the
// given string.  This is meant to be a more generic
// lexer, thus the keywords are passed in.
func Lex(s string, keywords []TokenDef) ([]Token, error) {
	initializeLexer()

	result := []Token{}

	currLoc := 0
	currentBuffer := ""
	var lastTokenSelection *Token = nil

	for currLoc < len(s) {
		currentBuffer += string(s[currLoc])

		// First let's trim out any whitespace
		if len(currentBuffer) == 1 && unicode.IsSpace(rune(currentBuffer[0])) {
			currLoc++
			currentBuffer = ""
			continue
		}

		currentSelection := reduceSelection(currentBuffer, tokenDefs)
		//fmt.Printf("Buffer=%s, selection=%v\n", currentBuffer, currentSelection)

		switch {
		case len(currentSelection) == 1:
			lastTokenSelection = NewToken(currentSelection[0].TypeID, currentBuffer, currentSelection[0].Name)
		case len(currentSelection) == 0 && lastTokenSelection != nil:
			result = append(result, *lastTokenSelection)
			lastTokenSelection = nil
			currentBuffer = ""
			currLoc--
		}
		currLoc++
	}

	if lastTokenSelection != nil {
		result = append(result, *lastTokenSelection)
	} else if len(currentBuffer) > 0 {
		return result, fmt.Errorf("Unknown token: %s", currentBuffer)
	}

	// A bit of a kludge; keywords will get set as identifiers
	// initally, now go through them to see if any are actually
	// keywords
	result = checkForKeywords(result, keywords)

	return result, nil
}

func findKeywordMatch(s string, keywords []TokenDef) *TokenDef {
	for _, k := range keywords {
		if s == k.Match {
			return &k
		}
	}

	return nil
}

func checkForKeywords(tokens []Token, keywords []TokenDef) []Token {
	result := []Token{}
	for _, v := range tokens {
		if v.TypeID == IDENTIFIER {
			if t := findKeywordMatch(v.Value, keywords); t != nil {
				v.TypeID = t.TypeID
				v.Name = "Keyword"
			}
		}
		result = append(result, v)
	}
	return result
}

func initializeLexer() {
	for i := range tokenDefs {
		tokenDefs[i].compile()
	}
}

func reduceSelection(s string, currentSelection []TokenDef) []TokenDef {
	result := []TokenDef{}
	for _, t := range currentSelection {
		r := t.exp.FindStringIndex(s)
		if len(r) > 0 && r[0] == 0 {
			result = append(result, t)
		}
	}
	return result
}

// Equals returns whether the two tokens are equal
func (t Token) Equals(t2 Token) bool {
	return t.TypeID == t2.TypeID && t.Value == t2.Value
}