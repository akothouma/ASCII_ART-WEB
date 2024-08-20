package main

import (
	"fmt"
	"net/http"

	handlers "ASCII-WEB/Handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomePageHandler)
	http.HandleFunc("/ascii-art", handlers.GenerateArt)
	fmt.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Server failed to start")
		return
	}
}
