package questions

import (
	"fmt"
	"runtime"
)

// Done Что выведет код ?
func Done() {
	runtime.GOMAXPROCS(1)

	done := false

	go func() {
		done = true
	}()

	for !done {
	}
	fmt.Println("finished")
}
