package main

import "fmt"

func main() {
	//Bucle for con autoincremento o autodecremento
	//Imprimir por pantalla los numeros del 1 al 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//Bucle while
	//Haz un bucle que diga si es mayor de edad y que reste años de 1 en 1
	edad := 22
	for edad >= 18 {
		fmt.Println("Eres mayor de edad")
		edad--
	}

	//Bucle do while
	//for{
	//	if(){
	//		//Condicion para romper el bucle
	//	}
	//}
}
