// Dynamicly building in C-shared mode
// $ go build -buildmode=c-shared .

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
