package rover

import "fmt"

type rover struct {
	point     point
	direction direction
}

func (r rover) String() string {
	return fmt.Sprintf("Direction: %v, %v", r.direction, r.point)
}

func receiveCommands(inputRover rover, commands []Command) rover {
	var r = inputRover
	for _, command := range commands {
		switch command {
		case F:
			r = move(r, coordinate(1))
		case B:
			r = move(r, coordinate(-1))
		case L:
			r = turnLeft(r)
		case R:
			r = turnRight(r)
		}
	}
	return r
}

func move(r rover, a coordinate) rover {
	var newPoint point
	switch r.direction {
	case N:
		newPoint = point{
			r.point.x, r.point.y + a,
		}
	case E:
		newPoint = point{
			r.point.x + a, r.point.y,
		}
	case S:
		newPoint = point{
			r.point.x, r.point.y - a,
		}
	case W:
		newPoint = point{
			r.point.x - a, r.point.y,
		}
	}

	return rover{
		point:     newPoint,
		direction: r.direction,
	}

}

func turnLeft(r rover) rover {
	return rover{
		direction: (r.direction + 3) % 4,
		point:     r.point,
	}
}

func turnRight(r rover) rover {
	return rover{
		direction: (r.direction + 1) % 4,
		point:     r.point,
	}
}
