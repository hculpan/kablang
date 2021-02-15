package ast

// NumberSymbol represents a symbol discovered
// in parsing
type NumberSymbol struct {
	Name       string
	NumberData NumberValue
	dataType   int
}

// NewNumberSymbol ...
func NewNumberSymbol(name string) *NumberSymbol {
	return &NumberSymbol{Name: name, dataType: TypeNumber, NumberData: NewIntNumber(0)}
}

// AsString returns a string representation of the
// symbol
func (s *NumberSymbol) AsString(indent string) string {
	return formatSymbolAsString(s, indent)
}

// ToString converts the data to a string
func (s *NumberSymbol) ToString() string {
	return s.NumberData.ToString()
}

// GetIntValue ...
func (s *NumberSymbol) GetIntValue() int64 {
	return s.NumberData.GetIntValue()
}

// GetFloatValue ...
func (s *NumberSymbol) GetFloatValue() float64 {
	return s.NumberData.GetFloatValue()
}

// GetName ...
func (s NumberSymbol) GetName() string {
	return s.Name
}

// GetDataType returns the data type identifier
func (s *NumberSymbol) GetDataType() int {
	return s.dataType
}

// SetValue value to this symbol; it will determine
// the type
func (s *NumberSymbol) SetValue(value interface{}) {
	s.NumberData.SetValue(value)
}

// Add ...
func (s *NumberSymbol) Add(n2 NumberValue) Number {
	return s.NumberData.Add(n2)
}

// Sub ...
func (s *NumberSymbol) Sub(n2 NumberValue) Number {
	return s.NumberData.Sub(n2)
}

// Mult ...
func (s *NumberSymbol) Mult(n2 NumberValue) Number {
	return s.NumberData.Mult(n2)
}

// Div ...
func (s *NumberSymbol) Div(n2 NumberValue) Number {
	return s.NumberData.Div(n2)
}
