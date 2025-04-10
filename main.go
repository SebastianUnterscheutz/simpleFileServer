package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve files from the "static" directory
	fileServer := http.FileServer(http.Dir("./static"))

	// Strip the "/static/" prefix when looking for files
	http.Handle("/", http.StripPrefix("/", fileServer))

	// Start the server on port 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
