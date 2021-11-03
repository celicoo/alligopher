package alligopher

import (
	"errors"
	"fmt"
	"os"
)

var (
	// errLowercaseTokenTypesKey is printed to stderr when a lowercase character
	// is used in the identifier of a token type.
	errLowercaseTokenTypesKey = errors.New("TokenType keys can't be lowercase characters")
)

// reportError prints err to stderr and exits the program with status code 1.
func reportError(err error) {
	_, _ = fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}
