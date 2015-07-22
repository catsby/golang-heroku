package main

import (
	"fmt"
	"net/http"
)

func handler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello –, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", nil)
}
