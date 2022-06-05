package main

import (
	"flag"
	"fmt"
	"net/http"
	"shorturl/handler"
)

func main() {

	filePath := flag.String("filePath", "paths.json", "provide filePath for additional configuration")
	flag.Parse()
	mux := defaultMux()

	redirectHandler, err := handler.FileHandler(*filePath, mux)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", redirectHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sorry, i dont know this page : '%s'.", r.URL.Path)
}
