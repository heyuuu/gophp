package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/heyuuu/gophp/php/ast"
	"github.com/heyuuu/gophp/php/ir"
	"github.com/heyuuu/gophp/php/parser"
	"github.com/heyuuu/gophp/php/printer"
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
	mode := request.FormValue("mode")

	astDump, printDump, irDump, parseErr := parseCode(input, mode)
	var parseErrStr string
	if parseErr != nil {
		parseErrStr = parseErr.Error()
	}

	content, err = json.Marshal(struct {
		Code  string
		Input string
		Ast   string
		Print string
		Ir    string
		Error string
	}{
		"api",
		input,
		astDump,
		printDump,
		irDump,
		parseErrStr,
	})

	return
}

func parseCode(code string, mode string) (astDump string, printDump string, irDump string, err error) {
	nodes, err := parser.ParseCode(code)
	if err != nil {
		return
	}

	astDump, err = ast.Sprint(nodes)
	if err != nil {
		return
	}

	printDump, err = printer.SprintFile(nodes)
	if err != nil {
		return
	}

	irNodes := ir.ParseAst(nodes)
	irDump, err = ir.Sprint(irNodes)
	if err != nil {
		return
	}

	return
}
