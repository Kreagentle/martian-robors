package robot

import (
	"martian-robots/internal/mapMars"
)

type Robot struct {
	Mars      mapMars.MarsField
	WalkPath  []rune
	X         int // current robot X coordinate position
	previousX int // previous robot X coordinate position
	Y         int // current robot Y coordinate position
	previousY int // previous robot Y coordinate position
	Direction int // robot Direction
	Lost      bool
}

func NewRobot(m *mapMars.MarsField) *Robot {
	return &Robot{
		Mars: *m,
		Lost: false,
	}
}
func (r *Robot) RobotWalk(counter int) {
	if counter == len(r.WalkPath) {
		return
	}

	if r.X > r.Mars.XMarsSize || r.Y > r.Mars.YMarsSize || r.Y < 0 || r.X < 0 {
		r.X, r.Y = r.previousX, r.previousY
		if !r.Mars.HasScent(r.X, r.Y) {
			r.Lost = true
			r.Mars.AddScent(r.X, r.Y)
			return
		}
	}

	r.previousX, r.previousY = r.X, r.Y

	switch r.WalkPath[counter] {
	case 'R':
		r.Direction += 90
		if r.Direction > 270 {
			r.Direction = 0
		}
	case 'L':
		r.Direction -= 90
		if r.Direction < 0 {
			r.Direction = 270
		}
	case 'F':
		switch r.Direction {
		case 0:
			r.X = r.X - 1
		case 90:
			r.Y = r.Y + 1
		case 180:
			r.X = r.X + 1
		case 270:
			r.Y = r.Y - 1
		}
	}
	r.RobotWalk(counter + 1)
}
