package main

import (
	"fmt"
	"strconv"
)

func inspect(v interface{}) {
	fmt.Printf("%T %#v \n", v, v)
}

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
	// Conversion: newType(var)
	// https://golang.org/doc/effective_go.html#conversions

	// Type conversion, changes underlying value (type casting)
	a := 6.4
	inspect(a)
	b := int(a)
	inspect(b)

	// Type conversion, no underlying change
	c := flint(4)
	inspect(c)
	d := int(c)
	inspect(d)

	// Interface Conversion aka Type Assertion: var.(interfaceName)
	// https://golang.org/doc/effective_go.html#interface_conversions
	// Type assertion can only be done on interfaces- that is, it can not be performed on a type.
	// Interfaces can be obtained by declaring a func return as that interface
	// or by declaring a variable to be that and then assigning a type to it.

	// obtain an interface from a function and then covert it
	e := c.Stringify()
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
