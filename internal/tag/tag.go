package tag

// New parses the input and returns, if successful, a Tag object that can be
// used to match against text.
func New(input string) Tag {
	Lexer.Init(input)
	return Parser.Parse()
}

// Tag is the representation of a compiled tag.
type Tag struct {
	Expressions []Expression
}

// AppendExpression adds one or more expressions to the end of t.Expressions.
// This method changes the length of t.Expressions.
func (t *Tag) AppendExpression(expressions ...Expression) {
	t.Expressions = append(t.Expressions, expressions...)
}

func (t *Tag) Matches() bool {
	// reggo.New(Lexer.Input).Matches(text)
	return false
}

