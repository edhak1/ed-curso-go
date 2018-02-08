package main

import "fmt"

func main() {
	/*
		anonima := func() {
			fmt.Println("Me imprimo estando en una variable anonima")
		}

		anonima()
	*/
	secuencia1 := secuencia()
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())

	fmt.Println("--------------")

	secuencia2 := secuencia()
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())

}

func secuencia() func() int { //retorna una función anonima() y esta función retorna un entero de 32bit
	var x int
	return func() int {
		x++
		return x
	}
}
