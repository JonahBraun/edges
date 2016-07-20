/*
Debug Print
https://play.golang.org/p/A_Kf2l5zdL
https://github.com/JonahBraun/edges/issues/2

Printing a variable is the most immediate debug method. Two ways are demonstrated.
*/

package main

import "fmt"

func main() {
	// The Builtin print() and println()
	// https://golang.org/ref/spec#Bootstrapping
	// It is often necessary to debug a file that has not imported fmt. This is possible
	// with the two builtin print functions.

	// println() is the more useful one as it inserts spaces between multiple variables.
	println("foo", "bar")
	//= foo bar

	println("biz", "baz")
	//= biz baz

	print("foo", "bar")
	print("biz", "baz", "\n")
	//= foobarbizbaz

	// Printing Type And Value
	// https://golang.org/pkg/fmt/
	// %T and %#v are useful verbs in the fmt package for inspecting variables.
	// The inspect() function defined below demonstrates their use. Let's see what
	// it looks like for a string, int, map and array.

	inspect("foo", 4, map[string]string{"foo": "bar"}, [3]int{1, 2, 3})
	//= string "foo"
	//= int 4
	//= map[string]string map[string]string{"foo":"bar"}
	//= [3]int [3]int{1, 2, 3}
}

func inspect(v ...interface{}) {
	for _, v := range v {
		fmt.Printf("%T %#v \n", v, v)
	}
}
