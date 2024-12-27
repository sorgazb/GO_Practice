package main

import "fmt"

var edad int
var nombre string
var dejasteResenia bool

func main() {
	fmt.Printf("%v,%v,%v\n", edad, nombre, dejasteResenia)
	edad = 22
	nombre = "Sergio"
	dejasteResenia = true
	fmt.Printf("%v,%v,%v\n", edad, nombre, dejasteResenia)
}
