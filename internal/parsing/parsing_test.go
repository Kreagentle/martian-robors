package parsing

import (
	"fmt"
	"testing"
)

func TestParseMap(t *testing.T) {
	tests := []struct {
		input         string
		expectedX     int
		expectedY     int
		expectedError error
	}{
		{"5 5", 5, 5, nil},
		{"10 20", 10, 20, nil},
		{"invalid", 0, 0, fmt.Errorf("invalid input format: expected two space-separated components")},
		{"5", 0, 0, fmt.Errorf("invalid input format: expected two space-separated components")},
	}

	for _, test := range tests {
		x, y, err := ParseMap(test.input)
		if x != test.expectedX || y != test.expectedY || err != nil && (test.expectedError == nil || err.Error() != test.expectedError.Error()) {
			t.Errorf("ParseMap(%s) = (%d, %d, %v), expected (%d, %d, %v)", test.input, x, y, err, test.expectedX, test.expectedY, test.expectedError)
		}
	}
}

func TestParseRobot(t *testing.T) {
	tests := []struct {
		input         string
		expectedX     int
		expectedY     int
		expectedAngle int
		expectedError error
	}{
		{"1 1 N", 1, 1, 90, nil},
		{"2 3 S", 2, 3, 270, nil},
		{"10 10 E", 10, 10, 180, nil},
		{"invalid", 0, 0, 0, fmt.Errorf("invalid input format: expected three space-separated components")},
		{"1 1 NS", 0, 0, 0, fmt.Errorf("invalid rune input: expected a single valid (N, E, W, S) rune")},
		{"1 1 Z", 0, 0, 0, fmt.Errorf("invalid rune input: expected a single valid (N, E, W, S) rune")},
	}

	for _, test := range tests {
		x, y, angle, err := ParseRobot(test.input)
		if x != test.expectedX || y != test.expectedY || angle != test.expectedAngle || err != nil && (test.expectedError == nil || err.Error() != test.expectedError.Error()) {
			t.Errorf("ParseRobot(%s) = (%d, %d, %d, %v), expected (%d, %d, %d, %v)", test.input, x, y, angle, err, test.expectedX, test.expectedY, test.expectedAngle, test.expectedError)
		}
	}
}
