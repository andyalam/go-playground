/*
# Serve the files in the "starting-files" folder

## To get your images to serve, use only this:

``` Go
	fs := http.FileServer(http.Dir("public"))
```

Hint: look to see what type FileServer returns, then think it through.
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
	http.Handle("/pics/", fs)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
