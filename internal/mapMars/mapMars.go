package mapMars

type MarsField struct {
	XMarsSize int
	YMarsSize int
	scents    []Coord
}

type Coord struct {
	X int
	Y int
}

func (m *MarsField) HasScent(x, y int) bool {
	for _, scent := range m.scents {
		if scent.X == x && scent.Y == y {
			return true
		}
	}
	return false
}

func (m *MarsField) AddScent(x, y int) {
	m.scents = append(m.scents, Coord{X: x, Y: y})
}
