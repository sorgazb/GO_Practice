package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x)
	fmt.Println(x[0])
	//Mostrar por rangos un slice
	fmt.Println(x[:])
	fmt.Println(x[1:8])
}
