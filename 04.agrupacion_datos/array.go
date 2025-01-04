package main

import "fmt"

func main() {
	//Declaracion de un array
	var anios [3]int
	fmt.Println(anios)

	anios[0] = 2001
	fmt.Println(anios)

	anios[2] = 2025
	fmt.Println(anios)

	hola := []string{"hola", "mundo", "!"}
	fmt.Println(hola)

	//Para ver la longitud de un array
	fmt.Println(len(hola))
}
