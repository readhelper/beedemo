package main

import (
	"net/http"
	"io"
	"log"
	"fmt"
)


func helloHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RequestURI, "hello, world!")
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/", helloHandle)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
