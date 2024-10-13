package main

import "iter"

type WorldMap [worldWidth * worldHeight]bool

// Add p to the map.
func (m *WorldMap) Add(p Point) {
	if p.X < 0 || p.X >= worldWidth || p.Y < 0 || p.Y >= worldHeight {
		return
	}
	m[p.X*worldHeight+p.Y] = true
}

// Contains returns true if p was added to the map.
func (m *WorldMap) Contains(p Point) bool {
	i := p.X*worldHeight + p.Y
	if i >= 0 && i < worldHeight*worldWidth {
		return m[i]
	}
	return false
}

// Neighbours returns true if the map contains a point one step away from p.
func (m *WorldMap) Neighbours(p Point) bool {
	return m.Contains(p.Translate(1, 0)) ||
		m.Contains(p.Translate(-1, 0)) ||
		m.Contains(p.Translate(0, 1)) ||
		m.Contains(p.Translate(0, -1))
}

// All iterates over all the points contained in the map.
func (m *WorldMap) All() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for i, added := range m {
			if added {
				p := Point{X: i / worldHeight, Y: i % worldHeight}
				if !yield(p) {
					return
				}
			}
		}
	}
}
