package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := Person{"Bob", 40}

	tpl.ExecuteTemplate(w, "index.gohtml", p1)
}

func about(w http.ResponseWriter, r *http.Request) {
	xi := []int{1, 10, 50, 40, 20}
	tpl.ExecuteTemplate(w, "about.gohtml", xi)
}

func contact(w http.ResponseWriter, r *http.Request) {
	msi := map[string]int{"w": 5, "h": 10, "z": 4}
	tpl.ExecuteTemplate(w, "contact.gohtml", msi)
}
