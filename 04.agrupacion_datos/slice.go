package main

import "fmt"

func main() {
	//Declaracion de un slice
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)

	//Recorrer un slice o un array con un bucle
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	//Este bucle nos indica el indice en el que estamos y el valor de ese indice
	for indice, valor := range x {
		fmt.Println(indice, valor)
	}

	//Si no queremos que nos indique el indice, debemos poner _ en la parte del indice
	for _, valor := range x {
		fmt.Println(valor)
	}
}
