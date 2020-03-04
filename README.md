# queue

queue is a threaded job library. It allows for an easy way to cap a batch of actions to a given number of goroutines. Reducing the overhead associated with spinning up a new goroutine. 

## Usage
### Queue.New
```go

func ExampleQueue_New() {
	var wg sync.WaitGroup
	// Initialize a queue with two threads and a length of four
	q := New(2, 4)
	// Set waitgroup for eight jobs
	wg.Add(8)

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
```