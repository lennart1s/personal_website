package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"personal_website/githubapi"
	"strings"
	"log"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

var router *mux.Router
var database *sql.DB
var sqlLogin = "phpmyadmin:argumentkillanalysis@tcp(127.0.0.1:3306)/sandbothe_dev"

var data = githubapi.GetLatestGithubRepos("lennart1s", 4)

func main() {
	var err error
	database, err = sql.Open("mysql", sqlLogin)
	check(err)

	router = mux.NewRouter()

	router.HandleFunc("/", indexHandler)

	router.HandleFunc("/contact", contactHandler)
	router.HandleFunc("/contact/submitrequest", submitRequest).Methods("POST")
	router.HandleFunc("/contact/{response}", contactResponseHandler)

	router.HandleFunc("/impressum", impressumHandler)

	router.PathPrefix("/res/").Handler(http.FileServer(http.Dir("./")))

	http.Handle("/", router)
	//err = http.ListenAndServe(":8080", nil)
	err = http.ListenAndServeTLS(":8004", "/etc/ssl/certs/sandbothe.dev.cer", "/etc/ssl/private/sandbothe.dev.key", nil)
	check(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := loadWithDependencies("./index.html")

	tmpl.Execute(w, data)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := loadWithDependencies("./contact/index.html")

	tmpl.Execute(w, nil)
}

func contactResponseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tmpl := loadWithDependencies("./contact/" + vars["response"])

	tmpl.Execute(w, nil)
}

func submitRequest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/contact/submit_error.html", http.StatusSeeOther)
		log.Print(err)
		return
	}

	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	msg := r.FormValue("message")
	var grec GrecResponse
	json.Unmarshal([]byte(r.FormValue("g-recaptcha-response")), &grec)
	println(grec.Success)

	_, err = database.Exec("INSERT INTO `contact_submissions`(`lastname`, `firstname`, `email`, `message`, `recaptcha_success`) VALUES (?, ?, ?, ?, ?)",
		lastname, firstname, email, msg, grec.Success)
	if err != nil {
		http.Redirect(w, r, "/contact/submit_error.html", http.StatusSeeOther)
		log.Print(err)
		return
	}

	http.Redirect(w, r, "/contact/submit_success.html", http.StatusSeeOther)
}

func impressumHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := loadWithDependencies("./impressum/index.html")

	tmpl.Execute(w, nil)
}

//////////////////////////////////////

func loadWithDependencies(htmlPath string) *template.Template {
	data, err := ioutil.ReadFile(htmlPath)
	check(err)
	html := string(data)

	dependencies := []string{htmlPath}

	for _, p := range strings.Split(html, "{{") {
		if strings.HasPrefix(p, "template") {
			name := "./res/html/"
			inside := false
			for _, c := range p {
				if c == '"' && !inside {
					inside = true
				} else if c == '"' {
					break
				} else if inside {
					name += string(c)
				}
			}
			dependencies = append(dependencies, name)
		}
	}

	tmpl, err := template.ParseFiles(dependencies...)
	check(err)

	return tmpl
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type GrecResponse struct {
	Success bool `json:"success"`
}
