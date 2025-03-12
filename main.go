package main

import (
	"io"
	"log"
	"net/http"
)

func parrot(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var message string
	if q.Get("message") == "" {
		message = "send a message in the query string"
	} else {
		message = q.Get("message")
	}
	io.WriteString(w, message)
}

func main() {
	http.HandleFunc("/", parrot)
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
