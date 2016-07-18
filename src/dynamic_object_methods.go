/*
Limitations of Interfaces Without Fields
https://github.com/JonahBraun/edges/issues
https://github.com/JonahBraun/edges/issues
*/

package main

import "fmt"

// Interfaces in Go all you to specify expected behavior. Unfortunately, the interface
// type can only specify functions, not fields.

// This ommission is noticeable when considering that channels provide functionality!
// They are quite similar to functions in that they can receive and return values.
// For example, consider an interface for broadcasting events. The sender doesn't
// care about a succesful return, a simple channel is all that is necessary.

type Evented interface {
	// Regular method of an interface:
	Broadcast() string

	// Not allowed:
	//Broadcast chan string
	//= syntax error: unexpected chan, expecting semicolon, newline, or }
}

// A real world example is the fsnotify library: https://github.com/howeyc/fsnotify
// It makes two channels available to listen to events on. This behavior implementation
// is impossible to abstract as an interface. The work around is to write extraneous
// getter functions.

type FSNotify interface {
	Event() (Event chan fmt.Stringer)
}

// Closures

// Closures (AKA anonymous functions, functions as values, first class functions)
// are essential for functional programming. Closures are useful because they
// encapsulate variables in scope at definition.

// A struct can have function variables, but although functions, these are still
// variables and so can not constitute an interface.

type Logger interface {

	// Can't have variables, even if they are a function definition. :/
	//Log func(...interface{})
	//= syntax error: unexpected chan, expecting semicolon, newline, or }
}

func main() {
}
