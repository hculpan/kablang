package executor

import (
	"fmt"

	"github.com/hculpan/kablang/ast"
)

// Executor contains the execution
// environment for this interpreter
type Executor struct {
	Errors []error
}

// NewExecutor ...
func NewExecutor() *Executor {
	result := &Executor{}
	result.Reset()
	return result
}

// Reset sets the execution environment
// back to it's initial state
func (e *Executor) Reset() {
	e.Errors = []error{}
}

// Execute executes the supplies AST
func (e *Executor) Execute(program *ast.Program) {
	block := program.BlockNode
	if block == nil {
		e.addError(fmt.Errorf("Program must contain a statement block"))
		return
	}

	if block.StatementsNode == nil {
		e.addError(fmt.Errorf("Program block must contain statements"))
		return
	}

	stmts := block.StatementsNode
	for _, s := range stmts.StatementListNode {
		switch s.(type) {
		case *ast.NullStatement:
			// do nothing
		case *ast.PrintStatement:
			e.executePrint(s.(*ast.PrintStatement))
		case *ast.PrintlnStatement:
			e.executePrintln(s.(*ast.PrintlnStatement))
		}
	}
}

func (e *Executor) executePrint(s *ast.PrintStatement) {
	switch s.ExpressionTypeID {
	case ast.NumExpressionType:
	case ast.StringExpressionType:
		exprResult := e.evaluateStringExpression(s.StringExpressionNode)
		fmt.Print(exprResult.Value)
	}
}

func (e *Executor) executePrintln(s *ast.PrintlnStatement) {
	switch s.ExpressionTypeID {
	case ast.NumExpressionType:
	case ast.EmptyExpressionType:
		fmt.Println()
	case ast.StringExpressionType:
		exprResult := e.evaluateStringExpression(s.StringExpressionNode)
		fmt.Println(exprResult.Value)
	}
}

func (e *Executor) evaluateStringExpression(exp *ast.StringExpression) *ast.String {
	result := exp.StringNode

	if exp.StringExpressionNode != nil {
		r2 := e.evaluateStringExpression(exp.StringExpressionNode.(*ast.StringExpression))
		result.Value += r2.Value
	}

	return result
}

func (e *Executor) addError(err error) {
	e.Errors = append(e.Errors, err)
}
