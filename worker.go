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
			randDirSrc.Next()
			p = p.Move(randDirSrc.Get()).Clamp()
		}
		worldMap.Add(p)
		addPoint(p)
	}
}

type RandDirSource struct {
	r uint64
	i int
}

func (s *RandDirSource) Next() {
	if s.i == 0 {
		s.i = 31
		s.r = rand.Uint64()
	} else {
		s.i--
		s.r >>= 2
	}
}

func (s *RandDirSource) Get() int {
	return int(s.r % 4)
}
