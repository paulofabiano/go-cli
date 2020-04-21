// Testing coverage
// $ go test -cover ./mylib/...
// $ go test -coverpkg mylib,fmt ./mylib
// $ go test -coverprofile cover.out ./mylib/...
// $ go tool cover -html=cover.out
// $ go test -covermode count -coverprofile cover.out ./mylib/...

package mylib

import "fmt"

func adder(l, r int) int {
	fmt.Printf("Adding %v and %v", l, r)
	return l + r
}

func subractor(l, r int) int {
	return l - r
}
