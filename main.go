package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"personal_website/githubapi"
	"personal_website/properties"
	"strings"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

var router *mux.Router
var database *sql.DB

var data = githubapi.GetLatestGithubRepos(properties.GithubListUser, properties.NumGithubCards)

func main() {
	var err error
	properties.Load(os.Args[1])

	if properties.UseDatabase {
		database, err = sql.Open("mysql", properties.SQLLogin)
		check(err)
	}

	router = mux.NewRouter()

	router.HandleFunc("/", indexHandler)

	router.HandleFunc("/contact", contactHandler)
	router.HandleFunc("/contact/submitrequest", submitRequest).Methods("POST")
	router.HandleFunc("/contact/{response}", contactResponseHandler)

	router.HandleFunc("/impressum", impressumHandler)

	router.PathPrefix("/res/").Handler(http.FileServer(http.Dir("./")))

	http.Handle("/", router)
	if !properties.UseSSL {
		err = http.ListenAndServe(":"+properties.Port, nil)
	} else {
		err = http.ListenAndServeTLS(":"+properties.Port, properties.CertPath, properties.KeyPath, nil)
	}
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
	userGrecToken := r.FormValue("g-recaptcha-response")

	var grecResp grecResponse
	grecResp.GetResp(userGrecToken)

	if grecResp.Success {

		if properties.UseDatabase {
			sqlQuery := "INSERT INTO `contact_submissions`(`lastname`, `firstname`, `email`, `message`, `recaptcha_success`) VALUES (?, ?, ?, ?, ?)"
			_, err = database.Exec(sqlQuery, lastname, firstname, email, msg, grecResp.Success)
			check(err) // TODO: stable error handling, send user to /contact/submit_error.html (optionally: error codes)
		} else { // TODO: Logging
			fmt.Println("\n----------------------------------------")
			fmt.Println("Contact Submission: (" + time.Now().String() + ")")
			fmt.Println(lastname + ", " + firstname)
			fmt.Println(email)
			fmt.Println(msg)
		}

	} else {
		fmt.Println("Invalid Contact submission (recaptcha) - " + time.Now().String())
	}

	/*if err != nil {
		http.Redirect(w, r, "/contact/submit_error.html", http.StatusSeeOther)
		log.Print(err)
		return
	}*/

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

type grecResponse struct {
	Success bool `json:"success"`
}

func (g *grecResponse) GetResp(userGrecToken string) (err error) {
	apiPrefix := "https://www.google.com/recaptcha/api/siteverify?secret="
	grecRespData, err := http.Post(apiPrefix+properties.GrecSecret+"&response="+userGrecToken, "application/x-www-form-urlencoded", nil)
	b, err := ioutil.ReadAll(grecRespData.Body)
	json.Unmarshal(b, &g)

	return err
}
