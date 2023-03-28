package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"gophp/php/ast"
	"gophp/php/parser"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	//go:embed static/index.html
	indexHtml []byte
)

var devMode bool

func main() {
	// parse args
	var port int
	flag.IntVar(&port, "p", 8081, "port")
	flag.BoolVar(&devMode, "dev", false, "open dev mode")
	flag.Parse()

	// addr
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	fmt.Printf("Web UI Url:  http://%s/\n\n", addr)

	// start server
	if devMode {
		wd, _ := os.Getwd()
		http.Handle("/", http.FileServer(http.Dir(filepath.Join(wd, "static"))))
	} else {
		http.HandleFunc("/", staticHandler(indexHtml))
	}
	http.HandleFunc("/api", wrapHandler(apiHandler))
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

func apiHandler(request *http.Request) (content []byte, err error) {
	err = request.ParseForm()
	if err != nil {
		return
	}

	if devMode {
		fmt.Printf("%+v\n", request.Form)
	}

	input := request.FormValue("input")

	output, parseErr := parseCode(input)
	var parseErrStr string
	if parseErr != nil {
		parseErrStr = parseErr.Error()
	}

	content, err = json.Marshal(struct {
		Code   string
		Input  string
		Output string
		Error  string
	}{
		"api",
		input,
		output,
		parseErrStr,
	})

	return
}

func parseCode(code string) (string, error) {
	nodes, err := parser.ParseCode(code)
	if err != nil {
		return "", err
	}

	return ast.Sprint(nodes)
}
