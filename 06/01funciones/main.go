package main

import "fmt"

func saludar(nombre string, edad uint8) {
	fmt.Printf("Hola %s tienes %d\n", nombre, edad)
	if edad > 30 {
		return // corta el programa si esta la condicional es verdadera
	}
	fmt.Println("Eres Joven")
}

func main() {
	saludar("Alexys", 25)
}
