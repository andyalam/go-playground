package main

import (
	"fmt"
	"html/template"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

struct book {
	
}

var tpl *template.Template
var DB *mgo.Database
var Book *mgo.Collection

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	s, err := mgo.Dial("mongodb://localhost/bookstore")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("bookstore")
	Books = DB.C("books")

	fmt.Println("Connected to mongodb")
}

func main() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	err := Books.Find(bson.M())

	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
