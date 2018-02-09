package main

import "fmt"

func main() {
	fmt.Println("Conectando a la base de datos")
	defer cerrarDB() //cuando termine el main se ejecuta

	fmt.Println("Ceonsultar la base de datos")
	defer cerrarSetDatos() // se ejecutar primero (se ejecuta en orden inverso)

	//cerrarSetDatos() el difere anterior es equivalente al que tenemos ahora
	//cerrarDB()
}

func cerrarDB() {
	fmt.Println("Cerrar la base de datos")
}

func cerrarSetDatos() {
	fmt.Println("Cerrar set de datos")
}
