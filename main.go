package main

import (
	"fmt"
	"sync"
)

func calcularFactorial(n int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() 

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}

	ch <- factorial 
}

func main() {
	numeros := []int{5, 9, 11, 4, 8} 

	resultados := make(chan int, len(numeros))

	var wg sync.WaitGroup

	for _, num := range numeros {
		wg.Add(1)
		go calcularFactorial(num, resultados, &wg)
	}

	go func() {
		wg.Wait()       
		close(resultados) 
	}()

	
	for res := range resultados {
		fmt.Println("El factorial es:", res)
	}
}