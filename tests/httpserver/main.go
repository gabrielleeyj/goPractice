package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// gorilla/mux way to handlefunc
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've request the book: %s on page %s\n", title, page)
	})

	// net/http method handler function
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, you are at the %s path \n", r.URL.Path)
	// })

	// creating a http fileserver
	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Starting Server at Port:80")
	// start the server referencing to gorilla/mux
	http.ListenAndServe(":80", r)

	// Original ListenAndServe method
	// http.ListenAndServe(":80", nil)
}
