package main

import "fmt"

func main() {
	/*
		numeros := []string{"cero", "uno", "dos", "tres"} //slice

		for _, v := range numeros { //range devuelve dos valores el indice y el valor(el valor se puede omitir y indice no)
			fmt.Println(v)
		}
	*/
	/*
		nombres := map[string]string{"es": "España", "co": "Colombia", "br": "Brasil"}
		for _, v := range nombres {
			fmt.Println(v)
		}
	*/
	frase := "Hola Mundo"

	for _, letra := range frase {
		fmt.Println(string(letra))
	}

	for _, letra := range "Go desde cero" {
		fmt.Println(string(letra))
	}
}
