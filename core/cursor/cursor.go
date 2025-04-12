package cursor

// / ( String in memory )
// /
// / input:      [  H  |  e  |  l  |  l  |  o  ]
// /               0      1      2      3      4
// /               ^      ^
// /               |      |
// /      currIdx ─┘      └─ nextIdx
// /
// / currChar = 'H'
type Cursor struct {
	input    []rune
	currIdx  int
	nextIdx  int
	currChar rune
}

func New(input string) *Cursor {
	c := &Cursor{
		input: []rune(input),
	}

	return c
}

func (c *Cursor) ReadChar() rune {
	if c.nextIdx >= len(c.input) {
		c.currChar = 0
	} else {
		c.currChar = c.input[c.nextIdx]
	}

	c.currIdx = c.nextIdx

	c.nextIdx++
	return c.currChar
}

func (c *Cursor) Peek() rune {
	if c.nextIdx >= len(c.input) {
		return 0
	}

	return c.input[c.nextIdx]
}

func (c *Cursor) CurrChar() rune {
	return c.currChar
}
