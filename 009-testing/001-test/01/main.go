package main

import "testing"

//TestMiSuma ... {
func TestMiSuma(t *testing.T) {
	v := MiSuma(1, 3)
	if v != 4 {
		t.Error("Expected", 4, "Got", v)
	}
}
