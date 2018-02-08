package main

import (
	"fmt"

	"github.com/edhak1/ed-curso-go/06/07paquetes/saludar"

	"github.com/edhak1/ed-curso-go/06/07paquetes/despedida"
) // se ccoloca el path des de src en adelante

func main() {
	nombre := "Gente del futuro"
	saludar.Saludar(nombre)
	saludar.MeVes = "Esto es un texto asigando desde el main"
	fmt.Println(saludar.MeVes)

	despedida.Despedirse("Gente del futuro ")
}
