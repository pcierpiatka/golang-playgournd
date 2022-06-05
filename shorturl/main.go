package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"shorturl/handler"
	"strings"
)

func main() {

	filePath := flag.String("filePath", "paths.json", "provide filePath for additional configuration")
	flag.Parse()
	mux := defaultMux()

	//TODO
	//able to determine what type of filePath
	//extract into function

	extension := strings.Split(*filePath, ".")

	println(extension[1])
	paths, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//TODO
	//select handler base on filePath type

	//bleHandler, err := switch extension[0] {
	//case "json":
	//	handler.JSONHandler(paths, mux)
	//default:
	//	handler.YAMLHandler(paths, mux)
	//}

	yamlHandler, err := handler.JSONHandler(paths, mux)
	if err != nil {
		panic(err)
	}

	//TODO write tests?
	//TODO add database support
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

//func fileHandler(fileExtension string) (http.Handler, error){
//	switch fileExtension {
//	case "json":
//		handler.JSONHandler(paths, mux)
//
//	}
//	return handler.YAMLHandler(paths, mux)
//}
