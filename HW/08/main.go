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
	http.HandleFunc("/about", about)
	http.HandleFunc("/faq", faq)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", "index12312")
}

func about(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", "about12312312")
}

func faq(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "faq.gohtml", "faq12312312")
}
