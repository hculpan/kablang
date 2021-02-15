package ast

import (
	"fmt"
	"strconv"
)

// Constants for different types
// of numbers
const (
	IntType = iota
	FloatType
)

// NumberValue is for any value that
// can stand in place of a number
// Current implementers:
//    Number
//    NumberSymbol
type NumberValue interface {
	GetDataType() int
	GetIntValue() int64
	GetFloatValue() float64
	SetValue(value interface{})
	AsString(indent string) string
	ToString() string

	Add(n2 NumberValue) Number
	Sub(n2 NumberValue) Number
	Mult(n2 NumberValue) Number
	Div(n2 NumberValue) Number
}

// Number represents both a float
// and a integer type
type Number struct {
	valueInt   int64
	valueFloat float64
	dataType   int
	numberType int
}

// NewIntNumber creates an int number
func NewIntNumber(v int64) *Number {
	return &Number{valueInt: v, dataType: TypeNumber, numberType: IntType}
}

// NewFloatNumber creates a float number
func NewFloatNumber(v float64) *Number {
	return &Number{valueFloat: v, dataType: TypeNumber, numberType: FloatType}
}

// GetIntValue returns the value as an int64, regardless of type
// Note that if the underlying value is float, this will
// truncate it to an int
func (n *Number) GetIntValue() int64 {
	if n.numberType == IntType {
		return n.valueInt
	}
	return int64(n.valueFloat)
}

// GetFloatValue returns the value as a float64, regardless
// of the underlying type.
func (n *Number) GetFloatValue() float64 {
	if n.numberType == IntType {
		return float64(n.valueInt)
	}
	return n.valueFloat
}

// GetDataType ...
func (n *Number) GetDataType() int {
	return n.dataType
}

// SetValue ...
func (n *Number) SetValue(value interface{}) {
	switch value.(type) {
	case *Number:
		n.valueInt = value.(*Number).valueInt
		n.valueFloat = value.(*Number).valueFloat
		n.numberType = value.(*Number).numberType
	case Number:
		n.valueInt = value.(Number).valueInt
		n.valueFloat = value.(Number).valueFloat
		n.numberType = value.(Number).numberType
	case int:
		n.valueInt = int64(value.(int))
		n.numberType = IntType
	case byte:
		n.valueInt = int64(value.(byte))
		n.numberType = IntType
	case int8:
		n.valueInt = int64(value.(int8))
		n.numberType = IntType
	case int16:
		n.valueInt = int64(value.(int16))
		n.numberType = IntType
	case int32:
		n.valueInt = int64(value.(int32))
		n.numberType = IntType
	case int64:
		n.valueInt = value.(int64)
		n.numberType = IntType
	case float32:
		n.valueFloat = float64(value.(float32))
	case float64:
		n.valueFloat = value.(float64)
	default:
		panic("Invalid data for number")
	}
}

// ToString returns the number formatted as a string
func (n *Number) ToString() string {
	if n.numberType == IntType {
		return strconv.FormatInt(n.valueInt, 10)
	}
	return strconv.FormatFloat(n.valueFloat, 'f', -1, 64)
}

// Add adds two numbers, irrespective of type
func (n Number) Add(n2 NumberValue) Number {
	if n.GetDataType() == FloatType || n2.GetDataType() == FloatType {
		return *NewFloatNumber(n.GetFloatValue() + n2.GetFloatValue())
	}
	return *NewIntNumber(n.GetIntValue() + n2.GetIntValue())
}

// Sub subtracts two numbers
func (n Number) Sub(n2 NumberValue) Number {
	if n.GetDataType() == FloatType || n2.GetDataType() == FloatType {
		return *NewFloatNumber(n.GetFloatValue() - n2.GetFloatValue())
	}
	return *NewIntNumber(n.GetIntValue() - n2.GetIntValue())
}

// Mult multiplies two numbers
func (n Number) Mult(n2 NumberValue) Number {
	if n.GetDataType() == FloatType || n2.GetDataType() == FloatType {
		return *NewFloatNumber(n.GetFloatValue() * n2.GetFloatValue())
	}
	return *NewIntNumber(n.GetIntValue() * n2.GetIntValue())
}

// Div divides two numbers
func (n Number) Div(n2 NumberValue) Number {
	return *NewFloatNumber(n.GetFloatValue() / n2.GetFloatValue())
}

// AsString returns a string representation of the node
func (n *Number) AsString(indent string) string {
	return indent + fmt.Sprintf("Signed number: '%s'", n.ToString())
}
