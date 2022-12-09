package main

import (
	"fmt"
	"sync"
)

/*
Implement the dining philosopher’s problem with the following constraints/modifications.
- There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
- Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
- The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
- In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
- The host allows no more than 2 philosophers to eat concurrently.
- Each philosopher is numbered, 1 through 5.
- When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a
  line by itself, where <number> is the number of the philosopher.
- When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a
  line by itself, where <number> is the number of the philosopher.
*/

// Create chopstick structure
type ChopS struct{ sync.Mutex }

// Create philosopher structure
type Philo struct {
	leftCS, rightCS *ChopS
	id              int
}

func eat(p *Philo, wg *sync.WaitGroup) {
	defer wg.Done()
	// A philosopher can eat up to 3x max
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("Starting to eat", p.id)
		fmt.Println("Finishing eating", p.id)
		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
}

func main() {
	// Create waitgroup (such that not all philosophers eat at once)
	// This acts as the 'host'
	var wg sync.WaitGroup

	// Create array of 5 chopsticks
	Csticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		Csticks[i] = new(ChopS)
	}

	// Create array of 5 philosophers
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{Csticks[i], Csticks[(i+1)%5], i}
	}

	// Let the philosophers eat (2 at a time, at most)
	for i := 0; i < 5; i++ {
		go eat(philos[i], &wg)
		wg.Add(1)
	}
	wg.Wait()
}
