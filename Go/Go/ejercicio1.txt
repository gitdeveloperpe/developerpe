package main

import (
	"fmt"
)

func main(){

	cabecera := `
	***************************************************
	Cantidad de numero impares en un determinado rango
	***************************************************
	`
	//Imprimimos nuestra cabecera
	fmt.Println(cabecera)

	//Solicitamos el rango inferior
	fmt.Println("Ingrese el rango inferior: ")
	var numero1 int
	fmt.Scanln(&numero1)

	//Solicitamos el rango superior
	fmt.Println("Ingrese el rango superior")
	var numero2 int
	fmt.Scanln(&numero2)

	//contador de numeros impares
	var contador int

	for i := numero1; i<= numero2; i++{
		//hallamos el numero impar
		if i%2 !=0{
			contador++
			fmt.Printf("El numero %d es un numero impar \n",i)
		}
	}

	fmt.Println("\n\n ************RESULTADO *************")

	//imprimimos la cantidad de numeros impares hallados

	fmt.Printf("Entre el %d y el %d hay %d numero impares",numero1,numero2,contador)


}