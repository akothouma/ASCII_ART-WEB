package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main(){
	http.HandleFunc("/index.html",HomePageHandler)
	fmt.Println("Starting server on port 8000")
	if err:=http.ListenAndServe("8000",nil);err !=nil{
		fmt.Println("Server failed to start")
	}

	
}

func HomePageHandler(w http.ResponseWriter,r *http.Request){
tmpl:=template.Must(template.ParseFiles("index.html"))
tmpl.Execute(w,nil)
return
}