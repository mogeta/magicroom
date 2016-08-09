package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/mogeta/magicroom/firebase"
)

type Data struct {
	Title string
	Body  string
}

var datas map[string]firebase.Irdata

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.tpl")
	p := &Data{"title data", "hoge"}
	t.Execute(w, p)
}

func handleIRDatas(rw http.ResponseWriter, req *http.Request) {

	list := firebase.GetList()
	outjson, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(rw, string(outjson))
}

func main() {
	datas = firebase.GetList()
	http.HandleFunc("/", handler)
	http.HandleFunc("/irdata/list", handleIRDatas)
	http.ListenAndServe(":8080", nil)
}
