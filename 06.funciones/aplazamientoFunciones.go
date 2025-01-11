package main

import "fmt"

func main() {
	//Defer aplaza la ejecucion hasta el final de la funcion que contiene a la funcion que contiene a la
	//afectada por el defer
	defer hola1()
	adios1()
}

func hola1() {
	fmt.Println("Hola")
}

func adios1() {
	fmt.Println("Adios")
}
