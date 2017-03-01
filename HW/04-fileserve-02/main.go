package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/1", serve1)
	http.HandleFunc("/2", serve2)
	http.HandleFunc("/3", serve3)
	http.HandleFunc("/4", serve4)
	http.HandleFunc("/5", serve5)
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func serve1(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "1.gohtml", nil)
}

func serve2(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "2.gohtml", nil)
}

func serve3(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "3.gohtml", nil)
}

func serve4(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "4.gohtml", nil)
}

func serve5(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "5.gohtml", nil)
}
