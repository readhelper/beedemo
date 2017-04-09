package main

import (
	"net/http"
	"io"
	"log"
	"fmt"
)

var server = &http.Server{}

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RequestURI, "hello, world!")
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
