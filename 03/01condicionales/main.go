package main

import "fmt"

func main() {
	isValid := true

	if edad, nombre := 10, "edwin"; edad < 14 {
		fmt.Println(nombre, "es niño")
	} else if edad < 18 {
		fmt.Println(nombre, "es Adolescente")
	} else if edad < 30 {
		fmt.Println(nombre, "Es un adulto")
	} else {
		fmt.Println(nombre, "Es un adulto grandesito ")
	}

	if 5 > 10 {
		fmt.Println("Esto esta en el bloque verdadero")
		fmt.Printf("El tipo de isValid %T\n", isValid)
		if 5 < 10 {
			fmt.Println("5 es menor a 10")
		} else {
			fmt.Println("5 no es menor a 10 ")
			fmt.Println(isValid)
		}
	} else {
		fmt.Println("Esto esta en el bloque falso")
	}

	fmt.Println("se acabo el programa ")
}
