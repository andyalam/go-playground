package main

import "fmt"

type person struct {
	fName string
	lName string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

// create an interface type that both person and secretAgent implement
type human interface {
	speak() string
}

// declare a func with a parameter of the interface’s type
func logStuff(i human) {
	fmt.Println(i.speak())
}

func (p person) speak() string {
	return "uptown, func you up, uptown func you up"
}

func (sa secretAgent) speak() string {
	return "shaken, not stirred"
}

func main() {
	p1 := person{"Nina", "Simone", 25}
	sa1 := secretAgent{person{"Ian", "Fleming", 42}, false}

	// call that func in main and pass in a value of type person
	logStuff(p1)

	// call that func in main and pass in a value of type secretAgent
	logStuff(sa1)

}

// Console Output
/*
uptown, func you up, uptown func you up
shaken, not stirred
*/
