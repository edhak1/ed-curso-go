package main

import "fmt"

func main() {
	//pasar por valor no tranforma el valor de la variable
	//pasar or puntero transforma el valor

	a := 3
	fmt.Println("Antes de duplicar a =", a)
	duplicar(&a) //¬dirección de memoria
	fmt.Println("Después de duplicar vale a = ", a)

}

func duplicar(a *int) {
	*a = *a * 2 // valor de la dirección de memoria(*x)
	fmt.Println("Dentro de la funcion duplicar vale a = ", *a)
}
