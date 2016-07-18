/*
Changing Types
https://github.com/JonahBraun/edges/issues
https://github.com/JonahBraun/edges/issues

The different ways a type can be changed in Go.

*/

package main

import (
	"fmt"
	"log"
	"strconv"
)

type flint int

func (f flint) String() string {
	return strconv.Itoa(int(f))
}

func (f flint) Stringify() fmt.Stringer {
	return f
}

func (f flint) Interfacify() interface{} {
	return f
}

func main() {

	// Type Conversion
	// Syntax: newType(var)
	// https://golang.org/doc/effective_go.html#conversions

	// Type conversion changes the type and possibly the underlying value. In other
	// languages this is called type casting.
	a := 6.4
	inspect(a)
	//= float64 6.4

	b := int(a)
	inspect(b)
	//= int 6

	// Type conversion can be used to change from custom to native types.
	c := flint(4)
	inspect(c)
	//= main.flint 4

	d := int(c)
	inspect(d)
	//= int 4

	// Type Assertion (Interface Conversion)
	// Sytnax: var.(interfaceName)
	// https://golang.org/doc/effective_go.html#interface_conversions

	// Type assertion can only be performed on the interface type.
	// obtain an interface from a function and then covert it
	var e fmt.Stringer = c
	//e := c.Stringify()
	inspect(e)
	f := e.(interface{})

	// declare an interface, assign to it and then convert it
	var g fmt.Stringer
	g = c
	inspect(g)
	inspect(g.(interface{}))

	// Type Switch
	// A special type of type assertion, can only be used on interfaces!
	// https://golang.org/doc/effective_go.html#type_switch
	// First case that matches is used, below both fmt.Stringer and interface{} will match,
	// but "stringer" is printed
	switch f := f.(type) {
	case fmt.Stringer:
		fmt.Println("stringer", f)
	case interface{}:
		fmt.Println("interface{}", f)
	default:
		fmt.Println("throw an error or something")
	}

	// function types act a lot like interfaces
}

func init() {
	// Set up log to print file line numbers.
	log.SetFlags(log.Lshortfile)
}

// This convenience function prints a Go representation of the type and value.
func inspect(v interface{}) {
	log.Output(2, fmt.Sprintf("%T %#v \n", v, v))
}
