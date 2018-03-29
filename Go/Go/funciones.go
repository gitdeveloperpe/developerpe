package main

import "fmt"

func imprimirNombre(nombre string){
	fmt.Println("El nombre es "+nombre)
	fmt.Println(calcularSuma(8,10))
}

func calcularSuma(a int, b int) int {
	return a+b
}


func main() {
	imprimirNombre("Jose")
	a:=calcularSuma(10,8)
	fmt.Println(a)
}