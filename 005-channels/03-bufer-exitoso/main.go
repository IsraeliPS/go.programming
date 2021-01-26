package main

import "fmt"

func main() {
	ca := make(chan int, 1)

	ca <- 42

	fmt.Println(<-ca)

}
