package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func set(w http.ResponseWriter, r *http.Request) {

	// Set the cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "customCookie",
		Value: "Any value you want",
	})

	// Just some help text
	fmt.Fprintln(w, "Cookie written - check your browser")
	fmt.Fprintln(w, "in chrome go to : dev tools / application / cookies")
}

func read(w http.ResponseWriter, r *http.Request) {

	// Retrieve cookie off the request
	c, err := r.Cookie("customCookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Your cookies, sir/ma'am: ", c)
}
