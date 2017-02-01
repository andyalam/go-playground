package main

import "fmt"

type person struct {
	last string
	age  int
}

type secretAgent struct {
	person
	ltk bool
}

// create a struct that holds person fields
type personHolder struct {
	people []person
}

// create a struct that holds secret agent fields and embeds person type
type team struct {
	agents []secretAgent
	boss   person
}

func main() {
	// create a variable of type person
	p1 := person{"Alameddine", 21}

	// create a variable of type secret agent
	sa1 := secretAgent{
		person{"Doe", 30},
		true,
	}

	// print a field from person
	fmt.Println(p1.last)

	// run pSpeak attached to the variable of type person
	p1.pSpeak()

	// print a field from secret agent
	fmt.Println(sa1.ltk)

	// run saSpeak attached to the variable of type secret agent
	sa1.saSpeak()

	// run pSpeak attached to the variable of type secret agent
	sa1.pSpeak()
}

// attach a method to person: pSpeak
func (p person) pSpeak() {
	fmt.Println("I am a person with the last name", p.last)
}

// attach a method to secret agent: saSpeak
func (sa secretAgent) saSpeak() {
	fmt.Println("I am a secret agent named", sa.person.last)
}

// Console Outpout
/*
Alameddine
I am a person with the last name Alameddine
true
I am a secret agent named Doe
I am a person with the last name Doe
*/
