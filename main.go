package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	fmt.Println("Starting server on port 8080  http://localhost:8080")

	mux.Handle("/", http.FileServer(http.Dir("./")))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}

}