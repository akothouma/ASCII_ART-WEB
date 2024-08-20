package pkg

import (
	"net/http"
	"text/template"
)

var banners = []string{"shadow.txt", "standard.txt", "thinkertoy.txt"}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
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
