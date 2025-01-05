package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
	estatura float32
}

func main() {
	persona1 := persona{
		nombre:   "Juan",
		apellido: "Perez",
		edad:     25,
		estatura: 1.75,
	}

	persona2 := persona{
		nombre:   "Pedro",
		apellido: "Gonzalez",
		edad:     30,
		estatura: 1.80,
	}

	fmt.Println(persona1)
	fmt.Println(persona2)
}
