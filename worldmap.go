package main

import "iter"

type WorldMap map[Point]struct{}

// Add p to the map.
func (m WorldMap) Add(p Point) {
	m[p] = struct{}{}
}

// Contains returns true if p was added to the map.
func (m WorldMap) Contains(p Point) bool {
	_, ok := m[p]
	return ok
}

// Neighbours returns true if the map contains a point one step away from p.
func (m WorldMap) Neighbours(p Point) bool {
	return m.Contains(p.Translate(1, 0)) ||
		m.Contains(p.Translate(-1, 0)) ||
		m.Contains(p.Translate(0, 1)) ||
		m.Contains(p.Translate(0, -1))
}

// All iterates over all the points contained in the map.
func (m WorldMap) All() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for p := range m {
			if !yield(p) {
				return
			}
		}
	}
}
