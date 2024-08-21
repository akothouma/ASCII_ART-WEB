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
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "templates/405.html")
		return
	}
	if r.URL.Path != "/ascii-art" {
		http.ServeFile(w, r, "templates/404.html")
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	text := r.FormValue("userInput")
	banner := r.FormValue("banners")

	if text == "" {
		http.ServeFile(w, r, "templates/400.html")
		return
	}

	result1, err := printingasciipackage.PrintingAscii(text, banner)
	if err != nil {
		http.ServeFile(w, r, "templates/404.html")
		return
	}

	neededData := ArtDetails{
		UserInput:  text,
		BannerFile: banner,
		Result:     result1,
	}

	err = tmpl.Execute(w, neededData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
