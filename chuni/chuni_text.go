package chuni

// Text define a chuni text struct.
type Text struct {
	plain string
	chuni string
}

// Set take a plain text and return a chuni text struct.
func (c *Text) Set(s string) *Text {
	c.plain = s
	c.convert()
	return c
}

// TODO
// To convert a plain text to chuni text.
func (c *Text) convert() {
	c.chuni = c.plain
}

func (c *Text) String() string {
	return c.chuni
}
