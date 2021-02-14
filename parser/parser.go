package parser

import (
	"fmt"

	"github.com/hculpan/kablang/ast"
	"github.com/hculpan/kablang/lexer"
)

// Keyword constantsw
const (
	NEWLINE = lexer.END_TOKEN_LIST + 1
	PRINTLN = lexer.END_TOKEN_LIST + 2
	PRINT   = lexer.END_TOKEN_LIST + 3
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
		case NEWLINE:
			done = true
		case PRINT:
			p.lexerHandler.Push()
			stmts = append(stmts, p.parsePrintStatement())
			if !p.lexerHandler.Swallow(NEWLINE) {
				p.errors = append(p.errors, fmt.Errorf("Expected NEWLINE, found '%s'", t.Value))
			}
		case PRINTLN:
			p.lexerHandler.Push()
			stmts = append(stmts, p.parsePrintlnStatement())
			if !p.lexerHandler.Swallow(NEWLINE) {
				p.errors = append(p.errors, fmt.Errorf("Expected NEWLINE, found '%s'", t.Value))
			}
		default:
			p.errors = append(p.errors, fmt.Errorf("Unexpected token: '%s'", t.Value))
			done = true
		}
	}

	return ast.NewStatements(stmts)
}

func (p *Parser) parsePrintStatement() *ast.PrintStatement {
	p.swallow(PRINT)

	t := p.lexerHandler.Peek()
	switch t.TypeID {
	case NEWLINE:
		p.addError(fmt.Errorf("Expected expression, found NEWLINE"))
	case lexer.END_TOKEN_LIST:
		p.addError(fmt.Errorf("Expected expression, found End of tokens"))
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
	case NEWLINE:
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
			p.addError(fmt.Errorf("Expecting string, found '%s'", t2.Name))
		}
		result.StringExpressionNode = p.parseStringExpression()
	default:
		p.lexerHandler.Push()
	}

	return result
}

func (p *Parser) swallow(typeID int) bool {
	if !p.lexerHandler.Swallow(typeID) {
		p.addError(fmt.Errorf("Expecting type %d, found type %d", typeID, p.lexerHandler.Peek().TypeID))
		return false
	}

	return true
}

func (p *Parser) addError(e error) {
	p.errors = append(p.errors, e)
}
