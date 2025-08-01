package main

import (
	"fmt"
	"net/http"
)

func testEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Api is funktion")
}

func main() {
	http.HandleFunc("/", testEndpoint)
	http.ListenAndServe(":8000", nil)
}
