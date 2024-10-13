package main

import "iter"

type WorldMap map[Point]struct{}

func (m WorldMap) Add(p Point) {
	m[p] = struct{}{}
}

func (m WorldMap) Contains(p Point) bool {
	_, ok := m[p]
	return ok
}

func (m WorldMap) Neighbours(p Point) bool {
	return m.Contains(p.Translate(1, 0)) ||
		m.Contains(p.Translate(-1, 0)) ||
		m.Contains(p.Translate(0, 1)) ||
		m.Contains(p.Translate(0, -1))
}

func (m WorldMap) All() iter.Seq[Point] {
	return func(yield func(Point) bool) {
		for p := range m {
			if !yield(p) {
				return
			}
		}
	}
}
