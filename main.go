package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello –, %s!", request.URL.Path[1:])
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8001"
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
