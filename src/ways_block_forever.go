/*
Blocking Forever
https://play.golang.org/p/wNYrob63tp

A go routine can be blocked from running, typically when waiting to send or
receive on a channel. Here we explore all the ways to block forever…

This script results in a deadlock. The dump for each channel will explain why
the channel is blocked.
*/

package main

import "runtime"

func main() {

	// We will start with the obvious ones.
	// Send on a channel with no receive.
	go func() {
		<-make(chan struct{})
	}()

	// Receive on a channel with no send.
	go func() {
		make(chan struct{}) <- struct{}{}
	}()

	// Send on a nil channel (the channel has not been initialized).
	go func() {
		var a chan bool
		a <- true
	}()

	// Receive on a nil channel.
	go func() {
		var a chan bool
		<-a
	}()

	// Range on a nil channel.
	go func() {
		var a chan bool
		for range a {
		}
	}()

	// Select with only nil channels.
	go func() {
		var a chan bool
		select {
		case a <- true:
		case <-a:
		}
	}()

	// Here's my favorite so far, a simple one liner: the empty channel selector!
	go func() {
		select {}
	}()

	// To get some descriptive dumps, we cause the the current goroutine to exit,
	// without exiting func main! This causes a fatal error because all
	// go routines are asleep, never to be awoken from their eternal blocks.
	runtime.Goexit()
}
