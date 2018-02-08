package main

import "fmt"

func main() {
	//const nombre = "Pedro"
	//fmt.Println(nombre)
	var a int
	var b int8
	n := "Pedro"
	m := "Goméz"
	a = 121212
	b = 5

	//casting
	c := a + int(b)
	fmt.Println(c)
	fmt.Printf("hola %s %s %d cómo estas?\n", n, m, b)
	fmt.Printf("C es de tipo: %T\n", b)
	//prioridad aritmetica
	//+ - * /
	// () * / +

	d := 6 + 6*6 - 6
	fmt.Println(d)

	var nombre string
	fmt.Println(nombre)
}
