/*
Variable Encapsulation of Closures
https://play.golang.org/p/mM2rOkI4-n
https://github.com/JonahBraun/edges/issues/1

Closures enclose variables at declaration and not at call time.
To obtain call time vars, pass parameters.
*/

package main

// Use log to print file numbers to make output easier to understand.
import "log"

var foo = "outside"

// Our first closure. It prints the previous declared foo.
var printFoo = func() {
	log.Println(foo)
}

// A closure factory, this produces functions that prints both the outside foo and
// whatever string is passed when creating the new function.
func factory(a string) func() {
	return func() {
		log.Println(foo, a)
	}
}

func main() {
	// Set up log to print file line numbers.
	log.SetFlags(log.Lshortfile)

	// As expected.
	printFoo()
	//= outside

	// We can change the global variable foo. printFoo() references the
	// global variable foo (the variable in scope at decleration).
	foo = "changed"
	printFoo()
	//= changed

	// We create a new foo that shadows the global foo.
	foo := "foo B"
	log.Println(foo)
	//= foo B

	// printFoo() continues to reference the global foo.
	printFoo()
	//= changed

	// This closure is immediately executed. It prints the foo in scope now.
	func() {
		log.Println(foo)
		//= foo B
	}()

	// Let's declare a new printFoo(), and call it B. Just like the pervious closure,
	// it refers to the foo currently in scope.
	B := func() {
		log.Println(foo)
	}

	B()
	//= foo B

	// We can change foo, and see the results
	foo = "foo B changed"
	B()
	//= foo B changed

	// Putting it Together

	// Let's use our factory next.
	a := "Paramastring"
	fA := factory(a)

	// This references the original global foo and the recently passed variable.
	fA()
	//= changed Paramastring

	// We can change the variable we passed, but it was passed by copy (not by reference)
	// and so still has the original value.
	a = "Never Printed"
	fA()
	//= changed Paramastring

	// To explore further, try changing factory() to take a string pointer (*string)
	// Because it holds the pointer, you will be able to change the value after declaration.
}
