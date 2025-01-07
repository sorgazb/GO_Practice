package main

import "fmt"

func main() {
	hola()
	holaTexto("Sergio")
	adios()

	fmt.Println(sumaNumeros(1, 2))
	sumaTotal := sumaNumeros(1, 2)
	fmt.Println(sumaTotal)
}

func hola() {
	fmt.Println("Hola desde funcion hola")
}

func holaTexto(texto string) {
	fmt.Println("Hola", texto)
}

func adios() {
	fmt.Println("Adios desde funcion adios")
}

func sumaNumeros(numero1 int, numero2 int) int {
	return numero1 + numero2
}
