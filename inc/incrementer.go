package inc

import "net/http"

// Incrementer defines operations used for
// conncurently incrementing integers
type Incrementer interface {
	// Start launches process
	Start()
	// Switch pauses/starts nth routine
	Switch(n int) error
	// Peek at value incremented by nth routine
	Peek(n int) (int, error)
	// Size returns number of launched routines
	Size() int
}

type IncrementerServer interface {
	Incrementer
	http.Handler
}
