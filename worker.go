package main

import (
	"log"

	"math/rand"
)

// AggregatePoints uses pickPoint to choose a starting point, moves it randomly
// until it aggregates, then registers it with addPoint.  It goes forever (or at
// least until it can no longer pick a point not on the map)
func AggregatePoints(
	workerNumber int,
	worldMap *WorldMap,
	pickPoint func() Point,
	addPoint func(Point),
) {
	var randDirSrc RandDirSource
	for {
		p := pickPoint()
		i := 0
		for worldMap.Contains(p) {
			i++
			if i == 100 {
				log.Printf("Worker %d stopping", workerNumber)
				return
			}
			p = pickPoint()
		}
		for !worldMap.Neighbours(p) {
			p = p.Move(randDirSrc.Dir()).Clamp()
		}
		worldMap.Add(p)
		addPoint(p)
	}
}

type RandDirSource struct {
	r uint64
	i int
}

func (s *RandDirSource) Dir() int {
	if s.i == 0 {
		s.i = 32
		s.r = rand.Uint64()
	}
	dir := int(s.r % 4)
	s.i--
	s.r >>= 2
	return dir
}
