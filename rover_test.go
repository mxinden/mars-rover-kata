package rover

import (
	"errors"
	"testing"
)

func TestRoverInitialization(t *testing.T) {
	r := rover{}

	if r.direction != N {
		t.Fatal(errors.New("facing wrong direction"))
	}

	if r.point.x != 0 || r.point.y != 0 {
		t.Fatal(errors.New("starting from wrong position"))
	}
}

func TestMoveForward(t *testing.T) {
	r := rover{}
	r = receiveCommands(r, []Command{F})

	if r.point.x != 0 || r.point.y != 1 {
		t.Fatal(r.point)
	}
}

func TestTurnRight(t *testing.T) {
	r := rover{}
	r = receiveCommands(r, []Command{R})

	if r.direction != E {
		t.Fatal(r.direction)
	}
}

func TestTurnRightMoveForward(t *testing.T) {
	r := rover{}
	r = receiveCommands(r, []Command{R, F})
	if r.direction != E || r.point.x != 1 {
		t.Fatal(r)
	}
}

func TestTurnRightMoveBackwards(t *testing.T) {
	r := rover{}
	r = receiveCommands(r, []Command{R, B})
	if r.direction != E || r.point.x != -1 {
		t.Fatal(r)
	}
}

func TestTurnLeftMoveForward(t *testing.T) {
	r := rover{}
	r = receiveCommands(r, []Command{L, F})
	if r.direction != W || r.point.x != -1 {
		t.Fatal(r)
	}
}
