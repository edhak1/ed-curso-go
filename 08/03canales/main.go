package main

import "fmt"

func main() {
	bodegaOrigen := []string{"php", "javascript", "html", "css", "java", "base de datos", "git"}

	bodegaDestino := []string{}

	fotocopiadora := make(chan string)
	fotocopiado := make(chan string)

	//Se carga la fotocopiadora
	go func() {
		for _, libro := range bodegaOrigen {
			fotocopiadora <- libro
		}
	}()
	//Se deja en la bodega bodegaDestino
	go func() {
		for {
			//entrega el contenido de la fotocopiadora a la varible libre
			libro := <-fotocopiadora
			fotocopiado <- libro

			//agregar el libro a la bodega destino
			bodegaDestino = append(bodegaDestino, libro)
			if len(bodegaOrigen) == len(bodegaDestino) {
				//cerrar canal del fotocopiado
				close(fotocopiado)
			}
		}
	}()

	fmt.Println("**listado de lisbros fotocopiados **")
	for {
		if libro, ok := <-fotocopiado; ok {
			fmt.Println(libro)
		} else {
			break
		}
	}
}
