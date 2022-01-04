package main

import "net/http"

func main() {
	http.HandleFunc("/api/hello", handler)

	http.ListenAndServe(":9090", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{
		"error": false,
		"message": "Hello World3!"
	}`))
}
