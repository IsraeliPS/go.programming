package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Sistema Operativo: \t", runtime.GOOS, "\nArquitectura: \t\t", runtime.GOARCH)
}
