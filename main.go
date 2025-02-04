package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	log.Fatal(http.ListenAndServe(server.Addr, server.Handler))
}
