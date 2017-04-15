/*
# Serve the files in the "starting-files" folder

## To get your images to serve, use:

``` Go
	func StripPrefix(prefix string, h Handler) Handler
	func FileServer(root FileSystem) Handler
```

Constraint: you are not allowed to change the route being used for images in the template file
*/

package main

import (
	"html/template"
	"net/http"
)

func main() {
	ts := http.FileServer(http.Dir("starting-files/public"))
	http.Handle("/public/", http.StripPrefix("/public", ts))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("starting-files/index.html")
	t.Execute(w, nil)
}
