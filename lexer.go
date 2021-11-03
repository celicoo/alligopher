package alligopher

var Lexer lexer

type lexer struct {
	input      string
	tokenTypes TokenTypes
	position   position
}

// Init initializes the Lexer with the input and resets its position to its
// initial state.
func (l *lexer) Init(input string) {
	l.input = input
	l.position.Reset()
}

func (l *lexer) SetTokenTypes(tokenTypes TokenTypes) {
	tokenTypes.Validate()
	l.tokenTypes = tokenTypes
}

// Scan scans the input and returns the next token. The end of the input is indicated
// by the EOFToken.
// TODO: Apply maximal munch here!
func (l *lexer) Scan() *token {
	if l.IsAtEnd() {
		return nil
	}
	//for k, v := range l.tokenTypes {
	//	if v.Rule == nil {
	//		continue
	//	}
	//	location := v.Rule(l.input)
	//	// Only if first character is included in the matched substring.
	//	if len(location) > 0 && location[0] == 0 {
	//		// TODO(celicoo): check if won't blow.
	//		value := l.input[location[0]:location[1]]
	//		l.consume(value)
	//		if v.Transform != nil {
	//			value = v.Transform(value)
	//		}
	//		return &token{k, value}
	//	}
	//}
	//err := newSyntaxError("LexerError", "invalid character found", l.input, l.position)
	//reportError(err)
	return nil
}

// IsAtEnd returns whether the Lexer has consumed the whole input.
func (l *lexer) IsAtEnd() bool {
	return len(l.input) == 0
}

// consume consumes the value and updates the position accordingly.
func (l *lexer) consume(value string) {
	if l.IsAtEnd() {
		return
	}
	for _, v := range value {
		l.position.Column++
		if v == '\n' {
			l.position.Line++
			l.position.Column = 0
		}
	}
	// Update the input with the remaining of it.
	l.input = l.input[len(value):len(l.input)]
}
