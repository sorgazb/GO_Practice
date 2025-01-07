package main

import "fmt"

func main() {
	fmt.Println(variosReturn(1, "Hola"))
}

func variosReturn(a int, b string) (valor1 int, valor2 string) {
	return a, b
}
