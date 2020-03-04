package queue

import (
	"runtime/debug"
)

func newThread(j jobs, w PanicWriter) *thread {
	var t thread
	t.j = j
	t.w = w
	return &t
}

type thread struct {
	j jobs
	w PanicWriter
}

func (t *thread) run(j Job) {
	defer t.recover()
	j()
}

func (t *thread) recover() {
	if r := recover(); r == nil {
		return
	}

	t.w("panic recovered\n" + string(debug.Stack()))
}

func (t *thread) work() {
	for job := range t.j {
		t.run(job)
	}
}
