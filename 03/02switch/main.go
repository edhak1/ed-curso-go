package main

import "fmt"

func main() {
	var a uint8
	a = 1
	switch a {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sábado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("No es un día de semana")
	}
	/*
		switch a := 3 {
		case 1:
			fallthrough //si a es uno continua al siguente
		case 2:
			fallthrough // se a es dos continua al siguiente
		case 3:
			fallthrough // lo mismo anterior
		case 4:
			fallthrough //lo mismo anterior
		case 5:
			fmt.Println("Estas entre semana")
		case 6:
			fallthrough
		case 7:
			fmt.Println("Estas en fin de semana")
		default:
			fmt.Println("No es un día de semana")
		}
	*/

	switch a := 3; {
	case a >= 0 && a <= 5:
		fmt.Println("Estas entre semana")
	case a >= 6 && a <= 7:
		fmt.Println("Estás en fin de semana")
	default:
		fmt.Println("No es un día valido")
	}

}
