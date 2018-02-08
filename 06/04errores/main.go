package main

import (
	"errors"
	"fmt"
)

func main() {
	r, err := division(100, 0)
	if err != nil {
		fmt.Println(err)
		return //no ejecute nada más de las siguientes intrucciones
	}
	fmt.Println(r)
}

func division(dividendo, divisor float64) (resultado float64, err error) {
	if divisor == 0 {
		err = errors.New("No se puede dividir por cero")
	} else {
		resultado = dividendo / divisor
	}
	return
}
