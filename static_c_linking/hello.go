// Statically building in C-archive mode
// $ go build -buildmode=c-archive .

package main

import (
	"C"
	"fmt"
)

func main() {}

//export Hello
func Hello() {
	fmt.Println("Hello!")
}
