package main

import (
	"fmt"
	"httpServer/pkg"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", pkg.HomeHandler)
	http.HandleFunc("/ascii-art", pkg.AsciiArtHandler)
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("error: Starting server:", err)
	}
}
