package main

import "net/http"

func main() {
	//responsável por attachar as urls(handles)
	mux := http.NewServeMux()

	mux.Handle("/blog", blog{Title: "my blog"})
	mux.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}

type blog struct {
	Title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.Title))
}
