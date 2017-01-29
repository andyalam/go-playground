package main

import (
	"fmt"

	"github.com/andyalam/go-playground/02_packageexample/util"
)

func main() {
	fmt.Println("test")
	util.Foo()

	// not allowed, bar is unexported
	//util.bar()
}
