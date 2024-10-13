package main

type Point struct {
	X, Y int
}

// Move the point by one step (dir is expected to be 0, 1, 2 or 3)
func (p Point) Move(dir int) Point {
	a := dir & 1
	b := dir >> 1
	return Point{
		X: p.X + a - b,
		Y: p.Y + a + b - 1,
	}
}

// Clamp to the point to within the confines of the world
func (p Point) Clamp() Point {
	if p.X < 0 {
		p.X = 0
	} else if p.X >= worldWidth {
		p.X = worldWidth - 1
	}
	if p.Y < 0 {
		p.Y = 0
	} else if p.Y >= worldHeight {
		p.Y = worldHeight - 1
	}
	return p
}

// Translate the point by x and y
func (p Point) Translate(x, y int) Point {
	return Point{X: p.X + x, Y: p.Y + y}
}
