package main

import "fmt"

func main() {

	//puntero es una direccion de memoria, direccion en la cual esta alojado un valor

	valor := 20

	var puntero1,puntero2 *int

	puntero1 = &valor
	puntero2 = &valor

	fmt.Println(puntero1)
	fmt.Println(puntero2)

	fmt.Println(*puntero1)
	fmt.Println(*puntero2)

	*puntero2 = 300
	fmt.Println(puntero1)
	fmt.Println(puntero2)

	fmt.Println(*puntero1)
	fmt.Println(*puntero2)
	
}