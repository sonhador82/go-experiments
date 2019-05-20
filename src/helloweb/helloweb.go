package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	log.Println("Test")


	http.HandleFunc("/hello", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "World")
		fmt.Fprintf(w, r.UserAgent())
		})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

