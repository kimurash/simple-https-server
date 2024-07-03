package main

import (
	"fmt"
	"net/http"
)

func startHTTPSServer() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting HTTPS server on :8443")
	if err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil); err != nil {
		fmt.Println("Error starting HTTPS server:", err)
	}
}
