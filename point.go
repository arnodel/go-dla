package main

type Point struct {
	X, Y int
}

// Move the point by one step (dir is expected to be 0, 1, 2 or 3) and clamp it
// to the confines of the world. It is assumed that the initial value of p is
// within the confines of the world.
func (p Point) MoveAndClamp(dir int) Point {
	a := dir & 1
	b := dir >> 1
	x := p.X + a - b
	y := p.Y + a + b - 1
	if uint(x) < worldWidth {
		// We get to this branch if and only if 0 <= x < worldwidth as if x < 0
		// then uint(x) wraps around to a very big number (bigger than
		// worldWidth)
		p.X = x
	}
	if uint(y) < worldHeight {
		// As above we get to this branch if and only if 0 <= y < worldHeight
		p.Y = y
	}
	return p
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
