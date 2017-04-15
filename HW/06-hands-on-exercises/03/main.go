/*
Serve the files in the "starting-files" folder
Use "http.FileServer"
*/

package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("starting-files"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
