package main

import (
	"html/template"
	"net/http"
)

const cookieKey string = "myCookie"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/forum/faq", faq)
	http.HandleFunc("/clear", clear)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie(cookieKey)
	c = &http.Cookie{
		Name:     cookieKey,
		Value:    "Hello world",
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, c)
	tpl.ExecuteTemplate(w, "index.gohtml", c)
}

func about(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie(cookieKey)
	tpl.ExecuteTemplate(w, "about.gohtml", c)
}

func faq(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie(cookieKey)
	c = &http.Cookie{
		Name:     cookieKey,
		Value:    "You visited the FAQ page",
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, c)
	tpl.ExecuteTemplate(w, "faq.gohtml", c)
}

func clear(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie(cookieKey)
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
