package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { 
		_ =  http.ListenAndServe(":8080", nil);
	})
}