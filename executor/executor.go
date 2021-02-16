package executor

import (
	"fmt"

	"github.com/hculpan/kablang/ast"
)

// Executor contains the execution
// environment for this interpreter
type Executor struct {
	Errors []error
	blocks *ast.BlockStack
}

// NewExecutor ...
func NewExecutor() *Executor {
	result := &Executor{blocks: ast.NewBlockStack()}
	result.Reset()
	return result
}

// CurrentBlock returns the current execution block
func (e *Executor) CurrentBlock() *ast.Block {
	return e.blocks.Peek()
}

// Reset sets the execution environment
// back to it's initial state
func (e *Executor) Reset() {
	e.Errors = []error{}
}

// Execute executes the supplies AST
func (e *Executor) Execute(program *ast.Program) {
	if program == nil {
		e.addError(fmt.Errorf("Invalid program"))
		return
	}

	if program.BlockNode == nil {
		e.addError(fmt.Errorf("Program must contain a statement block"))
		return
	}

	e.executeBlock(program.BlockNode)
}

func (e *Executor) executeBlock(block *ast.Block) {
	e.blocks.Push(block)
	if e.CurrentBlock().StatementsNode == nil {
		return
	}

	stmts := e.CurrentBlock().StatementsNode
	for _, s := range stmts.StatementListNode {
		switch s.(type) {
		case *ast.NullStatement:
			// do nothing
		case *ast.PrintStatement:
			e.executePrint(s.(*ast.PrintStatement))
		case *ast.AssignStatement:
			e.executeAssignment(s.(*ast.AssignStatement))
		case *ast.VarStatement:
			e.executeVar(s.(*ast.VarStatement))
		case *ast.Block:
			e.executeBlock(s.(*ast.Block))
		}
	}
	e.blocks.Pop()
}

func (e *Executor) executeVar(s *ast.VarStatement) {
	if symbol := s.SymbolNode; symbol != nil && s.ExpressionNode != nil {
		if symbol, exists := e.CurrentBlock().Symbols.Get(s.SymbolNode.GetName()); exists {
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

	}
}

func (e *Executor) executeAssignment(s *ast.AssignStatement) {
	if s.SymbolNode == nil || s.ExpressionNode == nil {
		e.addError(fmt.Errorf("Attempted invalid assignment operation"))
		return
	}

	if symbol, exists := e.CurrentBlock().Symbols.GetLocal(s.SymbolNode.GetName()); exists {
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
