package main

import (
	"fmt"

	"gitlab.com/israeli/go.programming/008-documentation/perro"
)

type canino struct {
	nombre string
	edad   int
}

func main() {
	c1 := canino{
		nombre: "Chester",
		edad:   perro.Edad(2),
	}
	fmt.Println(c1)
}
