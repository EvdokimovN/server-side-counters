package inc

import (
	"fmt"
	"time"
)

// inc implements IncrementerServer interface
type inc struct {
	pauses []chan bool
	nums   []chan int
}

var errOutOfBounds = fmt.Errorf("input must be in range [1,n]")

// NewIncrementer returns Incrementer object
// with i routines assigned
func NewIncrementer(i int) Incrementer {
	if i <= 0 {
		panic("input must be greater than 0")
	}
	return inc{
		pauses: make([]chan bool, i),
		nums:   make([]chan int, i),
	}
}

func (I inc) Start() {
	for i := range I.nums {
		n := make(chan int, 1)
		p := make(chan bool)
		I.nums[i] = n
		I.pauses[i] = p
		go I.increment(n, p)
		n <- 0
	}
}

func (I inc) increment(nums chan int, pause chan bool) {
	// track wheather routine is paused
	paused := false
	for {
		select {
		case <-pause:
			paused = !paused
		default:
			if paused {
				// break and waste processor time
				break
			}
			i := <-nums
			i = i + 1
			nums <- i
			// increment every second to allow viewing of progress
			time.Sleep(1 * time.Second)
		}
	}
}

func (I inc) checkBounds(i int) error {
	if i < 1 || i > len(I.pauses) {
		return errOutOfBounds
	}
	return nil
}

func (I inc) Switch(i int) error {
	err := I.checkBounds(i)
	if err != nil {
		return err
	}
	I.pauses[i-1] <- true
	return nil
}

func (I inc) Peek(i int) (int, error) {
	err := I.checkBounds(i)
	if err != nil {
		return -1, err
	}
	return <-I.nums[i-1], nil
}

func (I inc) Size() int {
	return len(I.nums)
}
