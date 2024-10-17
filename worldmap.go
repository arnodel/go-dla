package main

import "iter"

type locationProps uint8 // Properties of a location on the map

const (
	hasPoint   locationProps = 1 << iota // A point was added here
	isAdjacent                           // A point was added next to here
)

type WorldMap [worldWidth * worldHeight]locationProps

func (m *WorldMap) merge(p Point, props locationProps) {
	if p.X < 0 || p.X >= worldWidth || p.Y < 0 || p.Y >= worldHeight {
		return
	}
	m[p.X*worldHeight+p.Y] |= props
}

func (m *WorldMap) get(p Point) locationProps {
	return m[p.X*worldHeight+p.Y]
}

// Add p to the map.
func (m *WorldMap) Add(p Point) {
	m.merge(p, hasPoint)
	m.merge(p.Translate(1, 0), isAdjacent)
	m.merge(p.Translate(-1, 0), isAdjacent)
	m.merge(p.Translate(0, 1), isAdjacent)
	m.merge(p.Translate(0, -1), isAdjacent)
}

// Contains returns true if p was added to the map.
func (m *WorldMap) Contains(p Point) bool {
	return m.get(p)&hasPoint != 0
}

// Neighbours returns true if the map contains a point one step away from p.
func (m *WorldMap) Neighbours(p Point) bool {
	return m.get(p)&isAdjacent != 0
}

// All iterates over all the points contained in the map.
func (m *WorldMap) All() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for i, props := range m {
			if props&hasPoint != 0 {
				p := Point{X: i / worldHeight, Y: i % worldHeight}
				if !yield(p) {
					return
				}
			}
		}
	}
}
