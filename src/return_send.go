/*
Channel Sending is Second Class

Sending on a channel does not return a value (not even nil). This means it can
not be assigned, passed or returned.

This program does not compile.
*/

package main

import "fmt"

func inspect(v interface{}) {
	fmt.Printf("%T %#v \n", v, v)
}

func foo() {
	a := make(chan bool)

	// Fail:
	b := a<-true
	//= a <- true used as value

	// Fail:
	inspect(a<-true)
	//= a <- true used as value

	// Fail:
	return a<-true
	//= a <- true used as value
	//= too many arguments to return
}

func main() {
	foo()
}
