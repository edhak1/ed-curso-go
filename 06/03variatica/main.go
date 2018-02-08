package main

import "fmt"

func main() {
	saludarVarios(20, "Alexys", "Pedro", "Daniel", "Alvaro")
}

func saludarVarios(edad uint8, nombres ...string) {
	fmt.Printf("%T\n", nombres)
	for _, v := range nombres {
		fmt.Println("HOla", v, "tienes", edad, " años de edad")
	}
}
