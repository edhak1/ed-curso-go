package main

import (
	"fmt"
	"time"
)

func main() {
	var h string
	go MostrarNumeros() //goroutine(proceso que se ejecuta de manera independiente) (no parallismo  no hilo)
	fmt.Print("Digita algo: ")
	fmt.Scan(&h)
	fmt.Println("Digitaste", h)
}

func MostrarNumeros() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
