package main

import "fmt"

type Persona struct {
	nombre string
	edad   int8
}

type Numero int

func (n Numero) presentarse() {
	fmt.Println("Hola, yo soy un número")
}
func (p Persona) saludar() {
	fmt.Println("Hola soy una persona")
}

func (p *Persona) asignaredad(i int8) {
	if i > 0 {
		p.edad = i
	} else {
		fmt.Println("No es valida la edad es negativa")
	}
}

func main() {
	var persona Persona
	var numero Numero
	numero.presentarse()
	persona.saludar()
	persona.asignaredad(-35)
	fmt.Println(persona)
}
