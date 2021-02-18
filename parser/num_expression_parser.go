package parser

import (
	"fmt"
	"strconv"

	"github.com/hculpan/kablang/ast"
	"github.com/hculpan/kablang/lexer"
)

func (p *Parser) parseNumExpression() *ast.NumExpression {
	result, _ := ast.NewNumExpression(nil, nil)

	result.TermNode = p.term()
	t := p.lexerHandler.Pop()
	switch t.TypeID {
	case lexer.Plus:
		result.Operator = ast.PlusOperator
		result.NumExpressionNode = p.parseNumExpression()
	case lexer.Dash:
		result.Operator = ast.MinusOperator
		result.NumExpressionNode = p.parseNumExpression()
	default:
		p.lexerHandler.Push()
	}

	return result
}

func (p *Parser) term() *ast.Term {
	result := ast.NewTerm()
	result.FactorNode = p.factor()

	t := p.lexerHandler.Pop()
	switch t.TypeID {
	case lexer.Mult:
		result.Operator = ast.MultOperator
		result.TermNode = p.term()
	case lexer.Div:
		result.Operator = ast.DivOperator
		result.TermNode = p.term()
	default:
		p.lexerHandler.Push()
	}

	return result
}

func (p *Parser) factor() *ast.Factor {
	result := ast.NewFactor()

	t := p.lexerHandler.Pop()
	switch t.TypeID {
	case lexer.Integer, lexer.Float:
		result.NumberNode = p.number(&t)
	case lexer.Dash:
		t = p.lexerHandler.Pop()
		if t.TypeID != lexer.Integer && t.TypeID != lexer.Float {
			p.lexerHandler.Push()
			p.addExpectedErrorForString("Expected number", t)
			return result
		}
		result.NumberNode = p.number(&t)
		r := result.NumberNode.Mult(ast.NewIntNumber(-1))
		result.NumberNode = &r
	case lexer.LeftParen:
		result.ParenNode = p.parseNumExpression()
		p.swallow(lexer.RightParen)
	case lexer.Identifier:
		if symbol, exists := p.currentBlock().Symbols.Get(t.Value); exists {
			switch symbol.(type) {
			case *ast.NumberSymbol:
				result.NumberNode = symbol.(*ast.NumberSymbol)
			default:
				p.addError(fmt.Errorf("Cannot assign type %T to variable '%s' of type number", symbol, t.Value))
				return nil
			}
		} else {
			p.addError(fmt.Errorf("Undeclared variable '%s' at %d:%d", t.Value, t.Line, t.Col))
			return nil
		}
	default:
		p.lexerHandler.Push()
		p.addExpectedErrorForString("Expected number", t)
		return nil
	}

	return result
}

func (p *Parser) number(t *lexer.Token) *ast.Number {
	var result *ast.Number

	switch t.TypeID {
	case lexer.Integer:
		n, _ := strconv.Atoi(t.Value)
		result = ast.NewIntNumber(int64(n))
	case lexer.Float:
		n, _ := strconv.ParseFloat(t.Value, 64)
		result = ast.NewFloatNumber(n)
	}

	return result
}
