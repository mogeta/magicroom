package main

import (
	"html/template"
	"net/http"

	"github.com/mogeta/magicroom/firebase"
)

type Data struct {
	Title string
	Body  string
}

var data string

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.tpl")
	p := &Data{"title data", data}
	t.Execute(w, p)
}

func main() {
	data = firebase.Get("tv")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
