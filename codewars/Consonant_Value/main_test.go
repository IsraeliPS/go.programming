package main

import "testing"

func TestConsonantValue(t *testing.T) {

	type prueba struct {
		cadena    string
		respuesta int
	}

	pruebas := []prueba{
		{"a", 0},
		{"bcd", 9},
		{"zea", 26},
		{"az", 26},
		{"baz", 26},
		{"aeiou", 0},
		{"uaoczei", 29},
		{"abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes", 143},
		{"codewars", 37},
		{"bup", 16},
	}

	for _, x := range pruebas {
		v := ConsonantValue(x.cadena...)
		if v != x.respuesta {
			t.Error("Expected", x.respuesta, "Got", v)
		}
	}
}
