package main

import (
	"HostingCoder/src/code"
	"fmt"
	"net/http"
)

func testEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "api: Go")
}

func codeSaveAndRun(w http.ResponseWriter, r *http.Request) {
	code.Save_and_run_file("with open('test.txt', 'w') as file: file.write('HelloWorld!')", "python3", "test.py")
}

func main() {
	http.HandleFunc("/", testEndpoint)
	http.HandleFunc("/deploy", codeSaveAndRun)
	http.ListenAndServe(":8000", nil)
}
