package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type homeData struct {
	Title string
	User string
}
var tmpl = template.Must(template.ParseFiles("templates/home.html"))

func home(w http.ResponseWriter, r*http.Request) {
	data := homeData{Title: "go fist project", User: "paul"}
	tmpl.Execute(w, data)
}
func dev(w http.ResponseWriter, r*http.Request) {
	fmt.Fprintln(w, "dev")
}

func main() {
// when someone visits the site use HandleFunc
http.HandleFunc("/", home) 
http.HandleFunc("/dev", dev)
fmt.Println("server running on http://localhost:8080/dev")
http.ListenAndServe(":8080", nil)
	}
