package main

import "fmt"

func main() {
	//buffered channel (canal con bufer)
	c := make(chan int)

	//enviar
	go enviar(c)

	//recibir
	recibir(c)

	fmt.Println("Finalizado")
}

func enviar(c chan<- int) {
	c <- 42
}

func recibir(c <-chan int) {
	fmt.Println(<-c)
}
