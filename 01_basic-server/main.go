package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	title := "hello"
	body := "asdsadasdsa"
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", title, body)
}
