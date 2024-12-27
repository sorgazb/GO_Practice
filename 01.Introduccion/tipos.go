package main

import "fmt"

func main() {
	var x int

	type edad uint8
	var y edad

	fmt.Printf("Tipo: %T, Valor: %v", x, x)
	fmt.Printf("Tipo: %T, Valor: %v", y, y)
}
