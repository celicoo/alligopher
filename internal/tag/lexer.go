package tag

import "fmt"

var Lexer lexer

type lexer struct {
	Input    string
	Position uint
}

// Init sets the Lexer input and resets its position to its initial state.
func (l *lexer) Init(input string) {
	l.Input, l.Position = input, 0
}

// Next returns the next token and advances the Lexer Position by 1. If the
// Lexer has run out of characters to tokenize, the Position is not updated
// and the returning kind is EOF.
func (l *lexer) Next() (Kind, Value) {
	for {
		// INFO(celicoo): the call stack isn't really necessary here, and in
		// fact hits the performance of the lexer, but the current strategy
		// should be easier to maintain.
		if l.isAtEnd() {
			return EOF, nil
		} else if l.peek().Is(Whitespace) {
			l.advance()
		} else {
			return l.kind()
		}
	}
}

func (l *lexer) NextIs(kinds ...Kind) bool {
	defer func(position uint) { l.Position = position }(l.Position)
	k, _ := l.Next()
	return k.Is(kinds...)
}

func (l *lexer) kind() (Kind, Value) {
	c := l.peek()
	if c.IsKind() {
		l.advance()
		return Kind(c), nil
	}
	return l.string()
}

func (l *lexer) string() (Kind, Value) {
	c := l.peek()
	if c.Is(SingleQuote) {
		var v Value
		l.advance()
		for {
			c = l.peek()
			if l.isAtEnd() || c.Is(SingleQuote) {
				break
			}
			l.advance()
			v.Append(c)
		}
		if l.isAtEnd() {
			s := fmt.Sprintf("SyntaxError: error parsing tag: missing closing ': %q", l.Input)
			panic(s)
		}
		l.advance()
		return String, v
	}
	return l.identifier()
}

func (l *lexer) identifier() (Kind, Value) {
	c := l.peek()
	if c.IsUppercase() {
		var v Value
		for {
			if l.isAtEnd() {
				break
			}
			c = l.peek()
			if c.IsUppercase() {
				v.Append(c)
				l.advance()
				continue
			}
			break
		}
		return Identifier, v
	}
	s := fmt.Sprintf("SyntaxError: error parsing tag: invalid character %q: %q", c, l.Input)
	panic(s)
}

func (l *lexer) advance() {
	if l.isAtEnd() {
		return
	}
	l.Position++
}

// peek returns the next character the Lexer has yet to tokenize.
func (l *lexer) peek() Character {
	u := l.Input[l.Position]
	return Character(u)
}

// isAtEnd returns whether the Lexer has run out of characters to tokenize.
func (l *lexer) isAtEnd() bool {
	return int(l.Position) == len(l.Input)
}
