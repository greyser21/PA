package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("web")))

	log.Println("UpcycleConnect server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}