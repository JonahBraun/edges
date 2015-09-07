package main

import "fmt"

func main() {
	a := make(chan nil)

	go func() {
		a <- nil
	}()

	fmt.Println(<-a)
}
