package main

import (
	"fmt"
	"net/http"
)

func startHTTPServer() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting HTTP server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting HTTP server:", err)
	}
}
