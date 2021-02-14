package ast

import "strconv"

// Constants for different types
// of numbers
const (
	IntType = iota
	FloatType
)

// SignedNumber represents both a float
// and a integer type
type SignedNumber struct {
	valueInt   int64
	valueFloat float64
	TypeID     int
}

// NewIntNumber creates an int number
func NewIntNumber(v int64) SignedNumber {
	return SignedNumber{valueInt: v, TypeID: IntType}
}

// NewFloatNumber creates a float number
func NewFloatNumber(v float64) SignedNumber {
	return SignedNumber{valueFloat: v, TypeID: FloatType}
}

// AsInt returns the value as an int64, regardless of type
// Note that if the underlying value is float, this will
// truncate it to an int
func (n *SignedNumber) AsInt() int64 {
	if n.TypeID == IntType {
		return n.valueInt
	}
	return int64(n.valueFloat)
}

// AsFloat returns the value as a float64, regardless
// of the underlying type.
func (n *SignedNumber) AsFloat() float64 {
	if n.TypeID == IntType {
		return float64(n.valueInt)
	}
	return n.valueFloat
}

// AsString return the node as a string
func (n *SignedNumber) AsString() string {
	if n.TypeID == IntType {
		return strconv.FormatInt(n.valueInt, 10)
	}
	return strconv.FormatFloat(n.valueFloat, 'f', -1, 64)
}

// Add adds two numbers, irrespective of type
func (n SignedNumber) Add(n2 SignedNumber) SignedNumber {
	if n.TypeID == FloatType || n2.TypeID == FloatType {
		return NewFloatNumber(n.AsFloat() + n2.AsFloat())
	}
	return NewIntNumber(n.valueInt + n2.valueInt)
}

// Sub subtracts two numbers
func (n SignedNumber) Sub(n2 SignedNumber) SignedNumber {
	if n.TypeID == FloatType || n2.TypeID == FloatType {
		return NewFloatNumber(n.AsFloat() - n2.AsFloat())
	}
	return NewIntNumber(n.valueInt - n2.valueInt)
}

// Mult multiplies two numbers
func (n SignedNumber) Mult(n2 SignedNumber) SignedNumber {
	if n.TypeID == FloatType || n2.TypeID == FloatType {
		return NewFloatNumber(n.AsFloat() * n2.AsFloat())
	}
	return NewIntNumber(n.valueInt * n2.valueInt)
}

// Div divides two numbers
func (n SignedNumber) Div(n2 SignedNumber) SignedNumber {
	if n.TypeID == FloatType || n2.TypeID == FloatType {
		return NewFloatNumber(n.AsFloat() / n2.AsFloat())
	}
	return NewIntNumber(n.valueInt / n2.valueInt)
}
