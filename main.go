// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// handler function responds with the current date and time
func handler(w http.ResponseWriter, r *http.Request) {
	// Get the current time and format it
	currentTime := time.Now().Format(time.RFC1123) // e.g., "Wed, 17 Jul 2025 18:30:00 IST"
	fmt.Fprintf(w, "Current Date & Time: %s\n", currentTime)
	log.Printf("Request received from %s - served current time: %s", r.RemoteAddr, currentTime)
}

func main() {
	// Register the handler for the root path "/"
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 8080
	log.Println("Server starting on port 8080...")
	// log.Fatal will log the error and exit if the server fails to start
	log.Fatal(http.ListenAndServe(":8080", nil))
}