package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	//go:embed static/index.html
	indexHtml []byte
)

func main() {
	// parse args
	var port int
	flag.IntVar(&port, "p", 8081, "port")
	flag.Parse()

	// addr
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	fmt.Printf("Web UI Url:  http://%s/\n\n", addr)

	// start server
	http.HandleFunc("/", staticHandler(indexHtml))
	http.HandleFunc("/api", ApiWrapHandler(apiHandler))
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func wrapHandler(handler func(*http.Request) ([]byte, error)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("%s %s\n", request.Method, request.URL.String())
		content, err := handler(request)
		if err != nil {
			writer.WriteHeader(500)
			_, _ = writer.Write([]byte("Server Error: " + err.Error()))
		} else {
			writer.WriteHeader(200)
			_, _ = writer.Write(content)
		}
	}
}

func staticHandler(content []byte) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		_, _ = writer.Write(content)
	}
}
