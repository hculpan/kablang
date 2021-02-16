package ast

import "fmt"

// Data types
const (
	TypeString = iota
	TypeNumber
)

var typeNames []string = []string{
	"string",
	"number",
}

// GetTypeName ...
func GetTypeName(dataType int) string {
	if dataType >= 0 && dataType < len(typeNames) {
		return typeNames[dataType]
	}

	return "unknown"
}

// Symbol interface represents any
// type of symbol
type Symbol interface {
	GetName() string
	GetDataType() int
	SetValue(value interface{})
	AsString(indent string) string
	ToString() string
}

// NewSymbol ...
func NewSymbol(name string, typeID int) Symbol {
	switch typeID {
	case TypeNumber:
		return NewNumberSymbol(name)
	case TypeString:
		return NewStringSymbol(name)
	default:
		panic(fmt.Errorf("Attempt to create symbol with unrecognized type '%d'", typeID))
	}
}

func formatSymbolAsString(s Symbol, indent string) string {
	if len(s.ToString()) > 0 {
		return indent + fmt.Sprintf("Symbol: %-20s  %-12s '%s'", s.GetName(), typeNames[s.GetDataType()], s.ToString())
	}
	return indent + fmt.Sprintf("Symbol: %-20s  %-12s", s.GetName(), typeNames[s.GetDataType()])
}
