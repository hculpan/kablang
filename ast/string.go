package ast

import (
	"fmt"
	"strings"
)

// StringValue is for any value that
// can stand in place of a string
// Current implementers:
//    String
//    StringSymbol
type StringValue interface {
	GetValue() string
	SetValue(value interface{})
	AsString(indent string) string
}

// String represents a string terminal
type String struct {
	value string
}

// NewString ...
func NewString(value string) *String {
	return &String{value: strings.Trim(value, "\"")}
}

// AsString return the node as a string
func (s *String) AsString(indent string) string {
	return indent + fmt.Sprintf("String: '%s'", s.value)
}

// GetValue returns the value of this string
func (s *String) GetValue() string {
	return s.value
}

// SetValue sets the literal value of this string
func (s *String) SetValue(value interface{}) {
	switch value.(type) {
	case *String:
		s.value = value.(*String).value
	case String:
		s.value = value.(String).value
	case string:
		s.value = value.(string)
	default:
		panic(fmt.Errorf("Invalid data type for assignment to string : %T", value))
	}

	s.value = strings.Trim(s.value, "\"")
}
