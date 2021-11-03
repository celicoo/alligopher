package tag

// Expression is a Tag expression.
type Expression struct {
	Kind
	Value        []interface{}
	Alternative *Expression
	Options
}

// AppendValue adds one or more values to the end of e.Value.
// This method changes the length of e.Value.
func (e *Expression) AppendValue(values ...interface{})  {
	e.Value = append(e.Value, values...)
}
