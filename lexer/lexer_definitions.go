package lexer

import (
	"regexp"
)

/****************************************************************
*
* To add a new token:
*   1) Add constant to bottom of constants list
*   2) If it's a keyword, add to keywords below
*      Otheriwse add definition to tokenDefs
*        Make sure the regex is unique from all other tokens
*
****************************************************************/

// TokenType is the id for the types of tokens
type TokenType int

// Token constants
const (
	Identifier TokenType = iota
	Println
	Print
	Var
	StringType
	NumberType
	For
	If
	Else
	Integer
	Float
	Percent
	Dash
	Plus
	PlusEquals
	DoublePlus
	Mult
	Div
	Equals
	String
	LeftCurlyBrace
	RightCurlyBrace
	LeftParen
	RightParen
	LessThanEquals
	LessThan
	GreaterThanEquals
	GreaterThan
	DoubleEquals
	Not
	NotEquals
	Period
	Newline
	Hash
	EndTokenList
)

// TokenDef contains definition of an
// individual type of token
type TokenDef struct {
	TypeID  TokenType
	Match   string
	Name    string
	Keyword bool
	exp     *regexp.Regexp
}

var tokenDefs []TokenDef = []TokenDef{
	{TypeID: Identifier, Match: `^[a-zA-Z][a-zA-Z_0-9]*$`, Name: "Identifier"},
	{TypeID: Println, Match: "^println$", Name: "Println", Keyword: true},
	{TypeID: Print, Match: "^print$", Name: "Print", Keyword: true},
	{TypeID: Var, Match: "^var$", Name: "Var", Keyword: true},
	{TypeID: StringType, Match: "^string$", Name: "String", Keyword: true},
	{TypeID: NumberType, Match: "^number$", Name: "Number", Keyword: true},
	{TypeID: For, Match: `^for$`, Name: "For", Keyword: true},
	{TypeID: If, Match: `^if$`, Name: "If", Keyword: true},
	{TypeID: Else, Match: `^else$`, Name: "Else", Keyword: true},
	{TypeID: Integer, Match: `^[0-9]+$`, Name: "Integer"},
	{TypeID: Float, Match: `^[0-9]+\.[0-9]*$`, Name: "Float"},
	{TypeID: Percent, Match: `^%$`, Name: "Percent"},
	{TypeID: Dash, Match: `^-$`, Name: "Dash"},
	{TypeID: Plus, Match: `^\+$`, Name: "Plus"},
	{TypeID: PlusEquals, Match: `^\+=$`, Name: "Plus Equals"},
	{TypeID: DoublePlus, Match: `^\+\+$`, Name: "Double Plus"},
	{TypeID: Mult, Match: `^\*$`, Name: "Mult"},
	{TypeID: Div, Match: `^/$`, Name: "Div"},
	{TypeID: Equals, Match: `^=$`, Name: "Equals"},
	{TypeID: String, Match: `^\".*\"$`, Name: "String"},
	{TypeID: LeftCurlyBrace, Match: `^\{$`, Name: "Left Curly Brace"},
	{TypeID: RightCurlyBrace, Match: `^\}$`, Name: "Right Curly Brace"},
	{TypeID: LeftParen, Match: `^\($`, Name: "Left Paren"},
	{TypeID: RightParen, Match: `^\)$`, Name: "Right Parent"},
	{TypeID: LessThanEquals, Match: `^<=$`, Name: "Less Than or Equals"},
	{TypeID: LessThan, Match: `^<$`, Name: "Less Than"},
	{TypeID: GreaterThanEquals, Match: `^>=$`, Name: "Greater Than or Equals"},
	{TypeID: GreaterThan, Match: `^>$`, Name: "Greater Than"},
	{TypeID: DoubleEquals, Match: `^==$`, Name: "Double Equals"},
	{TypeID: Not, Match: `^!$`, Name: "Not"},
	{TypeID: NotEquals, Match: `^!=$`, Name: "Not Equals"},
	{TypeID: Period, Match: `^\.$`, Name: "Period"},
	{TypeID: Newline, Match: ``, Name: "Newline"},
	{TypeID: Hash, Match: `^#$`, Name: "Hash"},
	{TypeID: EndTokenList, Match: ``, Name: "End of tokens"},
}

// GetTokenDef returns the token definition
// for the specified type id
func GetTokenDef(typeID TokenType) *TokenDef {
	var result *TokenDef = nil

	for _, v := range tokenDefs {
		if v.TypeID == typeID {
			result = &v
			break
		}
	}

	return result
}

func (t *TokenDef) compile() {
	if len(t.Match) > 0 {
		t.exp = regexp.MustCompile(t.Match)
	}
}
