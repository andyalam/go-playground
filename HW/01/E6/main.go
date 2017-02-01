// exercise 6
package main

import (
	"fmt"
)

type person struct {
	first string
	age   int
}

func main() {
	p1 := person{"Moneypenny", 24}
	fmt.Println(p1)
	p1.foo()
	p1.walk()
}

func (p person) walk() {
	fmt.Println(p.first + " is walking.")
}

func (p person) foo() {
	fmt.Println("Hello from foo")
}

// func (receiver) identifier(parameters) (returns) {}

// Console Output
/*
{Moneypenny 24}
Hello from foo
Moneypenny is walking.
*/
