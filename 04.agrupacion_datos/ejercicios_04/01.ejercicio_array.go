//Enunciado:
/*
Crea un array que contiene 5 numeros enteros. (pedidos por el usuario)
Imprime los valores con su indice correspondiente
*/
package main

import "fmt"

func main() {
	var numeros [5]int
	for i := 0; i < len(numeros); i++ {
		fmt.Println("Introduce el numero de la casilla", i)
		fmt.Scan(&numeros[i])
	}
	fmt.Println(numeros)
}
