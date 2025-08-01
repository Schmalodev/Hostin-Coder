package main

import "testing"

func TestDummy(t *testing.T) {
	expected := 2
	calculation := 1 + 1
	if calculation != expected {
		t.Errorf("Greet(\"Go\") = %q; want %q", calculation, expected)
	}
}
