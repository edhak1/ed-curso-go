package main

import "fmt"

func main() {
	/*
		idiomas := make(map[string]string) //declaración map[tipo de llave key] tipo de dato
		idiomas["es"] = "español"
		idiomas["en"] = "Inglés"
		idiomas["fr"] = "Francés"
	*/

	/*
	   	idiomas := map[string]string{
	   		"es": "español",
	   		"en": "Inglés",
	   		"fr": "frances",
	   		"pt": "Portugues",
	   	}


	     //
	   	//	fmt.Println(idiomas["fr"])
	   	//delete(idiomas, "en")
	   	//	fmt.Println(idiomas)


	   	if idioma, ok := idiomas["pt"]; ok { //devuelve un array y un valor boleano (map devuelve dos valores)
	   		fmt.Println("La posición pt si existe y vale", idioma)
	   	} else {
	   		fmt.Println("La posición pt no Existe")
	   	}
	*/

	numeros := map[int]string{
		1:    "Uno es un numero",
		2016: "Esto esotro número",
		-4:   "Aqui la llave no es negativa",
	}

	fmt.Println(numeros)
}
