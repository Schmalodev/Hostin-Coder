package main

import (
	"HostingCoder/src/code/FileSaver"
	"encoding/json"
	"fmt"
	"net/http"
)

type CodeSaverData struct {
	Code     string
	Language string
	Data     string
	FileName string
}

func testEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "api: Go")
}

func codeSaveAndRun(w http.ResponseWriter, r *http.Request) {
	var codeSaver CodeSaverData

	err := json.NewDecoder(r.Body).Decode(&codeSaver)
	if err != nil {
		fmt.Println("ERROR! no response")
	}

	FileSaver.Save_and_run_file(codeSaver.Code, codeSaver.Language, codeSaver.Data, codeSaver.FileName)
}

func main() {
	http.HandleFunc("/", testEndpoint)
	http.HandleFunc("/deploy", codeSaveAndRun)
	http.ListenAndServe(":8000", nil)
}
