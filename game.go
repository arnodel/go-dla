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
	points [batchSize]Point
	steps  int
}

// NewGame returns a new game where the seeds are initialPoints.
func NewGame(initialPoints iter.Seq[Point], maxPoints int) *Game {
	game := &Game{
		worldImage: ebiten.NewImage(worldWidth, worldHeight),
		pending:    make(chan pendingItem, maxPendingPoints/batchSize),
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
	for k := 0; k%100 != 0 || time.Since(t0) < 10*time.Millisecond; k++ {
		// ^ k%100 is there so don't call time.Since(t0) too much
		select {
		case item := <-g.pending:
			for _, p := range item.points {
				// Start white and gradually fade out to black
				g.worldImage.Set(
					p.X, p.Y,
					color.Gray16{Y: uint16(0xFFFF * (n - i) / n)},
				)
			}
			s += item.steps
			i += batchSize
			if i%10000 < batchSize {
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
func (g *Game) AddBatch(points *[batchSize]Point, steps int) {
	g.pending <- pendingItem{
		points: *points,
		steps:  steps,
	}
}

type PointBatcher struct {
	game   *Game
	points [batchSize]Point
	steps  int
	i      int
}

func newPointBatcher(game *Game) *PointBatcher {
	return &PointBatcher{game: game}
}
func (b *PointBatcher) AddPoint(p Point, steps int) {
	b.points[b.i] = p
	b.steps += steps
	b.i++
	if b.i == batchSize {
		b.game.AddBatch(&b.points, b.steps)
		b.i = 0
		b.steps = 0
	}
}
