package main

import "fmt"

func main() {
	/*
		var n1, n2 int8
		n1 = 15
		n2 = 5

		r := suma(n1, n2)
		fmt.Println(r)
	*/
	/*
		  var edad int8
			edad = 30
			fmt.Println(tipoEdad(edad))
	*/
	n := []int8{52, 63, 47, 5, 123, 45, 67, 90}
	maximo, minimo := maxymin(n)
	fmt.Println("Maximo", maximo)
	fmt.Println("Minimo", minimo)
}

func maxymin(numeros []int8) (max int8, min int8) {
	for _, v := range numeros {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return
}

func suma(a, b int8) int8 { //(variables de entrada) variable de salida
	return a + b
}

func tipoEdad(edad int8) string {
	var tipo string
	switch {
	case edad < 12:
		tipo = "Niñez"
	case edad < 18:
		tipo = "Adolescente"
	default:
		tipo = "Madurez"
	}
	return tipo
}
