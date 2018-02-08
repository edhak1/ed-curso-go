package main

import "fmt"

func main() {
	//arrays
	//guarda solo de un tipo
	//tamaño fijo

	var nombres [3]string
	/*
		nombres[0] = "Daniel"
		nombres[1] = "Jose"
		nombres[2] = "Jesus"
	*/
	nombres = [3]string{
		"Daniel",
		"alvaro",
		"Alexys",
	}
	//nombre := nombres[1]
	size := len(nombres)

	fmt.Println("El tamaño del array es ", size)
	//fmt.Println(nombres[inicio:final(excluyentes)])
	fmt.Println(nombres)
	//fmt.Println(nombres[1:2])
}
