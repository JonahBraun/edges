package main

import "runtime"

// every way I've found so far to block forever is with channels
//
// this script should result in a deadlock with nothing being printed
// the error dump for each goroutine will describe why the channel is blocked
func main() {
	// send on a channel with no receive
	go func() {
		<-make(chan struct{})
	}()

	// vice versa
	go func() {
		make(chan struct{}) <- struct{}{}
	}()

	// send on a nil channel
	go func() {
		var a chan bool
		a <- true
	}()

	// vice vesra
	go func() {
		var a chan bool
		<-a
	}()

	// range on a nil channel
	go func() {
		var a chan bool
		for range a {
		}
	}()

	// select with only nil channels
	go func() {
		var a chan bool
		select {
		case a <- true:
		case <-a:
		}
	}()

	// my favorite so far… empty chan selector
	go func() {
		select {}
	}()

	// causes the current goroutine to exit, without exiting func main!
	// this only causes a crash b/c all other goroutines are asleep.
	// we could have done real work and then called os.Exit() in another goroutine,
	// this would be fine.
	runtime.Goexit()
}
