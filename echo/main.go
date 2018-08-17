package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	//A simple echo server

	http.HandleFunc("/", handleDefault)

	http.ListenAndServe(":8080", nil)
}

func handleDefault(w http.ResponseWriter, r *http.Request) {

	buf := new([]byte)

	io.ReadFull(r.Body, *buf)

	answer := string(*buf)

	fmt.Printf("ECHO: %s\n", answer)
	fmt.Fprintf(w, "%s\n", answer)
}
