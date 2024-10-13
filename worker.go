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
			p = p.Move(rand.Intn(4)).Clamp()
		}
		worldMap.Add(p)
		addPoint(p)
	}
}
