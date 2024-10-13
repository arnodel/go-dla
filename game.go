package main

import (
	"image/color"
	"iter"
	"log"
	"runtime/pprof"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	worldImage *ebiten.Image
	pending    chan Point
	pointCount int
	maxPoints  int
	start      time.Time
}

func newGame(initialPoints iter.Seq[Point], maxPoints int) *Game {
	game := &Game{
		worldImage: ebiten.NewImage(worldWidth, worldHeight),
		pending:    make(chan Point, maxPendingPoints),
		maxPoints:  maxPoints,
	}
	for p := range initialPoints {
		game.worldImage.Set(p.X, p.Y, color.White)
	}
	game.start = time.Now()
	return game
}

func (g *Game) Update() error {
	if g.pointCount >= g.maxPoints {
		return nil
	}
	t0 := time.Now()
	for {
		for i := 0; i < 100; i++ {
			select {
			case p := <-g.pending:
				g.worldImage.Set(p.X, p.Y, color.Gray16{Y: uint16(65535 * (g.maxPoints - g.pointCount) / g.maxPoints)})
				g.pointCount++
				if g.pointCount%1000 == 0 {
					log.Printf("Points: %d - %s", g.pointCount, time.Since(g.start))
				}
				if g.pointCount >= g.maxPoints {
					pprof.StopCPUProfile()
					return nil
				}
			default:
				return nil
			}
		}
		t := time.Since(t0)
		if t > 10*time.Millisecond {
			break
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.worldImage, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return worldWidth, worldHeight
}

func (g *Game) addPoint(p Point) {
	g.pending <- p
}
