/*
Title
https://github.com/JonahBraun/edges/issues
https://github.com/JonahBraun/edges/issues
*/
package main

import (
	"fmt"
)

type A struct {
	bar string
}

func (A) foo() {
	fmt.Println("A")
}

type B A

func (b B) foo() {
	fmt.Println("B")

	// The first two don't work:

	//b.A.foo()
	// Error: b.A undefined (type B has no field or method A)

	//B.A.foo()
	// Error: B.A undefined (type B has no method A)

	// But you can convert the type first:
	A(b).foo()

	b.bar = "bar b"
	fmt.Println(b.bar)

	// is a a clone of b?
	a := A(b)
	a.bar = "bar a"
	fmt.Println(A(b).bar)
	fmt.Println(b.bar)
	fmt.Println(a.bar)
}

type C struct {
	A
}

func (c C) foo() {
	fmt.Println("C")

	// “When we embed a type, the methods of that type become methods of the outer type, but when they are invoked, the receiver of the method is the inner type, not the outer one.” - Effective Go
	// Let's access the embedded type:
	c.A.foo()

	// Type casting does not work:
	//A(c).foo()
}

func main() {
	b := new(B)
	b.foo()

	c := new(C)
	c.foo()
}
