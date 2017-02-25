package rover

import "fmt"

type point struct {
	x coordinate
	y coordinate
}

func (p point) String() string {
	return fmt.Sprintf("X: %v, Y: %v", p.x, p.y)
}

type coordinate int
