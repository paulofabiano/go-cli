// Testing go cli:
// - go run & args
// - go build & args
// - go install & args

package main

import "github.com/pluralsight/go/cli/hello"

//export main
func main() {
	hello.SayHello()
}
