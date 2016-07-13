// Closing a channel is the only method of simulatenously communicating to
// all listening goroutines
//
// Best line from the spec: https://golang.org/ref/spec#Send_statements
// "A send on a closed channel proceeds by causing a run-time panic."
//
// This script purposely ends in a deadlock.
package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		<-c
		fmt.Println("close a")
	}()
	// pass the channel instead of enclose to demonostrate no magic
	go func(c chan struct{}) {
		defer wg.Done()
		<-c
		fmt.Println("close c")
	}(c)

	// this sends the zero value of the channel
	close(c)
	wg.Wait()

	// beware not to close again
	//close(c)
	// panic: close of closed channel

	// but you can read from it again, as much as you like
	<-c
	<-c

	c = make(chan struct{})
	wg.Add(1)
	// When doing this, do not test for a zero value to detect a close!
	// Instead, use the comma-ok idiom
	// the second boolean yields false for a nil or closed channel
	go func() {
		defer wg.Done()
		_, ok := <-c
		fmt.Println(ok)
	}()
	close(c)

	// the above method is the only method doing this. A naïve approach of
	// many routines listening to one channel does not work

	s := make(chan string)
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("First goroutine.", <-s)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Second goroutine.", <-s)
	}()
	s <- "This will only print once"

	// We will wait in vain, this script ends in a fatal error
	wg.Wait()
	// fatal error: all goroutines are asleep - deadlock!
}
