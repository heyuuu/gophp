package main

import (
	"embed"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

var (
	//go:embed static
	staticFS embed.FS
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
	http.HandleFunc("/", staticHandler())
	http.HandleFunc("/api", ApiWrapHandler(apiHandler))
	http.HandleFunc("/api/test/list", ApiWrapHandler(apiTestListHandler))
	http.HandleFunc("/api/test/run", ApiWrapHandler(apiTestRunHandler))

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

var defaultContentType = map[string]string{
	".html": "text/html",
	".js":   "text/javascript",
	".css":  "text/css",
}

func staticHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[1:]
		if path == "" {
			path = "static/index.html"
		}

		content, err := staticFS.ReadFile(path)
		if err != nil {
			writer.WriteHeader(404)
			_, _ = writer.Write([]byte("404 page not found"))
			return
		}

		if contentType, ok := defaultContentType[filepath.Ext(path)]; ok {
			writer.Header().Set("Content-Type", contentType)
		}

		writer.WriteHeader(200)
		_, _ = writer.Write(content)
	}
}
