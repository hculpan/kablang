package parser

import (
	"fmt"

	"github.com/hculpan/kablang/ast"
	"github.com/hculpan/kablang/lexer"
)

// Parser ...
type Parser struct {
	errors       []error
	lexerHandler *LexerHandler
	blockStack   *ast.BlockStack
}

// NewParser creates a new parser and returns
// a list of errors, if any
func NewParser() Parser {
	return Parser{blockStack: ast.NewBlockStack()}
}

// Parse parses the program send in in the lines.
// It will return the top node of the generated AST and
// and errors that were found
func (p *Parser) Parse(lines []string) (*ast.Program, []error) {
	p.errors = []error{}
	p.lexerHandler = NewLexerHandler(lines)
	if len(p.lexerHandler.Errors) != 0 {
		fmt.Println()
		return nil, []error{p.lexerHandler.Errors[0]}
	}

	result := p.parseProgram()

	return result, p.errors
}

func (p *Parser) printTokens() {
	for i, t := range p.lexerHandler.tokens {
		fmt.Printf("%2d: %s\n", i, t.Name)
	}
}

func (p *Parser) parseProgram() *ast.Program {
	return ast.NewProgram(p.parseBlock(nil))
}

func (p *Parser) parseBlock(parent *ast.Block) *ast.Block {
	p.swallow(lexer.CURLY_BRACE_LEFT)
	result := ast.NewBlock(parent)
	p.blockStack.Push(result)
	result.StatementsNode = p.parseStatements()
	p.swallow(lexer.CURLY_BRACE_RIGHT)
	p.blockStack.Pop()
	return result
}

func (p *Parser) currentBlock() *ast.Block {
	return p.blockStack.Peek()
}

func (p *Parser) parseStatements() *ast.Statements {
	stmts := []ast.Statement{}

	done := false
	for !done {
		var stmt ast.Statement = nil
		t := p.lexerHandler.Pop()
		switch t.TypeID {
		case lexer.NEWLINE:
			continue
		case lexer.CURLY_BRACE_LEFT:
			p.lexerHandler.Push()
			stmt = p.parseBlock(p.blockStack.Peek())
		case lexer.CURLY_BRACE_RIGHT:
			p.lexerHandler.Push()
			done = true
		case lexer.END_TOKEN_LIST:
			done = true
		case lexer.HASH:
			for {
				t = p.lexerHandler.Pop()
				if t.TypeID == lexer.NEWLINE || t.TypeID == lexer.END_TOKEN_LIST {
					break
				}
			}
		case lexer.VAR:
			a, err := p.parseVarStatement()
			if err == nil {
				stmt = a
				symbol := stmt.(*ast.VarStatement).SymbolNode
				if _, exists := p.currentBlock().Symbols.GetLocal(symbol.GetName()); exists {
					p.addError(fmt.Errorf("Redefinition of variable '%s' at %d:%d", symbol.GetName(), t.Line, t.Col))
					return nil
				}
				p.currentBlock().AddSymbol(symbol)
			}
			p.swallow(lexer.NEWLINE)
		case lexer.IDENTIFIER:
			stmt = p.parseAssignStatement(&t)
			p.swallow(lexer.NEWLINE)
		case lexer.PRINT:
			p.lexerHandler.Push()
			stmt = p.parsePrintStatement(false)
			if !p.lexerHandler.Swallow(lexer.NEWLINE) {
				p.addExpectedErrorForTypeID(lexer.NEWLINE, t)
			}
		case lexer.PRINTLN:
			p.lexerHandler.Push()
			stmt = p.parsePrintStatement(true)
			p.swallow(lexer.NEWLINE)
		default:
			p.addError(fmt.Errorf("Unexpected token: '%s' at line %d:%d", t.Value, t.Line, t.Col))
			done = true
		}
		if stmt != nil {
			stmts = append(stmts, stmt)
		}
	}

	return ast.NewStatements(stmts)
}

func (p *Parser) parseAssignStatement(t *lexer.Token) *ast.AssignStatement {
	if t.TypeID != lexer.IDENTIFIER {
		return nil
	}

	p.swallow(lexer.EQUALS)

	if symbol, exists := p.currentBlock().Symbols.GetLocal(t.Value); exists {
		stmt := ast.NewAssignStatement(symbol)
		switch symbol.GetDataType() {
		case ast.TypeString:
			stmt.ExpressionNode = p.parseStringExpression()
			return stmt
		case ast.TypeNumber:
			stmt.ExpressionNode = p.parseNumExpression()
			return stmt
		default:
			p.addError(fmt.Errorf("Unsupported data type for variable assignment at line %d:%d", t.Line, t.Col))
			return nil
		}
	}

	p.addError(fmt.Errorf("Assignment without declaration for variable '%s' at line %d:%d", t.Value, t.Line, t.Col))
	return nil
}

func (p *Parser) parseVarStatement() (*ast.VarStatement, error) {
	nameToken := p.lexerHandler.Pop()
	if nameToken.TypeID != lexer.IDENTIFIER {
		p.lexerHandler.Push()
		p.addExpectedErrorForTypeID(lexer.IDENTIFIER, nameToken)
		return nil, fmt.Errorf("")
	}

	typeToken := p.lexerHandler.Pop()
	if typeToken.TypeID != lexer.STRING_TYPE && typeToken.TypeID != lexer.NUMBER_TYPE {
		p.lexerHandler.Push()
		p.addExpectedErrorForString("Expecting data type indicator", typeToken)
		return nil, fmt.Errorf("")
	}

	var result *ast.VarStatement = nil

	switch typeToken.Value {
	case "string":
		result = ast.NewVarStatement(nameToken.Value, ast.TypeString)
	case "number":
		result = ast.NewVarStatement(nameToken.Value, ast.TypeNumber)
	default:
		p.addError(fmt.Errorf("Invalid data type: %s", typeToken.Value))
		return nil, fmt.Errorf("")
	}

	t := p.lexerHandler.Peek()
	if t.TypeID == lexer.EQUALS {
		p.swallow(lexer.EQUALS)
		switch result.SymbolNode.GetDataType() {
		case ast.TypeString:
			result.ExpressionNode = p.parseStringExpression()
		case ast.TypeNumber:
			result.ExpressionNode = p.parseNumExpression()
		default:
			p.addError(fmt.Errorf("Invalid data type assigned to variable '%s' of type '%s' at %d:%d",
				result.SymbolNode.GetName(), ast.GetTypeName(result.SymbolNode.GetDataType()), t.Line, t.Col))
			return nil, fmt.Errorf("")
		}
	}

	return result, nil
}

func (p *Parser) parsePrintStatement(endline bool) *ast.PrintStatement {
	if endline {
		p.swallow(lexer.PRINTLN)
	} else {
		p.swallow(lexer.PRINT)
	}

	t := p.lexerHandler.Peek()
	switch t.TypeID {
	case lexer.NEWLINE:
		return ast.NewEmptyPrintStatement(endline)
	case lexer.END_TOKEN_LIST:
		return ast.NewEmptyPrintStatement(endline)
	case lexer.STRING:
		return ast.NewStringPrintStatement(p.parseStringExpression(), endline)
	case lexer.INTEGER, lexer.FLOAT, lexer.DASH, lexer.PAREN_LEFT:
		return ast.NewNumPrintStatement(p.parseNumExpression(), endline)
	case lexer.IDENTIFIER:
		if symbol, exists := p.currentBlock().Symbols.Get(t.Value); exists {
			switch symbol.GetDataType() {
			case ast.TypeString:
				return ast.NewStringPrintStatement(p.parseStringExpression(), endline)
			case ast.TypeNumber:
				return ast.NewNumPrintStatement(p.parseNumExpression(), endline)
			default:
				p.addError(fmt.Errorf("Invalid data type for variable '%s' at %d:%d", t.Value, t.Line, t.Col))
				return nil
			}
		} else {
			p.addError(fmt.Errorf("Undeclared variable '%s' at %d:%d", t.Value, t.Line, t.Col))
			return nil
		}
	}

	return nil
}

func (p *Parser) parseString() ast.StringValue {
	var result ast.StringValue

	t := p.lexerHandler.Pop()
	switch t.TypeID {
	case lexer.IDENTIFIER:
		if symbol, exists := p.currentBlock().Symbols.Get(t.Value); exists {
			switch symbol.(type) {
			case *ast.StringSymbol:
				result = symbol.(*ast.StringSymbol)
			default:
				p.addError(fmt.Errorf("Cannot assign type %T to variable '%s' of type string", symbol, t.Value))
				return nil
			}
		} else {
			p.addError(fmt.Errorf("Undeclared variable '%s' at %d:%d", t.Value, t.Line, t.Col))
			return nil
		}
	case lexer.STRING:
		result = ast.NewString(t.Value)
	}

	return result
}

func (p *Parser) parseStringExpression() *ast.StringExpression {
	var result *ast.StringExpression = ast.NewStringExpression()

	result.StringNode = p.parseString()

	t := p.lexerHandler.Pop()
	switch t.TypeID {
	case lexer.PLUS:
		t2 := p.lexerHandler.Peek()
		if t2.TypeID != lexer.STRING && t2.TypeID != lexer.IDENTIFIER {
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
	tokenDef := p.GetTokenDef(expected)
	if tokenDef != nil {
		p.addExpectedError(*tokenDef, actual)
	}
}

func (p *Parser) addExpectedErrorForString(msg string, actual lexer.Token) {
	p.addError(fmt.Errorf("%s, found %s at line %d:%d", msg, actual.Name, actual.Line, actual.Col))
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
	return lexer.GetTokenDef(typeID)
}
