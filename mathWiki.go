package main

import (
	"html/template"
	"log"
	"net/http"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		return
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", indexHandle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}