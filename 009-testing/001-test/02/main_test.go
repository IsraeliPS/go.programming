package main

import "testing"

func TestMiSuma(t *testing.T) {
	type prueba struct {
		datos     []int
		respuesta int
	}

	pruebas := []prueba{
		{[]int{2, 4, 6}, 12},
		{[]int{1, 3, 5}, 9},
		{[]int{0, 20, 12}, 32},
		{[]int{-10, 4, 20}, 14},
	}

	for _, x := range pruebas {
		v := miSuma(x.datos...)
		if v != x.respuesta {
			t.Error("Expected", x.respuesta, "Got", v)
		}
	}
}
