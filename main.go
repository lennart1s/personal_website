package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./index.html")
	check(err)
	tmpl.Execute(w, nil)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
