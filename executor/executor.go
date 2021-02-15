package executor

import (
	"fmt"

	"github.com/hculpan/kablang/ast"
)

// Executor contains the execution
// environment for this interpreter
type Executor struct {
	Errors       []error
	CurrentBlock *ast.Block
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
	e.CurrentBlock = program.BlockNode
	if e.CurrentBlock == nil {
		e.addError(fmt.Errorf("Program must contain a statement block"))
		return
	}

	if e.CurrentBlock.StatementsNode == nil {
		e.addError(fmt.Errorf("Program block must contain statements"))
		return
	}

	stmts := e.CurrentBlock.StatementsNode
	for _, s := range stmts.StatementListNode {
		switch s.(type) {
		case *ast.NullStatement:
			// do nothing
		case *ast.PrintStatement:
			e.executePrint(s.(*ast.PrintStatement))
		case *ast.AssignStatement:
			e.executeAssignment(s.(*ast.AssignStatement))
		}
	}
}

func (e *Executor) executeAssignment(s *ast.AssignStatement) {
	if s.SymbolNode == nil || s.ExpressionNode == nil {
		e.addError(fmt.Errorf("Attempted invalid assignment operation"))
		return
	}

	if symbol, exists := e.CurrentBlock.Symbols[s.SymbolNode.GetName()]; exists {
		switch symbol.GetDataType() {
		case ast.TypeString:
			symbol.SetValue(e.evaluateStringExpression(s.ExpressionNode.(*ast.StringExpression)))
			return
		case ast.TypeNumber:
			symbol.SetValue(e.evaluateNumExpression(s.ExpressionNode.(*ast.NumExpression)))
			return
		default:
			e.addError(fmt.Errorf("Unrecognized data type for variable %s", s.SymbolNode.GetName()))
			return
		}
	}

	e.addError(fmt.Errorf("Attempted assignment to undeclared variable %s", s.SymbolNode.GetName()))
}

func (e *Executor) executePrint(s *ast.PrintStatement) {
	switch s.ExpressionTypeID {
	case ast.NumExpressionType:
		exprResult := e.evaluateNumExpression(s.NumExpressionNode)
		fmt.Print(exprResult.ToString())
	case ast.StringExpressionType:
		exprResult := e.evaluateStringExpression(s.StringExpressionNode)
		fmt.Print(exprResult.GetValue())
	}

	if s.WithEndline {
		fmt.Println()
	}
}

func (e *Executor) evaluateStringExpression(exp *ast.StringExpression) ast.StringValue {
	result := exp.StringNode

	if exp.StringExpressionNode != nil {
		r2 := e.evaluateStringExpression(exp.StringExpressionNode.(*ast.StringExpression))
		result.SetValue(result.GetValue() + r2.GetValue())
	}

	return result
}

func (e *Executor) addError(err error) {
	e.Errors = append(e.Errors, err)
}
