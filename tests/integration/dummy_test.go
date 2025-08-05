package integration

import (
	"testing"
)

func TestDummy(t *testing.T) {
	result := 2 + 2
	expected := 4

	if result != expected {
		t.Errorf("2+2 = %d, erwartet %d", result, expected)
	}
}
