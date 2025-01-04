package main

import "fmt"

func main() {
	mapa := map[string]int{
		"Sergio": 12345,
		"Pedro":  1213123,
	}

	fmt.Println(mapa)

	if _, ok := mapa["Pedro"]; ok {
		fmt.Println("Pedro esta dentro del mapa")
	}

	if _, ok := mapa["Luis"]; ok {
		fmt.Println("Luis esta dentro del mapa")
	} else {
		fmt.Println("Luis no esta dentro del mapa")
	}

	//Recorrer un mapa
	for clave, valor := range mapa {
		fmt.Println(clave, valor)
	}

	//Borrar elementos de un mapa, se borra por la clave
	delete(mapa, "Pedro")
	fmt.Println(mapa)
}
