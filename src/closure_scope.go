// Closures enclose the vars available at declaration and not at call time.
// To obtain call time vars, pass parameters.
package main

import "log"

var foo = "outside"
var A = func() {
	log.Println(foo)
}

func factory(a string) func() {
	return func() {
		log.Println(a)
	}
}

func main() {
	log.SetFlags(log.Lshortfile)

	foo = "main"
	A()
	//= main

	// foo is shadowed next while A() enclosed the outside foo
	foo := "main 2"
	A()
	//= main

	log.Println(foo)
	//= main 2

	func() {
		log.Println(foo)
		//= main 2
	}()

	B := func() {
		log.Println(foo)
	}

	B()
	//= main 2

	foo = "main 3"
	B()
	//= main 3

	a := "a string"
	fA := factory(a)
	fA()
	//= a string

	a = "b string"
	fA()
	//= a string

	func(a string) {
		log.Println(a)
	}("new string")
	//= new string
}
