/*
Title
https://github.com/JonahBraun/edges/issues
https://github.com/JonahBraun/edges/issues
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	p := make([]byte, 1)
	_, err := os.Stdin.Read(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}
