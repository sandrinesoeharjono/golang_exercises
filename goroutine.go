package main

import (
	"fmt"
	"sync"
)

// Write two goroutines which have a race condition when executed concurrently.
// Explain what the race condition is and how it can occur.

// global variable
var a int

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	a = a + 1
}

func main() {
	// Set value of 1 to a
	a = 1

	// Create a collection of goroutines
	var go_wg sync.WaitGroup

	// Alternate between adding to the wait group counter & incrementing a+=1
	go_wg.Add(1)
	go increment(&go_wg)
	go_wg.Add(1)
	go increment(&go_wg)
	go_wg.Wait()
	fmt.Println(a)
}

/*
A race condition is when two concurrent contexts simultaneously read and/or write to the same variables,
unexpectedly changing the outputs. In this case, the two concurrent goroutines access the same variable
when run, thereby stepping on each other's toes and ruining each other's work.

In this case, the final value of 'a' could be 2 or 3.
*/
