//Enunciado
/*
Usando el mapa del ejercicio anterior, agrega un nuevo contacto con el nombre
Juan y los telefonos 1234567,12313131313
Borra el contacto de Juan anterior antes de agregar el nuevo contacto
Imprime el mapa actualizado

Pidele al usuario el nombre del contacto a borrar y luego imprime el mapa acualizado
*/
package main

import "fmt"

func main() {
	mapa := map[string][]int{
		"Sergio": {1212121, 1212121},
		"Pedro":  {1212121, 1212121},
		"Juan":   {1234567, 12313131313},
	}

	fmt.Println(mapa)

	delete(mapa, "Juan")

	fmt.Println(mapa)
}
