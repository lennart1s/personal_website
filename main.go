package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"personal_website/githubapi"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var router *mux.Router

var data = githubapi.GetLatestGithubRepos("lennart1s", 4)

func main() {
	loadWithDependencies("./index.html")

	router = mux.NewRouter()

	router.HandleFunc("/", indexHandler)

	router.HandleFunc("/contact", contactHandler)
	router.HandleFunc("/contact/submitrequest", submitRequest).Methods("POST")
	router.HandleFunc("/contact/{response}", contactResponseHandler)

	router.PathPrefix("/res/").Handler(http.FileServer(http.Dir("./")))

	http.Handle("/", router)
	//err := http.ListenAndServe(":8080", nil)
	err := http.ListenAndServeTLS(":8004", "/etc/ssl/certs/sandbothe.dev.cer", "/etc/ssl/private/sandbothe.dev.key", nil)
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
	check(err)

	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	msg := r.FormValue("message")
	var grec GrecResponse
	json.Unmarshal([]byte(r.FormValue("g-recaptcha-response")), &grec)

	f, err := os.OpenFile("./contactsubmission.txt", os.O_APPEND, 0600)
	check(err)
	defer f.Close()
	f.WriteString(time.Now().String() + "\n")
	f.WriteString("New Contact-Form-Submission: (" + strconv.FormatBool(grec.Success) + ")\n")
	f.WriteString(lastname + ", " + firstname + "\n")
	f.WriteString(email + "\n")
	f.WriteString(msg + "\n\n")

	http.Redirect(w, r, "/contact/submit_success.html", http.StatusSeeOther)
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
