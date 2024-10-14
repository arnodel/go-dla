package main

import (
	"flag"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime/pprof"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	var (
		cpuprofile string
		npoints    int
		methodName string
	)
	flag.StringVar(&cpuprofile, "cpuprofile", "", "write cpu profile to file")
	flag.IntVar(&npoints, "npoints", 300000, "number of points to draw")
	flag.StringVar(&methodName, "method", "circle", "method")
	flag.Parse()

	method := mustFindMethod(methodName)

	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	worldMap := &WorldMap{}

	method.init(worldMap)

	game := NewGame(worldMap.All(), npoints)

	go AggregatePoints(worldMap, method.pickPoint, game.AddPoint)

	ebiten.SetWindowSize(worldHeight, worldWidth)
	ebiten.SetWindowTitle("Diffraction-limited aggregation")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func mustFindMethod(name string) methodSpec {
	for _, method := range methods {
		if method.name == name {
			return method
		}
	}
	panic("invalid method")
}

// A few ways
type methodSpec struct {
	name      string          // what to call it on the command line
	init      func(*WorldMap) // function to add seeds to the world map
	pickPoint func() Point    // function to pick a random point
}

var methods = []methodSpec{
	{
		name:      "point",
		init:      DrawHorizontalPoints(1),
		pickPoint: RandomPoint,
	},
	{
		name:      "point2",
		init:      DrawHorizontalPoints(2),
		pickPoint: RandomPoint,
	},
	{
		name:      "circle",
		init:      DrawCircle,
		pickPoint: RandomPointInCircle,
	},
	{
		name:      "hline",
		init:      DrawHorizontalLine,
		pickPoint: RandomPoint,
	},
}

func RandomPointInCircle() Point {
	a := rand.Float64() * math.Pi * 2
	r := math.Sqrt(rand.Float64()) * worldWidth / 2
	return Point{
		X: worldWidth/2 + int(math.Cos(a)*r),
		Y: worldHeight/2 + int(math.Sin(a)*r),
	}
}

func RandomPoint() Point {
	return Point{
		X: rand.Intn(worldWidth),
		Y: rand.Intn(worldHeight),
	}
}

func DrawHorizontalLine(m *WorldMap) {
	for x := 0; x < worldWidth; x++ {
		m.Add(Point{X: x, Y: worldHeight / 2})
	}
}

func DrawHorizontalPoints(nPoints int) func(*WorldMap) {
	return func(m *WorldMap) {
		for i := 1; i <= nPoints; i++ {
			m.Add(Point{X: i * worldWidth / (nPoints + 1), Y: worldHeight / 2})
		}
	}
}

func DrawCircle(m *WorldMap) {
	const N = 2000
	for i := 0; i < N; i++ {
		a := math.Pi * 2 * float64(i) / N
		m.Add(Point{
			X: int(worldWidth / 2 * (1 + math.Cos(a))),
			Y: int(worldHeight / 2 * (1 + math.Sin(a))),
		}.Clamp())
	}
}
