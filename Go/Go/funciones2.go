package main

import "fmt"

func suma(numeros ...int) int{
	resultado:=0
	for _,num := range numeros{
		resultado+=num
	}
	return resultado
}

func main() {
	fmt.Println(suma(1,10,56,89,34))
	fmt.Println(suma(1,10,56,89,34,89,168))
}