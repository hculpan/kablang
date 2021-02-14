package lexer

import (
	"fmt"
	"testing"
)

const (
	FOR     = 1 + END_TOKEN_LIST
	IF      = 2 + END_TOKEN_LIST
	ELSE    = 3 + END_TOKEN_LIST
	PRINTLN = 4 + END_TOKEN_LIST
	PRINT   = 5 + END_TOKEN_LIST
)

var keywords []TokenDef = []TokenDef{
	{TypeID: FOR, Match: "for"},
	{TypeID: IF, Match: "if"},
	{TypeID: ELSE, Match: "else"},
	{TypeID: PRINTLN, Match: "println"},
	{TypeID: PRINT, Match: "print"},
}

func TestLexer1(t *testing.T) {
	r, err := Lex("1+2.3", []TokenDef{}, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(r) != 3 {
		t.Log(fmt.Sprintf("Expected 3 tokens, found %d", len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: INTEGER, Value: "1"})
		testToken(t, r[1], Token{TypeID: PLUS, Value: "+"})
		testToken(t, r[2], Token{TypeID: FLOAT, Value: "2.3"})
	}
}

func TestLexer2(t *testing.T) {
	r, err := Lex("1 + 2.3", []TokenDef{}, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(r) != 3 {
		t.Log(fmt.Sprintf("Expected 3 tokens, found %d", len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: INTEGER, Value: "1"})
		testToken(t, r[1], Token{TypeID: PLUS, Value: "+"})
		testToken(t, r[2], Token{TypeID: FLOAT, Value: "2.3"})
	}
}

func TestLexer3(t *testing.T) {
	r, err := Lex("1 +=     2.3", []TokenDef{}, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(r) != 3 {
		t.Log(fmt.Sprintf("Expected 3 tokens, found %d", len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: INTEGER, Value: "1"})
		testToken(t, r[1], Token{TypeID: PLUS_EQUALS, Value: "+="})
		testToken(t, r[2], Token{TypeID: FLOAT, Value: "2.3"})
	}
}

func TestLexer4(t *testing.T) {
	r, err := Lex("1 +=     2.3++", []TokenDef{}, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(r) != 4 {
		t.Log(fmt.Sprintf("Expected 4 tokens, found %d", len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: INTEGER, Value: "1"})
		testToken(t, r[1], Token{TypeID: PLUS_EQUALS, Value: "+="})
		testToken(t, r[2], Token{TypeID: FLOAT, Value: "2.3"})
		testToken(t, r[3], Token{TypeID: DOUBLE_PLUS, Value: "++"})
	}
}

func TestLexer5(t *testing.T) {
	r, err := Lex("  my_number1=1+ 0.3  ", []TokenDef{}, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(r) != 5 {
		t.Log(fmt.Sprintf("Expected 5 tokens, found %d", len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: IDENTIFIER, Value: "my_number1"})
		testToken(t, r[1], Token{TypeID: EQUALS, Value: "="})
		testToken(t, r[2], Token{TypeID: INTEGER, Value: "1"})
		testToken(t, r[3], Token{TypeID: PLUS, Value: "+"})
		testToken(t, r[4], Token{TypeID: FLOAT, Value: "0.3"})
	}
}

func TestLexer6_ErrorTesting(t *testing.T) {
	_, err := Lex("  my_number1 = ~1+ .3", []TokenDef{}, 1)
	if err == nil {
		t.Log("Should have received error for unknown token '~'")
		t.Fail()
		return
	}
}

func TestLexer6_StringTokenizing(t *testing.T) {
	r, err := Lex(`  my_string = "This is a test string"`, []TokenDef{}, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	expectedCount := 3
	if len(r) != expectedCount {
		t.Log(fmt.Sprintf("Expected %d tokens, found %d", expectedCount, len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: IDENTIFIER, Value: "my_string"})
		testToken(t, r[1], Token{TypeID: EQUALS, Value: "="})
		testToken(t, r[2], Token{TypeID: STRING, Value: `"This is a test string"`})
	}
}

func TestLexer7_ForKeywordTest(t *testing.T) {
	r, err := Lex(`  for (a<2) {}`, []TokenDef{
		{TypeID: FOR, Match: "for"},
	}, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	expectedCount := 8
	if len(r) != expectedCount {
		t.Log(fmt.Sprintf("Expected %d tokens, found %d", expectedCount, len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: FOR, Value: "for"})
		testToken(t, r[1], Token{TypeID: PAREN_LEFT, Value: "("})
		testToken(t, r[2], Token{TypeID: IDENTIFIER, Value: `a`})
		testToken(t, r[3], Token{TypeID: LESS_THAN, Value: `<`})
		testToken(t, r[4], Token{TypeID: INTEGER, Value: `2`})
		testToken(t, r[5], Token{TypeID: PAREN_RIGHT, Value: `)`})
		testToken(t, r[6], Token{TypeID: CURLY_BRACE_LEFT, Value: `{`})
		testToken(t, r[7], Token{TypeID: CURLY_BRACE_RIGHT, Value: `}`})
	}
}

func TestLexer8(t *testing.T) {
	r, err := Lex(`  if (a>=2) {} else if (fora > 3.09) {}`, []TokenDef{
		{TypeID: FOR, Match: "for"},
		{TypeID: IF, Match: "if"},
		{TypeID: ELSE, Match: "else"},
	}, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	expectedCount := 17
	if len(r) != expectedCount {
		t.Log(fmt.Sprintf("Expected %d tokens, found %d", expectedCount, len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: IF, Value: "if"})
		testToken(t, r[1], Token{TypeID: PAREN_LEFT, Value: "("})
		testToken(t, r[2], Token{TypeID: IDENTIFIER, Value: `a`})
		testToken(t, r[3], Token{TypeID: GREATER_THAN_EQUALS, Value: `>=`})
		testToken(t, r[4], Token{TypeID: INTEGER, Value: `2`})
		testToken(t, r[5], Token{TypeID: PAREN_RIGHT, Value: `)`})
		testToken(t, r[6], Token{TypeID: CURLY_BRACE_LEFT, Value: `{`})
		testToken(t, r[7], Token{TypeID: CURLY_BRACE_RIGHT, Value: `}`})
		testToken(t, r[8], Token{TypeID: ELSE, Value: `else`})
		testToken(t, r[9], Token{TypeID: IF, Value: `if`})
		testToken(t, r[10], Token{TypeID: PAREN_LEFT, Value: `(`})
		testToken(t, r[11], Token{TypeID: IDENTIFIER, Value: `fora`})
		testToken(t, r[12], Token{TypeID: GREATER_THAN, Value: `>`})
		testToken(t, r[13], Token{TypeID: FLOAT, Value: `3.09`})
		testToken(t, r[14], Token{TypeID: PAREN_RIGHT, Value: `)`})
		testToken(t, r[15], Token{TypeID: CURLY_BRACE_LEFT, Value: `{`})
		testToken(t, r[16], Token{TypeID: CURLY_BRACE_RIGHT, Value: `}`})
	}
}

func TestLexer9(t *testing.T) {
	r, err := Lex(`  if (a>=2) {} else if (fora == "A test") {}`, []TokenDef{
		{TypeID: FOR, Match: "for"},
		{TypeID: IF, Match: "if"},
		{TypeID: ELSE, Match: "else"},
	}, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	expectedCount := 17
	if len(r) != expectedCount {
		t.Log(fmt.Sprintf("Expected %d tokens, found %d", expectedCount, len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: IF, Value: "if"})
		testToken(t, r[1], Token{TypeID: PAREN_LEFT, Value: "("})
		testToken(t, r[2], Token{TypeID: IDENTIFIER, Value: `a`})
		testToken(t, r[3], Token{TypeID: GREATER_THAN_EQUALS, Value: `>=`})
		testToken(t, r[4], Token{TypeID: INTEGER, Value: `2`})
		testToken(t, r[5], Token{TypeID: PAREN_RIGHT, Value: `)`})
		testToken(t, r[6], Token{TypeID: CURLY_BRACE_LEFT, Value: `{`})
		testToken(t, r[7], Token{TypeID: CURLY_BRACE_RIGHT, Value: `}`})
		testToken(t, r[8], Token{TypeID: ELSE, Value: `else`})
		testToken(t, r[9], Token{TypeID: IF, Value: `if`})
		testToken(t, r[10], Token{TypeID: PAREN_LEFT, Value: `(`})
		testToken(t, r[11], Token{TypeID: IDENTIFIER, Value: `fora`})
		testToken(t, r[12], Token{TypeID: DOUBLE_EQUALS, Value: `==`})
		testToken(t, r[13], Token{TypeID: STRING, Value: `"A test"`})
		testToken(t, r[14], Token{TypeID: PAREN_RIGHT, Value: `)`})
		testToken(t, r[15], Token{TypeID: CURLY_BRACE_LEFT, Value: `{`})
		testToken(t, r[16], Token{TypeID: CURLY_BRACE_RIGHT, Value: `}`})
	}
}

func TestLexer10(t *testing.T) {
	var keywords []TokenDef = []TokenDef{
		{TypeID: FOR, Match: "for"},
		{TypeID: IF, Match: "if"},
		{TypeID: ELSE, Match: "else"},
		{TypeID: PRINTLN, Match: "println"},
		{TypeID: PRINT, Match: "print"},
	}

	r, err := Lex(`  if (  a!= 2  ) {

	   } else if (!fora.Equals  ("A test")) {

	   }`, keywords, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	expectedCount := 21
	if len(r) != expectedCount {
		t.Log(fmt.Sprintf("Expected %d tokens, found %d", expectedCount, len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: IF, Value: "if"})
		testToken(t, r[1], Token{TypeID: PAREN_LEFT, Value: "("})
		testToken(t, r[2], Token{TypeID: IDENTIFIER, Value: `a`})
		testToken(t, r[3], Token{TypeID: NOT_EQUALS, Value: `!=`})
		testToken(t, r[4], Token{TypeID: INTEGER, Value: `2`})
		testToken(t, r[5], Token{TypeID: PAREN_RIGHT, Value: `)`})
		testToken(t, r[6], Token{TypeID: CURLY_BRACE_LEFT, Value: `{`})
		testToken(t, r[7], Token{TypeID: CURLY_BRACE_RIGHT, Value: `}`})
		testToken(t, r[8], Token{TypeID: ELSE, Value: `else`})
		testToken(t, r[9], Token{TypeID: IF, Value: `if`})
		testToken(t, r[10], Token{TypeID: PAREN_LEFT, Value: `(`})
		testToken(t, r[11], Token{TypeID: NOT, Value: `!`})
		testToken(t, r[12], Token{TypeID: IDENTIFIER, Value: `fora`})
		testToken(t, r[13], Token{TypeID: PERIOD, Value: `.`})
		testToken(t, r[14], Token{TypeID: IDENTIFIER, Value: `Equals`})
		testToken(t, r[15], Token{TypeID: PAREN_LEFT, Value: `(`})
		testToken(t, r[16], Token{TypeID: STRING, Value: `"A test"`})
		testToken(t, r[17], Token{TypeID: PAREN_RIGHT, Value: `)`})
		testToken(t, r[18], Token{TypeID: PAREN_RIGHT, Value: `)`})
		testToken(t, r[19], Token{TypeID: CURLY_BRACE_LEFT, Value: `{`})
		testToken(t, r[20], Token{TypeID: CURLY_BRACE_RIGHT, Value: `}`})
	}
}

func TestLexer11_PrintlnTest(t *testing.T) {
	r, err := Lex(`println "Hello" + " world"`, keywords, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	expectedCount := 4
	if len(r) != expectedCount {
		t.Log(fmt.Sprintf("Expected %d tokens, found %d", expectedCount, len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: PRINTLN, Value: "println"})
		testToken(t, r[1], Token{TypeID: STRING, Value: `"Hello"`})
		testToken(t, r[2], Token{TypeID: PLUS, Value: `+`})
		testToken(t, r[3], Token{TypeID: STRING, Value: `" world"`})
	}
}

func testToken(t *testing.T, token Token, expected Token) {
	if !token.Equals(expected) {
		t.Log(fmt.Sprintf("Expected %v, found %v", expected, token))
		t.Fail()
	}
}
