package ast

import (
	"fmt"
	"strings"
)

// String represents a string terminal
type String struct {
	Value string
}

// NewString ...
func NewString(s string) *String {
	return &String{Value: strings.Trim(s, "\"")}
}

// AsString return the node as a string
func (s *String) AsString(indent string) string {
	return indent + fmt.Sprintf("String: '%s'", s.Value)
}
