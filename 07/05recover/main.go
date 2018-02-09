package main

import "fmt"

func main() {
	f()
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%T\n", r)
			fmt.Println("Recuperado en f: ", r)
		}
	}() //con dos parentesis se autoejecuta
	fmt.Println("Lamando a g.")
	g(5)
	fmt.Println("Retorna normalmente desde g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("AAAAAAAAaaaaaaaaa")
		panic("Número no puede ser mayor que 3")
	}
	defer fmt.Println("Defer en la funcion g")
	fmt.Println("Imprimiendo en g", i)
}
