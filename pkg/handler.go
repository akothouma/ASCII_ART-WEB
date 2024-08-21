package pkg

import (
	"net/http"
	"text/template"
)

var banners = []string{"shadow.txt", "standard.txt", "thinkertoy.txt"}

// handles http requests and responses
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		http.ServeFile(w, r, "404.html")
		return
	}
	//check the HTTP method
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "internal server Error", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, banners)
}

// ASCII art handler
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
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

	result := GenerateASCIIArt(text, banner)

	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		http.Error(w, "Internal Server Error popo", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, result)
}
