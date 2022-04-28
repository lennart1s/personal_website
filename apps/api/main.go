package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// http.HandleFunc("/api/hello", handler)
	// http.HandleFunc("/api/auth/callback", authCallbackHandler)
	// http.HandleFunc("")
	r.HandleFunc("/api/hello", handler)
	r.HandleFunc("/api/users", userPOSTHandler)

	// http.Handle("/", r)
	http.ListenAndServe(":9090", r)
}

func userPOSTHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New User registered")
	fmt.Printf("%s\n\n", formatRequest(r))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{
		"error": false,
		"message": "Hello World!"
	}`))
}

func authCallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)

	var data []byte
	buffer := make([]byte, 512)
	for n, err := r.Body.Read(buffer); n > 0 && err != io.EOF; n, err = r.Body.Read(buffer) {
		data = append(data, buffer[:n]...)
	}
	fmt.Println(string(data))

	url := "http://localhost:8080"
	code, ok := r.URL.Query()["code"]
	if ok {
		url += "?code=" + code[0]
	}

	http.Redirect(w, r, url, 301)
}

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}

	var data []byte
	buffer := make([]byte, 512)
	for n, err := r.Body.Read(buffer); n > 0; n, err = r.Body.Read(buffer) {
		data = append(data, buffer[:n]...)
		if err != nil && err != io.EOF {
			fmt.Println(err.Error())
			return "error"
		}
	}
	request = append(request, string(data))

	// Return the request as a string
	return strings.Join(request, "\n")
}
