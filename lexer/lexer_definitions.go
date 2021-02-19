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
	Exponent
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
	TypeID   TokenType
	Match    string
	Name     string
	Priority int
	exp      *regexp.Regexp
}

type byPriority []TokenDef

func (a byPriority) Len() int           { return len(a) }
func (a byPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPriority) Less(i, j int) bool { return a[i].Priority > a[j].Priority }

var keywords []TokenDef = []TokenDef{
	newTokenDef(Println, "println", "Println"),
	newTokenDef(Print, "print", "Print"),
	newTokenDef(Var, "var", "Var"),
	newTokenDef(For, `for`, "For"),
	newTokenDef(If, `if`, "If"),
	newTokenDef(Else, `else`, "Else"),
}

var tokenDefs []TokenDef = []TokenDef{
	newTokenDef(Identifier, `^[a-zA-Z][a-zA-Z_0-9]*`, "Identifier"),
	newTokenDef(StringType, "^string", "String"),
	newTokenDef(NumberType, "^number", "Number"),
	newTokenDef(Integer, `^[0-9]+`, "Integer"),
	newTokenDef(Float, `^[0-9]+\.[0-9]*`, "Float"),
	newTokenDef(Percent, `^%`, "Percent"),
	newTokenDef(Dash, `^-`, "Dash"),
	newTokenDef(Exponent, `^\^`, "Exponent"),
	newTokenDef(Plus, `^\+`, "Plus"),
	newTokenDef(PlusEquals, `^\+=`, "Plus Equals"),
	newTokenDef(DoublePlus, `^\+\+`, "Double Plus"),
	newTokenDef(Mult, `^\*`, "Mult"),
	newTokenDef(Div, `^/`, "Div"),
	newTokenDef(Equals, `^=`, "Equals"),
	newTokenDef(String, `^\"[^\"]*\"`, "String"),
	newTokenDef(LeftCurlyBrace, `^\{`, "Left Curly Brace"),
	newTokenDef(RightCurlyBrace, `^\}`, "Right Curly Brace"),
	newTokenDef(LeftParen, `^\(`, "Left Paren"),
	newTokenDef(RightParen, `^\)`, "Right Parent"),
	newTokenDef(LessThanEquals, `^<=`, "Less Than or Equals"),
	newTokenDef(LessThan, `^<`, "Less Than"),
	newTokenDef(GreaterThanEquals, `^>=`, "Greater Than or Equals"),
	newTokenDef(GreaterThan, `^>`, "Greater Than"),
	newTokenDef(DoubleEquals, `^==`, "Double Equals"),
	newTokenDef(Not, `^!`, "Not"),
	newTokenDef(NotEquals, `^!=`, "Not Equals"),
	newTokenDef(Period, `^\.`, "Period"),
	newTokenDef(Newline, `^[\n]+`, "Newline"),
	newTokenDef(Hash, `^#`, "Hash"),
	newTokenDef(EndTokenList, ``, "End of tokens"),
}

func newTokenDef(typeID TokenType, match string, name string) TokenDef {
	var result TokenDef

	if typeID == Identifier || typeID == Integer || typeID == String {
		result = TokenDef{TypeID: typeID, Match: match, Name: name, Priority: 0}
	} else if typeID == Float {
		result = TokenDef{TypeID: typeID, Match: match, Name: name, Priority: 1}
	} else {
		result = TokenDef{TypeID: typeID, Match: match, Name: name, Priority: len(match)}
	}
	result.compile()

	return result
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
