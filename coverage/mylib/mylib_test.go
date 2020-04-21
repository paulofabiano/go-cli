// Benchmarking
// $ go test -bench Adder ./mylib
// $ go test -bench . ./mylib
// $ go test -benchtime 5s -bench . ./mylib
// $ go test -bench . -benchmem ./mylib

package mylib

import (
	"testing"
)

func TestAdder(t *testing.T) {
	if adder(2, 5) != 7 {
		t.Fail()
	}
}
