package main

import "fmt"

type Person struct {
	first     string
	last      string
	age       int
	prefs     []string
	inventory map[int]string
}

func (p Person) print() {
	fmt.Println(p)
}

// pointer, modifies original object
func (p *Person) setFirst(newFirst string) {
	p.first = newFirst
}

// this function is not tied to the Person struct
// however it can still take in a pointer to person
// and modify it
func clearOutStuff(p *Person) {
	p.first = ""
	p.last = ""
}

type secretAgent struct {
	Person
	ltk bool
}

func main() {

	p1 := Person{
		"James",
		"Bond",
		32,
		[]string{"martini", "women"},
		map[int]string{1: "adasdsa", 2: "12312312"},
	}

	sa1 := secretAgent{
		p1,
		false,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.first)

	p1.print()
	p1.setFirst("Jimmy")
	p1.print()
	clearOutStuff(&p1)
	p1.print()

}
