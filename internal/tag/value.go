package tag

type Value []Character

func (v *Value) Append(characters ...Character) {
	*v = append(*v, characters...)
}
