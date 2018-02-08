package saludar

import "fmt"

// MeVes espara utilizar desde otro paquete
var MeVes string
var NoMeVes string

func Saludar(nombre string) { // S mayuscula (MAYUSCULAS se usa pra saber si es una función exportada)
	fmt.Println("Hola", nombre)
}

func noVisible() { // nobres que empiezan con minuscula, entocnes no son visibles, por otros programas
	fmt.Println("No es visible")
}
