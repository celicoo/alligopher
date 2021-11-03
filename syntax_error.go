package alligopher

import (
	"bytes"
	"strings"
	"text/template"
)

// newSyntaxError constructs and returns a new syntaxError.
func newSyntaxError(input string, position position) syntaxError {
	var indentation, i uint
	if position.Column == 0 {
		i = 1
	} else {
		i = position.Column
	}
	for i != 0 {
		i /= 10
		indentation++
	}
	indentation++
	return syntaxError{input, indentation, position}
}

// syntaxError is used by the Lexer when an invalid character is found or
// by the Parser when it cannot continue to construct the abstract syntax tree
// when an unexpected token is found.
type syntaxError struct {
	Input       string
	Indentation uint
	position
}

func (p syntaxError) Error() string {
	var (
		t template.Template
		b bytes.Buffer
	)
	t.Funcs(template.FuncMap{ "split": strings.Split }).Parse(syntaxErrorText)
	t.Execute(&b, p)
	return b.String()
}

var syntaxErrorText = `
SyntaxError: invalid character found
{{printf "%*[2]s" .Indentation ""    -}} > {{.Column}}
{{printf "%*[2]s" .Indentation ""    -}} |
{{printf "%*[2]d" .Indentation .Line -}} | {{index (split .Input "\n") 0}}
{{printf "%*[2]s" .Indentation ""    -}} | ^
`
