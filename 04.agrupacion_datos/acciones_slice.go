package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	//Agregar elementos a un array
	x = append(x, 6)
	fmt.Println(x)

	//Borrar elementos de un slice
	//Para ello lo debemos hacer en otro slice en el que pasaremos
	//el rango de valores que deseamos borrar
	w := append(x[:1], x[3:]...)
	fmt.Println(w)

	//Make
	//Primero se indica el tipo de dato del slice, despues la longitud y despues la capacidad
	sliceNuevo := make([]int, 10, 100)
	fmt.Println(sliceNuevo)

	//Slice bidimensional
	abc := []string{"a", "b", "c"}
	def := []string{"d", "e", "f"}

	abcdef := [][]string{abc, def}
	fmt.Println(abcdef)
}
