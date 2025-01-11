package main

import "fmt"

func main() {
	sliceInt := []int{1, 2, 3, 4, 5}
	sumaNumeros2(sliceInt...)
	sumaNumeros3(sliceInt)
}

func sumaNumeros2(x ...int) {
	fmt.Println(x)
	suma := 0
	for _, v := range x {
		suma += v
	}
	fmt.Println(suma)
}

func sumaNumeros3(x []int) {
	fmt.Println(x)
	suma := 0
	for _, v := range x {
		suma += v
	}
	fmt.Println(suma)
}
