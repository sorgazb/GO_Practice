package main

import "fmt"

type persona struct {
	nombre string
	edad   int
}

type estudiante struct {
	persona
	id int
}

func (e estudiante) hablar() {
	fmt.Println("Hola, mi nombre es:", e.nombre, " mi id es:", e.id, " mi edad es:", e.edad)
}

func (p persona) caminar() {
	fmt.Println("Estoy caminando")
}

func (p persona) respirar() {
	fmt.Println("Estoy respirado")
}

type serVivo interface {
	respirar()
}

func main() {
	est1 := estudiante{
		persona: persona{
			nombre: "Sergio",
			edad:   21,
		},
		id: 1234,
	}

	est1.hablar()
	est1.respirar()
}
