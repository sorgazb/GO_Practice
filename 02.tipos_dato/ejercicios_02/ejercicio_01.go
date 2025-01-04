// Escribe un programa que imprima un numero entero, decimal, binario y hexadecimal
// Pideselos al usuario por consola
package main

import "fmt"

func main() {
	var edad int
	fmt.Println("Introduce tu edad:")
	fmt.Scanf("%d", &edad)
	fmt.Println("Tu edad es:", edad)
	fmt.Printf("%v,%d,%b,%x", edad, edad, edad, edad)
}
