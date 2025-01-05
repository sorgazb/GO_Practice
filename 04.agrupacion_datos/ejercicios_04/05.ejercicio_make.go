//Enunciado
/*
Crea un slice con nombre de alumnos, normalmente los grupos son de 5 personas
pero luego se inscriben mas personas, prevee eso, el maximo puede ser de 10
que el nombre de los alumnos los ingrese el usuario

Cual sera la capacidad del slice ?
Cual seria la longitud del slice ?
Imprime el contenido del slice con sus indices y sus valores
*/
package main

import "fmt"

func main() {
	sliceNombres := make([]string, 5, 10)
	fmt.Println(sliceNombres)

	for i := 0; i < len(sliceNombres); i++ {
		fmt.Println("Introduce el nombre de los alumnos del slice")
		fmt.Scan(&sliceNombres[i])
	}

	fmt.Println(sliceNombres)
}
