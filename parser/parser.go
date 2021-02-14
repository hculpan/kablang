package parser

import (
	"fmt"

	"github.com/hculpan/kablang/ast"
	"github.com/hculpan/kablang/lexer"
)

// Keyword constantsw
const (
	PRINTLN = lexer.END_TOKEN_LIST + 1
	PRINT   = lexer.END_TOKEN_LIST + 2
)

var keywords []lexer.TokenDef = []lexer.TokenDef{
	{TypeID: PRINTLN, Match: "println", Name: "Println"},
	{TypeID: PRINT, Match: "print", Name: "Print"},
}

// Parser ...
type Parser struct {
	errors       []error
	lexerHandler *LexerHandler
}

// NewParser creates a new parser and returns
// a list of errors, if any
func NewParser() Parser {
	return Parser{}
}

// Parse parses the program send in in the lines.
// It will return the top node of the generated AST and
// and errors that were found
func (p *Parser) Parse(lines []string) (*ast.Program, []error) {
	p.errors = []error{}
	fmt.Print("Lexing program...")
	p.lexerHandler = NewLexerHandler(lines, keywords)
	if len(p.lexerHandler.Errors) != 0 {
		fmt.Println()
		return nil, []error{p.lexerHandler.Errors[0]}
	}
	fmt.Println("done")

	//	p.printTokens()

	fmt.Print("Parsing program...")

	result := p.parseProgram()

	fmt.Println("done")
	return result, p.errors
}

func (p *Parser) printTokens() {
	for i, t := range p.lexerHandler.tokens {
		fmt.Printf("%2d: %s\n", i, t.Name)
	}
}

func (p *Parser) parseProgram() *ast.Program {
	return ast.NewProgram(p.parseBlock())
}

func (p *Parser) parseBlock() *ast.Block {
	return ast.NewBlock(p.parseStatements())
}

func (p *Parser) parseStatements() *ast.Statements {
	stmts := []ast.Statement{}

	done := false
	for !done {
		t := p.lexerHandler.Pop()
		switch t.TypeID {
		case lexer.END_TOKEN_LIST:
			done = true
		case lexer.NEWLINE:
			done = true
		case PRINT:
			p.lexerHandler.Push()
			stmts = append(stmts, p.parsePrintStatement())
			if !p.lexerHandler.Swallow(lexer.NEWLINE) {
				p.addExpectedErrorForTypeID(lexer.NEWLINE, t)
			}
		case PRINTLN:
			p.lexerHandler.Push()
			stmts = append(stmts, p.parsePrintlnStatement())
			if !p.lexerHandler.Swallow(lexer.NEWLINE) {
				p.addExpectedErrorForTypeID(lexer.NEWLINE, t)
			}
		default:
			p.addError(fmt.Errorf("Unexpected token: '%s' at line %d:%d", t.Value, t.Line, t.Col))
			done = true
		}
	}

	return ast.NewStatements(stmts)
}

func (p *Parser) parsePrintStatement() *ast.PrintStatement {
	p.swallow(PRINT)

	t := p.lexerHandler.Peek()
	switch t.TypeID {
	case lexer.NEWLINE:
		p.addError(fmt.Errorf("Expected expression, found NEWLINE at line %d:%d", t.Line, t.Col))
	case lexer.END_TOKEN_LIST:
		p.addError(fmt.Errorf("Expected expression, found End of tokens at line %d:%d", t.Line, t.Col))
	case lexer.STRING:
		return ast.NewStringPrintStatement(p.parseStringExpression())
	case lexer.INTEGER:
		//		return ast.NewNumPrintlnStatement(p.parseStringExpression())
	case lexer.FLOAT:
		//		return ast.NewNumPrintlnStatement(p.parseStringExpression())
	}

	return nil
}

func (p *Parser) parsePrintlnStatement() *ast.PrintlnStatement {
	p.swallow(PRINTLN)

	t := p.lexerHandler.Peek()
	switch t.TypeID {
	case lexer.NEWLINE:
		return ast.NewEmptyPrintlnStatement()
	case lexer.END_TOKEN_LIST:
		return ast.NewEmptyPrintlnStatement()
	case lexer.STRING:
		return ast.NewStringPrintlnStatement(p.parseStringExpression())
	case lexer.INTEGER:
		//		return ast.NewNumPrintlnStatement(p.parseStringExpression())
	case lexer.FLOAT:
		//		return ast.NewNumPrintlnStatement(p.parseStringExpression())
	}

	return nil
}

func (p *Parser) parseStringExpression() *ast.StringExpression {
	t := p.lexerHandler.Pop()
	result, err := ast.NewStringExpression(ast.NewString(t.Value), nil)
	if err != nil {
		p.addError(err)
	}

	t = p.lexerHandler.Pop()
	switch t.TypeID {
	case lexer.PLUS:
		t2 := p.lexerHandler.Peek()
		if t2.TypeID != lexer.STRING {
			p.addExpectedErrorForTypeID(lexer.STRING, t2)
		}
		result.StringExpressionNode = p.parseStringExpression()
	default:
		p.lexerHandler.Push()
	}

	return result
}

func (p *Parser) swallow(typeID int) bool {
	if !p.lexerHandler.Swallow(typeID) {
		p.addExpectedErrorForTypeID(typeID, p.lexerHandler.Peek())
		return false
	}

	return true
}

func (p *Parser) addExpectedErrorForTypeID(expected int, actual lexer.Token) {
	tokenDef := lexer.GetTokenDef(expected)
	if tokenDef != nil {
		p.addExpectedError(*tokenDef, actual)
	}
}

func (p *Parser) addExpectedError(expected lexer.TokenDef, actual lexer.Token) {
	p.addError(fmt.Errorf("Expected %s, found %s at line %d:%d", expected.Name, actual.Name, actual.Line, actual.Col))
}

func (p *Parser) addError(e error) {
	p.errors = append(p.errors, e)
}

// GetTokenDef tries to get the token definition from the
// lexer definition, and if that fails, tries to get it
// from the keywords
func (p *Parser) GetTokenDef(typeID int) *lexer.TokenDef {
	result := lexer.GetTokenDef(typeID)
	if result == nil && typeID < lexer.END_TOKEN_LIST+len(keywords) {
		result = &keywords[typeID-lexer.END_TOKEN_LIST]
	}
	return result
}
