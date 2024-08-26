package main

import (
	"fmt"
	"net/http"
	"os"

	handlers "ASCII-WEB/Handlers"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("usage: go run main.go")
		os.Exit(0)
	}
	http.HandleFunc("/", handlers.HomePageHandler)
	fmt.Println("Starting server on http://localhost:5000")
	http.HandleFunc("/ascii-art", handlers.GenerateArt)
	port := ":5000"
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("\nServer failed to start \n %s\n", err)
		return
	}
}
