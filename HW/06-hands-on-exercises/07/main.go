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

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("starting-files/templates/*.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("starting-files/public"))
	http.Handle("/resources/", http.StripPrefix("/resources", fs))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
