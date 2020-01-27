package main

import (
	"net/http"
	"html/template"
)

// var temp will save all template pages 
// pkg `template` == encasule all templates and retur template or err 
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main()  {
	http.HandleFunc("/", index) // call index
	http.ListenAndServe(":8000", nil)  // call server 
}

func index(w http.ResponseWriter, r *http.Request) { 
	temp.ExecuteTemplate(w, "Index", nil)
}