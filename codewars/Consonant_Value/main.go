package main

import (
	"fmt"
	"strings"
)

func main() {
	val := "abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes"
	r := solve(val)
	fmt.Println("\n", r)
}

func ConsonantValue(cadena ...string) int {
	var suma, val int = 0, 0
	var nueva string = cadena
	var vocal []string = []string{"a", "e", "i", "o", "u"}
	var letras = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26}

	for x := 0; x < 5; x++ { //remove vowels
		nueva = strings.ReplaceAll(nueva, vocal[x], " ")
	}
	cambio := strings.Fields(nueva)
	for _, v := range cambio {
		suma = 0
		for _, x := range v {
			suma += letras[string(x)]
		}
		if val < suma {
			val = suma
		}
	}
	return val
}
