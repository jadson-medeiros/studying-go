package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) { 
		w.Write([]byte("hello world!"))
	})
	mux.Handle("/blog", blog{})

	http.ListenAndServe(":8080", mux)
}

type blog struct {}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("blog"))
}