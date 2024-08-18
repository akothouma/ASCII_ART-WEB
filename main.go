package main

import (
	 "ASCII-WEB/ascii-art/ascii-art/printingasciipackage"
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/ascii-art", GenerateArt)
	fmt.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
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

	if r.Method != http.MethodGet{
		http.ServeFile(w,r,"405.html")
		return
	}

	tmpl.Execute(w, nil)
}

func GenerateArt(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("index.html")
	if err !=nil{
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if r.Method !=http.MethodPost{
		http.ServeFile(w,r,"405.html")
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	text := r.FormValue("userInput")
	banner := r.FormValue("banners")

	if text == "" {
		return
	}
	result1 := printingasciipackage.PrintingAscii(text, banner)
	fmt.Println(result1)

	neededData := ArtDetails{
		UserInput:  text,
		BannerFile: banner,
		Result:     result1,
	}

	err = tmpl.Execute(w, neededData) 
	if err!=nil{
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Template execution error:", err)
		return
	}
}
