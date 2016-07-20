/*
Checking If A Channel Is Ready
https://github.com/JonahBraun/edges/issues
https://github.com/JonahBraun/edges/issues

You can check if a channel is ready with a default case on a channel selector.
*/

package main

import (
	"fmt"
)

func checkChan(c chan struct{}) {
	select {
	case <-c:
		fmt.Println("we got value")
	default:
	}
	fmt.Println("not ready")
}

func main() {
	c := make(chan struct{}, 1)
	checkChan(c)

	c <- struct{}{}

	checkChan(c)
}
