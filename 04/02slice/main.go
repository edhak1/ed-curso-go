package main

import "fmt"

func main() {
	//slice
	//1.- apuntador a un array
	//2.- tamaño
	//3.- Capacidad
	//make (tipo de dato del sline, tamaño inicial, capacidad inicial)

	nombres := make([]string, 0)
	/*
		var nombres []string
		nombres = append(nombres, "Daniel")
		fmt.Printf("Su tamaño es %d y su capacidad %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Alvaro")
		fmt.Printf("Su tamaño es %d y su capacidad %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Alexys")
		fmt.Printf("Su tamaño es %d y su capacidad %d\n", len(nombres), cap(nombres))
		nombres = append(nombres, "Pedro")
		fmt.Printf("Su tamaño es %d y su capacidad %d\n", len(nombres), cap(nombres))
	*/
	nombres = []string{
		"Alvaro",
		"Daniel",
		"Alexys",
	}
	fmt.Println(nombres)
}
