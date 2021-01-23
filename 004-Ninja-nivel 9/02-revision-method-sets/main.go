package main

import (
	"fmt"
)

type persona struct {
	nombre string
}

type humano interface {
	hablar()
}

func (p *persona) hablar() {
	fmt.Println("Hola soy:", p.nombre)
}

func diAlgo(h humano) {
	h.hablar()
}

func main() {
	p1 := persona{
		nombre: "israel",
	}
	diAlgo(&p1)
}
