package code

import (
	"os"
	"os/exec"
)

func Save_and_run_file(event string, language string, file string) {
	_save_file(event)
	_run_file(language, file)
}

func _save_file(msg string) {
	text := msg
	err := os.WriteFile("test.py", []byte(text), 0644)
	if err != nil {
		println("WriteFileError", err)
	}
}

func _run_file(language string, file string) {
	run := language
	com := exec.Command(run, file)
	com.Run()
}
