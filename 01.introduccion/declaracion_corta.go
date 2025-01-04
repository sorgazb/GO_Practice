package main

import (
	"fmt"
)

func main() {
	//Declaracion corta
	nombre := "Sergio"

	//Declaracion tradicional
	var apellido string = "Orgaz"

	edad := 22
	var anio int = 2002

	fmt.Println(nombre, apellido, edad, anio)

	nombre = "Raul"
	fmt.Println(nombre)

	//Declarar una variable sin inicializar
	var escritorio string
	escritorio = "madera"
	fmt.Println(escritorio)

}
