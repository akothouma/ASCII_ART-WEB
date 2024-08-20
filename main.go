package main

import (
	"fmt"
	"httpServer/pkg"
	"net/http"
)

// Home page handler

// Main function
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", pkg.HomeHandler)
	http.HandleFunc("/ascii-art", pkg.AsciiArtHandler)

	// fmt.Println("Server is running on :8080")
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
