package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Count int
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"こんにちは おはようございます GOマスターになりたい！", 20230531}
	template, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}

	err = template.Execute(w, page)
	if err != nil {
		panic(err)
	}

}

func main() {
	http.HandleFunc("/", ViewHandler)
	http.ListenAndServe(":8080", nil)
}
