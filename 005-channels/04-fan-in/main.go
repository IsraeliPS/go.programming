package main

import "fmt"

func main() {

	salir := make(chan int)
	par := make(chan int)
	impar := make(chan int)

	//enviar
	go enviar(par, impar, salir)

	//recibir
	recibir(par, impar, salir)

	fmt.Println("Finalizado")
}

func enviar(p, i, s chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}

	s <- 0
}

func recibir(p, i, s <-chan int) {
	for {
		select {
		case v := <-p:
			fmt.Println("Desde el canal Par\t\t", v)
		case v := <-i:
			fmt.Println("Desde el canal Impar\t\t", v)
		case v := <-s:
			fmt.Println("Desde el canal Salir\t\t", v)
			return
		}
	}
}
