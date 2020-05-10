package main

import (
	"html/template"
	"net/http"
	"personal_website/githubapi"

	"github.com/gorilla/mux"
)

var data []githubapi.Repository = githubapi.GetLatestGithubRepos("lennart1s", 4)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.PathPrefix("/res/").Handler(http.FileServer(http.Dir("./")))

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./index.html")
	check(err)

	tmpl.Execute(w, data)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
