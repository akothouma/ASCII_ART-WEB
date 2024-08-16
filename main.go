package main

import (
	"fmt"
	"net/http"
	"text/template"

	"ASCII-WEB/ascii-art/printingasciipackage"
)

func main() {
	http.HandleFunc("/", HomePageHandler)
	fmt.Println("Starting server on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println("Server failed to start")
		return
	}
}

type ArtDetails struct {
	UserInput  string
	BannerFile string
	Result     string
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodPost {
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Println("Template execution error:", err)
			return
		}
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Template execution error:", err)
		return
	}

	text := r.FormValue("userInput")
	banner := r.FormValue("banners")

	if text == "" {
		return
	}
	result1 := printingasciipackage.PrintingAscii(text, banner)

	neededData := ArtDetails{
		UserInput:  text,
		BannerFile: banner,
		Result:     result1,
	}

	if err := tmpl.Execute(w, neededData); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Template execution error:", err)
		return
	}
}
