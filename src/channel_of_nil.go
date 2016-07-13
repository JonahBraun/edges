/*
Channel of Nil

You can have nil channels (a channel not yet initialized), and then you can have
an initialized, working channel of nil. This demonstrates the latter.
*/

package main

import "fmt"

func main() {
	a := make(chan nil)

	go func() {
		a <- nil
	}()

	fmt.Println(<-a)
}
