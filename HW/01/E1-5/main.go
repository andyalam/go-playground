package main

import (
	"fmt"
	"strconv"
)

// Exercise 3
type person struct {
	fName string
	lName string
}

// Exercise 5
type personF struct {
	person
	favFood []string
}

func main() {

	// Exercise 1
	fmt.Println("Exercise 1:")

	xi := []int{4, 5, 90, 40, 2}
	fmt.Println(xi)
	for k := range xi {
		fmt.Println(k)
	}
	for k, i := range xi {
		fmt.Println(k, i)
	}

	// Exercise 2
	fmt.Println("\n\nExercise 2:")

	msi := map[string]int{
		"red":   5,
		"blue":  15,
		"green": 0,
	}
	fmt.Println(msi)
	for k := range msi {
		fmt.Println(k)
	}
	for k, i := range msi {
		fmt.Println(k + ":" + strconv.Itoa(i))
	}

	// Exercise 4
	fmt.Println("\n\nExercise 4:")
	p1 := person{"Andy", "Alameddine"}
	fmt.Println(p1)
	fmt.Println(p1.fName)

	// Exercise 5
	fmt.Println("\n\nExercise 5:")
	myFavFood := []string{"sushi", "pizza", "potato salad"}
	p2 := personF{p1, myFavFood}
	fmt.Println(p2.favFood)
	for _, food := range p2.favFood {
		fmt.Println(food)
	}

}

// Console Output
/*
Exercise 1:
[4 5 90 40 2]
0
1
2
3
4
0 4
1 5
2 90
3 40
4 2


Exercise 2:
map[red:5 blue:15 green:0]
blue
green
red
green:0
red:5
blue:15


Exercise 4:
{Andy Alameddine}
Andy


Exercise 5:
[sushi pizza potato salad]
sushi
pizza
potato salad
*/
