// Profiling
// $ go tool pprof -web mylib.test mem.out
// $ go test -memprofile mem.out -memprofilerate 1 ./mylib
// $ go test -cpuprofile cpu.out -count 1000000 ./mylib
// $ go test -trace trace.out ./mylib
// $ go tool trace trace.out

package mylib

import (
	"fmt"
	"time"
)

type User struct {
	id        int
	name      string
	lastLogin *time.Time
}

func messageWriter(greeting, name string) string {
	message := fmt.Sprintf("%v, %v", greeting, name)
	return message
}
