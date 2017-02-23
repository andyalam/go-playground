package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

type person struct {
	First string
	Last  string
}

type pet struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/test", pic)
	http.HandleFunc("/test.jpg", servePic)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	data := struct {
		People []person
		Pets   []pet
	}{
		[]person{person{"h", "h"}},
		[]pet{pet{"1", 1}},
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func pic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/test.jpg"/>`)
}

func servePic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/img/test.jpg")
}
