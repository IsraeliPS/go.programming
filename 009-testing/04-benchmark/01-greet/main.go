package main

import "fmt"

func main() {
	fmt.Println(Greet("James"))
}

//Greet funcion que solo saluda
func Greet(s string) string {
	return fmt.Sprint("Hello my dear, ", s)
}
