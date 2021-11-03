package alligopher

// position is used by the Lexer to update its current position when consuming
// the input.
type position struct {
	Line   uint
	Column uint
}

// Reset sets the position Line and Column to their initial value.
func (p position) Reset() {
	p.Line, p.Column = 0, 0
}
