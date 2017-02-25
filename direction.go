package rover

type direction int

func (d direction) String() string {
	switch d {
	case N:
		return "North"
	case S:
		return "South"
	case W:
		return "West"
	case E:
		return "East"
	}
	return ""
}

const (
	N direction = 0
	E direction = 1
	S direction = 2
	W direction = 3
)
