package main

import (
	"html/template"
	"net/http"
)

type person struct {
	First string
	Last  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/1", handler1)
	http.HandleFunc("/2", handler2)
	http.HandleFunc("/3", handler3)
	http.HandleFunc("/4", handler4)
	http.HandleFunc("/5", handler5)
	http.HandleFunc("/6", handler6)
	http.HandleFunc("/7", handler7)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

// pass a value of type int to a template and display the value in the template
func handler1(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "1.gohtml", 42)
}

// pass a value of type string to a template and display the value in the template
func handler2(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "2.gohtml", "Hello World")
}

// pass a value of type bool to a template and display the value in the template
func handler3(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "3.gohtml", true)
}

// pass a value of type []string to a template and display the values in the template
func handler4(w http.ResponseWriter, r *http.Request) {
	xs := []string{"hello", "world", "i", "am", "computer"}
	tpl.ExecuteTemplate(w, "4.gohtml", xs)
}

// pass a value of type map to a template and display the values in the template
func handler5(w http.ResponseWriter, r *http.Request) {
	msi := map[string]int{
		"wat":      5,
		"no":       10,
		"lkasdk":   2,
		"asdsadsa": 0,
	}
	tpl.ExecuteTemplate(w, "5.gohtml", msi)
}

// pass a value of type map to a template and display the keys and values in the template
func handler6(w http.ResponseWriter, r *http.Request) {
	msi := map[string]int{
		"wat":      5,
		"no":       10,
		"lkasdk":   2,
		"asdsadsa": 0,
	}
	tpl.ExecuteTemplate(w, "6.gohtml", msi)
}

// create a struct person. create some values of type person.
// add those values to a slice of person.
// pass the value of type []person to a template and display all of the
// values in the template
func handler7(w http.ResponseWriter, r *http.Request) {
	p1 := person{"Andy", "A"}
	p2 := person{"Joe", "J"}
	p3 := person{"Jane", "J"}

	xpeople := []person{p1, p2, p3}

	tpl.ExecuteTemplate(w, "7.gohtml", xpeople)
}
