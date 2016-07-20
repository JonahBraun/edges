/*
Channel Broadcast by Sending Close
https://github.com/JonahBraun/edges/issues
https://github.com/JonahBraun/edges/issues
*/
// Sending a value on a channel will result in only one receive, regardless of how
// many goroutines are listening on the channel. This makes broadcasting to all
// listening channels impossible without a loop.

// However, closing a channel will result in all listeners receiving the zero value
// for the channel type. For the purpose of signaling, "sending a close" is an effective
// way of broadcast.

// Be careful not to send after you have closed the channel, for:
// "A send on a closed channel proceeds by causing a run-time panic."
// https://golang.org/ref/spec#Send_statements

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

	// Pass the channel instead of enclose to demonostrate no magic.
	go func(c chan struct{}) {
		defer wg.Done()
		<-c
		fmt.Println("close c")
	}(c)

	// This sends the zero value of the channel.
	close(c)
	wg.Wait()

	// Beware not to close again:.
	//close(c)
	//= panic: close of closed channel

	// But you can read from it again, as much as you like.
	<-c
	<-c

	c = make(chan struct{})
	wg.Add(1)

	// When doing this, do not test for a zero value to detect a close! This can be
	// confused with sending the zero value. Instead, use the comma-ok idiom. The
	// second value will be false for a nil or closed channel.
	go func() {
		defer wg.Done()
		_, ok := <-c
		fmt.Println(ok)
	}()
	close(c)

	// The above method is the only method doing this. A naïve approach of
	// many routines listening to one channel does not work.

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

	// We will wait in vain, this script ends in a fatal error.
	wg.Wait()
	//= fatal error: all goroutines are asleep - deadlock!
}
