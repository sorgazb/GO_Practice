/*
Crea 4 variables de diferentes tipos
Asignale a una tu edad,
a otra tu nombre,
a otra si te esta gustando el curso
y por ultimo el numero pi con 4 decimales
*/
package main

import "fmt"

func main() {
	var edad int
	var nombre string
	var gustoCurso bool
	var pi float32
	edad = 22
	nombre = "Sergio"
	gustoCurso = true
	pi = 3.1416

	fmt.Println(edad, nombre, gustoCurso, pi)
}
