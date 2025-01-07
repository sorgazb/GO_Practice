package main

import "fmt"

func main() {
	recibirEnteros(1)
	recibirEnteros(2, 3, 4, 5)
	fmt.Println(sumarEnteros(1, 2, 3, 4, 5, 6, 7, 9, 10))
}

// Los ... significan que puede recibir de 0 a N valores por parametros
func recibirEnteros(n ...int) {
	fmt.Println(n)
	//Obtenemos el tipo de dato
	fmt.Printf("%T\n", n)
}

func sumarEnteros(n ...int) int {
	suma := 0
	for _, i := range n {
		suma = suma + i
	}
	return suma
}
