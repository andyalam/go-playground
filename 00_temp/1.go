package main

import "fmt"

type Point struct {
	X, Y int
}

var (
	p1       = Point{1, 2}
	p2       = Point{X: 10}
	p3       = Point{}
	pointerP = &Point{3, 2}
)

func main() {
	a := 42
	b := '.'
	fmt.Printf("%T, %T \n", a, b)
	adder()
	namePrinter()
	fmt.Println(add(5, 10))

	// func that returns multiple results
	test1, test2 := swap("hello", "bye")
	fmt.Println(test1, test2)

	// function that has predefined return values
	fmt.Println(sum(10, 25))

	fmt.Println(p1.X + p2.X)

	fmt.Println()
	fmt.Println()

	// SLICE - keep a list of values of the same type
	// Similar to list in python, vector in C++, etc.
	xs := []string{"hello", "this", "self", "f"}
	fmt.Println(xs)
	fmt.Println(xs[3])
	for i, v := range xs {
		fmt.Println(i, v)
	}

	fmt.Println()
	fmt.Println()

	// MAP - stores key, value pairs
	m := map[string]int{"foo": 1, "bar": 101}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}

/*
signature for a function

func (receiver) identifier(parameters) (returns) {

}

*/

func adder() {
	a := 3
	b := 5
	fmt.Println(a, b)
}

func namePrinter() {
	fmt.Println("Andy")
}

// function that returns a specific type
func add(x, y int) int {
	return x + y
}

// function with multiple return values
func swap(a, b string) (string, string) {
	return b, a
}

// function that has predefined return values
func sum(a, b int) (sum int) {
	sum = a + b
	return sum
}
