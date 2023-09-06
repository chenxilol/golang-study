package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
}

func (engine *Engine) ServerHttp(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range w.Header() {
			fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
		}
	default:
		fmt.Fprintf(w, "不知道其类型")
	}

}

func main() {
	engine := new(Engine)

	http.HandleFunc("/", engine.ServerHttp)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// handler echoes r.URL.Path
