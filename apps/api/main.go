package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", handler)
	http.HandleFunc("/api/auth/callback", authCallbackHandler)

	http.ListenAndServe(":9090", nil)
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
