package tag

import "unicode"

type Character rune

const (
	Whitespace  Character = 32
	SingleQuote Character = 39
)

// IsUppercase returns whether c is an upper case letter.
func (c Character) IsUppercase() bool {
	r := rune(c)
	return unicode.IsUpper(r)
}

func (c Character) Is(characters ...Character) bool {
	for i := range characters {
		if c == characters[i] {
			return true
		}
	}
	return false
}

// IsKind returns whether c can be considered a Kind.
func (c Character) IsKind() bool {
	k := Kind(c)
	return k.Is(LeftParenthesis, RightParenthesis, Star, Plus, Question, At, VerticalBar)
}
