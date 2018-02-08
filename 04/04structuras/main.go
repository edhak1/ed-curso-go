package main

import "fmt"

//Persona es un tipo de dato estructura,
type Persona struct {
	Nombre string // se imprimen en el mismo orden en que se declaran
	Edad   uint8
	Emails []string
}

func main() {
	/*
		var persona1 Persona
		persona1.Nombre = "Pedro"
		persona1.Edad = 20
		fmt.Println(persona1)
	*/
	//emails := string{"a@b.com", "z@b.com"}

	persona2 := Persona{
		"Pablo",
		33,
		[]string{"a@g.com", "b@f.com"},
	}

	fmt.Println(persona2)
}
