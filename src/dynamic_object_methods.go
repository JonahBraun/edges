/*
Limitations of Interfaces Without Fields
*/

package main

// Interfaces in Go all you to specify expected behavior. Unfortunately, the interface
// type can only specify functions, not fields.

// This ommission is noticeable when considering that channels provide functionality!
// They are quite similar to functions in that they can receive and return values.

type Evented interface {
	// Regular method of an interface:
	Broadcast() string

	// Not allowed:
	//Broadcast chan string
	//= syntax error: unexpected chan, expecting semicolon, newline, or }
}

// Closures (AKA anonymous functions, functions as values, first class functions)
// are essential for functional programming. Closures are useful because they
// encapsulate variables in scope at definition.

// A struct can have function variables, but, unfortunately, this behavior can not
// be exported in an abstract interfcase because interfaces can not have fields
// — even when that field is a function!

type Logger interface {

	// Can't have variables, even if they are a function definition. :/
	//Log func(...interface{})
	//= syntax error: unexpected chan, expecting semicolon, newline, or }
}

func main() {
}
