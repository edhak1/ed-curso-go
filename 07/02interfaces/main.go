package main

import (
	"github.com/edhak1/ed-curso-go/07/02interfaces/animales"
)

func main() {
	var p animales.Perro
	var g animales.Gato

	p.Nombre = "firlais"
	g.Nombre = "chispar"

	//AdoptarPerro(p)
	//AdoptarGato(g)
	AdoptarMarcota(p)
	AdoptarMarcota(g)

}

/*
func AdoptarPerro(p animales.Perro) {
	p.Comunicarse()
}

func AdoptarGato(g animales.Gato) {
	g.Comunicarse()
}
*/

func AdoptarMarcota(m animales.Mascota) {
	m.Comunicarse()
}
