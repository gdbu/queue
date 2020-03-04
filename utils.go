package queue

import "os"

func stderrWriter(str string) {
	os.Stderr.WriteString(str)
}

// PanicWriter is used when an output other than stderr is desirable
type PanicWriter func(string)
