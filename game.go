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
	worldImage *ebiten.Image    // We draw the world here
	pending    chan pendingItem // Channel where workers put points to draw
	pointCount int              // Points drawn so far
	stepCount  int              // Steps made so far
	maxPoints  int              // Total number of points to draw
	start      time.Time        // When the game started (to log timings)
}

type pendingItem struct {
	Point
	steps int
}

// NewGame returns a new game where the seeds are initialPoints.
func NewGame(initialPoints iter.Seq[Point], maxPoints int) *Game {
	game := &Game{
		worldImage: ebiten.NewImage(worldWidth, worldHeight),
		pending:    make(chan pendingItem, maxPendingPoints),
		maxPoints:  maxPoints,
	}
	for p := range initialPoints {
		game.worldImage.Set(p.X, p.Y, color.White)
	}
	game.start = time.Now()
	return game
}

// Update implements ebitengine.Game.Update().
func (g *Game) Update() error {
	var (
		i  = g.pointCount
		s  = g.stepCount
		n  = g.maxPoints
		t0 = time.Now()
	)
	if i >= n {
		return nil
	}
	defer func() {
		g.pointCount = i
		g.stepCount = s
	}()
	for i%100 != 0 || time.Since(t0) < 10*time.Millisecond {
		// ^ i%100 is there so don't call time.Since(t0) too much
		select {
		case p := <-g.pending:
			// Start white and gradually fade out to black
			g.worldImage.Set(
				p.X, p.Y,
				color.Gray16{Y: uint16(0xFFFF * (n - i) / n)},
			)
			i++
			s += p.steps
			if i%10000 == 0 {
				log.Printf("Points: %d - Steps = %d, %s", i, s, time.Since(g.start))
			}
			if i >= n {
				pprof.StopCPUProfile()
				return nil
			}
		default:
			return nil
		}
	}
	log.Printf("Ran out of time in Update loop after %d points", i-g.pointCount)
	return nil
}

// Draw implements ebitengine.Game.Draw().
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.worldImage, nil)
}

// Layout implements ebitengine.Game.Layout().
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return worldWidth, worldHeight
}

// AddPoint adds a point to the pending points to draw.
func (g *Game) AddPoint(p Point, steps int) {
	g.pending <- pendingItem{
		Point: p,
		steps: steps,
	}
}
