package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
		"/go-website": "https://go.dev"}
	mapHandler := MapHandler(pathsToUrls, mux)
	fmt.Println("Starting Server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
