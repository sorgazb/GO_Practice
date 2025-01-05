//Enunciado
/*
Crea un mapa con tus contactos que tengan nombre y telefono
Para cada contacto, imprime el nombre y el telefono
Guarda mas de 2 numeros por persona
Pista, usa un slice para guardar los telefonos
*/
package main

import "fmt"

func main() {
	mapa := map[string][]int{
		"Sergio": {1212121, 1212121},
		"Pedro":  {1212121, 1212121},
	}

	fmt.Println(mapa)
}
