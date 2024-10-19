package main

const (
	worldWidth  = 1000
	worldHeight = 1000

	maxPendingPoints = 20000
	batchSize        = 20

	// Size of a "free walk" - a walk that we can perform in one go because
	// there are no obstacles within that number of steps.  It must be at most
	// 32.
	freeWalkSize = 16
)
