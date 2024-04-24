package mapMars

import "testing"

func TestHasScent(t *testing.T) {
	field := MarsField{scents: []Coord{{X: 1, Y: 1}, {X: 3, Y: 3}}}

	tests := []struct {
		x        int
		y        int
		expected bool
	}{
		{1, 1, true},
		{3, 3, true},
		{2, 2, false},
		{0, 0, false},
	}

	for _, test := range tests {
		if result := field.HasScent(test.x, test.y); result != test.expected {
			t.Errorf("HasScent(%d, %d) = %t, expected %t", test.x, test.y, result, test.expected)
		}
	}
}

func TestAddScent(t *testing.T) {
	field := MarsField{}

	field.AddScent(1, 1)
	field.AddScent(2, 2)

	expectedScents := []Coord{{X: 1, Y: 1}, {X: 2, Y: 2}}

	if len(field.scents) != len(expectedScents) {
		t.Errorf("Number of scents is %d, expected %d", len(field.scents), len(expectedScents))
	}

	for i, scent := range field.scents {
		if scent != expectedScents[i] {
			t.Errorf("Scent at index %d is %v, expected %v", i, scent, expectedScents[i])
		}
	}
}
