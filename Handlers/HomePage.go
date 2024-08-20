package handlers

import (
	"html/template"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/" {
		http.ServeFile(w, r, "404.html")
		return
	}

	if r.Method != http.MethodGet {
		http.ServeFile(w, r, "405.html")
		return
	}
	tmpl.Execute(w, nil)
}
