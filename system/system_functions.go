package system

import (
	"fmt"

	"github.com/hculpan/kablang/ast"
)

// BuiltInFunctions contains a list of all built-in functions
var BuiltInFunctions map[string]*ast.Function

// NewSystemFunction creates a new built-in function
func NewSystemFunction(
	name string,
	params []ast.Parameter,
	returnType int,
	functionCall ast.SystemFunctionCall) *ast.Function {
	result := &ast.Function{
		Name:           name,
		Parameters:     params,
		ReturnDataType: returnType,
		FunctionCall:   functionCall,
	}
	BuiltInFunctions[name] = result
	return result
}

// InitSystemFunctions loads all the system functions
func InitSystemFunctions() {
	NewSystemFunction("PrintHello", []ast.Parameter{}, ast.TypeString, func([]interface{}) interface{} {
		fmt.Println("Hello")
		return 0
	})
}
