package queue

// New returns a new Queue
func New(threads, length int) *Queue {
	return NewWithWriter(threads, length, stderrWriter)
}

// NewWithWriter accepts a custom PanicWriter and returns a new Queue
func NewWithWriter(threads, length int, w PanicWriter) *Queue {
	var q Queue
	q.j = make(jobs, length)
	q.w = w
	q.spawnThreads(threads)
	return &q
}

// Queue manages workers
type Queue struct {
	j jobs

	w PanicWriter
}

func (q *Queue) spawnThreads(threads int) {
	for i := 0; i < threads; i++ {
		t := newThread(q.j, q.w)
		go t.work()
	}
}

// New will insert a new Job into the queue
func (q *Queue) New(js ...Job) {
	for _, j := range js {
		q.j <- j
	}
}
