package robot

import (
	"martian-robots/internal/mapMars"
	"testing"
)

func TestRobotWalk(t *testing.T) {
	var mars mapMars.MarsField
	mars.XMarsSize, mars.YMarsSize = 5, 3

	robot := NewRobot(&mars)
	robot.X, robot.Y, robot.Direction = 0, 0, 90

	robot.WalkPath = []rune{'F', 'R', 'F', 'L', 'F'}
	robot.RobotWalk(0)
	if robot.X != 1 || robot.Y != 2 || robot.Direction != 90 || robot.Lost {
		t.Errorf("Test case 1 failed: Expected (2, 3, 180, false), got (%d, %d, %d, %t)", robot.X, robot.Y, robot.Direction, robot.Lost)
	}

	// Movement outside the map boundaries
	robot.X, robot.Y, robot.Direction = 0, 3, 0
	robot.WalkPath = []rune{'F', 'F', 'F', 'F', 'F', 'R', 'F', 'F', 'F', 'F', 'F'}
	robot.RobotWalk(0)
	if !robot.Lost {
		t.Errorf("Test case 2 failed: Robot was expected to be lost but it wasn't")
	}

	robot.X, robot.Y, robot.Direction, robot.Lost = 1, 1, 90, false
	robot.WalkPath = []rune{'F', 'L', 'F', 'F', 'R', 'F', 'F', 'L', 'F'}
	robot.RobotWalk(0)
	if robot.X != 0 || robot.Y != 2 || robot.Direction != 0 || !robot.Lost {
		t.Errorf("Test case 3 failed: Expected (2, 2, 270, false), got (%d, %d, %d, %t)", robot.X, robot.Y, robot.Direction, robot.Lost)
	}
}
