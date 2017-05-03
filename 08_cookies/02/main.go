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
	fmt.Println(http.MethodPost)
	http.HandleFunc("/read", readCookie)
	http.HandleFunc("/write", writeCookie)
	http.ListenAndServe(":9000", nil)
}

func readCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("myCookie")
	if err != nil {
		// no cookie found
		fmt.Println(err)
	}

	tpl.ExecuteTemplate(w, "read.gohtml", c)
}

func writeCookie(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:     "myCookie",
		Value:    "the cookie has been set...",
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(w, c)
	tpl.ExecuteTemplate(w, "write.gohtml", c)
}
