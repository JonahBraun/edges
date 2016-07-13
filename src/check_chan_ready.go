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
