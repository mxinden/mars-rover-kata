package rover

type Command int

func (c Command) String() string {
	switch c {
	case F:
		return "forward"
	case B:
		return "backwards"
	case L:
		return "left"
	case R:
		return "right"
	}
	return ""
}

const (
	F Command = iota
	B Command = iota
	L Command = iota
	R Command = iota
)
