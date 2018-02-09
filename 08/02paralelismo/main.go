package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup

	numbers := []uint32{123453, 134135, 235, 1241387, 32135647, 41788}
	wg.Add(len(numbers)) //tamaño a procesar las cosas

	fmt.Println("Comienza el proceso...")
	for _, v := range numbers {
		go func(a uint32) {
			defer wg.Done() // quita del tamaño que falta procesara
			fmt.Println(a, EsPrimo(a))
		}(v)
	}
	wg.Wait()
	fmt.Println("Termino el proceso")
}

func EsPrimo(a uint32) bool {
	c := 0
	var i uint32
	for i = 1; i <= a; i++ {
		if a%i == 0 {
			c++
		}
	}
	if c == 2 {
		return true
	}
	return false
}
