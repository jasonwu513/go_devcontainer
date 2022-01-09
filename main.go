package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, hello my go pracice</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "go blog practice"+
			"<a href=\"mailto:jasonwu513@gmail.com\">jasonwu513@gmail.com</a>")
	} else {
		fmt.Fprint(w, "<h1> page not found :(</h1>"+
			"<p> if any problem please contact us。</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
