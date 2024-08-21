package handlers

import (
	"html/template"
	"net/http"

	"ASCII-WEB/ascii-art1/printingasciipackage"
)

type ArtDetails struct {
	UserInput  string
	BannerFile string
	Result     string
}

func GenerateArt(w http.ResponseWriter, r *http.Request) {
	// Parse the template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Check for valid HTTP method
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check for correct URL path
	if r.URL.Path != "/ascii-art" {
		http.NotFound(w, r)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Retrieve user input and banner selection
	text := r.FormValue("userInput")
	banner := r.FormValue("banners")

	// Validate user input and banner selection
	if text == "" || banner == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Generate ASCII art
	result1, err := printingasciipackage.PrintingAscii(text, banner)
	if err != nil {
		if err.Error() == "Character not supported" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		http.NotFound(w, r)
		return
	}

	// Prepare data for the template
	neededData := ArtDetails{
		UserInput:  text,
		BannerFile: banner,
		Result:     result1,
	}

	// Execute the template and write the response
	if err := tmpl.Execute(w, neededData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
