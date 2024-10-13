package main

import (
	"log"

	"math/bits"
	"math/rand"
)

// AggregatePoints uses pickPoint to choose a starting point, moves it randomly
// until it aggregates, then registers it with addPoint.  It goes forever (or at
// least until it can no longer pick a point not on the map)
func AggregatePoints(
	workerNumber int,
	worldMap *WorldMap,
	pickPoint func() Point,
	addPoint func(Point, int),
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
		steps := 0
		for !worldMap.CannotFreeWalkFrom(p) {
			p = p.Translate(RandomFreeWalk()).Clamp()
			steps += freeWalkSize
		}
		for !worldMap.Neighbours(p) {
			randDirSrc.Next()
			p = p.MoveAndClamp(randDirSrc.Get())
			steps++
		}
		worldMap.Add(p)
		addPoint(p, steps)
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

// Find in one go the displacement (x, y) after N random steps of a 2D random
// walk.
//
// N is a constant integer, 1 <= N <= 32 (32 because we need 2N random bits from
// a total of 64 bits).
func RandomFreeWalk() (x int, y int) {
	const N = freeWalkSize
	r := rand.Uint64()
	k := bits.OnesCount32(uint32(r << (32 - N)))
	l := bits.OnesCount32(uint32(r >> (64 - N)))
	x = N - k - l
	y = l - k
	return
}
