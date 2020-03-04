package queue

import (
	"fmt"
	"sync"
	"testing"
)

func TestQueue_New(t *testing.T) {
	var wg sync.WaitGroup
	// Set waitgroup for eight jobs
	wg.Add(8)

	// Initialize a queue with two threads and a length of four
	q := New(2, 4)

	for i := 0; i < 8; i++ {
		// Create job with closure (passing index)
		job := func(i int) Job {
			// Return job
			return func() {
				fmt.Printf("Job %d\n", i)
				wg.Done()
			}
		}(i)

		// Insert job
		q.New(job)
	}

	wg.Wait()
}

func TestQueue_New_with_panic(t *testing.T) {
	var wg sync.WaitGroup
	// Set waitgroup for one job
	wg.Add(1)

	// Create PanicWriter
	w := func(str string) {
		fmt.Println("Panic caught! :)")
		wg.Done()
	}

	// Initialize a queue with two threads and a length of four
	q := NewWithWriter(2, 4, w)

	// Create job
	job := func() {
		fmt.Printf("Panic Job\n")
		panic("job panic!")
	}

	// Insert job
	q.New(job)

	// Wait for panic writer to catch panic
	wg.Wait()
}

func ExampleQueue_New() {
	var wg sync.WaitGroup
	// Set waitgroup for eight jobs
	wg.Add(8)

	// Initialize a queue with two threads and a length of four
	q := New(2, 4)

	// Iterate eight times
	for i := 0; i < 8; i++ {
		// Create job with a closure to maintain the value of i
		job := func(i int) Job {
			// Return job
			return func() {
				// Print our stored i value
				fmt.Printf("Job %d\n", i)
				// Notify waitgroup that we have finished our task
				wg.Done()
			}
		}(i)

		// Push job into queue
		q.New(job)
	}

	// Wait until all of our jobs have finished
	wg.Wait()
}
