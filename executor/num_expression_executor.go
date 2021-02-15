package executor

import "github.com/hculpan/kablang/ast"

func (e *Executor) evaluateNumExpression(exp *ast.NumExpression) ast.NumberValue {
	var result ast.NumberValue
	if exp.TermNode != nil {
		result = e.evaluateTerm(exp.TermNode)
	}

	switch exp.Operator {
	case ast.PlusOperator:
		expValue := e.evaluateNumExpression(exp.NumExpressionNode.(*ast.NumExpression))
		r := result.Add(expValue)
		result = &r
	case ast.MinusOperator:
		expValue := e.evaluateNumExpression(exp.NumExpressionNode.(*ast.NumExpression))
		r := result.Sub(expValue)
		result = &r
	}

	return result
}

func (e *Executor) evaluateTerm(term *ast.Term) ast.NumberValue {
	var result ast.NumberValue
	if term.FactorNode != nil {
		result = e.evaluateFactor(term.FactorNode)
	}

	switch term.Operator {
	case ast.MultOperator:
		termValue := e.evaluateTerm(term.TermNode)
		r := result.Mult(termValue)
		result = &r
	case ast.DivOperator:
		termValue := e.evaluateTerm(term.TermNode)
		r := result.Div(termValue)
		result = &r
	}

	return result
}

func (e *Executor) evaluateFactor(factor *ast.Factor) ast.NumberValue {
	if factor.NumberNode != nil {
		return factor.NumberNode
	} else if factor.ParenNode != nil {
		return e.evaluateNumExpression(factor.ParenNode)
	}

	return nil
}
