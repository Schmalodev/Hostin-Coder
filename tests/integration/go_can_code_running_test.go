package integration

import (
	"net/http"
	"testing"
)

func TestThatCodeIsSendAndRun(t *testing.T) {
	http.Get("http://localhost:8080/deploy")
}
