//Enunciado
/*
slice := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
Agrega el elemento 11 al final del slice
Imprime el slice
Agrega los valores del slice a un nuevo slice con valores agregados de 12, 13, 14, 15
Imprime el nuevo slice
*/
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice = append(slice, 11)
	fmt.Println(slice)

	sliceNuevo := []int{12, 13, 14, 15}
	sliceNuevo = append(sliceNuevo, slice...)
	fmt.Println(sliceNuevo)
}
