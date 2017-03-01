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
	http.HandleFunc("/1", pic1)
	http.HandleFunc("/2", pic2)
	http.HandleFunc("/3", pic3)
	http.HandleFunc("/4", pic4)
	http.HandleFunc("/5", pic5)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func pic1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/1.png")
}

func pic2(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/2.png")
}

func pic3(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/3.png")
}

func pic4(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/4.png")
}

func pic5(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/5.png")
}
