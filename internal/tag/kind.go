package tag

type Kind Character

// The allowed kinds.
const (
	EOF              Kind = -1
	Identifier       Kind = -2
	String           Kind = -3
	LeftParenthesis  Kind = 40
	RightParenthesis Kind = 41
	Star             Kind = 42
	Plus             Kind = 43
	Question         Kind = 63
	At               Kind = 64
	VerticalBar      Kind = 124
)

func (k Kind) Is(kinds ...Kind) bool {
	for i := range kinds {
		if k == kinds[i] {
			return true
		}
	}
	return false
}
