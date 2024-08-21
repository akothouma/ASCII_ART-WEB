package handlers

import (
	"fmt"
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
		w.WriteHeader(405)
		http.ServeFile(w, r, "templates/405.html")
		return
	}
	if r.URL.Path != "/ascii-art" {
		w.WriteHeader(404)
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
		w.WriteHeader(400)
		http.ServeFile(w, r, "templates/400.html")
		return
	}
	if banner == "" {
		w.WriteHeader(400)
		http.ServeFile(w, r, "templates/400.html")
		return
	}

	result1, err := printingasciipackage.PrintingAscii(text, banner)
	fmt.Println(err)
	if err != nil {
		if err.Error()=="Character not supported"{
			w.WriteHeader(400)
			http.ServeFile(w,r,"templates/400.html")
			return
		}
		w.WriteHeader(404)
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
