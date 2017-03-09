package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9000", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("weight")
	if v != "" {
		floatConv, _ := strconv.ParseFloat(v, 64)
		floatConv = floatConv / 2.2
		v = "Your weight in kg is approx. " + strconv.FormatFloat(floatConv, 'f', 2, 64)
		fmt.Println("1")
	}

	fmt.Println("2")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
    <form method="POST">
      <label for="weight">Enter your weight in lbs</label>
      <br>
      <input type="text" name="weight" id="weight">
      <input type="submit">
    </form>
  <br>`+v)
}
