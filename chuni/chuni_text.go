package chuni

// Text define a chuni text struct.
type Text struct {
	plain string
	chuni string
}

func (t *Text) String() string {
	return t.chuni
}

// Set take a plain text and return a chuni text struct.
func (t *Text) Set(p string) *Text {
	t.plain = p
	t.convert()
	return t
}

// TODO
// To convert a plain text to chuni text.
func (t *Text) convert() {
	t.chuni = t.plain
}
