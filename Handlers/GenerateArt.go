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
		http.Error(w, "Bad Request, input text and select banner file", http.StatusBadRequest)
		return
	}
	// check if banner file exists
	// if _, err := os.Stat("path/to/banner/files/" + banner); os.IsNotExist(err) {
	// 	http.NotFound(w, r) // 404 if banner file is missing
	// 	return
	// }

	// Generate ASCII art
	result1, err := printingasciipackage.PrintingAscii(text, banner)
	if err != nil {
		if err.Error() == "character not supported" {
			http.Error(w, "Bad Request: special characters not allowed(\\v,\\b,\\a,\\r,\\f) ", http.StatusBadRequest)
			// return
		} else {
			// w.Header()
			http.Error(w, "Oops something went wrong on our end. Try again later", http.StatusInternalServerError)
		}
		return
	}

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
