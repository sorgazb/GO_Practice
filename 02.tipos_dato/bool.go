package main

import "fmt"

func main() {
	x := 10
	y := 20

	fmt.Println("X:", x, "Y:", y)
	fmt.Println("X == Y:", x == y) //Igual
	fmt.Println("X != Y:", x != y) //Distinto
	fmt.Println("X < Y:", x < y)   //Menor
	fmt.Println("X > Y:", x > y)   //Mayor
	fmt.Println("X <= Y:", x <= y) //Menor o igual
	fmt.Println("X >= Y:", x >= y) //Mayor o igual

	//No se puede comparar datos de distinto tipo

	//var w int
	//w = 10
	//type NewInt int
	//var z NewInt
	//z = 10
	//fmt.Println("W == Z", w == z)
}
