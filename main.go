package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var banners = []string{"shadow.txt", "standard.txt", "thinkertoy.txt"}

func GetLine(filename string, num int) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1 // Line numbers start from 1
	for scanner.Scan() {
		if lineNum == num {
			return scanner.Text()
		}
		lineNum++
	}

	return ""
}

func generateASCIIArt(inputs, banner string) string {
	var result strings.Builder
	inputs = strings.ReplaceAll(inputs, "\n", "\\n")
	slice := strings.Split(inputs, "\\n")
	// Split directly on newline

	// Iterate over each line of input
	for _, input := range slice {
		// if input == "\n" {
		// 	result.WriteString("\n")
		// }
		// Iterate over each line (8 lines per character)
		for line := 1; line < 9; line++ {
			// Iterate over each character in input
			for _, char := range input {
				pos := 1 + int(char-' ')*9 + line
				lineContent := GetLine(banner, pos)
				result.WriteString(lineContent)
			}
			result.WriteString("\n") // New line after each set of characters for this line
		}
	}

	return result.String()
}
// Home page handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, banners)
}

// ASCII art handler
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	result := generateASCIIArt(text, banner)

	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, result)
}

// Main function
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	// fmt.Println("Server is running on :8080")
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
