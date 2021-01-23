package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("Número de CPUs al inicio: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Gorutinas al inicio: %v\n", runtime.NumGoroutine())
	var wg sync.WaitGroup

	contador := 0

	const gs int = 100
	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			x := contador
			x++
			contador = x
			fmt.Println(contador)
			m.Unlock()
			wg.Done()
		}()
		fmt.Printf("Número de CPUs en medio: %v\n", runtime.NumCPU())
		fmt.Printf("Número de Gorutinas en medio: %v\n", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Printf("Cuenta: %v\n", contador)
	fmt.Printf("Número de CPUs al final: %v\n", runtime.NumCPU())
	fmt.Printf("Número de Gorutinas al final: %v\n", runtime.NumGoroutine())
}
