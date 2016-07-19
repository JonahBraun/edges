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

func (f flint) Foo() {
	fmt.Println("Foo!")
}

type FooBar interface {
	Foo()
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
	// First we declare a fmt.Stringer interface variable, then assign to it.
	var e fmt.Stringer = c
	// Note the underlying type and value do not change. Interfaces only declare a
	// set of functions, not the type or contents.
	inspect(e)
	//= main.flint 4

	// We can convert the interface to another interface, since the underlying type
	// and value is not changing, all we are really doing is asserting that the variable
	// has a certain type. We can do this since our flint type has the Foo() function
	// required for the FooBar interface.
	f := e.(FooBar)

	// Again, the type and value have not changed.
	inspect(f)
	//= main.flint 4

	// Type Switch
	// A special type of type assertion, can only be used on interfaces.
	// https://golang.org/doc/effective_go.html#type_switch
	// First case that matches is used, below both fmt.Stringer and interface{} will match,
	// but "stringer" is printed.
	switch g := f.(type) {
	case fmt.Stringer:
		fmt.Println("It's a Stringer", g)
	case FooBar:
		fmt.Println("It's a FooBar", g)
	default:
		fmt.Println("You could throw an error here")
	}
	//= It's a Stringer 4

	// Function types act a lot like interfaces. More on this another time…
}

func init() {
	// Set up log to print file line numbers.
	log.SetFlags(log.Lshortfile)
}

// This convenience function prints a Go representation of the type and value.
func inspect(v interface{}) {
	log.Output(2, fmt.Sprintf("%T %#v \n", v, v))
}
