package main

import "iter"

const (
	isContained = 1 << iota
	isAdjacent
)

type WorldMap [worldWidth * worldHeight]uint8

func (m *WorldMap) merge(p Point, val uint8) {
	if p.X < 0 || p.X >= worldWidth || p.Y < 0 || p.Y >= worldHeight {
		return
	}
	m[p.X*worldHeight+p.Y] |= val
}

func (m *WorldMap) get(p Point) uint8 {
	return m[p.X*worldHeight+p.Y]
}

// Add p to the map.
func (m *WorldMap) Add(p Point) {
	m.merge(p, isContained)
	m.merge(p.Translate(1, 0), isAdjacent)
	m.merge(p.Translate(-1, 0), isAdjacent)
	m.merge(p.Translate(0, 1), isAdjacent)
	m.merge(p.Translate(0, -1), isAdjacent)
}

// Contains returns true if p was added to the map.
func (m *WorldMap) Contains(p Point) bool {
	return m.get(p)&isContained != 0
}

// Neighbours returns true if the map contains a point one step away from p.
func (m *WorldMap) Neighbours(p Point) bool {
	return m.get(p)&isAdjacent != 0
}

// All iterates over all the points contained in the map.
func (m *WorldMap) All() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for i, status := range m {
			if status&isContained != 0 {
				p := Point{X: i / worldHeight, Y: i % worldHeight}
				if !yield(p) {
					return
				}
			}
		}
	}
}
