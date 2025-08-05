package code

import (
	"os"
	"os/exec"
)

func Save_and_run_file(code string, language string, file string) {
	_save_file(code)
	_run_file(language, file)
}

func _save_file(code string) {
	err := os.WriteFile("test.py", []byte(code), 0644)
	if err != nil {
		println("WriteFileError", err)
	}
}

func _run_file(language string, file string) {
	com := exec.Command(language, file)
	com.Run()
}
