// Using plugins
// $ go build -buildmode=plugin -o=plugin.so ./plugin/plugin.go
package main

import "fmt"

//ThingToDo is just a simple function
func ThingToDo() {
	fmt.Println("Executing action...")
}
