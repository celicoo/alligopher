package alligopher

import "strings"

// TokenType is...
type TokenType struct {
	Rule       func(string) []int
	Transform *func(string) interface{}
}

// TokenTypes is ...
type TokenTypes map[string]TokenType

// Validate validates if the token types keys are all uppercase.
func (t TokenTypes) Validate() {
	for k := range t {
		if k != strings.ToUpper(k) {
			panic(errLowercaseTokenTypesKey)
		}
	}
}
