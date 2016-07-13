// you can't return a send
package main

import "fmt"

func foo() {
	a := make(chan bool)
	return a<-true
}
func main() {
	fmt.Println(foo())
}
