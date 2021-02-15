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
	IDENTIFIER = iota
	PRINTLN
	PRINT
	VAR
	STRING_TYPE
	NUMBER_TYPE
	FOR
	IF
	ELSE
	INTEGER
	FLOAT
	PERCENT
	DASH
	PLUS
	PLUS_EQUALS
	DOUBLE_PLUS
	MULT
	DIV
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
	NEWLINE
	HASH
	END_TOKEN_LIST
)

// TokenDef contains definition of an
// individual type of token
type TokenDef struct {
	TypeID  int
	Match   string
	Name    string
	Keyword bool
	exp     *regexp.Regexp
}

var tokenDefs []TokenDef = []TokenDef{
	{TypeID: IDENTIFIER, Match: `^[a-zA-Z][a-zA-Z_0-9]*$`, Name: "Identifier"},
	{TypeID: PRINTLN, Match: "^println$", Name: "Println", Keyword: true},
	{TypeID: PRINT, Match: "^print$", Name: "Print", Keyword: true},
	{TypeID: VAR, Match: "^var$", Name: "Var", Keyword: true},
	{TypeID: STRING_TYPE, Match: "^string$", Name: "String", Keyword: true},
	{TypeID: NUMBER_TYPE, Match: "^number$", Name: "Number", Keyword: true},
	{TypeID: FOR, Match: `^for$`, Name: "For", Keyword: true},
	{TypeID: IF, Match: `^if$`, Name: "If", Keyword: true},
	{TypeID: ELSE, Match: `^else$`, Name: "Else", Keyword: true},
	{TypeID: INTEGER, Match: `^[0-9]+$`, Name: "Integer"},
	{TypeID: FLOAT, Match: `^[0-9]+\.[0-9]*$`, Name: "Float"},
	{TypeID: PERCENT, Match: `^%$`, Name: "Percent"},
	{TypeID: DASH, Match: `^-$`, Name: "Dash"},
	{TypeID: PLUS, Match: `^\+$`, Name: "Plus"},
	{TypeID: PLUS_EQUALS, Match: `^\+=$`, Name: "Plus Equals"},
	{TypeID: DOUBLE_PLUS, Match: `^\+\+$`, Name: "Double Plus"},
	{TypeID: MULT, Match: `^\*$`, Name: "Mult"},
	{TypeID: DIV, Match: `^/$`, Name: "Div"},
	{TypeID: EQUALS, Match: `^=$`, Name: "Equals"},
	{TypeID: STRING, Match: `^\".*\"$`, Name: "String"},
	{TypeID: CURLY_BRACE_LEFT, Match: `^\{$`, Name: "Left Curly Brace"},
	{TypeID: CURLY_BRACE_RIGHT, Match: `^\}$`, Name: "Right Curly Brace"},
	{TypeID: PAREN_LEFT, Match: `^\($`, Name: "Left Paren"},
	{TypeID: PAREN_RIGHT, Match: `^\)$`, Name: "Right Parent"},
	{TypeID: LESS_THAN_EQUALS, Match: `^<=$`, Name: "Less Than or Equals"},
	{TypeID: LESS_THAN, Match: `^<$`, Name: "Less Than"},
	{TypeID: GREATER_THAN_EQUALS, Match: `^>=$`, Name: "Greater Than or Equals"},
	{TypeID: GREATER_THAN, Match: `^>$`, Name: "Greater Than"},
	{TypeID: DOUBLE_EQUALS, Match: `^==$`, Name: "Double Equals"},
	{TypeID: NOT, Match: `^!$`, Name: "Not"},
	{TypeID: NOT_EQUALS, Match: `^!=$`, Name: "Not Equals"},
	{TypeID: PERIOD, Match: `^\.$`, Name: "Period"},
	{TypeID: NEWLINE, Match: ``, Name: "Newline"},
	{TypeID: HASH, Match: `^#$`, Name: "Hash"},
	{TypeID: END_TOKEN_LIST, Match: ``, Name: "End of tokens"},
}

// GetTokenDef returns the token definition
// for the specified type id
func GetTokenDef(typeID int) *TokenDef {
	if typeID > END_TOKEN_LIST {
		return nil
	}
	return &tokenDefs[typeID]
}

func (t *TokenDef) compile() {
	if len(t.Match) > 0 {
		t.exp = regexp.MustCompile(t.Match)
	}
}
