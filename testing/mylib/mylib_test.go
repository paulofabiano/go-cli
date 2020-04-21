// Testing
// $ go test ./mylib/...
// $ go test -list Basics ./mylib/...
// $ go test -run Basic mylib/...
// $ go test -v ./mylib/...
// $ go test -count 10 ./mylib/...

package mylib

import (
	"testing"
)

func Test_BasicChecks(t *testing.T) {
	t.Parallel()
	t.Run("Go can add", func(t *testing.T) {
		if 1+1 != 2 {
			t.Fail()
		}
	})
	t.Run("Go can concatenate strings", func(t *testing.T) {
		if "Hello, " + "Go" != "Hello, Go" {
			t.Fail()
		}
	})
}
