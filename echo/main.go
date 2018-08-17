package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	//A simple echo server

	hostname, _ := os.Hostname()

	fmt.Printf("Starting echo server on %s\n", hostname)

	http.HandleFunc("/", handleDefault)
	http.ListenAndServe(":8080", nil)
}

func handleDefault(w http.ResponseWriter, r *http.Request) {

	var answer string
	message := r.URL.Query().Get("msg")

	if message != "" {
		answer = message
	} else {
		buf, err := ioutil.ReadAll(r.Body)

		if err != nil {
			fmt.Printf("ERROR: %s\n", err)
			fmt.Fprintf(w, "ERROR: %s\n", err)
			return
		}

		answer = string(buf)
	}

	fmt.Printf("ECHO: %s\n", answer)
	fmt.Fprintf(w, "%s\n", answer)
}
