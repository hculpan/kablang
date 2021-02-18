package lexer

import (
	"fmt"
	"testing"
)

func TestLexer1(t *testing.T) {
	r, err := Lex("1+2.3", 1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	fmt.Printf("%+v\n", r)
	if len(r) != 3 {
		t.Log(fmt.Sprintf("Expected 3 tokens, found %d", len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: Integer, Value: "1"})
		testToken(t, r[1], Token{TypeID: Plus, Value: "+"})
		testToken(t, r[2], Token{TypeID: Float, Value: "2.3"})
	}
}

func TestLexer2(t *testing.T) {
	r, err := Lex("1 + 2.3", 1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(r) != 3 {
		t.Log(fmt.Sprintf("Expected 3 tokens, found %d", len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: Integer, Value: "1"})
		testToken(t, r[1], Token{TypeID: Plus, Value: "+"})
		testToken(t, r[2], Token{TypeID: Float, Value: "2.3"})
	}
}

func TestLexer3(t *testing.T) {
	r, err := Lex("1 +=     2.3", 1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(r) != 3 {
		t.Log(fmt.Sprintf("Expected 3 tokens, found %d", len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: Integer, Value: "1"})
		testToken(t, r[1], Token{TypeID: PlusEquals, Value: "+="})
		testToken(t, r[2], Token{TypeID: Float, Value: "2.3"})
	}
}

func TestLexer4(t *testing.T) {
	r, err := Lex("1 +=     2.3++", 1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(r) != 4 {
		t.Log(fmt.Sprintf("Expected 4 tokens, found %d", len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: Integer, Value: "1"})
		testToken(t, r[1], Token{TypeID: PlusEquals, Value: "+="})
		testToken(t, r[2], Token{TypeID: Float, Value: "2.3"})
		testToken(t, r[3], Token{TypeID: DoublePlus, Value: "++"})
	}
}

func TestLexer5(t *testing.T) {
	r, err := Lex("  my_number1=1+ 0.3  ", 1)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(r) != 5 {
		t.Log(fmt.Sprintf("Expected 5 tokens, found %d", len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: Identifier, Value: "my_number1"})
		testToken(t, r[1], Token{TypeID: Equals, Value: "="})
		testToken(t, r[2], Token{TypeID: Integer, Value: "1"})
		testToken(t, r[3], Token{TypeID: Plus, Value: "+"})
		testToken(t, r[4], Token{TypeID: Float, Value: "0.3"})
	}
}

func TestLexer6_ErrorTesting(t *testing.T) {
	_, err := Lex("  my_number1 = ~1+ .3", 1)
	if err == nil {
		t.Log("Should have received error for unknown token '~'")
		t.Fail()
		return
	}
}

func TestLexer6_StringTokenizing(t *testing.T) {
	r, err := Lex(`  my_string = "This is a test string"`, 1)
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
		testToken(t, r[0], Token{TypeID: Identifier, Value: "my_string"})
		testToken(t, r[1], Token{TypeID: Equals, Value: "="})
		testToken(t, r[2], Token{TypeID: String, Value: `"This is a test string"`})
	}
}

func TestLexer7_ForKeywordTest(t *testing.T) {
	r, err := Lex(`  for (a<2) {}`, 1)
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
		testToken(t, r[0], Token{TypeID: For, Value: "for"})
		testToken(t, r[1], Token{TypeID: LeftParen, Value: "("})
		testToken(t, r[2], Token{TypeID: Identifier, Value: `a`})
		testToken(t, r[3], Token{TypeID: LessThan, Value: `<`})
		testToken(t, r[4], Token{TypeID: Integer, Value: `2`})
		testToken(t, r[5], Token{TypeID: RightParen, Value: `)`})
		testToken(t, r[6], Token{TypeID: LeftCurlyBrace, Value: `{`})
		testToken(t, r[7], Token{TypeID: RightCurlyBrace, Value: `}`})
	}
}

func TestLexer8(t *testing.T) {
	r, err := Lex(`  if (a>=2) {} else if (fora > 3.09) {}`, 1)
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
		testToken(t, r[0], Token{TypeID: If, Value: "if"})
		testToken(t, r[1], Token{TypeID: LeftParen, Value: "("})
		testToken(t, r[2], Token{TypeID: Identifier, Value: `a`})
		testToken(t, r[3], Token{TypeID: GreaterThanEquals, Value: `>=`})
		testToken(t, r[4], Token{TypeID: Integer, Value: `2`})
		testToken(t, r[5], Token{TypeID: RightParen, Value: `)`})
		testToken(t, r[6], Token{TypeID: LeftCurlyBrace, Value: `{`})
		testToken(t, r[7], Token{TypeID: RightCurlyBrace, Value: `}`})
		testToken(t, r[8], Token{TypeID: Else, Value: `else`})
		testToken(t, r[9], Token{TypeID: If, Value: `if`})
		testToken(t, r[10], Token{TypeID: LeftParen, Value: `(`})
		testToken(t, r[11], Token{TypeID: Identifier, Value: `fora`})
		testToken(t, r[12], Token{TypeID: GreaterThan, Value: `>`})
		testToken(t, r[13], Token{TypeID: Float, Value: `3.09`})
		testToken(t, r[14], Token{TypeID: RightParen, Value: `)`})
		testToken(t, r[15], Token{TypeID: LeftCurlyBrace, Value: `{`})
		testToken(t, r[16], Token{TypeID: RightCurlyBrace, Value: `}`})
	}
}

func TestLexer9(t *testing.T) {
	r, err := Lex(`  if (a>=2) {} else if (fora == "A test") {}`, 1)
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
		testToken(t, r[0], Token{TypeID: If, Value: "if"})
		testToken(t, r[1], Token{TypeID: LeftParen, Value: "("})
		testToken(t, r[2], Token{TypeID: Identifier, Value: `a`})
		testToken(t, r[3], Token{TypeID: GreaterThanEquals, Value: `>=`})
		testToken(t, r[4], Token{TypeID: Integer, Value: `2`})
		testToken(t, r[5], Token{TypeID: RightParen, Value: `)`})
		testToken(t, r[6], Token{TypeID: LeftCurlyBrace, Value: `{`})
		testToken(t, r[7], Token{TypeID: RightCurlyBrace, Value: `}`})
		testToken(t, r[8], Token{TypeID: Else, Value: `else`})
		testToken(t, r[9], Token{TypeID: If, Value: `if`})
		testToken(t, r[10], Token{TypeID: LeftParen, Value: `(`})
		testToken(t, r[11], Token{TypeID: Identifier, Value: `fora`})
		testToken(t, r[12], Token{TypeID: DoubleEquals, Value: `==`})
		testToken(t, r[13], Token{TypeID: String, Value: `"A test"`})
		testToken(t, r[14], Token{TypeID: RightParen, Value: `)`})
		testToken(t, r[15], Token{TypeID: LeftCurlyBrace, Value: `{`})
		testToken(t, r[16], Token{TypeID: RightCurlyBrace, Value: `}`})
	}
}

func TestLexer10(t *testing.T) {
	r, err := Lex(`  if (  a!= 2  ) {

	   } else if (!fora.Equals  ("A test")) {

	   }`, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	//	fmt.Printf("%+v\n", r)
	expectedCount := 23
	if len(r) != expectedCount {
		t.Log(fmt.Sprintf("Expected %d tokens, found %d", expectedCount, len(r)))
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: If, Value: "if"})
		testToken(t, r[1], Token{TypeID: LeftParen, Value: "("})
		testToken(t, r[2], Token{TypeID: Identifier, Value: `a`})
		testToken(t, r[3], Token{TypeID: NotEquals, Value: `!=`})
		testToken(t, r[4], Token{TypeID: Integer, Value: `2`})
		testToken(t, r[5], Token{TypeID: RightParen, Value: `)`})
		testToken(t, r[6], Token{TypeID: LeftCurlyBrace, Value: `{`})
		testToken(t, r[7], Token{TypeID: Newline, Value: `\n`})
		testToken(t, r[8], Token{TypeID: RightCurlyBrace, Value: `}`})
		testToken(t, r[9], Token{TypeID: Else, Value: `else`})
		testToken(t, r[10], Token{TypeID: If, Value: `if`})
		testToken(t, r[11], Token{TypeID: LeftParen, Value: `(`})
		testToken(t, r[12], Token{TypeID: Not, Value: `!`})
		testToken(t, r[13], Token{TypeID: Identifier, Value: `fora`})
		testToken(t, r[14], Token{TypeID: Period, Value: `.`})
		testToken(t, r[15], Token{TypeID: Identifier, Value: `Equals`})
		testToken(t, r[16], Token{TypeID: LeftParen, Value: `(`})
		testToken(t, r[17], Token{TypeID: String, Value: `"A test"`})
		testToken(t, r[18], Token{TypeID: RightParen, Value: `)`})
		testToken(t, r[19], Token{TypeID: RightParen, Value: `)`})
		testToken(t, r[20], Token{TypeID: LeftCurlyBrace, Value: `{`})
		testToken(t, r[21], Token{TypeID: Newline, Value: `\n`})
		testToken(t, r[22], Token{TypeID: RightCurlyBrace, Value: `}`})
	}
}

func TestLexer11_PrintlnTest(t *testing.T) {
	r, err := Lex(`println "Hello" + " world"`, 1)
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
		testToken(t, r[0], Token{TypeID: Println, Value: "println"})
		testToken(t, r[1], Token{TypeID: String, Value: `"Hello"`})
		testToken(t, r[2], Token{TypeID: Plus, Value: `+`})
		testToken(t, r[3], Token{TypeID: String, Value: `" world"`})
	}
}

func TestLexer12_KeywordTest(t *testing.T) {
	r, err := Lex(`println var if else string number`, 1)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	expectedCount := 6
	if len(r) != expectedCount {
		t.Log(fmt.Sprintf("Expected %d tokens, found %d", expectedCount, len(r)))
		fmt.Printf("%+v\n", r[0])
		t.Fail()
	} else {
		testToken(t, r[0], Token{TypeID: Println, Value: "println"})
		testToken(t, r[1], Token{TypeID: Var, Value: "var"})
		testToken(t, r[2], Token{TypeID: If, Value: "if"})
		testToken(t, r[3], Token{TypeID: Else, Value: "else"})
		testToken(t, r[4], Token{TypeID: StringType, Value: "string"})
		testToken(t, r[5], Token{TypeID: NumberType, Value: "number"})
	}
}

func testToken(t *testing.T, token Token, expected Token) {
	if !token.Equals(expected) {
		t.Log(fmt.Sprintf("Expected %s, found %s [%s]", expected.TypeID.String(), token.TypeID.String(), token.Value))
		t.Fail()
	}
}
