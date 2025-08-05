package integration

import (
	"HostingCoder/src/code"
	"fmt"
	"os"
	"testing"
)

func TestThatCodeIsRunning(t *testing.T) {
	code.Save_and_run_file("with open('test.txt', 'w') as file: file.write('Hello World!')", "python3", "test.py", "test.py")

	actuel, err := os.ReadFile("test.txt")
	exp := "Hello World!"

	if err != nil {
		fmt.Println("File not Exist")
	}

	if string(actuel) != exp {
		t.Errorf("Save_and_run_file() = %s, erwartet %s", exp, string(actuel))
	}
}
