package ast

// StringSymbol represents a symbol discovered
// in parsing
type StringSymbol struct {
	Name       string
	StringData StringValue
	dataType   int
}

// NewStringSymbol ...
func NewStringSymbol(name string) *StringSymbol {
	return &StringSymbol{Name: name, dataType: TypeString, StringData: NewString("")}
}

// AsString returns a string representation of the
// symbol
func (s *StringSymbol) AsString(indent string) string {
	return formatSymbolAsString(s, indent)
}

// GetValue returns the value of this string
func (s *StringSymbol) GetValue() string {
	return s.StringData.GetValue()
}

// GetName ...
func (s StringSymbol) GetName() string {
	return s.Name
}

// GetDataType returns the data type identifier
func (s *StringSymbol) GetDataType() int {
	return s.dataType
}

// SetValue value to this symbol; it will determine
// the type
func (s *StringSymbol) SetValue(value interface{}) {
	s.StringData.SetValue(value)
}

// ToString ...
func (s *StringSymbol) ToString() string {
	return s.StringData.GetValue()
}
