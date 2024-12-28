package main

import "fmt"

const (
	nombre   string  = "Sergio"
	edad     int     = 22
	pi       float32 = 3.1416
	apellido string  = "Orgaz"
)

func main() {

	//var (
	//	nombre   string  = "Sergio"
	//	edad     int     = 22
	//	pi       float32 = 3.1416
	//	apellido string  = "Orgaz"
	//)

	fmt.Println(nombre)
	fmt.Println(apellido)
	fmt.Println(edad)
	fmt.Println(pi)
	fmt.Printf("%T\n", nombre)
	fmt.Printf("%T\n", apellido)
	fmt.Printf("%T\n", edad)
	fmt.Printf("%T\n", pi)
}
