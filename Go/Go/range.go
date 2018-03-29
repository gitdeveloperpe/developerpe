package main 

import "fmt"

func main() {
	
	apellidos := [] string{
		"Sandoval",
		"Hernandez",
		"Tandazo",
		"Arevalo",
		"Quezada",
		"Palomino",
	}

	for _,apellido := range apellidos{
		fmt.Println("El apellido es ",apellido)
	}
}