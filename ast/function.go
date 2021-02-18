package ast

// Parameter is the definition
// for a function parameter
type Parameter struct {
	Name     string
	DataType int
}

// SystemFunctionCall is the function signature for built-in functions
type SystemFunctionCall func([]interface{}) interface{}

// Function represents a function
// definition
type Function struct {
	Name           string
	Parameters     []Parameter
	ReturnDataType int
	FunctionCall   SystemFunctionCall
}

// NewFunction ...
func NewFunction() *Function {
	return &Function{}
}
