package main

import (
	"log"

	"math/rand"
)

func calcPointsMap(worldMap *WorldMap, pickPoint func() Point, addPoint func(Point)) {
	for {
		p := pickPoint()
		i := 0
		for worldMap.Contains(p) {
			i++
			if i == 100 {
				log.Printf("Stopping")
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
