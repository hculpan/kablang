package lexer

import "regexp"

/****************************************************************
*
* To add a new token:
*   1) Add constant to bottom of constants list
*   2) If it's a keyword, add to keywords below
*      Otheriwse add definition to tokenDefs
*        Make sure the regex is unique from all other tokens
*
****************************************************************/

// Token constants
const (
	INTEGER = iota
	FLOAT
	PLUS
	PLUS_EQUALS
	DOUBLE_PLUS
	IDENTIFIER
	EQUALS
	STRING
	CURLY_BRACE_LEFT
	CURLY_BRACE_RIGHT
	PAREN_LEFT
	PAREN_RIGHT
	LESS_THAN_EQUALS
	LESS_THAN
	GREATER_THAN_EQUALS
	GREATER_THAN
	DOUBLE_EQUALS
	NOT
	NOT_EQUALS
	PERIOD
	END_TOKEN_LIST
)

// TokenDef contains definition of an
// individual type of token
type TokenDef struct {
	TypeID int
	Match  string
	Name   string
	exp    *regexp.Regexp
}

var tokenDefs []TokenDef = []TokenDef{
	{TypeID: INTEGER, Match: `^[0-9]+$`, Name: "Integer"},
	{TypeID: FLOAT, Match: `^[0-9]+\.[0-9]*$`, Name: "Float"},
	{TypeID: PLUS, Match: `^\+$`, Name: "Plus"},
	{TypeID: EQUALS, Match: `^=$`, Name: "Equals"},
	{TypeID: DOUBLE_EQUALS, Match: `^==$`, Name: "Double Equals"},
	{TypeID: PLUS_EQUALS, Match: `^\+=$`, Name: "Plus Equals"},
	{TypeID: DOUBLE_PLUS, Match: `^\+\+$`, Name: "Double Plus"},
	{TypeID: NOT, Match: `^!$`, Name: "Not"},
	{TypeID: NOT_EQUALS, Match: `^!=$`, Name: "Not Equals"},
	{TypeID: IDENTIFIER, Match: `^[a-zA-Z][a-zA-Z_0-9]*$`, Name: "Identifier"},
	{TypeID: STRING, Match: `^\".*\"$`, Name: "String"},
	{TypeID: CURLY_BRACE_LEFT, Match: `^{$`, Name: "Left Curly Brace"},
	{TypeID: CURLY_BRACE_RIGHT, Match: `^}$`, Name: "Right Curly Brace"},
	{TypeID: PAREN_LEFT, Match: `^\($`, Name: "Left Paren"},
	{TypeID: PAREN_RIGHT, Match: `^\)$`, Name: "Right Parent"},
	{TypeID: LESS_THAN, Match: `^<$`, Name: "Less Than"},
	{TypeID: LESS_THAN_EQUALS, Match: `^<=$`, Name: "Less Than or Equals"},
	{TypeID: GREATER_THAN, Match: `^>$`, Name: "Greater Than"},
	{TypeID: GREATER_THAN_EQUALS, Match: `^>=$`, Name: "Greater Than or Equals"},
	{TypeID: PERIOD, Match: `^\.$`, Name: "Period"},
}

func (t *TokenDef) compile() {
	t.exp = regexp.MustCompile(t.Match)
}
